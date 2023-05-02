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

// AccountUpdateEmojiStatusRequest represents TL type `account.updateEmojiStatus#fbd3de6b`.
// Set an emoji status¹
//
// Links:
//  1. https://core.telegram.org/api/emoji-status
//
// See https://core.telegram.org/method/account.updateEmojiStatus for reference.
type AccountUpdateEmojiStatusRequest struct {
	// Emoji status¹ to set
	//
	// Links:
	//  1) https://core.telegram.org/api/emoji-status
	EmojiStatus EmojiStatusClass
}

// AccountUpdateEmojiStatusRequestTypeID is TL type id of AccountUpdateEmojiStatusRequest.
const AccountUpdateEmojiStatusRequestTypeID = 0xfbd3de6b

// Ensuring interfaces in compile-time for AccountUpdateEmojiStatusRequest.
var (
	_ bin.Encoder     = &AccountUpdateEmojiStatusRequest{}
	_ bin.Decoder     = &AccountUpdateEmojiStatusRequest{}
	_ bin.BareEncoder = &AccountUpdateEmojiStatusRequest{}
	_ bin.BareDecoder = &AccountUpdateEmojiStatusRequest{}
)

func (u *AccountUpdateEmojiStatusRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.EmojiStatus == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateEmojiStatusRequest) String() string {
	if u == nil {
		return "AccountUpdateEmojiStatusRequest(nil)"
	}
	type Alias AccountUpdateEmojiStatusRequest
	return fmt.Sprintf("AccountUpdateEmojiStatusRequest%+v", Alias(*u))
}

// FillFrom fills AccountUpdateEmojiStatusRequest from given interface.
func (u *AccountUpdateEmojiStatusRequest) FillFrom(from interface {
	GetEmojiStatus() (value EmojiStatusClass)
}) {
	u.EmojiStatus = from.GetEmojiStatus()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdateEmojiStatusRequest) TypeID() uint32 {
	return AccountUpdateEmojiStatusRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdateEmojiStatusRequest) TypeName() string {
	return "account.updateEmojiStatus"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdateEmojiStatusRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updateEmojiStatus",
		ID:   AccountUpdateEmojiStatusRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "EmojiStatus",
			SchemaName: "emoji_status",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUpdateEmojiStatusRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateEmojiStatus#fbd3de6b as nil")
	}
	b.PutID(AccountUpdateEmojiStatusRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdateEmojiStatusRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateEmojiStatus#fbd3de6b as nil")
	}
	if u.EmojiStatus == nil {
		return fmt.Errorf("unable to encode account.updateEmojiStatus#fbd3de6b: field emoji_status is nil")
	}
	if err := u.EmojiStatus.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateEmojiStatus#fbd3de6b: field emoji_status: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateEmojiStatusRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateEmojiStatus#fbd3de6b to nil")
	}
	if err := b.ConsumeID(AccountUpdateEmojiStatusRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateEmojiStatus#fbd3de6b: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdateEmojiStatusRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateEmojiStatus#fbd3de6b to nil")
	}
	{
		value, err := DecodeEmojiStatus(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.updateEmojiStatus#fbd3de6b: field emoji_status: %w", err)
		}
		u.EmojiStatus = value
	}
	return nil
}

// GetEmojiStatus returns value of EmojiStatus field.
func (u *AccountUpdateEmojiStatusRequest) GetEmojiStatus() (value EmojiStatusClass) {
	if u == nil {
		return
	}
	return u.EmojiStatus
}

// GetEmojiStatusAsNotEmpty returns mapped value of EmojiStatus field.
func (u *AccountUpdateEmojiStatusRequest) GetEmojiStatusAsNotEmpty() (NotEmptyEmojiStatus, bool) {
	return u.EmojiStatus.AsNotEmpty()
}

// AccountUpdateEmojiStatus invokes method account.updateEmojiStatus#fbd3de6b returning error if any.
// Set an emoji status¹
//
// Links:
//  1. https://core.telegram.org/api/emoji-status
//
// See https://core.telegram.org/method/account.updateEmojiStatus for reference.
func (c *Client) AccountUpdateEmojiStatus(ctx context.Context, emojistatus EmojiStatusClass) (bool, error) {
	var result BoolBox

	request := &AccountUpdateEmojiStatusRequest{
		EmojiStatus: emojistatus,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
