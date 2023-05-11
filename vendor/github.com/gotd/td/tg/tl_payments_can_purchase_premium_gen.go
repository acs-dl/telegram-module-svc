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

// PaymentsCanPurchasePremiumRequest represents TL type `payments.canPurchasePremium#9fc19eb6`.
// Checks whether Telegram Premium purchase is possible. Must be called before in-store
// Premium purchase, official apps only.
//
// See https://core.telegram.org/method/payments.canPurchasePremium for reference.
type PaymentsCanPurchasePremiumRequest struct {
	// Payment purpose
	Purpose InputStorePaymentPurposeClass
}

// PaymentsCanPurchasePremiumRequestTypeID is TL type id of PaymentsCanPurchasePremiumRequest.
const PaymentsCanPurchasePremiumRequestTypeID = 0x9fc19eb6

// Ensuring interfaces in compile-time for PaymentsCanPurchasePremiumRequest.
var (
	_ bin.Encoder     = &PaymentsCanPurchasePremiumRequest{}
	_ bin.Decoder     = &PaymentsCanPurchasePremiumRequest{}
	_ bin.BareEncoder = &PaymentsCanPurchasePremiumRequest{}
	_ bin.BareDecoder = &PaymentsCanPurchasePremiumRequest{}
)

func (c *PaymentsCanPurchasePremiumRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Purpose == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *PaymentsCanPurchasePremiumRequest) String() string {
	if c == nil {
		return "PaymentsCanPurchasePremiumRequest(nil)"
	}
	type Alias PaymentsCanPurchasePremiumRequest
	return fmt.Sprintf("PaymentsCanPurchasePremiumRequest%+v", Alias(*c))
}

// FillFrom fills PaymentsCanPurchasePremiumRequest from given interface.
func (c *PaymentsCanPurchasePremiumRequest) FillFrom(from interface {
	GetPurpose() (value InputStorePaymentPurposeClass)
}) {
	c.Purpose = from.GetPurpose()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsCanPurchasePremiumRequest) TypeID() uint32 {
	return PaymentsCanPurchasePremiumRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsCanPurchasePremiumRequest) TypeName() string {
	return "payments.canPurchasePremium"
}

// TypeInfo returns info about TL type.
func (c *PaymentsCanPurchasePremiumRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.canPurchasePremium",
		ID:   PaymentsCanPurchasePremiumRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Purpose",
			SchemaName: "purpose",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *PaymentsCanPurchasePremiumRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode payments.canPurchasePremium#9fc19eb6 as nil")
	}
	b.PutID(PaymentsCanPurchasePremiumRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *PaymentsCanPurchasePremiumRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode payments.canPurchasePremium#9fc19eb6 as nil")
	}
	if c.Purpose == nil {
		return fmt.Errorf("unable to encode payments.canPurchasePremium#9fc19eb6: field purpose is nil")
	}
	if err := c.Purpose.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.canPurchasePremium#9fc19eb6: field purpose: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *PaymentsCanPurchasePremiumRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode payments.canPurchasePremium#9fc19eb6 to nil")
	}
	if err := b.ConsumeID(PaymentsCanPurchasePremiumRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.canPurchasePremium#9fc19eb6: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *PaymentsCanPurchasePremiumRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode payments.canPurchasePremium#9fc19eb6 to nil")
	}
	{
		value, err := DecodeInputStorePaymentPurpose(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.canPurchasePremium#9fc19eb6: field purpose: %w", err)
		}
		c.Purpose = value
	}
	return nil
}

// GetPurpose returns value of Purpose field.
func (c *PaymentsCanPurchasePremiumRequest) GetPurpose() (value InputStorePaymentPurposeClass) {
	if c == nil {
		return
	}
	return c.Purpose
}

// PaymentsCanPurchasePremium invokes method payments.canPurchasePremium#9fc19eb6 returning error if any.
// Checks whether Telegram Premium purchase is possible. Must be called before in-store
// Premium purchase, official apps only.
//
// See https://core.telegram.org/method/payments.canPurchasePremium for reference.
func (c *Client) PaymentsCanPurchasePremium(ctx context.Context, purpose InputStorePaymentPurposeClass) (bool, error) {
	var result BoolBox

	request := &PaymentsCanPurchasePremiumRequest{
		Purpose: purpose,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}