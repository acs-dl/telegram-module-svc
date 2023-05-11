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

// BotCommandVector is a box for Vector<BotCommand>
type BotCommandVector struct {
	// Elements of Vector<BotCommand>
	Elems []BotCommand
}

// BotCommandVectorTypeID is TL type id of BotCommandVector.
const BotCommandVectorTypeID = bin.TypeVector

// Ensuring interfaces in compile-time for BotCommandVector.
var (
	_ bin.Encoder     = &BotCommandVector{}
	_ bin.Decoder     = &BotCommandVector{}
	_ bin.BareEncoder = &BotCommandVector{}
	_ bin.BareDecoder = &BotCommandVector{}
)

func (vec *BotCommandVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *BotCommandVector) String() string {
	if vec == nil {
		return "BotCommandVector(nil)"
	}
	type Alias BotCommandVector
	return fmt.Sprintf("BotCommandVector%+v", Alias(*vec))
}

// FillFrom fills BotCommandVector from given interface.
func (vec *BotCommandVector) FillFrom(from interface {
	GetElems() (value []BotCommand)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandVector) TypeID() uint32 {
	return BotCommandVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *BotCommandVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   BotCommandVectorTypeID,
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
func (vec *BotCommandVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<BotCommand> as nil")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *BotCommandVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<BotCommand> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<BotCommand>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *BotCommandVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<BotCommand> to nil")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *BotCommandVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<BotCommand> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<BotCommand>: field Elems: %w", err)
		}

		if headerLen > 0 {
			vec.Elems = make([]BotCommand, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BotCommand
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<BotCommand>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *BotCommandVector) GetElems() (value []BotCommand) {
	if vec == nil {
		return
	}
	return vec.Elems
}
