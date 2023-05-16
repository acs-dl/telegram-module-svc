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

// MessagesGetMessageReactionsListRequest represents TL type `messages.getMessageReactionsList#461b3f48`.
// Get message reaction¹ list, along with the sender of each reaction.
//
// Links:
//  1. https://core.telegram.org/api/reactions
//
// See https://core.telegram.org/method/messages.getMessageReactionsList for reference.
type MessagesGetMessageReactionsListRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Peer
	Peer InputPeerClass
	// Message ID
	ID int
	// Get only reactions of this type (UTF8 emoji)
	//
	// Use SetReaction and GetReaction helpers.
	Reaction ReactionClass
	// Offset (typically taken from the next_offset field of the returned messages
	// MessageReactionsList¹)
	//
	// Links:
	//  1) https://core.telegram.org/type/messages.MessageReactionsList
	//
	// Use SetOffset and GetOffset helpers.
	Offset string
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// MessagesGetMessageReactionsListRequestTypeID is TL type id of MessagesGetMessageReactionsListRequest.
const MessagesGetMessageReactionsListRequestTypeID = 0x461b3f48

// Ensuring interfaces in compile-time for MessagesGetMessageReactionsListRequest.
var (
	_ bin.Encoder     = &MessagesGetMessageReactionsListRequest{}
	_ bin.Decoder     = &MessagesGetMessageReactionsListRequest{}
	_ bin.BareEncoder = &MessagesGetMessageReactionsListRequest{}
	_ bin.BareDecoder = &MessagesGetMessageReactionsListRequest{}
)

func (g *MessagesGetMessageReactionsListRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.ID == 0) {
		return false
	}
	if !(g.Reaction == nil) {
		return false
	}
	if !(g.Offset == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetMessageReactionsListRequest) String() string {
	if g == nil {
		return "MessagesGetMessageReactionsListRequest(nil)"
	}
	type Alias MessagesGetMessageReactionsListRequest
	return fmt.Sprintf("MessagesGetMessageReactionsListRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetMessageReactionsListRequest from given interface.
func (g *MessagesGetMessageReactionsListRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetID() (value int)
	GetReaction() (value ReactionClass, ok bool)
	GetOffset() (value string, ok bool)
	GetLimit() (value int)
}) {
	g.Peer = from.GetPeer()
	g.ID = from.GetID()
	if val, ok := from.GetReaction(); ok {
		g.Reaction = val
	}

	if val, ok := from.GetOffset(); ok {
		g.Offset = val
	}

	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetMessageReactionsListRequest) TypeID() uint32 {
	return MessagesGetMessageReactionsListRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetMessageReactionsListRequest) TypeName() string {
	return "messages.getMessageReactionsList"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetMessageReactionsListRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getMessageReactionsList",
		ID:   MessagesGetMessageReactionsListRequestTypeID,
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
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Reaction",
			SchemaName: "reaction",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *MessagesGetMessageReactionsListRequest) SetFlags() {
	if !(g.Reaction == nil) {
		g.Flags.Set(0)
	}
	if !(g.Offset == "") {
		g.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (g *MessagesGetMessageReactionsListRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMessageReactionsList#461b3f48 as nil")
	}
	b.PutID(MessagesGetMessageReactionsListRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetMessageReactionsListRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMessageReactionsList#461b3f48 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getMessageReactionsList#461b3f48: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getMessageReactionsList#461b3f48: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getMessageReactionsList#461b3f48: field peer: %w", err)
	}
	b.PutInt(g.ID)
	if g.Flags.Has(0) {
		if g.Reaction == nil {
			return fmt.Errorf("unable to encode messages.getMessageReactionsList#461b3f48: field reaction is nil")
		}
		if err := g.Reaction.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.getMessageReactionsList#461b3f48: field reaction: %w", err)
		}
	}
	if g.Flags.Has(1) {
		b.PutString(g.Offset)
	}
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetMessageReactionsListRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMessageReactionsList#461b3f48 to nil")
	}
	if err := b.ConsumeID(MessagesGetMessageReactionsListRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getMessageReactionsList#461b3f48: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetMessageReactionsListRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMessageReactionsList#461b3f48 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getMessageReactionsList#461b3f48: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessageReactionsList#461b3f48: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessageReactionsList#461b3f48: field id: %w", err)
		}
		g.ID = value
	}
	if g.Flags.Has(0) {
		value, err := DecodeReaction(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessageReactionsList#461b3f48: field reaction: %w", err)
		}
		g.Reaction = value
	}
	if g.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessageReactionsList#461b3f48: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessageReactionsList#461b3f48: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetMessageReactionsListRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetID returns value of ID field.
func (g *MessagesGetMessageReactionsListRequest) GetID() (value int) {
	if g == nil {
		return
	}
	return g.ID
}

// SetReaction sets value of Reaction conditional field.
func (g *MessagesGetMessageReactionsListRequest) SetReaction(value ReactionClass) {
	g.Flags.Set(0)
	g.Reaction = value
}

// GetReaction returns value of Reaction conditional field and
// boolean which is true if field was set.
func (g *MessagesGetMessageReactionsListRequest) GetReaction() (value ReactionClass, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.Reaction, true
}

// SetOffset sets value of Offset conditional field.
func (g *MessagesGetMessageReactionsListRequest) SetOffset(value string) {
	g.Flags.Set(1)
	g.Offset = value
}

// GetOffset returns value of Offset conditional field and
// boolean which is true if field was set.
func (g *MessagesGetMessageReactionsListRequest) GetOffset() (value string, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(1) {
		return value, false
	}
	return g.Offset, true
}

// GetLimit returns value of Limit field.
func (g *MessagesGetMessageReactionsListRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// MessagesGetMessageReactionsList invokes method messages.getMessageReactionsList#461b3f48 returning error if any.
// Get message reaction¹ list, along with the sender of each reaction.
//
// Links:
//  1. https://core.telegram.org/api/reactions
//
// Possible errors:
//
//	403 BROADCAST_FORBIDDEN: Participants of polls in channels should stay anonymous.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//
// See https://core.telegram.org/method/messages.getMessageReactionsList for reference.
func (c *Client) MessagesGetMessageReactionsList(ctx context.Context, request *MessagesGetMessageReactionsListRequest) (*MessagesMessageReactionsList, error) {
	var result MessagesMessageReactionsList

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
