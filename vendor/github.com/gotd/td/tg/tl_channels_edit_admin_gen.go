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

// ChannelsEditAdminRequest represents TL type `channels.editAdmin#d33c8902`.
// Modify the admin rights of a user in a supergroup/channel¹.
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.editAdmin for reference.
type ChannelsEditAdminRequest struct {
	// The supergroup/channel¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
	// The ID of the user whose admin rights should be modified
	UserID InputUserClass
	// The admin rights
	AdminRights ChatAdminRights
	// Indicates the role (rank) of the admin in the group: just an arbitrary string
	Rank string
}

// ChannelsEditAdminRequestTypeID is TL type id of ChannelsEditAdminRequest.
const ChannelsEditAdminRequestTypeID = 0xd33c8902

// Ensuring interfaces in compile-time for ChannelsEditAdminRequest.
var (
	_ bin.Encoder     = &ChannelsEditAdminRequest{}
	_ bin.Decoder     = &ChannelsEditAdminRequest{}
	_ bin.BareEncoder = &ChannelsEditAdminRequest{}
	_ bin.BareDecoder = &ChannelsEditAdminRequest{}
)

func (e *ChannelsEditAdminRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Channel == nil) {
		return false
	}
	if !(e.UserID == nil) {
		return false
	}
	if !(e.AdminRights.Zero()) {
		return false
	}
	if !(e.Rank == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *ChannelsEditAdminRequest) String() string {
	if e == nil {
		return "ChannelsEditAdminRequest(nil)"
	}
	type Alias ChannelsEditAdminRequest
	return fmt.Sprintf("ChannelsEditAdminRequest%+v", Alias(*e))
}

// FillFrom fills ChannelsEditAdminRequest from given interface.
func (e *ChannelsEditAdminRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetUserID() (value InputUserClass)
	GetAdminRights() (value ChatAdminRights)
	GetRank() (value string)
}) {
	e.Channel = from.GetChannel()
	e.UserID = from.GetUserID()
	e.AdminRights = from.GetAdminRights()
	e.Rank = from.GetRank()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsEditAdminRequest) TypeID() uint32 {
	return ChannelsEditAdminRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsEditAdminRequest) TypeName() string {
	return "channels.editAdmin"
}

// TypeInfo returns info about TL type.
func (e *ChannelsEditAdminRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.editAdmin",
		ID:   ChannelsEditAdminRequestTypeID,
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
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "AdminRights",
			SchemaName: "admin_rights",
		},
		{
			Name:       "Rank",
			SchemaName: "rank",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *ChannelsEditAdminRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editAdmin#d33c8902 as nil")
	}
	b.PutID(ChannelsEditAdminRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *ChannelsEditAdminRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editAdmin#d33c8902 as nil")
	}
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.editAdmin#d33c8902: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editAdmin#d33c8902: field channel: %w", err)
	}
	if e.UserID == nil {
		return fmt.Errorf("unable to encode channels.editAdmin#d33c8902: field user_id is nil")
	}
	if err := e.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editAdmin#d33c8902: field user_id: %w", err)
	}
	if err := e.AdminRights.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editAdmin#d33c8902: field admin_rights: %w", err)
	}
	b.PutString(e.Rank)
	return nil
}

// Decode implements bin.Decoder.
func (e *ChannelsEditAdminRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editAdmin#d33c8902 to nil")
	}
	if err := b.ConsumeID(ChannelsEditAdminRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.editAdmin#d33c8902: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *ChannelsEditAdminRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editAdmin#d33c8902 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editAdmin#d33c8902: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editAdmin#d33c8902: field user_id: %w", err)
		}
		e.UserID = value
	}
	{
		if err := e.AdminRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.editAdmin#d33c8902: field admin_rights: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.editAdmin#d33c8902: field rank: %w", err)
		}
		e.Rank = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (e *ChannelsEditAdminRequest) GetChannel() (value InputChannelClass) {
	if e == nil {
		return
	}
	return e.Channel
}

// GetUserID returns value of UserID field.
func (e *ChannelsEditAdminRequest) GetUserID() (value InputUserClass) {
	if e == nil {
		return
	}
	return e.UserID
}

// GetAdminRights returns value of AdminRights field.
func (e *ChannelsEditAdminRequest) GetAdminRights() (value ChatAdminRights) {
	if e == nil {
		return
	}
	return e.AdminRights
}

// GetRank returns value of Rank field.
func (e *ChannelsEditAdminRequest) GetRank() (value string) {
	if e == nil {
		return
	}
	return e.Rank
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (e *ChannelsEditAdminRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return e.Channel.AsNotEmpty()
}

// ChannelsEditAdmin invokes method channels.editAdmin#d33c8902 returning error if any.
// Modify the admin rights of a user in a supergroup/channel¹.
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// Possible errors:
//
//	400 ADMINS_TOO_MUCH: There are too many admins.
//	400 ADMIN_RANK_EMOJI_NOT_ALLOWED: An admin rank cannot contain emojis.
//	400 ADMIN_RANK_INVALID: The specified admin rank is invalid.
//	400 BOTS_TOO_MUCH: There are too many bots in this chat/channel.
//	400 BOT_CHANNELS_NA: Bots can't edit admin privileges.
//	400 BOT_GROUPS_BLOCKED: This bot can't be added to groups.
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	403 CHAT_ADMIN_INVITE_REQUIRED: You do not have the rights to do this.
//	403 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	406 FRESH_CHANGE_ADMINS_FORBIDDEN: You were just elected admin, you can't add or modify other admins yet.
//	400 INPUT_USER_DEACTIVATED: The specified user was deleted.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	403 RIGHT_FORBIDDEN: Your admin rights do not allow you to do this.
//	400 USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example).
//	400 USER_BLOCKED: User blocked.
//	403 USER_CHANNELS_TOO_MUCH: One of the users you tried to add is already in too many channels/supergroups.
//	400 USER_CREATOR: You can't leave this channel, because you're its creator.
//	400 USER_ID_INVALID: The provided user ID is invalid.
//	403 USER_NOT_MUTUAL_CONTACT: The provided user is not a mutual contact.
//	403 USER_PRIVACY_RESTRICTED: The user's privacy settings do not allow you to do this.
//	403 USER_RESTRICTED: You're spamreported, you can't create channels or chats.
//
// See https://core.telegram.org/method/channels.editAdmin for reference.
// Can be used by bots.
func (c *Client) ChannelsEditAdmin(ctx context.Context, request *ChannelsEditAdminRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}