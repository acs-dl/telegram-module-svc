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

// AccountGetDefaultProfilePhotoEmojisRequest represents TL type `account.getDefaultProfilePhotoEmojis#e2750328`.
//
// See https://core.telegram.org/method/account.getDefaultProfilePhotoEmojis for reference.
type AccountGetDefaultProfilePhotoEmojisRequest struct {
	// Hash field of AccountGetDefaultProfilePhotoEmojisRequest.
	Hash int64
}

// AccountGetDefaultProfilePhotoEmojisRequestTypeID is TL type id of AccountGetDefaultProfilePhotoEmojisRequest.
const AccountGetDefaultProfilePhotoEmojisRequestTypeID = 0xe2750328

// Ensuring interfaces in compile-time for AccountGetDefaultProfilePhotoEmojisRequest.
var (
	_ bin.Encoder     = &AccountGetDefaultProfilePhotoEmojisRequest{}
	_ bin.Decoder     = &AccountGetDefaultProfilePhotoEmojisRequest{}
	_ bin.BareEncoder = &AccountGetDefaultProfilePhotoEmojisRequest{}
	_ bin.BareDecoder = &AccountGetDefaultProfilePhotoEmojisRequest{}
)

func (g *AccountGetDefaultProfilePhotoEmojisRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetDefaultProfilePhotoEmojisRequest) String() string {
	if g == nil {
		return "AccountGetDefaultProfilePhotoEmojisRequest(nil)"
	}
	type Alias AccountGetDefaultProfilePhotoEmojisRequest
	return fmt.Sprintf("AccountGetDefaultProfilePhotoEmojisRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetDefaultProfilePhotoEmojisRequest from given interface.
func (g *AccountGetDefaultProfilePhotoEmojisRequest) FillFrom(from interface {
	GetHash() (value int64)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetDefaultProfilePhotoEmojisRequest) TypeID() uint32 {
	return AccountGetDefaultProfilePhotoEmojisRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetDefaultProfilePhotoEmojisRequest) TypeName() string {
	return "account.getDefaultProfilePhotoEmojis"
}

// TypeInfo returns info about TL type.
func (g *AccountGetDefaultProfilePhotoEmojisRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getDefaultProfilePhotoEmojis",
		ID:   AccountGetDefaultProfilePhotoEmojisRequestTypeID,
	}
	if g == nil {
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
func (g *AccountGetDefaultProfilePhotoEmojisRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getDefaultProfilePhotoEmojis#e2750328 as nil")
	}
	b.PutID(AccountGetDefaultProfilePhotoEmojisRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetDefaultProfilePhotoEmojisRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getDefaultProfilePhotoEmojis#e2750328 as nil")
	}
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetDefaultProfilePhotoEmojisRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getDefaultProfilePhotoEmojis#e2750328 to nil")
	}
	if err := b.ConsumeID(AccountGetDefaultProfilePhotoEmojisRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getDefaultProfilePhotoEmojis#e2750328: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetDefaultProfilePhotoEmojisRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getDefaultProfilePhotoEmojis#e2750328 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.getDefaultProfilePhotoEmojis#e2750328: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *AccountGetDefaultProfilePhotoEmojisRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// AccountGetDefaultProfilePhotoEmojis invokes method account.getDefaultProfilePhotoEmojis#e2750328 returning error if any.
//
// See https://core.telegram.org/method/account.getDefaultProfilePhotoEmojis for reference.
func (c *Client) AccountGetDefaultProfilePhotoEmojis(ctx context.Context, hash int64) (EmojiListClass, error) {
	var result EmojiListBox

	request := &AccountGetDefaultProfilePhotoEmojisRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.EmojiList, nil
}
