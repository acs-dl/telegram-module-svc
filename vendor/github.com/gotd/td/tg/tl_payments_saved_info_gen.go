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

// PaymentsSavedInfo represents TL type `payments.savedInfo#fb8fe43c`.
// Saved server-side order information
//
// See https://core.telegram.org/constructor/payments.savedInfo for reference.
type PaymentsSavedInfo struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the user has some saved payment credentials
	HasSavedCredentials bool
	// Saved server-side order information
	//
	// Use SetSavedInfo and GetSavedInfo helpers.
	SavedInfo PaymentRequestedInfo
}

// PaymentsSavedInfoTypeID is TL type id of PaymentsSavedInfo.
const PaymentsSavedInfoTypeID = 0xfb8fe43c

// Ensuring interfaces in compile-time for PaymentsSavedInfo.
var (
	_ bin.Encoder     = &PaymentsSavedInfo{}
	_ bin.Decoder     = &PaymentsSavedInfo{}
	_ bin.BareEncoder = &PaymentsSavedInfo{}
	_ bin.BareDecoder = &PaymentsSavedInfo{}
)

func (s *PaymentsSavedInfo) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.HasSavedCredentials == false) {
		return false
	}
	if !(s.SavedInfo.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PaymentsSavedInfo) String() string {
	if s == nil {
		return "PaymentsSavedInfo(nil)"
	}
	type Alias PaymentsSavedInfo
	return fmt.Sprintf("PaymentsSavedInfo%+v", Alias(*s))
}

// FillFrom fills PaymentsSavedInfo from given interface.
func (s *PaymentsSavedInfo) FillFrom(from interface {
	GetHasSavedCredentials() (value bool)
	GetSavedInfo() (value PaymentRequestedInfo, ok bool)
}) {
	s.HasSavedCredentials = from.GetHasSavedCredentials()
	if val, ok := from.GetSavedInfo(); ok {
		s.SavedInfo = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsSavedInfo) TypeID() uint32 {
	return PaymentsSavedInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsSavedInfo) TypeName() string {
	return "payments.savedInfo"
}

// TypeInfo returns info about TL type.
func (s *PaymentsSavedInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.savedInfo",
		ID:   PaymentsSavedInfoTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "HasSavedCredentials",
			SchemaName: "has_saved_credentials",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "SavedInfo",
			SchemaName: "saved_info",
			Null:       !s.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *PaymentsSavedInfo) SetFlags() {
	if !(s.HasSavedCredentials == false) {
		s.Flags.Set(1)
	}
	if !(s.SavedInfo.Zero()) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *PaymentsSavedInfo) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.savedInfo#fb8fe43c as nil")
	}
	b.PutID(PaymentsSavedInfoTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *PaymentsSavedInfo) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.savedInfo#fb8fe43c as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.savedInfo#fb8fe43c: field flags: %w", err)
	}
	if s.Flags.Has(0) {
		if err := s.SavedInfo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.savedInfo#fb8fe43c: field saved_info: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *PaymentsSavedInfo) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.savedInfo#fb8fe43c to nil")
	}
	if err := b.ConsumeID(PaymentsSavedInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.savedInfo#fb8fe43c: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *PaymentsSavedInfo) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.savedInfo#fb8fe43c to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.savedInfo#fb8fe43c: field flags: %w", err)
		}
	}
	s.HasSavedCredentials = s.Flags.Has(1)
	if s.Flags.Has(0) {
		if err := s.SavedInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.savedInfo#fb8fe43c: field saved_info: %w", err)
		}
	}
	return nil
}

// SetHasSavedCredentials sets value of HasSavedCredentials conditional field.
func (s *PaymentsSavedInfo) SetHasSavedCredentials(value bool) {
	if value {
		s.Flags.Set(1)
		s.HasSavedCredentials = true
	} else {
		s.Flags.Unset(1)
		s.HasSavedCredentials = false
	}
}

// GetHasSavedCredentials returns value of HasSavedCredentials conditional field.
func (s *PaymentsSavedInfo) GetHasSavedCredentials() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(1)
}

// SetSavedInfo sets value of SavedInfo conditional field.
func (s *PaymentsSavedInfo) SetSavedInfo(value PaymentRequestedInfo) {
	s.Flags.Set(0)
	s.SavedInfo = value
}

// GetSavedInfo returns value of SavedInfo conditional field and
// boolean which is true if field was set.
func (s *PaymentsSavedInfo) GetSavedInfo() (value PaymentRequestedInfo, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.SavedInfo, true
}