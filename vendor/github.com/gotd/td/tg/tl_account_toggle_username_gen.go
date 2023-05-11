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

// AccountToggleUsernameRequest represents TL type `account.toggleUsername#58d6b376`.
// Associate or dissociate a purchased fragment.com¹ username to the currently logged-in
// user.
//
// Links:
//  1. https://fragment.com
//
// See https://core.telegram.org/method/account.toggleUsername for reference.
type AccountToggleUsernameRequest struct {
	// Username
	Username string
	// Whether to associate or dissociate it
	Active bool
}

// AccountToggleUsernameRequestTypeID is TL type id of AccountToggleUsernameRequest.
const AccountToggleUsernameRequestTypeID = 0x58d6b376

// Ensuring interfaces in compile-time for AccountToggleUsernameRequest.
var (
	_ bin.Encoder     = &AccountToggleUsernameRequest{}
	_ bin.Decoder     = &AccountToggleUsernameRequest{}
	_ bin.BareEncoder = &AccountToggleUsernameRequest{}
	_ bin.BareDecoder = &AccountToggleUsernameRequest{}
)

func (t *AccountToggleUsernameRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Username == "") {
		return false
	}
	if !(t.Active == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *AccountToggleUsernameRequest) String() string {
	if t == nil {
		return "AccountToggleUsernameRequest(nil)"
	}
	type Alias AccountToggleUsernameRequest
	return fmt.Sprintf("AccountToggleUsernameRequest%+v", Alias(*t))
}

// FillFrom fills AccountToggleUsernameRequest from given interface.
func (t *AccountToggleUsernameRequest) FillFrom(from interface {
	GetUsername() (value string)
	GetActive() (value bool)
}) {
	t.Username = from.GetUsername()
	t.Active = from.GetActive()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountToggleUsernameRequest) TypeID() uint32 {
	return AccountToggleUsernameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountToggleUsernameRequest) TypeName() string {
	return "account.toggleUsername"
}

// TypeInfo returns info about TL type.
func (t *AccountToggleUsernameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.toggleUsername",
		ID:   AccountToggleUsernameRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Username",
			SchemaName: "username",
		},
		{
			Name:       "Active",
			SchemaName: "active",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *AccountToggleUsernameRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.toggleUsername#58d6b376 as nil")
	}
	b.PutID(AccountToggleUsernameRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *AccountToggleUsernameRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.toggleUsername#58d6b376 as nil")
	}
	b.PutString(t.Username)
	b.PutBool(t.Active)
	return nil
}

// Decode implements bin.Decoder.
func (t *AccountToggleUsernameRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.toggleUsername#58d6b376 to nil")
	}
	if err := b.ConsumeID(AccountToggleUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.toggleUsername#58d6b376: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *AccountToggleUsernameRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.toggleUsername#58d6b376 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.toggleUsername#58d6b376: field username: %w", err)
		}
		t.Username = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.toggleUsername#58d6b376: field active: %w", err)
		}
		t.Active = value
	}
	return nil
}

// GetUsername returns value of Username field.
func (t *AccountToggleUsernameRequest) GetUsername() (value string) {
	if t == nil {
		return
	}
	return t.Username
}

// GetActive returns value of Active field.
func (t *AccountToggleUsernameRequest) GetActive() (value bool) {
	if t == nil {
		return
	}
	return t.Active
}

// AccountToggleUsername invokes method account.toggleUsername#58d6b376 returning error if any.
// Associate or dissociate a purchased fragment.com¹ username to the currently logged-in
// user.
//
// Links:
//  1. https://fragment.com
//
// Possible errors:
//
//	400 USERNAME_INVALID: The provided username is not valid.
//
// See https://core.telegram.org/method/account.toggleUsername for reference.
func (c *Client) AccountToggleUsername(ctx context.Context, request *AccountToggleUsernameRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}