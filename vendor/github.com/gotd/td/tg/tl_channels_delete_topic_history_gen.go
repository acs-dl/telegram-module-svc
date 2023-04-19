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

// ChannelsDeleteTopicHistoryRequest represents TL type `channels.deleteTopicHistory#34435f2d`.
//
// See https://core.telegram.org/method/channels.deleteTopicHistory for reference.
type ChannelsDeleteTopicHistoryRequest struct {
	// Channel field of ChannelsDeleteTopicHistoryRequest.
	Channel InputChannelClass
	// TopMsgID field of ChannelsDeleteTopicHistoryRequest.
	TopMsgID int
}

// ChannelsDeleteTopicHistoryRequestTypeID is TL type id of ChannelsDeleteTopicHistoryRequest.
const ChannelsDeleteTopicHistoryRequestTypeID = 0x34435f2d

// Ensuring interfaces in compile-time for ChannelsDeleteTopicHistoryRequest.
var (
	_ bin.Encoder     = &ChannelsDeleteTopicHistoryRequest{}
	_ bin.Decoder     = &ChannelsDeleteTopicHistoryRequest{}
	_ bin.BareEncoder = &ChannelsDeleteTopicHistoryRequest{}
	_ bin.BareDecoder = &ChannelsDeleteTopicHistoryRequest{}
)

func (d *ChannelsDeleteTopicHistoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Channel == nil) {
		return false
	}
	if !(d.TopMsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *ChannelsDeleteTopicHistoryRequest) String() string {
	if d == nil {
		return "ChannelsDeleteTopicHistoryRequest(nil)"
	}
	type Alias ChannelsDeleteTopicHistoryRequest
	return fmt.Sprintf("ChannelsDeleteTopicHistoryRequest%+v", Alias(*d))
}

// FillFrom fills ChannelsDeleteTopicHistoryRequest from given interface.
func (d *ChannelsDeleteTopicHistoryRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetTopMsgID() (value int)
}) {
	d.Channel = from.GetChannel()
	d.TopMsgID = from.GetTopMsgID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsDeleteTopicHistoryRequest) TypeID() uint32 {
	return ChannelsDeleteTopicHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsDeleteTopicHistoryRequest) TypeName() string {
	return "channels.deleteTopicHistory"
}

// TypeInfo returns info about TL type.
func (d *ChannelsDeleteTopicHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.deleteTopicHistory",
		ID:   ChannelsDeleteTopicHistoryRequestTypeID,
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
		{
			Name:       "TopMsgID",
			SchemaName: "top_msg_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *ChannelsDeleteTopicHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteTopicHistory#34435f2d as nil")
	}
	b.PutID(ChannelsDeleteTopicHistoryRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *ChannelsDeleteTopicHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteTopicHistory#34435f2d as nil")
	}
	if d.Channel == nil {
		return fmt.Errorf("unable to encode channels.deleteTopicHistory#34435f2d: field channel is nil")
	}
	if err := d.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteTopicHistory#34435f2d: field channel: %w", err)
	}
	b.PutInt(d.TopMsgID)
	return nil
}

// Decode implements bin.Decoder.
func (d *ChannelsDeleteTopicHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteTopicHistory#34435f2d to nil")
	}
	if err := b.ConsumeID(ChannelsDeleteTopicHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.deleteTopicHistory#34435f2d: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *ChannelsDeleteTopicHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteTopicHistory#34435f2d to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteTopicHistory#34435f2d: field channel: %w", err)
		}
		d.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteTopicHistory#34435f2d: field top_msg_id: %w", err)
		}
		d.TopMsgID = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (d *ChannelsDeleteTopicHistoryRequest) GetChannel() (value InputChannelClass) {
	if d == nil {
		return
	}
	return d.Channel
}

// GetTopMsgID returns value of TopMsgID field.
func (d *ChannelsDeleteTopicHistoryRequest) GetTopMsgID() (value int) {
	if d == nil {
		return
	}
	return d.TopMsgID
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (d *ChannelsDeleteTopicHistoryRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return d.Channel.AsNotEmpty()
}

// ChannelsDeleteTopicHistory invokes method channels.deleteTopicHistory#34435f2d returning error if any.
//
// See https://core.telegram.org/method/channels.deleteTopicHistory for reference.
func (c *Client) ChannelsDeleteTopicHistory(ctx context.Context, request *ChannelsDeleteTopicHistoryRequest) (*MessagesAffectedHistory, error) {
	var result MessagesAffectedHistory

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
