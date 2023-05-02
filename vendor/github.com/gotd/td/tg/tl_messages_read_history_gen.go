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

// MessagesReadHistoryRequest represents TL type `messages.readHistory#e306d3a`.
// Marks message history as read.
//
// See https://core.telegram.org/method/messages.readHistory for reference.
type MessagesReadHistoryRequest struct {
	// Target user or group
	Peer InputPeerClass
	// If a positive value is passed, only messages with identifiers less or equal than the
	// given one will be read
	MaxID int
}

// MessagesReadHistoryRequestTypeID is TL type id of MessagesReadHistoryRequest.
const MessagesReadHistoryRequestTypeID = 0xe306d3a

// Ensuring interfaces in compile-time for MessagesReadHistoryRequest.
var (
	_ bin.Encoder     = &MessagesReadHistoryRequest{}
	_ bin.Decoder     = &MessagesReadHistoryRequest{}
	_ bin.BareEncoder = &MessagesReadHistoryRequest{}
	_ bin.BareDecoder = &MessagesReadHistoryRequest{}
)

func (r *MessagesReadHistoryRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.MaxID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReadHistoryRequest) String() string {
	if r == nil {
		return "MessagesReadHistoryRequest(nil)"
	}
	type Alias MessagesReadHistoryRequest
	return fmt.Sprintf("MessagesReadHistoryRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReadHistoryRequest from given interface.
func (r *MessagesReadHistoryRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMaxID() (value int)
}) {
	r.Peer = from.GetPeer()
	r.MaxID = from.GetMaxID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReadHistoryRequest) TypeID() uint32 {
	return MessagesReadHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReadHistoryRequest) TypeName() string {
	return "messages.readHistory"
}

// TypeInfo returns info about TL type.
func (r *MessagesReadHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.readHistory",
		ID:   MessagesReadHistoryRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesReadHistoryRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readHistory#e306d3a as nil")
	}
	b.PutID(MessagesReadHistoryRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReadHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readHistory#e306d3a as nil")
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.readHistory#e306d3a: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.readHistory#e306d3a: field peer: %w", err)
	}
	b.PutInt(r.MaxID)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReadHistoryRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readHistory#e306d3a to nil")
	}
	if err := b.ConsumeID(MessagesReadHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.readHistory#e306d3a: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReadHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readHistory#e306d3a to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.readHistory#e306d3a: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.readHistory#e306d3a: field max_id: %w", err)
		}
		r.MaxID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (r *MessagesReadHistoryRequest) GetPeer() (value InputPeerClass) {
	if r == nil {
		return
	}
	return r.Peer
}

// GetMaxID returns value of MaxID field.
func (r *MessagesReadHistoryRequest) GetMaxID() (value int) {
	if r == nil {
		return
	}
	return r.MaxID
}

// MessagesReadHistory invokes method messages.readHistory#e306d3a returning error if any.
// Marks message history as read.
//
// Possible errors:
//
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 CHAT_ID_INVALID: The provided chat id is invalid.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.readHistory for reference.
func (c *Client) MessagesReadHistory(ctx context.Context, request *MessagesReadHistoryRequest) (*MessagesAffectedMessages, error) {
	var result MessagesAffectedMessages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
