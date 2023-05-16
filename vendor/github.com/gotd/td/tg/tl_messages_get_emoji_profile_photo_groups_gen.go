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

// MessagesGetEmojiProfilePhotoGroupsRequest represents TL type `messages.getEmojiProfilePhotoGroups#21a548f3`.
// Represents a list of emoji categories¹, to be used when selecting custom emojis to
// set as profile picture².
//
// Links:
//  1. https://core.telegram.org/api/custom-emoji#emoji-categories
//  2. https://core.telegram.org/api/files#sticker-profile-pictures
//
// See https://core.telegram.org/method/messages.getEmojiProfilePhotoGroups for reference.
type MessagesGetEmojiProfilePhotoGroupsRequest struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// MessagesGetEmojiProfilePhotoGroupsRequestTypeID is TL type id of MessagesGetEmojiProfilePhotoGroupsRequest.
const MessagesGetEmojiProfilePhotoGroupsRequestTypeID = 0x21a548f3

// Ensuring interfaces in compile-time for MessagesGetEmojiProfilePhotoGroupsRequest.
var (
	_ bin.Encoder     = &MessagesGetEmojiProfilePhotoGroupsRequest{}
	_ bin.Decoder     = &MessagesGetEmojiProfilePhotoGroupsRequest{}
	_ bin.BareEncoder = &MessagesGetEmojiProfilePhotoGroupsRequest{}
	_ bin.BareDecoder = &MessagesGetEmojiProfilePhotoGroupsRequest{}
)

func (g *MessagesGetEmojiProfilePhotoGroupsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetEmojiProfilePhotoGroupsRequest) String() string {
	if g == nil {
		return "MessagesGetEmojiProfilePhotoGroupsRequest(nil)"
	}
	type Alias MessagesGetEmojiProfilePhotoGroupsRequest
	return fmt.Sprintf("MessagesGetEmojiProfilePhotoGroupsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetEmojiProfilePhotoGroupsRequest from given interface.
func (g *MessagesGetEmojiProfilePhotoGroupsRequest) FillFrom(from interface {
	GetHash() (value int)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetEmojiProfilePhotoGroupsRequest) TypeID() uint32 {
	return MessagesGetEmojiProfilePhotoGroupsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetEmojiProfilePhotoGroupsRequest) TypeName() string {
	return "messages.getEmojiProfilePhotoGroups"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetEmojiProfilePhotoGroupsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getEmojiProfilePhotoGroups",
		ID:   MessagesGetEmojiProfilePhotoGroupsRequestTypeID,
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
func (g *MessagesGetEmojiProfilePhotoGroupsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiProfilePhotoGroups#21a548f3 as nil")
	}
	b.PutID(MessagesGetEmojiProfilePhotoGroupsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetEmojiProfilePhotoGroupsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiProfilePhotoGroups#21a548f3 as nil")
	}
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetEmojiProfilePhotoGroupsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiProfilePhotoGroups#21a548f3 to nil")
	}
	if err := b.ConsumeID(MessagesGetEmojiProfilePhotoGroupsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getEmojiProfilePhotoGroups#21a548f3: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetEmojiProfilePhotoGroupsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiProfilePhotoGroups#21a548f3 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getEmojiProfilePhotoGroups#21a548f3: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *MessagesGetEmojiProfilePhotoGroupsRequest) GetHash() (value int) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetEmojiProfilePhotoGroups invokes method messages.getEmojiProfilePhotoGroups#21a548f3 returning error if any.
// Represents a list of emoji categories¹, to be used when selecting custom emojis to
// set as profile picture².
//
// Links:
//  1. https://core.telegram.org/api/custom-emoji#emoji-categories
//  2. https://core.telegram.org/api/files#sticker-profile-pictures
//
// See https://core.telegram.org/method/messages.getEmojiProfilePhotoGroups for reference.
// Can be used by bots.
func (c *Client) MessagesGetEmojiProfilePhotoGroups(ctx context.Context, hash int) (MessagesEmojiGroupsClass, error) {
	var result MessagesEmojiGroupsBox

	request := &MessagesGetEmojiProfilePhotoGroupsRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.EmojiGroups, nil
}
