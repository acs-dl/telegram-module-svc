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

// ChannelsViewSponsoredMessageRequest represents TL type `channels.viewSponsoredMessage#beaedb94`.
// Mark a specific sponsored message as read
//
// See https://core.telegram.org/method/channels.viewSponsoredMessage for reference.
type ChannelsViewSponsoredMessageRequest struct {
	// Peer
	Channel InputChannelClass
	// Message ID
	RandomID []byte
}

// ChannelsViewSponsoredMessageRequestTypeID is TL type id of ChannelsViewSponsoredMessageRequest.
const ChannelsViewSponsoredMessageRequestTypeID = 0xbeaedb94

// Ensuring interfaces in compile-time for ChannelsViewSponsoredMessageRequest.
var (
	_ bin.Encoder     = &ChannelsViewSponsoredMessageRequest{}
	_ bin.Decoder     = &ChannelsViewSponsoredMessageRequest{}
	_ bin.BareEncoder = &ChannelsViewSponsoredMessageRequest{}
	_ bin.BareDecoder = &ChannelsViewSponsoredMessageRequest{}
)

func (v *ChannelsViewSponsoredMessageRequest) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Channel == nil) {
		return false
	}
	if !(v.RandomID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *ChannelsViewSponsoredMessageRequest) String() string {
	if v == nil {
		return "ChannelsViewSponsoredMessageRequest(nil)"
	}
	type Alias ChannelsViewSponsoredMessageRequest
	return fmt.Sprintf("ChannelsViewSponsoredMessageRequest%+v", Alias(*v))
}

// FillFrom fills ChannelsViewSponsoredMessageRequest from given interface.
func (v *ChannelsViewSponsoredMessageRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetRandomID() (value []byte)
}) {
	v.Channel = from.GetChannel()
	v.RandomID = from.GetRandomID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsViewSponsoredMessageRequest) TypeID() uint32 {
	return ChannelsViewSponsoredMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsViewSponsoredMessageRequest) TypeName() string {
	return "channels.viewSponsoredMessage"
}

// TypeInfo returns info about TL type.
func (v *ChannelsViewSponsoredMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.viewSponsoredMessage",
		ID:   ChannelsViewSponsoredMessageRequestTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *ChannelsViewSponsoredMessageRequest) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode channels.viewSponsoredMessage#beaedb94 as nil")
	}
	b.PutID(ChannelsViewSponsoredMessageRequestTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *ChannelsViewSponsoredMessageRequest) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode channels.viewSponsoredMessage#beaedb94 as nil")
	}
	if v.Channel == nil {
		return fmt.Errorf("unable to encode channels.viewSponsoredMessage#beaedb94: field channel is nil")
	}
	if err := v.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.viewSponsoredMessage#beaedb94: field channel: %w", err)
	}
	b.PutBytes(v.RandomID)
	return nil
}

// Decode implements bin.Decoder.
func (v *ChannelsViewSponsoredMessageRequest) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode channels.viewSponsoredMessage#beaedb94 to nil")
	}
	if err := b.ConsumeID(ChannelsViewSponsoredMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.viewSponsoredMessage#beaedb94: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *ChannelsViewSponsoredMessageRequest) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode channels.viewSponsoredMessage#beaedb94 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.viewSponsoredMessage#beaedb94: field channel: %w", err)
		}
		v.Channel = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode channels.viewSponsoredMessage#beaedb94: field random_id: %w", err)
		}
		v.RandomID = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (v *ChannelsViewSponsoredMessageRequest) GetChannel() (value InputChannelClass) {
	if v == nil {
		return
	}
	return v.Channel
}

// GetRandomID returns value of RandomID field.
func (v *ChannelsViewSponsoredMessageRequest) GetRandomID() (value []byte) {
	if v == nil {
		return
	}
	return v.RandomID
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (v *ChannelsViewSponsoredMessageRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return v.Channel.AsNotEmpty()
}

// ChannelsViewSponsoredMessage invokes method channels.viewSponsoredMessage#beaedb94 returning error if any.
// Mark a specific sponsored message as read
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//
// See https://core.telegram.org/method/channels.viewSponsoredMessage for reference.
func (c *Client) ChannelsViewSponsoredMessage(ctx context.Context, request *ChannelsViewSponsoredMessageRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
