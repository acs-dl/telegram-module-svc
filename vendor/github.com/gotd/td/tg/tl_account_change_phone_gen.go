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

// AccountChangePhoneRequest represents TL type `account.changePhone#70c32edb`.
// Change the phone number of the current account
//
// See https://core.telegram.org/method/account.changePhone for reference.
type AccountChangePhoneRequest struct {
	// New phone number
	PhoneNumber string
	// Phone code hash received when calling account.sendChangePhoneCode¹
	//
	// Links:
	//  1) https://core.telegram.org/method/account.sendChangePhoneCode
	PhoneCodeHash string
	// Phone code received when calling account.sendChangePhoneCode¹
	//
	// Links:
	//  1) https://core.telegram.org/method/account.sendChangePhoneCode
	PhoneCode string
}

// AccountChangePhoneRequestTypeID is TL type id of AccountChangePhoneRequest.
const AccountChangePhoneRequestTypeID = 0x70c32edb

// Ensuring interfaces in compile-time for AccountChangePhoneRequest.
var (
	_ bin.Encoder     = &AccountChangePhoneRequest{}
	_ bin.Decoder     = &AccountChangePhoneRequest{}
	_ bin.BareEncoder = &AccountChangePhoneRequest{}
	_ bin.BareDecoder = &AccountChangePhoneRequest{}
)

func (c *AccountChangePhoneRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.PhoneNumber == "") {
		return false
	}
	if !(c.PhoneCodeHash == "") {
		return false
	}
	if !(c.PhoneCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AccountChangePhoneRequest) String() string {
	if c == nil {
		return "AccountChangePhoneRequest(nil)"
	}
	type Alias AccountChangePhoneRequest
	return fmt.Sprintf("AccountChangePhoneRequest%+v", Alias(*c))
}

// FillFrom fills AccountChangePhoneRequest from given interface.
func (c *AccountChangePhoneRequest) FillFrom(from interface {
	GetPhoneNumber() (value string)
	GetPhoneCodeHash() (value string)
	GetPhoneCode() (value string)
}) {
	c.PhoneNumber = from.GetPhoneNumber()
	c.PhoneCodeHash = from.GetPhoneCodeHash()
	c.PhoneCode = from.GetPhoneCode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountChangePhoneRequest) TypeID() uint32 {
	return AccountChangePhoneRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountChangePhoneRequest) TypeName() string {
	return "account.changePhone"
}

// TypeInfo returns info about TL type.
func (c *AccountChangePhoneRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.changePhone",
		ID:   AccountChangePhoneRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
		{
			Name:       "PhoneCodeHash",
			SchemaName: "phone_code_hash",
		},
		{
			Name:       "PhoneCode",
			SchemaName: "phone_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *AccountChangePhoneRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.changePhone#70c32edb as nil")
	}
	b.PutID(AccountChangePhoneRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *AccountChangePhoneRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.changePhone#70c32edb as nil")
	}
	b.PutString(c.PhoneNumber)
	b.PutString(c.PhoneCodeHash)
	b.PutString(c.PhoneCode)
	return nil
}

// Decode implements bin.Decoder.
func (c *AccountChangePhoneRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.changePhone#70c32edb to nil")
	}
	if err := b.ConsumeID(AccountChangePhoneRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.changePhone#70c32edb: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *AccountChangePhoneRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.changePhone#70c32edb to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.changePhone#70c32edb: field phone_number: %w", err)
		}
		c.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.changePhone#70c32edb: field phone_code_hash: %w", err)
		}
		c.PhoneCodeHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.changePhone#70c32edb: field phone_code: %w", err)
		}
		c.PhoneCode = value
	}
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (c *AccountChangePhoneRequest) GetPhoneNumber() (value string) {
	if c == nil {
		return
	}
	return c.PhoneNumber
}

// GetPhoneCodeHash returns value of PhoneCodeHash field.
func (c *AccountChangePhoneRequest) GetPhoneCodeHash() (value string) {
	if c == nil {
		return
	}
	return c.PhoneCodeHash
}

// GetPhoneCode returns value of PhoneCode field.
func (c *AccountChangePhoneRequest) GetPhoneCode() (value string) {
	if c == nil {
		return
	}
	return c.PhoneCode
}

// AccountChangePhone invokes method account.changePhone#70c32edb returning error if any.
// Change the phone number of the current account
//
// Possible errors:
//
//	400 PHONE_CODE_EMPTY: phone_code is missing.
//	400 PHONE_CODE_EXPIRED: The phone code you provided has expired.
//	406 PHONE_NUMBER_INVALID: The phone number is invalid.
//	400 PHONE_NUMBER_OCCUPIED: The phone number is already in use.
//
// See https://core.telegram.org/method/account.changePhone for reference.
func (c *Client) AccountChangePhone(ctx context.Context, request *AccountChangePhoneRequest) (UserClass, error) {
	var result UserBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.User, nil
}
