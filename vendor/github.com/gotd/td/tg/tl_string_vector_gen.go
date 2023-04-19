// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// StringVector is a box for Vector<string>
type StringVector struct {
	// Elements of Vector<string>
	Elems []string
}

// StringVectorTypeID is TL type id of StringVector.
const StringVectorTypeID = bin.TypeVector

// Ensuring interfaces in compile-time for StringVector.
var (
	_ bin.Encoder     = &StringVector{}
	_ bin.Decoder     = &StringVector{}
	_ bin.BareEncoder = &StringVector{}
	_ bin.BareDecoder = &StringVector{}
)

func (vec *StringVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *StringVector) String() string {
	if vec == nil {
		return "StringVector(nil)"
	}
	type Alias StringVector
	return fmt.Sprintf("StringVector%+v", Alias(*vec))
}

// FillFrom fills StringVector from given interface.
func (vec *StringVector) FillFrom(from interface {
	GetElems() (value []string)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StringVector) TypeID() uint32 {
	return StringVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*StringVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *StringVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   StringVectorTypeID,
	}
	if vec == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Elems",
			SchemaName: "Elems",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (vec *StringVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<string> as nil")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *StringVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<string> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for _, v := range vec.Elems {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *StringVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<string> to nil")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *StringVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<string> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<string>: field Elems: %w", err)
		}

		if headerLen > 0 {
			vec.Elems = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode Vector<string>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *StringVector) GetElems() (value []string) {
	if vec == nil {
		return
	}
	return vec.Elems
}
