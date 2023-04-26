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

// AccountResetAuthorizationRequest represents TL type `account.resetAuthorization#df77f3bc`.
// Log out an active authorized session¹ by its hash
//
// Links:
//  1. https://core.telegram.org/api/auth
//
// See https://core.telegram.org/method/account.resetAuthorization for reference.
type AccountResetAuthorizationRequest struct {
	// Session hash
	Hash int64
}

// AccountResetAuthorizationRequestTypeID is TL type id of AccountResetAuthorizationRequest.
const AccountResetAuthorizationRequestTypeID = 0xdf77f3bc

// Ensuring interfaces in compile-time for AccountResetAuthorizationRequest.
var (
	_ bin.Encoder     = &AccountResetAuthorizationRequest{}
	_ bin.Decoder     = &AccountResetAuthorizationRequest{}
	_ bin.BareEncoder = &AccountResetAuthorizationRequest{}
	_ bin.BareDecoder = &AccountResetAuthorizationRequest{}
)

func (r *AccountResetAuthorizationRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountResetAuthorizationRequest) String() string {
	if r == nil {
		return "AccountResetAuthorizationRequest(nil)"
	}
	type Alias AccountResetAuthorizationRequest
	return fmt.Sprintf("AccountResetAuthorizationRequest%+v", Alias(*r))
}

// FillFrom fills AccountResetAuthorizationRequest from given interface.
func (r *AccountResetAuthorizationRequest) FillFrom(from interface {
	GetHash() (value int64)
}) {
	r.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountResetAuthorizationRequest) TypeID() uint32 {
	return AccountResetAuthorizationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountResetAuthorizationRequest) TypeName() string {
	return "account.resetAuthorization"
}

// TypeInfo returns info about TL type.
func (r *AccountResetAuthorizationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.resetAuthorization",
		ID:   AccountResetAuthorizationRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *AccountResetAuthorizationRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetAuthorization#df77f3bc as nil")
	}
	b.PutID(AccountResetAuthorizationRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AccountResetAuthorizationRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetAuthorization#df77f3bc as nil")
	}
	b.PutLong(r.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResetAuthorizationRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetAuthorization#df77f3bc to nil")
	}
	if err := b.ConsumeID(AccountResetAuthorizationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resetAuthorization#df77f3bc: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AccountResetAuthorizationRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetAuthorization#df77f3bc to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.resetAuthorization#df77f3bc: field hash: %w", err)
		}
		r.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (r *AccountResetAuthorizationRequest) GetHash() (value int64) {
	if r == nil {
		return
	}
	return r.Hash
}

// AccountResetAuthorization invokes method account.resetAuthorization#df77f3bc returning error if any.
// Log out an active authorized session¹ by its hash
//
// Links:
//  1. https://core.telegram.org/api/auth
//
// Possible errors:
//
//	406 FRESH_RESET_AUTHORISATION_FORBIDDEN: You can't logout other sessions if less than 24 hours have passed since you logged on the current session.
//	400 HASH_INVALID: The provided hash is invalid.
//
// See https://core.telegram.org/method/account.resetAuthorization for reference.
func (c *Client) AccountResetAuthorization(ctx context.Context, hash int64) (bool, error) {
	var result BoolBox

	request := &AccountResetAuthorizationRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
