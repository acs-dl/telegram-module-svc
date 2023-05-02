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

// MessagesGetUnreadMentionsRequest represents TL type `messages.getUnreadMentions#f107e790`.
// Get unread messages where we were mentioned
//
// See https://core.telegram.org/method/messages.getUnreadMentions for reference.
type MessagesGetUnreadMentionsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Peer where to look for mentions
	Peer InputPeerClass
	// If set, considers only messages within the specified forum topic¹
	//
	// Links:
	//  1) https://core.telegram.org/api/forum#forum-topics
	//
	// Use SetTopMsgID and GetTopMsgID helpers.
	TopMsgID int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetID int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	AddOffset int
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
	// Maximum message ID to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	MaxID int
	// Minimum message ID to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	MinID int
}

// MessagesGetUnreadMentionsRequestTypeID is TL type id of MessagesGetUnreadMentionsRequest.
const MessagesGetUnreadMentionsRequestTypeID = 0xf107e790

// Ensuring interfaces in compile-time for MessagesGetUnreadMentionsRequest.
var (
	_ bin.Encoder     = &MessagesGetUnreadMentionsRequest{}
	_ bin.Decoder     = &MessagesGetUnreadMentionsRequest{}
	_ bin.BareEncoder = &MessagesGetUnreadMentionsRequest{}
	_ bin.BareDecoder = &MessagesGetUnreadMentionsRequest{}
)

func (g *MessagesGetUnreadMentionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.TopMsgID == 0) {
		return false
	}
	if !(g.OffsetID == 0) {
		return false
	}
	if !(g.AddOffset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.MaxID == 0) {
		return false
	}
	if !(g.MinID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetUnreadMentionsRequest) String() string {
	if g == nil {
		return "MessagesGetUnreadMentionsRequest(nil)"
	}
	type Alias MessagesGetUnreadMentionsRequest
	return fmt.Sprintf("MessagesGetUnreadMentionsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetUnreadMentionsRequest from given interface.
func (g *MessagesGetUnreadMentionsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetTopMsgID() (value int, ok bool)
	GetOffsetID() (value int)
	GetAddOffset() (value int)
	GetLimit() (value int)
	GetMaxID() (value int)
	GetMinID() (value int)
}) {
	g.Peer = from.GetPeer()
	if val, ok := from.GetTopMsgID(); ok {
		g.TopMsgID = val
	}

	g.OffsetID = from.GetOffsetID()
	g.AddOffset = from.GetAddOffset()
	g.Limit = from.GetLimit()
	g.MaxID = from.GetMaxID()
	g.MinID = from.GetMinID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetUnreadMentionsRequest) TypeID() uint32 {
	return MessagesGetUnreadMentionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetUnreadMentionsRequest) TypeName() string {
	return "messages.getUnreadMentions"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetUnreadMentionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getUnreadMentions",
		ID:   MessagesGetUnreadMentionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "TopMsgID",
			SchemaName: "top_msg_id",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "AddOffset",
			SchemaName: "add_offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
		{
			Name:       "MinID",
			SchemaName: "min_id",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *MessagesGetUnreadMentionsRequest) SetFlags() {
	if !(g.TopMsgID == 0) {
		g.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (g *MessagesGetUnreadMentionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getUnreadMentions#f107e790 as nil")
	}
	b.PutID(MessagesGetUnreadMentionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetUnreadMentionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getUnreadMentions#f107e790 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getUnreadMentions#f107e790: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getUnreadMentions#f107e790: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getUnreadMentions#f107e790: field peer: %w", err)
	}
	if g.Flags.Has(0) {
		b.PutInt(g.TopMsgID)
	}
	b.PutInt(g.OffsetID)
	b.PutInt(g.AddOffset)
	b.PutInt(g.Limit)
	b.PutInt(g.MaxID)
	b.PutInt(g.MinID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetUnreadMentionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getUnreadMentions#f107e790 to nil")
	}
	if err := b.ConsumeID(MessagesGetUnreadMentionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getUnreadMentions#f107e790: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetUnreadMentionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getUnreadMentions#f107e790 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadMentions#f107e790: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadMentions#f107e790: field peer: %w", err)
		}
		g.Peer = value
	}
	if g.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadMentions#f107e790: field top_msg_id: %w", err)
		}
		g.TopMsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadMentions#f107e790: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadMentions#f107e790: field add_offset: %w", err)
		}
		g.AddOffset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadMentions#f107e790: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadMentions#f107e790: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadMentions#f107e790: field min_id: %w", err)
		}
		g.MinID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetUnreadMentionsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// SetTopMsgID sets value of TopMsgID conditional field.
func (g *MessagesGetUnreadMentionsRequest) SetTopMsgID(value int) {
	g.Flags.Set(0)
	g.TopMsgID = value
}

// GetTopMsgID returns value of TopMsgID conditional field and
// boolean which is true if field was set.
func (g *MessagesGetUnreadMentionsRequest) GetTopMsgID() (value int, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.TopMsgID, true
}

// GetOffsetID returns value of OffsetID field.
func (g *MessagesGetUnreadMentionsRequest) GetOffsetID() (value int) {
	if g == nil {
		return
	}
	return g.OffsetID
}

// GetAddOffset returns value of AddOffset field.
func (g *MessagesGetUnreadMentionsRequest) GetAddOffset() (value int) {
	if g == nil {
		return
	}
	return g.AddOffset
}

// GetLimit returns value of Limit field.
func (g *MessagesGetUnreadMentionsRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetMaxID returns value of MaxID field.
func (g *MessagesGetUnreadMentionsRequest) GetMaxID() (value int) {
	if g == nil {
		return
	}
	return g.MaxID
}

// GetMinID returns value of MinID field.
func (g *MessagesGetUnreadMentionsRequest) GetMinID() (value int) {
	if g == nil {
		return
	}
	return g.MinID
}

// MessagesGetUnreadMentions invokes method messages.getUnreadMentions#f107e790 returning error if any.
// Get unread messages where we were mentioned
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.getUnreadMentions for reference.
func (c *Client) MessagesGetUnreadMentions(ctx context.Context, request *MessagesGetUnreadMentionsRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
