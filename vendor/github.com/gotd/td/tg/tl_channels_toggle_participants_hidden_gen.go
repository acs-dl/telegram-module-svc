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

// ChannelsToggleParticipantsHiddenRequest represents TL type `channels.toggleParticipantsHidden#6a6e7854`.
// Hide or display the participants list in a supergroup
//
// See https://core.telegram.org/method/channels.toggleParticipantsHidden for reference.
type ChannelsToggleParticipantsHiddenRequest struct {
	// Supergroup ID
	Channel InputChannelClass
	// If true, will hide the participants list; otherwise will unhide it.
	Enabled bool
}

// ChannelsToggleParticipantsHiddenRequestTypeID is TL type id of ChannelsToggleParticipantsHiddenRequest.
const ChannelsToggleParticipantsHiddenRequestTypeID = 0x6a6e7854

// Ensuring interfaces in compile-time for ChannelsToggleParticipantsHiddenRequest.
var (
	_ bin.Encoder     = &ChannelsToggleParticipantsHiddenRequest{}
	_ bin.Decoder     = &ChannelsToggleParticipantsHiddenRequest{}
	_ bin.BareEncoder = &ChannelsToggleParticipantsHiddenRequest{}
	_ bin.BareDecoder = &ChannelsToggleParticipantsHiddenRequest{}
)

func (t *ChannelsToggleParticipantsHiddenRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Channel == nil) {
		return false
	}
	if !(t.Enabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ChannelsToggleParticipantsHiddenRequest) String() string {
	if t == nil {
		return "ChannelsToggleParticipantsHiddenRequest(nil)"
	}
	type Alias ChannelsToggleParticipantsHiddenRequest
	return fmt.Sprintf("ChannelsToggleParticipantsHiddenRequest%+v", Alias(*t))
}

// FillFrom fills ChannelsToggleParticipantsHiddenRequest from given interface.
func (t *ChannelsToggleParticipantsHiddenRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetEnabled() (value bool)
}) {
	t.Channel = from.GetChannel()
	t.Enabled = from.GetEnabled()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsToggleParticipantsHiddenRequest) TypeID() uint32 {
	return ChannelsToggleParticipantsHiddenRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsToggleParticipantsHiddenRequest) TypeName() string {
	return "channels.toggleParticipantsHidden"
}

// TypeInfo returns info about TL type.
func (t *ChannelsToggleParticipantsHiddenRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.toggleParticipantsHidden",
		ID:   ChannelsToggleParticipantsHiddenRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Enabled",
			SchemaName: "enabled",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ChannelsToggleParticipantsHiddenRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleParticipantsHidden#6a6e7854 as nil")
	}
	b.PutID(ChannelsToggleParticipantsHiddenRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ChannelsToggleParticipantsHiddenRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleParticipantsHidden#6a6e7854 as nil")
	}
	if t.Channel == nil {
		return fmt.Errorf("unable to encode channels.toggleParticipantsHidden#6a6e7854: field channel is nil")
	}
	if err := t.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.toggleParticipantsHidden#6a6e7854: field channel: %w", err)
	}
	b.PutBool(t.Enabled)
	return nil
}

// Decode implements bin.Decoder.
func (t *ChannelsToggleParticipantsHiddenRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleParticipantsHidden#6a6e7854 to nil")
	}
	if err := b.ConsumeID(ChannelsToggleParticipantsHiddenRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.toggleParticipantsHidden#6a6e7854: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ChannelsToggleParticipantsHiddenRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleParticipantsHidden#6a6e7854 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleParticipantsHidden#6a6e7854: field channel: %w", err)
		}
		t.Channel = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleParticipantsHidden#6a6e7854: field enabled: %w", err)
		}
		t.Enabled = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (t *ChannelsToggleParticipantsHiddenRequest) GetChannel() (value InputChannelClass) {
	if t == nil {
		return
	}
	return t.Channel
}

// GetEnabled returns value of Enabled field.
func (t *ChannelsToggleParticipantsHiddenRequest) GetEnabled() (value bool) {
	if t == nil {
		return
	}
	return t.Enabled
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (t *ChannelsToggleParticipantsHiddenRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return t.Channel.AsNotEmpty()
}

// ChannelsToggleParticipantsHidden invokes method channels.toggleParticipantsHidden#6a6e7854 returning error if any.
// Hide or display the participants list in a supergroup
//
// See https://core.telegram.org/method/channels.toggleParticipantsHidden for reference.
// Can be used by bots.
func (c *Client) ChannelsToggleParticipantsHidden(ctx context.Context, request *ChannelsToggleParticipantsHiddenRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
