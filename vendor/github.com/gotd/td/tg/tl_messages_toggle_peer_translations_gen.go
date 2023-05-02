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

// MessagesTogglePeerTranslationsRequest represents TL type `messages.togglePeerTranslations#e47cb579`.
// Toggle real-time chat translation¹ for a certain chat
//
// Links:
//  1. https://core.telegram.org/api/translation
//
// See https://core.telegram.org/method/messages.togglePeerTranslations for reference.
type MessagesTogglePeerTranslationsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to disable or enable real-time chat translation
	Disabled bool
	// Peer where to enable or disable real-time chat translation
	Peer InputPeerClass
}

// MessagesTogglePeerTranslationsRequestTypeID is TL type id of MessagesTogglePeerTranslationsRequest.
const MessagesTogglePeerTranslationsRequestTypeID = 0xe47cb579

// Ensuring interfaces in compile-time for MessagesTogglePeerTranslationsRequest.
var (
	_ bin.Encoder     = &MessagesTogglePeerTranslationsRequest{}
	_ bin.Decoder     = &MessagesTogglePeerTranslationsRequest{}
	_ bin.BareEncoder = &MessagesTogglePeerTranslationsRequest{}
	_ bin.BareDecoder = &MessagesTogglePeerTranslationsRequest{}
)

func (t *MessagesTogglePeerTranslationsRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.Disabled == false) {
		return false
	}
	if !(t.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *MessagesTogglePeerTranslationsRequest) String() string {
	if t == nil {
		return "MessagesTogglePeerTranslationsRequest(nil)"
	}
	type Alias MessagesTogglePeerTranslationsRequest
	return fmt.Sprintf("MessagesTogglePeerTranslationsRequest%+v", Alias(*t))
}

// FillFrom fills MessagesTogglePeerTranslationsRequest from given interface.
func (t *MessagesTogglePeerTranslationsRequest) FillFrom(from interface {
	GetDisabled() (value bool)
	GetPeer() (value InputPeerClass)
}) {
	t.Disabled = from.GetDisabled()
	t.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesTogglePeerTranslationsRequest) TypeID() uint32 {
	return MessagesTogglePeerTranslationsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesTogglePeerTranslationsRequest) TypeName() string {
	return "messages.togglePeerTranslations"
}

// TypeInfo returns info about TL type.
func (t *MessagesTogglePeerTranslationsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.togglePeerTranslations",
		ID:   MessagesTogglePeerTranslationsRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Disabled",
			SchemaName: "disabled",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (t *MessagesTogglePeerTranslationsRequest) SetFlags() {
	if !(t.Disabled == false) {
		t.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (t *MessagesTogglePeerTranslationsRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.togglePeerTranslations#e47cb579 as nil")
	}
	b.PutID(MessagesTogglePeerTranslationsRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *MessagesTogglePeerTranslationsRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.togglePeerTranslations#e47cb579 as nil")
	}
	t.SetFlags()
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.togglePeerTranslations#e47cb579: field flags: %w", err)
	}
	if t.Peer == nil {
		return fmt.Errorf("unable to encode messages.togglePeerTranslations#e47cb579: field peer is nil")
	}
	if err := t.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.togglePeerTranslations#e47cb579: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *MessagesTogglePeerTranslationsRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.togglePeerTranslations#e47cb579 to nil")
	}
	if err := b.ConsumeID(MessagesTogglePeerTranslationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.togglePeerTranslations#e47cb579: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *MessagesTogglePeerTranslationsRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.togglePeerTranslations#e47cb579 to nil")
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.togglePeerTranslations#e47cb579: field flags: %w", err)
		}
	}
	t.Disabled = t.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.togglePeerTranslations#e47cb579: field peer: %w", err)
		}
		t.Peer = value
	}
	return nil
}

// SetDisabled sets value of Disabled conditional field.
func (t *MessagesTogglePeerTranslationsRequest) SetDisabled(value bool) {
	if value {
		t.Flags.Set(0)
		t.Disabled = true
	} else {
		t.Flags.Unset(0)
		t.Disabled = false
	}
}

// GetDisabled returns value of Disabled conditional field.
func (t *MessagesTogglePeerTranslationsRequest) GetDisabled() (value bool) {
	if t == nil {
		return
	}
	return t.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (t *MessagesTogglePeerTranslationsRequest) GetPeer() (value InputPeerClass) {
	if t == nil {
		return
	}
	return t.Peer
}

// MessagesTogglePeerTranslations invokes method messages.togglePeerTranslations#e47cb579 returning error if any.
// Toggle real-time chat translation¹ for a certain chat
//
// Links:
//  1. https://core.telegram.org/api/translation
//
// See https://core.telegram.org/method/messages.togglePeerTranslations for reference.
// Can be used by bots.
func (c *Client) MessagesTogglePeerTranslations(ctx context.Context, request *MessagesTogglePeerTranslationsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
