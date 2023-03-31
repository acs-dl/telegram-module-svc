package betterreflect

import (
	"encoding"
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type Struct struct {
	src reflect.Value
}

func NewStruct(src interface{}) *Struct {
	return &Struct{
		src: reflect.ValueOf(src),
	}
}

func NewStructFromValue(src reflect.Value) *Struct {
	return &Struct{
		src: src,
	}
}

func (s *Struct) srcvalue() reflect.Value {
	return reflect.Indirect(s.src)
}

func (s *Struct) srctype() reflect.Type {
	return s.srcvalue().Type()
}

func (s *Struct) NumField() int {
	return s.srcvalue().Type().NumField()
}

func (s *Struct) Tag(i int, key string) string {
	return s.srctype().Field(i).Tag.Get(key)
}

func (s *Struct) TagInfo(i int, key string) (string, bool) {
	value := s.srctype().Field(i).Tag.Get(key)
	if value == "" {
		return "", false
	}
	list := strings.SplitN(value, ",", 1)
	if len(list) > 2 || len(list) == 2 && list[1] != "required" {
		panic(fmt.Sprintf("urlval: unknown tag format of field %s, expected `filter:\"value[,required]\"`, "+
			"`url:\"value[,required]\"` or `page:\"value[,required]\", got %s", s.Field(i).Name, value))
	}
	required := false
	if len(list) == 2 {
		required = list[1] == "required"
	}
	return list[0], required
}

func (s *Struct) Type(i int) reflect.Type {
	return s.srctype().Field(i).Type
}

func (s *Struct) Value(i int) reflect.Value {
	return s.srcvalue().Field(i)
}

func (s *Struct) Field(i int) reflect.StructField {
	return s.srcvalue().Type().Field(i)
}

func (s *Struct) Set(i int, value interface{}) (err error) {
	switch v := value.(type) {
	case string:
		return parseIntoValue(s.Value(i), v)
	case []string:
		return parseIntoValue(s.Value(i), v...)
	}
	setValue(s.Value(i), value)
	return nil
}

// IsPrivate returns is the field private (unexported) or not.
func (s *Struct) IsPrivate(i int) bool {
	r, _ := utf8.DecodeRuneInString(s.Field(i).Name)
	return strings.ToLower(string(r)) == string(r)
}

// IsRequired returns if given tag has required flag set.
func (s *Struct) IsRequired(i int, key string) bool {
	_, required := s.TagInfo(i, key)
	return required
}

var ErrSingleValueExpected = errors.New("expected single value, got slice")

func parseIntoValue(dest reflect.Value, input ...string) error {
	if len(input) == 0 || !dest.CanSet() {
		return nil
	}

	t := dest.Type()
	kind := t.Kind()

	if kind == reflect.Ptr {
		v := dest
		if dest.IsNil() {
			v = reflect.New(t.Elem())
		}

		err := parseIntoValue(v.Elem(), input...)
		if err != nil {
			return err
		}
		dest.Set(v)
	}
	if un, impl := dest.Addr().Interface().(encoding.TextUnmarshaler); impl {
		err := un.UnmarshalText([]byte(input[0]))
		if err != nil {
			return errors.Wrap(err, "text unmarshaling failed")
		}
		return nil
	}

	if len(input) > 1 && dest.Kind() != reflect.Slice {
		return ErrSingleValueExpected
	}
	var convertedValue interface{}
	var err error
	switch kind {
	case reflect.String:
		convertedValue = input[0]
	case reflect.Bool:
		convertedValue, err = cast.ToBoolE(input[0])
	case reflect.Int:
		convertedValue, err = cast.ToIntE(input[0])
	case reflect.Int8:
		convertedValue, err = cast.ToInt8E(input[0])
	case reflect.Int16:
		convertedValue, err = cast.ToInt16E(input[0])
	case reflect.Int32:
		convertedValue, err = cast.ToInt32E(input[0])
	case reflect.Int64:
		convertedValue, err = cast.ToInt64E(input[0])
	case reflect.Uint:
		convertedValue, err = cast.ToUintE(input[0])
	case reflect.Uint8:
		convertedValue, err = cast.ToUint8E(input[0])
	case reflect.Uint16:
		convertedValue, err = cast.ToUint16E(input[0])
	case reflect.Uint32:
		convertedValue, err = cast.ToUint32E(input[0])
	case reflect.Uint64:
		convertedValue, err = cast.ToUint64E(input[0])
	case reflect.Float32:
		convertedValue, err = cast.ToFloat32E(input[0])
	case reflect.Float64:
		convertedValue, err = cast.ToFloat64E(input[0])
	case reflect.Slice:
		convertedValue, err = parseSlice(dest.Type().Elem(), input)
	case reflect.Complex64, reflect.Complex128, reflect.Chan,
		reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr,
		reflect.Array, reflect.Struct, reflect.UnsafePointer:
	default:
		panic(fmt.Sprintf("unknown field kind: %v", kind))
	}
	if err != nil {
		return errors.Wrapf(err, "expected value to be %s, got %s", t, input)
	}

	if convertedValue == nil {
		return nil
	}
	setValue(dest, convertedValue)
	return nil
}

// ConvertToUnderlyingType takes a value ot custom type T and converts it to
// it's underlying types. If value is of built-in type, ConvertToUnderlyingType
// just returns value as provided.
func ConvertToUnderlyingType(value interface{}) interface{} {
	var rType reflect.Type
	var rValue = reflect.ValueOf(value)
	var rKind = rValue.Kind()

	if rKind == reflect.Ptr {
		if !rValue.Elem().IsValid() {
			return nil
		}

		rValue = rValue.Elem()
		rKind = rValue.Kind()
	}

	switch rKind {
	case reflect.String:
		rType = reflect.TypeOf("")
	case reflect.Bool:
		rType = reflect.TypeOf(false)
	case reflect.Int:
		rType = reflect.TypeOf(0)
	case reflect.Int8:
		rType = reflect.TypeOf(int8(0))
	case reflect.Int16:
		rType = reflect.TypeOf(int16(0))
	case reflect.Int32:
		rType = reflect.TypeOf(int32(0))
	case reflect.Int64:
		rType = reflect.TypeOf(int64(0))
	case reflect.Uint:
		rType = reflect.TypeOf(uint(0))
	case reflect.Uint8:
		rType = reflect.TypeOf(uint8(0))
	case reflect.Uint16:
		rType = reflect.TypeOf(uint16(0))
	case reflect.Uint32:
		rType = reflect.TypeOf(uint32(0))
	case reflect.Uint64:
		rType = reflect.TypeOf(uint64(0))
	case reflect.Float32:
		rType = reflect.TypeOf(uint32(0))
	case reflect.Float64:
		rType = reflect.TypeOf(uint64(0))
	case reflect.Slice:
		rType = rValue.Type()
	case reflect.Complex64, reflect.Complex128, reflect.Chan,
		reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr,
		reflect.Array, reflect.Struct, reflect.UnsafePointer:
		panic(errors.Errorf("got %s, when ConvertToUnderlyingType works only with primitive types", rValue.Type()))
	default:
		panic(fmt.Sprintf("unknown rKind: %v", rKind))
	}

	return rValue.Convert(rType).Interface()
}

func setValue(dest reflect.Value, v interface{}) {
	// to assign type to their aliases:
	if reflect.TypeOf(v) != dest.Type() && reflect.TypeOf(v).Kind() == dest.Type().Kind() {
		dest.Set(reflect.ValueOf(v).Convert(dest.Type()))
		return
	}

	dest.Set(reflect.ValueOf(v))
}

func IsZero(value interface{}) bool {
	reflectValue := reflect.ValueOf(value)
	switch reflectValue.Type().Kind() {
	case reflect.Ptr:
		ptr := reflectValue
		return ptr.IsNil() || ptr.Elem() == reflect.Zero(reflectValue.Type())
	case reflect.Slice:
		slice := reflectValue
		return slice.Len() == 0
	case reflect.Complex64, reflect.Complex128, reflect.Chan,
		reflect.Func, reflect.Interface, reflect.Map,
		reflect.Array, reflect.Struct, reflect.UnsafePointer:
		return false
	default:
		return value == reflect.Zero(reflectValue.Type()).Interface()
	}
}

// parseSlice takes a slice of strings and converts it's items to the "destElemType".
// Returns error, if types are not convertible to each other.
func parseSlice(destElemType reflect.Type, sourceSlice []string) (interface{}, error) {
	destSliceType := reflect.SliceOf(destElemType)
	destSlice := reflect.MakeSlice(destSliceType, len(sourceSlice), len(sourceSlice))

	if len(sourceSlice) == 0 {
		return destSlice.Interface(), nil
	}

	if _, ok := destSlice.Index(0).Addr().Interface().(encoding.TextUnmarshaler); ok && destElemType.Kind() != reflect.Interface {
		for j := 0; j < len(sourceSlice); j++ {
			err := destSlice.Index(j).Addr().Interface().(encoding.TextUnmarshaler).UnmarshalText([]byte(sourceSlice[j]))
			if err != nil {
				return nil, err
			}
		}
		return destSlice.Interface(), nil
	}

	for j := 0; j < len(sourceSlice); j++ {
		err := parseIntoValue(destSlice.Index(j), sourceSlice[j])
		if err != nil {
			return nil, err
		}
	}

	return destSlice.Interface(), nil
}

func convertToString(value interface{}) (*string, error) {
	if value == nil {
		return nil, nil
	}
	var result string
	var err error
	if marshaler, ok := value.(encoding.TextMarshaler); ok {
		if marshaler == nil {
			return nil, nil
		}
		bytes, err := marshaler.MarshalText()
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal value")
		}
		result = string(bytes)
	} else {
		value = ConvertToUnderlyingType(value)
		// NOTE: allows to use Stringer too
		result, err = cast.ToStringE(value)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

// ConvertToStringSlice converts value into string slice using MarshalText implementation
// or conversion from primitives to string.
func ConvertToStringSlice(value interface{}) ([]string, error) {
	if slice := reflect.ValueOf(value); slice.Kind() == reflect.Slice {
		strSlice := make([]string, 0, slice.Len())
		for i := 0; i < slice.Len(); i++ {
			str, err := convertToString(slice.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			if str == nil {
				continue
			}
			strSlice = append(strSlice, *str)
		}
		return strSlice, nil
	}
	str, err := convertToString(value)
	if err != nil {
		return nil, err
	}
	if str == nil {
		return nil, nil
	}
	return []string{*str}, nil
}
