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

// PhonePhoneCall represents TL type `phone.phoneCall#ec82e140`.
// A VoIP phone call
//
// See https://core.telegram.org/constructor/phone.phoneCall for reference.
type PhonePhoneCall struct {
	// The VoIP phone call
	PhoneCall PhoneCallClass
	// VoIP phone call participants
	Users []UserClass
}

// PhonePhoneCallTypeID is TL type id of PhonePhoneCall.
const PhonePhoneCallTypeID = 0xec82e140

// Ensuring interfaces in compile-time for PhonePhoneCall.
var (
	_ bin.Encoder     = &PhonePhoneCall{}
	_ bin.Decoder     = &PhonePhoneCall{}
	_ bin.BareEncoder = &PhonePhoneCall{}
	_ bin.BareDecoder = &PhonePhoneCall{}
)

func (p *PhonePhoneCall) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.PhoneCall == nil) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhonePhoneCall) String() string {
	if p == nil {
		return "PhonePhoneCall(nil)"
	}
	type Alias PhonePhoneCall
	return fmt.Sprintf("PhonePhoneCall%+v", Alias(*p))
}

// FillFrom fills PhonePhoneCall from given interface.
func (p *PhonePhoneCall) FillFrom(from interface {
	GetPhoneCall() (value PhoneCallClass)
	GetUsers() (value []UserClass)
}) {
	p.PhoneCall = from.GetPhoneCall()
	p.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhonePhoneCall) TypeID() uint32 {
	return PhonePhoneCallTypeID
}

// TypeName returns name of type in TL schema.
func (*PhonePhoneCall) TypeName() string {
	return "phone.phoneCall"
}

// TypeInfo returns info about TL type.
func (p *PhonePhoneCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.phoneCall",
		ID:   PhonePhoneCallTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneCall",
			SchemaName: "phone_call",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhonePhoneCall) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phone.phoneCall#ec82e140 as nil")
	}
	b.PutID(PhonePhoneCallTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhonePhoneCall) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phone.phoneCall#ec82e140 as nil")
	}
	if p.PhoneCall == nil {
		return fmt.Errorf("unable to encode phone.phoneCall#ec82e140: field phone_call is nil")
	}
	if err := p.PhoneCall.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.phoneCall#ec82e140: field phone_call: %w", err)
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode phone.phoneCall#ec82e140: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phone.phoneCall#ec82e140: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PhonePhoneCall) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phone.phoneCall#ec82e140 to nil")
	}
	if err := b.ConsumeID(PhonePhoneCallTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.phoneCall#ec82e140: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhonePhoneCall) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phone.phoneCall#ec82e140 to nil")
	}
	{
		value, err := DecodePhoneCall(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.phoneCall#ec82e140: field phone_call: %w", err)
		}
		p.PhoneCall = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.phoneCall#ec82e140: field users: %w", err)
		}

		if headerLen > 0 {
			p.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode phone.phoneCall#ec82e140: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// GetPhoneCall returns value of PhoneCall field.
func (p *PhonePhoneCall) GetPhoneCall() (value PhoneCallClass) {
	if p == nil {
		return
	}
	return p.PhoneCall
}

// GetUsers returns value of Users field.
func (p *PhonePhoneCall) GetUsers() (value []UserClass) {
	if p == nil {
		return
	}
	return p.Users
}

// GetPhoneCallAsNotEmpty returns mapped value of PhoneCall field.
func (p *PhonePhoneCall) GetPhoneCallAsNotEmpty() (NotEmptyPhoneCall, bool) {
	return p.PhoneCall.AsNotEmpty()
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (p *PhonePhoneCall) MapUsers() (value UserClassArray) {
	return UserClassArray(p.Users)
}
