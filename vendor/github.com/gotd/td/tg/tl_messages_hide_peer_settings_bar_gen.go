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

// MessagesHidePeerSettingsBarRequest represents TL type `messages.hidePeerSettingsBar#4facb138`.
// Should be called after the user hides the report spam/add as contact bar of a new chat
// effectively prevents the user from executing the actions specified in the peer's
// settings¹.
//
// Links:
//  1. https://core.telegram.org/constructor/peerSettings
//
// See https://core.telegram.org/method/messages.hidePeerSettingsBar for reference.
type MessagesHidePeerSettingsBarRequest struct {
	// Peer
	Peer InputPeerClass
}

// MessagesHidePeerSettingsBarRequestTypeID is TL type id of MessagesHidePeerSettingsBarRequest.
const MessagesHidePeerSettingsBarRequestTypeID = 0x4facb138

// Ensuring interfaces in compile-time for MessagesHidePeerSettingsBarRequest.
var (
	_ bin.Encoder     = &MessagesHidePeerSettingsBarRequest{}
	_ bin.Decoder     = &MessagesHidePeerSettingsBarRequest{}
	_ bin.BareEncoder = &MessagesHidePeerSettingsBarRequest{}
	_ bin.BareDecoder = &MessagesHidePeerSettingsBarRequest{}
)

func (h *MessagesHidePeerSettingsBarRequest) Zero() bool {
	if h == nil {
		return true
	}
	if !(h.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (h *MessagesHidePeerSettingsBarRequest) String() string {
	if h == nil {
		return "MessagesHidePeerSettingsBarRequest(nil)"
	}
	type Alias MessagesHidePeerSettingsBarRequest
	return fmt.Sprintf("MessagesHidePeerSettingsBarRequest%+v", Alias(*h))
}

// FillFrom fills MessagesHidePeerSettingsBarRequest from given interface.
func (h *MessagesHidePeerSettingsBarRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
}) {
	h.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesHidePeerSettingsBarRequest) TypeID() uint32 {
	return MessagesHidePeerSettingsBarRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesHidePeerSettingsBarRequest) TypeName() string {
	return "messages.hidePeerSettingsBar"
}

// TypeInfo returns info about TL type.
func (h *MessagesHidePeerSettingsBarRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.hidePeerSettingsBar",
		ID:   MessagesHidePeerSettingsBarRequestTypeID,
	}
	if h == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (h *MessagesHidePeerSettingsBarRequest) Encode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode messages.hidePeerSettingsBar#4facb138 as nil")
	}
	b.PutID(MessagesHidePeerSettingsBarRequestTypeID)
	return h.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (h *MessagesHidePeerSettingsBarRequest) EncodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode messages.hidePeerSettingsBar#4facb138 as nil")
	}
	if h.Peer == nil {
		return fmt.Errorf("unable to encode messages.hidePeerSettingsBar#4facb138: field peer is nil")
	}
	if err := h.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.hidePeerSettingsBar#4facb138: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (h *MessagesHidePeerSettingsBarRequest) Decode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode messages.hidePeerSettingsBar#4facb138 to nil")
	}
	if err := b.ConsumeID(MessagesHidePeerSettingsBarRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.hidePeerSettingsBar#4facb138: %w", err)
	}
	return h.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (h *MessagesHidePeerSettingsBarRequest) DecodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode messages.hidePeerSettingsBar#4facb138 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.hidePeerSettingsBar#4facb138: field peer: %w", err)
		}
		h.Peer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (h *MessagesHidePeerSettingsBarRequest) GetPeer() (value InputPeerClass) {
	if h == nil {
		return
	}
	return h.Peer
}

// MessagesHidePeerSettingsBar invokes method messages.hidePeerSettingsBar#4facb138 returning error if any.
// Should be called after the user hides the report spam/add as contact bar of a new chat
// effectively prevents the user from executing the actions specified in the peer's
// settings¹.
//
// Links:
//  1. https://core.telegram.org/constructor/peerSettings
//
// See https://core.telegram.org/method/messages.hidePeerSettingsBar for reference.
func (c *Client) MessagesHidePeerSettingsBar(ctx context.Context, peer InputPeerClass) (bool, error) {
	var result BoolBox

	request := &MessagesHidePeerSettingsBarRequest{
		Peer: peer,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}