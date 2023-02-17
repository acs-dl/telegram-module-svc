package urlval

import (
	"encoding"
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/distributed_lab/urlval/internal/betterreflect"
)

// Encode encodes given struct with urlval tags into url query string.
func Encode(src interface{}) (string, error) {
	refsrc := betterreflect.NewStruct(src)
	values := url.Values{}

	err := populateValues(values, refsrc)
	if err != nil {
		return "", err
	}

	return values.Encode(), nil
}

//MustEncode - encodes given struct with urlval tags into url query string. Panics on failure
func MustEncode(src interface{}) string {
	res, err := Encode(src)
	if err != nil {
		panic(errors.Wrap(err, "failed to encode"))
	}

	return res
}

func populateValue(values url.Values, refsrc *betterreflect.Struct, i int) error {
	refValue := refsrc.Value(i)
	if _, implements := refValue.Interface().(encoding.TextMarshaler); refValue.Type().Kind() == reflect.Struct && !implements {
		err := populateValues(values, betterreflect.NewStructFromValue(refValue))
		if err != nil {
			return err
		}
		return nil
	}

	// ignore nil fields
	if refValue.Kind() == reflect.Ptr && refValue.IsNil() {
		return nil
	}

	switch name, tag := getTag(refsrc, i); name {
	case "include":
		if refValue.Bool() {
			includes := values.Get("include")
			if includes == "" {
				includes = tag
			} else {
				includes += "," + tag
			}
			setValue(values, "include", includes)
		}
	case "page":
		key := fmt.Sprintf("page[%s]", tag)
		value := refValue
		setValue(values, key, value)
	case "filter":
		fieldValue := refValue.Interface()
		key := fmt.Sprintf("filter[%s]", tag)

		list, err := betterreflect.ConvertToStringSlice(fieldValue)
		if err != nil {
			return err
		}
		setValue(values, key, list)
	case "url":
		list, err := betterreflect.ConvertToStringSlice(refValue.Interface())
		if err != nil {
			return err
		}
		setValue(values, tag, list)
	}
	return nil
}

func populateValues(values url.Values, refsrc *betterreflect.Struct) error {
	for i := 0; i < refsrc.NumField(); i++ {
		// ignore private
		if refsrc.IsPrivate(i) {
			continue
		}
		err := populateValue(values, refsrc, i)
		if err != nil {
			return errors.From(err, logan.F{"field": refsrc.Field(i).Name})
		}
	}
	return nil
}

func setValue(values url.Values, key string, value interface{}) {
	if strVal := toString(value); strVal != "" {
		values.Set(key, strVal)
	}
}

func toString(value interface{}) string {
	if v, ok := value.(reflect.Value); ok {
		value = v.Interface()
	}

	if v, ok := value.([]string); ok {
		value = strings.Join(v, ",")
	}

	// some magic to convert values of custom aliased types to their,
	// underlying type, because cast fails to do this:
	if value = betterreflect.ConvertToUnderlyingType(value); value == nil {
		return ""
	}

	return cast.ToString(value)
}

func getTag(refsrc *betterreflect.Struct, i int) (name, tag string) {
	var names = []string{
		"include",
		"page",
		"filter",
		"sort",
		"url",
	}
	for _, name = range names {
		if tag = refsrc.Tag(i, name); tag != "" {
			return name, tag
		}
	}

	return "", ""
}
