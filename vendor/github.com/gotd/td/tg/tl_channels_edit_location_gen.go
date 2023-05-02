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

// ChannelsEditLocationRequest represents TL type `channels.editLocation#58e63f6d`.
// Edit location of geogroup
//
// See https://core.telegram.org/method/channels.editLocation for reference.
type ChannelsEditLocationRequest struct {
	// Geogroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
	// New geolocation
	GeoPoint InputGeoPointClass
	// Address string
	Address string
}

// ChannelsEditLocationRequestTypeID is TL type id of ChannelsEditLocationRequest.
const ChannelsEditLocationRequestTypeID = 0x58e63f6d

// Ensuring interfaces in compile-time for ChannelsEditLocationRequest.
var (
	_ bin.Encoder     = &ChannelsEditLocationRequest{}
	_ bin.Decoder     = &ChannelsEditLocationRequest{}
	_ bin.BareEncoder = &ChannelsEditLocationRequest{}
	_ bin.BareDecoder = &ChannelsEditLocationRequest{}
)

func (e *ChannelsEditLocationRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Channel == nil) {
		return false
	}
	if !(e.GeoPoint == nil) {
		return false
	}
	if !(e.Address == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *ChannelsEditLocationRequest) String() string {
	if e == nil {
		return "ChannelsEditLocationRequest(nil)"
	}
	type Alias ChannelsEditLocationRequest
	return fmt.Sprintf("ChannelsEditLocationRequest%+v", Alias(*e))
}

// FillFrom fills ChannelsEditLocationRequest from given interface.
func (e *ChannelsEditLocationRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetGeoPoint() (value InputGeoPointClass)
	GetAddress() (value string)
}) {
	e.Channel = from.GetChannel()
	e.GeoPoint = from.GetGeoPoint()
	e.Address = from.GetAddress()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsEditLocationRequest) TypeID() uint32 {
	return ChannelsEditLocationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsEditLocationRequest) TypeName() string {
	return "channels.editLocation"
}

// TypeInfo returns info about TL type.
func (e *ChannelsEditLocationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.editLocation",
		ID:   ChannelsEditLocationRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "GeoPoint",
			SchemaName: "geo_point",
		},
		{
			Name:       "Address",
			SchemaName: "address",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *ChannelsEditLocationRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editLocation#58e63f6d as nil")
	}
	b.PutID(ChannelsEditLocationRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *ChannelsEditLocationRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editLocation#58e63f6d as nil")
	}
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.editLocation#58e63f6d: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editLocation#58e63f6d: field channel: %w", err)
	}
	if e.GeoPoint == nil {
		return fmt.Errorf("unable to encode channels.editLocation#58e63f6d: field geo_point is nil")
	}
	if err := e.GeoPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editLocation#58e63f6d: field geo_point: %w", err)
	}
	b.PutString(e.Address)
	return nil
}

// Decode implements bin.Decoder.
func (e *ChannelsEditLocationRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editLocation#58e63f6d to nil")
	}
	if err := b.ConsumeID(ChannelsEditLocationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.editLocation#58e63f6d: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *ChannelsEditLocationRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editLocation#58e63f6d to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editLocation#58e63f6d: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := DecodeInputGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editLocation#58e63f6d: field geo_point: %w", err)
		}
		e.GeoPoint = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.editLocation#58e63f6d: field address: %w", err)
		}
		e.Address = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (e *ChannelsEditLocationRequest) GetChannel() (value InputChannelClass) {
	if e == nil {
		return
	}
	return e.Channel
}

// GetGeoPoint returns value of GeoPoint field.
func (e *ChannelsEditLocationRequest) GetGeoPoint() (value InputGeoPointClass) {
	if e == nil {
		return
	}
	return e.GeoPoint
}

// GetAddress returns value of Address field.
func (e *ChannelsEditLocationRequest) GetAddress() (value string) {
	if e == nil {
		return
	}
	return e.Address
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (e *ChannelsEditLocationRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return e.Channel.AsNotEmpty()
}

// GetGeoPointAsNotEmpty returns mapped value of GeoPoint field.
func (e *ChannelsEditLocationRequest) GetGeoPointAsNotEmpty() (*InputGeoPoint, bool) {
	return e.GeoPoint.AsNotEmpty()
}

// ChannelsEditLocation invokes method channels.editLocation#58e63f6d returning error if any.
// Edit location of geogroup
//
// Possible errors:
//
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 CHAT_NOT_MODIFIED: The pinned message wasn't modified.
//	400 MEGAGROUP_REQUIRED: You can only use this method on a supergroup.
//
// See https://core.telegram.org/method/channels.editLocation for reference.
func (c *Client) ChannelsEditLocation(ctx context.Context, request *ChannelsEditLocationRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
