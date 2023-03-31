package urlval

import (
	"encoding"
	"net/url"
	"reflect"
	"strings"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/distributed_lab/urlval/internal"
	"gitlab.com/distributed_lab/urlval/internal/betterreflect"
)

var errNotSupportedParameter = errors.New("query parameter is not supported for this endpoint")
var errMultipleSortKeys = errors.New("only one sort key is supported")

// Decode is decodes provided url values to destination struct.
// Using Decode requires your request to follow JSON API spec -  If it's
// values contains anything except from "include", "sort", "search",  "filter", "page",
// or query has parameters that are not tagged in dest -  Decode still populates
// dest but also returns an error. The only error type it returns is errBadRequest
// that is compatible with ape (https://gitlab.com/distributed_lab/ape),
// so can be rendered directly to client.
func Decode(values url.Values, dest interface{}) error {
	tokens := internal.Tokenize(values)
	refdest := betterreflect.NewStruct(dest)
	setDefaults(refdest)
	errs := errBadRequest{}

	for token := range tokens {
		ok, err := decodeToken(token, refdest)
		if err != nil {
			errs[token.Raw] = err
			continue
		}
		if !ok {
			errs[token.Raw] = errNotSupportedParameter
		}
	}

	return errs.Filter()
}

// DecodeSilently decodes provided url values to destination struct.
// Using DecodeSilently requires your request to follow JSON API spec -  If it's
// values contains anything except from "include", "sort", "search",  "filter", "page",
// or query has parameters that are not tagged in dest -  DecodeSilently populates
// dest and does not return an error (where Decode returns). The only error type it returns is errBadRequest
// that is compatible with ape (https://gitlab.com/distributed_lab/ape),
// so can be rendered directly to client.
func DecodeSilently(values url.Values, dest interface{}) error {
	tokens := internal.Tokenize(values)
	refdest := betterreflect.NewStruct(dest)
	setDefaults(refdest)
	errs := errBadRequest{}

	for token := range tokens {
		_, err := decodeToken(token, refdest)
		if err != nil {
			errs[token.Raw] = err
			continue
		}
	}

	return errs.Filter()
}

func setDefaults(s *betterreflect.Struct) {
	for i := 0; i < s.NumField(); i++ {
		value := s.Value(i)
		if value.Kind() == reflect.Struct {
			nestedStruct := betterreflect.NewStructFromValue(value)
			setDefaults(nestedStruct)
			continue
		}

		if value.CanSet() && betterreflect.IsZero(value.Interface()) && s.Tag(i, "default") != "" {
			if err := s.Set(i, splitTokenValue(s.Tag(i, "default"))); err != nil {
				panic(errors.Wrap(err, "failed to set default value"))
			}
		}
	}
}

func decodeToken(token internal.Token, dest *betterreflect.Struct) (bool, error) {
	var decoded bool

	for i := 0; i < dest.NumField(); i++ {
		var ok bool
		var err error

		if dest.IsPrivate(i) {
			continue
		}

		// TODO: extract struct traversing (in case, if required feature is needed; it will be then easier to implement).
		value := dest.Value(i)
		_, implements := value.Interface().(encoding.TextUnmarshaler)
		_, implementsPtr := value.Addr().Interface().(encoding.TextUnmarshaler)
		if dest.Type(i).Kind() == reflect.Struct && !implements && !implementsPtr {
			if ok, err = decodeToken(token, betterreflect.NewStructFromValue(value)); err != nil {
				return false, err
			}
			if ok {
				if decoded {
					panic(errors.New("decoding same token twice - probably your struct has 2 or more similar tags"))
				}

				decoded = true
				// not returning here, because we still need to traverse
				// whole struct to ensure we don't decode same token twice.
			}

			continue
		}

		if ok, err = trySet(dest, i, token); err != nil {
			return false, err
		}

		if ok {
			if decoded {
				panic(errors.New("decoding same token twice - probably your struct has 2 or more similar tags"))
			}

			decoded = true
			// not returning here, because we still need to traverse
			// whole struct to ensure we don't decode same token twice.
		}
	}

	return decoded, nil
}

func trySet(dest *betterreflect.Struct, i int, token internal.Token) (bool, error) {
	keys := splitTokenValue(token.Value)
	var value interface{} = keys
	switch token.Type {
	case internal.TokenTypeInclude:
		if dest.Tag(i, "include") == token.Key {
			if dest.Type(i).Kind() != reflect.Bool {
				panic("invalid destination type, expected bool for include tags")
			}

			value = true
		} else {
			return false, nil
		}
	case internal.TokenTypeFilter:
		if dest.Tag(i, "filter") != token.Key {
			return false, nil
		}
		if dest.Type(i).Kind() != reflect.Ptr &&
			dest.Type(i).Kind() != reflect.Slice {
			panic("invalid destination type, expected pointer or slice for filter tags")
		}
	case internal.TokenTypePage:
		if dest.Tag(i, "page") != token.Key {
			return false, nil
		}
	case internal.TokenTypeCustomParameter:
		if dest.Tag(i, "url") != token.Key {
			return false, nil
		}
	default:
		panic(errors.Errorf("unknown token type: %d", token.Type))
	}

	if value == nil {
		return false, nil
	}

	if err := dest.Set(i, value); err != nil {
		return false, errors.From(err, logan.F{"field": dest.Field(i).Name})
	}

	return true, nil
}

func splitTokenValue(value string) []string {
	splitted := strings.Split(value, ",")
	result := splitted[:0]
	for _, s := range splitted {
		if s == "" {
			continue
		}
		result = append(result, s)
	}
	return result
}
