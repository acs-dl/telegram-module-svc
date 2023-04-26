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

// ChatlistsJoinChatlistInviteRequest represents TL type `chatlists.joinChatlistInvite#a6b1e39a`.
//
// See https://core.telegram.org/method/chatlists.joinChatlistInvite for reference.
type ChatlistsJoinChatlistInviteRequest struct {
	// Slug field of ChatlistsJoinChatlistInviteRequest.
	Slug string
	// Peers field of ChatlistsJoinChatlistInviteRequest.
	Peers []InputPeerClass
}

// ChatlistsJoinChatlistInviteRequestTypeID is TL type id of ChatlistsJoinChatlistInviteRequest.
const ChatlistsJoinChatlistInviteRequestTypeID = 0xa6b1e39a

// Ensuring interfaces in compile-time for ChatlistsJoinChatlistInviteRequest.
var (
	_ bin.Encoder     = &ChatlistsJoinChatlistInviteRequest{}
	_ bin.Decoder     = &ChatlistsJoinChatlistInviteRequest{}
	_ bin.BareEncoder = &ChatlistsJoinChatlistInviteRequest{}
	_ bin.BareDecoder = &ChatlistsJoinChatlistInviteRequest{}
)

func (j *ChatlistsJoinChatlistInviteRequest) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.Slug == "") {
		return false
	}
	if !(j.Peers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *ChatlistsJoinChatlistInviteRequest) String() string {
	if j == nil {
		return "ChatlistsJoinChatlistInviteRequest(nil)"
	}
	type Alias ChatlistsJoinChatlistInviteRequest
	return fmt.Sprintf("ChatlistsJoinChatlistInviteRequest%+v", Alias(*j))
}

// FillFrom fills ChatlistsJoinChatlistInviteRequest from given interface.
func (j *ChatlistsJoinChatlistInviteRequest) FillFrom(from interface {
	GetSlug() (value string)
	GetPeers() (value []InputPeerClass)
}) {
	j.Slug = from.GetSlug()
	j.Peers = from.GetPeers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatlistsJoinChatlistInviteRequest) TypeID() uint32 {
	return ChatlistsJoinChatlistInviteRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatlistsJoinChatlistInviteRequest) TypeName() string {
	return "chatlists.joinChatlistInvite"
}

// TypeInfo returns info about TL type.
func (j *ChatlistsJoinChatlistInviteRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatlists.joinChatlistInvite",
		ID:   ChatlistsJoinChatlistInviteRequestTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Slug",
			SchemaName: "slug",
		},
		{
			Name:       "Peers",
			SchemaName: "peers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (j *ChatlistsJoinChatlistInviteRequest) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode chatlists.joinChatlistInvite#a6b1e39a as nil")
	}
	b.PutID(ChatlistsJoinChatlistInviteRequestTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *ChatlistsJoinChatlistInviteRequest) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode chatlists.joinChatlistInvite#a6b1e39a as nil")
	}
	b.PutString(j.Slug)
	b.PutVectorHeader(len(j.Peers))
	for idx, v := range j.Peers {
		if v == nil {
			return fmt.Errorf("unable to encode chatlists.joinChatlistInvite#a6b1e39a: field peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatlists.joinChatlistInvite#a6b1e39a: field peers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (j *ChatlistsJoinChatlistInviteRequest) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode chatlists.joinChatlistInvite#a6b1e39a to nil")
	}
	if err := b.ConsumeID(ChatlistsJoinChatlistInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode chatlists.joinChatlistInvite#a6b1e39a: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *ChatlistsJoinChatlistInviteRequest) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode chatlists.joinChatlistInvite#a6b1e39a to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.joinChatlistInvite#a6b1e39a: field slug: %w", err)
		}
		j.Slug = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.joinChatlistInvite#a6b1e39a: field peers: %w", err)
		}

		if headerLen > 0 {
			j.Peers = make([]InputPeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatlists.joinChatlistInvite#a6b1e39a: field peers: %w", err)
			}
			j.Peers = append(j.Peers, value)
		}
	}
	return nil
}

// GetSlug returns value of Slug field.
func (j *ChatlistsJoinChatlistInviteRequest) GetSlug() (value string) {
	if j == nil {
		return
	}
	return j.Slug
}

// GetPeers returns value of Peers field.
func (j *ChatlistsJoinChatlistInviteRequest) GetPeers() (value []InputPeerClass) {
	if j == nil {
		return
	}
	return j.Peers
}

// MapPeers returns field Peers wrapped in InputPeerClassArray helper.
func (j *ChatlistsJoinChatlistInviteRequest) MapPeers() (value InputPeerClassArray) {
	return InputPeerClassArray(j.Peers)
}

// ChatlistsJoinChatlistInvite invokes method chatlists.joinChatlistInvite#a6b1e39a returning error if any.
//
// See https://core.telegram.org/method/chatlists.joinChatlistInvite for reference.
func (c *Client) ChatlistsJoinChatlistInvite(ctx context.Context, request *ChatlistsJoinChatlistInviteRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
