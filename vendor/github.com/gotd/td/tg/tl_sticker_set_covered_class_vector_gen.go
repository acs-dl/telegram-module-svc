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

// StickerSetCoveredClassVector is a box for Vector<StickerSetCovered>
type StickerSetCoveredClassVector struct {
	// Elements of Vector<StickerSetCovered>
	Elems []StickerSetCoveredClass
}

// StickerSetCoveredClassVectorTypeID is TL type id of StickerSetCoveredClassVector.
const StickerSetCoveredClassVectorTypeID = bin.TypeVector

// Ensuring interfaces in compile-time for StickerSetCoveredClassVector.
var (
	_ bin.Encoder     = &StickerSetCoveredClassVector{}
	_ bin.Decoder     = &StickerSetCoveredClassVector{}
	_ bin.BareEncoder = &StickerSetCoveredClassVector{}
	_ bin.BareDecoder = &StickerSetCoveredClassVector{}
)

func (vec *StickerSetCoveredClassVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *StickerSetCoveredClassVector) String() string {
	if vec == nil {
		return "StickerSetCoveredClassVector(nil)"
	}
	type Alias StickerSetCoveredClassVector
	return fmt.Sprintf("StickerSetCoveredClassVector%+v", Alias(*vec))
}

// FillFrom fills StickerSetCoveredClassVector from given interface.
func (vec *StickerSetCoveredClassVector) FillFrom(from interface {
	GetElems() (value []StickerSetCoveredClass)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickerSetCoveredClassVector) TypeID() uint32 {
	return StickerSetCoveredClassVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*StickerSetCoveredClassVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *StickerSetCoveredClassVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   StickerSetCoveredClassVectorTypeID,
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
func (vec *StickerSetCoveredClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<StickerSetCovered> as nil")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *StickerSetCoveredClassVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<StickerSetCovered> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode Vector<StickerSetCovered>: field Elems element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<StickerSetCovered>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *StickerSetCoveredClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<StickerSetCovered> to nil")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *StickerSetCoveredClassVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<StickerSetCovered> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<StickerSetCovered>: field Elems: %w", err)
		}

		if headerLen > 0 {
			vec.Elems = make([]StickerSetCoveredClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeStickerSetCovered(b)
			if err != nil {
				return fmt.Errorf("unable to decode Vector<StickerSetCovered>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *StickerSetCoveredClassVector) GetElems() (value []StickerSetCoveredClass) {
	if vec == nil {
		return
	}
	return vec.Elems
}

// MapElems returns field Elems wrapped in StickerSetCoveredClassArray helper.
func (vec *StickerSetCoveredClassVector) MapElems() (value StickerSetCoveredClassArray) {
	return StickerSetCoveredClassArray(vec.Elems)
}