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

// ChannelsDeleteChannelRequest represents TL type `channels.deleteChannel#c0111fe3`.
// Delete a channel/supergroup¹
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.deleteChannel for reference.
type ChannelsDeleteChannelRequest struct {
	// Channel/supergroup¹ to delete
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
}

// ChannelsDeleteChannelRequestTypeID is TL type id of ChannelsDeleteChannelRequest.
const ChannelsDeleteChannelRequestTypeID = 0xc0111fe3

// Ensuring interfaces in compile-time for ChannelsDeleteChannelRequest.
var (
	_ bin.Encoder     = &ChannelsDeleteChannelRequest{}
	_ bin.Decoder     = &ChannelsDeleteChannelRequest{}
	_ bin.BareEncoder = &ChannelsDeleteChannelRequest{}
	_ bin.BareDecoder = &ChannelsDeleteChannelRequest{}
)

func (d *ChannelsDeleteChannelRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Channel == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *ChannelsDeleteChannelRequest) String() string {
	if d == nil {
		return "ChannelsDeleteChannelRequest(nil)"
	}
	type Alias ChannelsDeleteChannelRequest
	return fmt.Sprintf("ChannelsDeleteChannelRequest%+v", Alias(*d))
}

// FillFrom fills ChannelsDeleteChannelRequest from given interface.
func (d *ChannelsDeleteChannelRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
}) {
	d.Channel = from.GetChannel()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsDeleteChannelRequest) TypeID() uint32 {
	return ChannelsDeleteChannelRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsDeleteChannelRequest) TypeName() string {
	return "channels.deleteChannel"
}

// TypeInfo returns info about TL type.
func (d *ChannelsDeleteChannelRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.deleteChannel",
		ID:   ChannelsDeleteChannelRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *ChannelsDeleteChannelRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteChannel#c0111fe3 as nil")
	}
	b.PutID(ChannelsDeleteChannelRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *ChannelsDeleteChannelRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteChannel#c0111fe3 as nil")
	}
	if d.Channel == nil {
		return fmt.Errorf("unable to encode channels.deleteChannel#c0111fe3: field channel is nil")
	}
	if err := d.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteChannel#c0111fe3: field channel: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *ChannelsDeleteChannelRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteChannel#c0111fe3 to nil")
	}
	if err := b.ConsumeID(ChannelsDeleteChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.deleteChannel#c0111fe3: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *ChannelsDeleteChannelRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteChannel#c0111fe3 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteChannel#c0111fe3: field channel: %w", err)
		}
		d.Channel = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (d *ChannelsDeleteChannelRequest) GetChannel() (value InputChannelClass) {
	if d == nil {
		return
	}
	return d.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (d *ChannelsDeleteChannelRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return d.Channel.AsNotEmpty()
}

// ChannelsDeleteChannel invokes method channels.deleteChannel#c0111fe3 returning error if any.
// Delete a channel/supergroup¹
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	406 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	406 CHANNEL_TOO_LARGE: Channel is too large to be deleted; this error is issued when trying to delete channels with more than 1000 members (subject to change).
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 CHAT_NOT_MODIFIED: The pinned message wasn't modified.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//
// See https://core.telegram.org/method/channels.deleteChannel for reference.
func (c *Client) ChannelsDeleteChannel(ctx context.Context, channel InputChannelClass) (UpdatesClass, error) {
	var result UpdatesBox

	request := &ChannelsDeleteChannelRequest{
		Channel: channel,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
