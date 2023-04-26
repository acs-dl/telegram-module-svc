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

// InputCheckPasswordEmpty represents TL type `inputCheckPasswordEmpty#9880f658`.
// There is no password
//
// See https://core.telegram.org/constructor/inputCheckPasswordEmpty for reference.
type InputCheckPasswordEmpty struct {
}

// InputCheckPasswordEmptyTypeID is TL type id of InputCheckPasswordEmpty.
const InputCheckPasswordEmptyTypeID = 0x9880f658

// construct implements constructor of InputCheckPasswordSRPClass.
func (i InputCheckPasswordEmpty) construct() InputCheckPasswordSRPClass { return &i }

// Ensuring interfaces in compile-time for InputCheckPasswordEmpty.
var (
	_ bin.Encoder     = &InputCheckPasswordEmpty{}
	_ bin.Decoder     = &InputCheckPasswordEmpty{}
	_ bin.BareEncoder = &InputCheckPasswordEmpty{}
	_ bin.BareDecoder = &InputCheckPasswordEmpty{}

	_ InputCheckPasswordSRPClass = &InputCheckPasswordEmpty{}
)

func (i *InputCheckPasswordEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputCheckPasswordEmpty) String() string {
	if i == nil {
		return "InputCheckPasswordEmpty(nil)"
	}
	type Alias InputCheckPasswordEmpty
	return fmt.Sprintf("InputCheckPasswordEmpty%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputCheckPasswordEmpty) TypeID() uint32 {
	return InputCheckPasswordEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputCheckPasswordEmpty) TypeName() string {
	return "inputCheckPasswordEmpty"
}

// TypeInfo returns info about TL type.
func (i *InputCheckPasswordEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputCheckPasswordEmpty",
		ID:   InputCheckPasswordEmptyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputCheckPasswordEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCheckPasswordEmpty#9880f658 as nil")
	}
	b.PutID(InputCheckPasswordEmptyTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputCheckPasswordEmpty) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCheckPasswordEmpty#9880f658 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputCheckPasswordEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCheckPasswordEmpty#9880f658 to nil")
	}
	if err := b.ConsumeID(InputCheckPasswordEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputCheckPasswordEmpty#9880f658: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputCheckPasswordEmpty) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCheckPasswordEmpty#9880f658 to nil")
	}
	return nil
}

// InputCheckPasswordSRP represents TL type `inputCheckPasswordSRP#d27ff082`.
// Constructor for checking the validity of a 2FA SRP password (see SRP¹)
//
// Links:
//  1. https://core.telegram.org/api/srp
//
// See https://core.telegram.org/constructor/inputCheckPasswordSRP for reference.
type InputCheckPasswordSRP struct {
	// SRP ID¹
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	SRPID int64
	// A parameter (see SRP¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	A []byte
	// M1 parameter (see SRP¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	M1 []byte
}

// InputCheckPasswordSRPTypeID is TL type id of InputCheckPasswordSRP.
const InputCheckPasswordSRPTypeID = 0xd27ff082

// construct implements constructor of InputCheckPasswordSRPClass.
func (i InputCheckPasswordSRP) construct() InputCheckPasswordSRPClass { return &i }

// Ensuring interfaces in compile-time for InputCheckPasswordSRP.
var (
	_ bin.Encoder     = &InputCheckPasswordSRP{}
	_ bin.Decoder     = &InputCheckPasswordSRP{}
	_ bin.BareEncoder = &InputCheckPasswordSRP{}
	_ bin.BareDecoder = &InputCheckPasswordSRP{}

	_ InputCheckPasswordSRPClass = &InputCheckPasswordSRP{}
)

func (i *InputCheckPasswordSRP) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.SRPID == 0) {
		return false
	}
	if !(i.A == nil) {
		return false
	}
	if !(i.M1 == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputCheckPasswordSRP) String() string {
	if i == nil {
		return "InputCheckPasswordSRP(nil)"
	}
	type Alias InputCheckPasswordSRP
	return fmt.Sprintf("InputCheckPasswordSRP%+v", Alias(*i))
}

// FillFrom fills InputCheckPasswordSRP from given interface.
func (i *InputCheckPasswordSRP) FillFrom(from interface {
	GetSRPID() (value int64)
	GetA() (value []byte)
	GetM1() (value []byte)
}) {
	i.SRPID = from.GetSRPID()
	i.A = from.GetA()
	i.M1 = from.GetM1()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputCheckPasswordSRP) TypeID() uint32 {
	return InputCheckPasswordSRPTypeID
}

// TypeName returns name of type in TL schema.
func (*InputCheckPasswordSRP) TypeName() string {
	return "inputCheckPasswordSRP"
}

// TypeInfo returns info about TL type.
func (i *InputCheckPasswordSRP) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputCheckPasswordSRP",
		ID:   InputCheckPasswordSRPTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SRPID",
			SchemaName: "srp_id",
		},
		{
			Name:       "A",
			SchemaName: "A",
		},
		{
			Name:       "M1",
			SchemaName: "M1",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputCheckPasswordSRP) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCheckPasswordSRP#d27ff082 as nil")
	}
	b.PutID(InputCheckPasswordSRPTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputCheckPasswordSRP) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCheckPasswordSRP#d27ff082 as nil")
	}
	b.PutLong(i.SRPID)
	b.PutBytes(i.A)
	b.PutBytes(i.M1)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputCheckPasswordSRP) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCheckPasswordSRP#d27ff082 to nil")
	}
	if err := b.ConsumeID(InputCheckPasswordSRPTypeID); err != nil {
		return fmt.Errorf("unable to decode inputCheckPasswordSRP#d27ff082: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputCheckPasswordSRP) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCheckPasswordSRP#d27ff082 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputCheckPasswordSRP#d27ff082: field srp_id: %w", err)
		}
		i.SRPID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputCheckPasswordSRP#d27ff082: field A: %w", err)
		}
		i.A = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputCheckPasswordSRP#d27ff082: field M1: %w", err)
		}
		i.M1 = value
	}
	return nil
}

// GetSRPID returns value of SRPID field.
func (i *InputCheckPasswordSRP) GetSRPID() (value int64) {
	if i == nil {
		return
	}
	return i.SRPID
}

// GetA returns value of A field.
func (i *InputCheckPasswordSRP) GetA() (value []byte) {
	if i == nil {
		return
	}
	return i.A
}

// GetM1 returns value of M1 field.
func (i *InputCheckPasswordSRP) GetM1() (value []byte) {
	if i == nil {
		return
	}
	return i.M1
}

// InputCheckPasswordSRPClassName is schema name of InputCheckPasswordSRPClass.
const InputCheckPasswordSRPClassName = "InputCheckPasswordSRP"

// InputCheckPasswordSRPClass represents InputCheckPasswordSRP generic type.
//
// See https://core.telegram.org/type/InputCheckPasswordSRP for reference.
//
// Example:
//
//	g, err := tg.DecodeInputCheckPasswordSRP(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.InputCheckPasswordEmpty: // inputCheckPasswordEmpty#9880f658
//	case *tg.InputCheckPasswordSRP: // inputCheckPasswordSRP#d27ff082
//	default: panic(v)
//	}
type InputCheckPasswordSRPClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputCheckPasswordSRPClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsNotEmpty tries to map InputCheckPasswordSRPClass to InputCheckPasswordSRP.
	AsNotEmpty() (*InputCheckPasswordSRP, bool)
}

// AsNotEmpty tries to map InputCheckPasswordEmpty to InputCheckPasswordSRP.
func (i *InputCheckPasswordEmpty) AsNotEmpty() (*InputCheckPasswordSRP, bool) {
	return nil, false
}

// AsNotEmpty tries to map InputCheckPasswordSRP to InputCheckPasswordSRP.
func (i *InputCheckPasswordSRP) AsNotEmpty() (*InputCheckPasswordSRP, bool) {
	return i, true
}

// DecodeInputCheckPasswordSRP implements binary de-serialization for InputCheckPasswordSRPClass.
func DecodeInputCheckPasswordSRP(buf *bin.Buffer) (InputCheckPasswordSRPClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputCheckPasswordEmptyTypeID:
		// Decoding inputCheckPasswordEmpty#9880f658.
		v := InputCheckPasswordEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputCheckPasswordSRPClass: %w", err)
		}
		return &v, nil
	case InputCheckPasswordSRPTypeID:
		// Decoding inputCheckPasswordSRP#d27ff082.
		v := InputCheckPasswordSRP{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputCheckPasswordSRPClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputCheckPasswordSRPClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputCheckPasswordSRP boxes the InputCheckPasswordSRPClass providing a helper.
type InputCheckPasswordSRPBox struct {
	InputCheckPasswordSRP InputCheckPasswordSRPClass
}

// Decode implements bin.Decoder for InputCheckPasswordSRPBox.
func (b *InputCheckPasswordSRPBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputCheckPasswordSRPBox to nil")
	}
	v, err := DecodeInputCheckPasswordSRP(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputCheckPasswordSRP = v
	return nil
}

// Encode implements bin.Encode for InputCheckPasswordSRPBox.
func (b *InputCheckPasswordSRPBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputCheckPasswordSRP == nil {
		return fmt.Errorf("unable to encode InputCheckPasswordSRPClass as nil")
	}
	return b.InputCheckPasswordSRP.Encode(buf)
}
