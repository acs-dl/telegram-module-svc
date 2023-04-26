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

// ChannelAdminLogEventsFilter represents TL type `channelAdminLogEventsFilter#ea107ae4`.
// Filter only certain admin log events
//
// See https://core.telegram.org/constructor/channelAdminLogEventsFilter for reference.
type ChannelAdminLogEventsFilter struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Join events¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channelAdminLogEventActionParticipantJoin
	Join bool
	// Leave events¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channelAdminLogEventActionParticipantLeave
	Leave bool
	// Invite events¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channelAdminLogEventActionParticipantInvite
	Invite bool
	// Ban events¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channelAdminLogEventActionParticipantToggleBan
	Ban bool
	// Unban events¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channelAdminLogEventActionParticipantToggleBan
	Unban bool
	// Kick events¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channelAdminLogEventActionParticipantToggleBan
	Kick bool
	// Unkick events¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channelAdminLogEventActionParticipantToggleBan
	Unkick bool
	// Admin promotion events¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channelAdminLogEventActionParticipantToggleAdmin
	Promote bool
	// Admin demotion events¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channelAdminLogEventActionParticipantToggleAdmin
	Demote bool
	// Info change events (when about¹, linked chat², location³, photo⁴, stickerset⁵,
	// title⁶ or username⁷ data of a channel gets modified)
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channelAdminLogEventActionChangeAbout
	//  2) https://core.telegram.org/constructor/channelAdminLogEventActionChangeLinkedChat
	//  3) https://core.telegram.org/constructor/channelAdminLogEventActionChangeLocation
	//  4) https://core.telegram.org/constructor/channelAdminLogEventActionChangePhoto
	//  5) https://core.telegram.org/constructor/channelAdminLogEventActionChangeStickerSet
	//  6) https://core.telegram.org/constructor/channelAdminLogEventActionChangeTitle
	//  7) https://core.telegram.org/constructor/channelAdminLogEventActionChangeUsername
	Info bool
	// Settings change events (invites¹, hidden prehistory², signatures³, default banned
	// rights⁴)
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channelAdminLogEventActionToggleInvites
	//  2) https://core.telegram.org/constructor/channelAdminLogEventActionTogglePreHistoryHidden
	//  3) https://core.telegram.org/constructor/channelAdminLogEventActionToggleSignatures
	//  4) https://core.telegram.org/constructor/channelAdminLogEventActionDefaultBannedRights
	Settings bool
	// Message pin events¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channelAdminLogEventActionUpdatePinned
	Pinned bool
	// Message edit events¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channelAdminLogEventActionEditMessage
	Edit bool
	// Message deletion events¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channelAdminLogEventActionDeleteMessage
	Delete bool
	// Group call events
	GroupCall bool
	// Invite events
	Invites bool
	// A message was posted in a channel
	Send bool
	// Forum¹-related events
	//
	// Links:
	//  1) https://core.telegram.org/api/forum#forum
	Forums bool
}

// ChannelAdminLogEventsFilterTypeID is TL type id of ChannelAdminLogEventsFilter.
const ChannelAdminLogEventsFilterTypeID = 0xea107ae4

// Ensuring interfaces in compile-time for ChannelAdminLogEventsFilter.
var (
	_ bin.Encoder     = &ChannelAdminLogEventsFilter{}
	_ bin.Decoder     = &ChannelAdminLogEventsFilter{}
	_ bin.BareEncoder = &ChannelAdminLogEventsFilter{}
	_ bin.BareDecoder = &ChannelAdminLogEventsFilter{}
)

func (c *ChannelAdminLogEventsFilter) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Join == false) {
		return false
	}
	if !(c.Leave == false) {
		return false
	}
	if !(c.Invite == false) {
		return false
	}
	if !(c.Ban == false) {
		return false
	}
	if !(c.Unban == false) {
		return false
	}
	if !(c.Kick == false) {
		return false
	}
	if !(c.Unkick == false) {
		return false
	}
	if !(c.Promote == false) {
		return false
	}
	if !(c.Demote == false) {
		return false
	}
	if !(c.Info == false) {
		return false
	}
	if !(c.Settings == false) {
		return false
	}
	if !(c.Pinned == false) {
		return false
	}
	if !(c.Edit == false) {
		return false
	}
	if !(c.Delete == false) {
		return false
	}
	if !(c.GroupCall == false) {
		return false
	}
	if !(c.Invites == false) {
		return false
	}
	if !(c.Send == false) {
		return false
	}
	if !(c.Forums == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelAdminLogEventsFilter) String() string {
	if c == nil {
		return "ChannelAdminLogEventsFilter(nil)"
	}
	type Alias ChannelAdminLogEventsFilter
	return fmt.Sprintf("ChannelAdminLogEventsFilter%+v", Alias(*c))
}

// FillFrom fills ChannelAdminLogEventsFilter from given interface.
func (c *ChannelAdminLogEventsFilter) FillFrom(from interface {
	GetJoin() (value bool)
	GetLeave() (value bool)
	GetInvite() (value bool)
	GetBan() (value bool)
	GetUnban() (value bool)
	GetKick() (value bool)
	GetUnkick() (value bool)
	GetPromote() (value bool)
	GetDemote() (value bool)
	GetInfo() (value bool)
	GetSettings() (value bool)
	GetPinned() (value bool)
	GetEdit() (value bool)
	GetDelete() (value bool)
	GetGroupCall() (value bool)
	GetInvites() (value bool)
	GetSend() (value bool)
	GetForums() (value bool)
}) {
	c.Join = from.GetJoin()
	c.Leave = from.GetLeave()
	c.Invite = from.GetInvite()
	c.Ban = from.GetBan()
	c.Unban = from.GetUnban()
	c.Kick = from.GetKick()
	c.Unkick = from.GetUnkick()
	c.Promote = from.GetPromote()
	c.Demote = from.GetDemote()
	c.Info = from.GetInfo()
	c.Settings = from.GetSettings()
	c.Pinned = from.GetPinned()
	c.Edit = from.GetEdit()
	c.Delete = from.GetDelete()
	c.GroupCall = from.GetGroupCall()
	c.Invites = from.GetInvites()
	c.Send = from.GetSend()
	c.Forums = from.GetForums()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelAdminLogEventsFilter) TypeID() uint32 {
	return ChannelAdminLogEventsFilterTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelAdminLogEventsFilter) TypeName() string {
	return "channelAdminLogEventsFilter"
}

// TypeInfo returns info about TL type.
func (c *ChannelAdminLogEventsFilter) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channelAdminLogEventsFilter",
		ID:   ChannelAdminLogEventsFilterTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Join",
			SchemaName: "join",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "Leave",
			SchemaName: "leave",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "Invite",
			SchemaName: "invite",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "Ban",
			SchemaName: "ban",
			Null:       !c.Flags.Has(3),
		},
		{
			Name:       "Unban",
			SchemaName: "unban",
			Null:       !c.Flags.Has(4),
		},
		{
			Name:       "Kick",
			SchemaName: "kick",
			Null:       !c.Flags.Has(5),
		},
		{
			Name:       "Unkick",
			SchemaName: "unkick",
			Null:       !c.Flags.Has(6),
		},
		{
			Name:       "Promote",
			SchemaName: "promote",
			Null:       !c.Flags.Has(7),
		},
		{
			Name:       "Demote",
			SchemaName: "demote",
			Null:       !c.Flags.Has(8),
		},
		{
			Name:       "Info",
			SchemaName: "info",
			Null:       !c.Flags.Has(9),
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
			Null:       !c.Flags.Has(10),
		},
		{
			Name:       "Pinned",
			SchemaName: "pinned",
			Null:       !c.Flags.Has(11),
		},
		{
			Name:       "Edit",
			SchemaName: "edit",
			Null:       !c.Flags.Has(12),
		},
		{
			Name:       "Delete",
			SchemaName: "delete",
			Null:       !c.Flags.Has(13),
		},
		{
			Name:       "GroupCall",
			SchemaName: "group_call",
			Null:       !c.Flags.Has(14),
		},
		{
			Name:       "Invites",
			SchemaName: "invites",
			Null:       !c.Flags.Has(15),
		},
		{
			Name:       "Send",
			SchemaName: "send",
			Null:       !c.Flags.Has(16),
		},
		{
			Name:       "Forums",
			SchemaName: "forums",
			Null:       !c.Flags.Has(17),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *ChannelAdminLogEventsFilter) SetFlags() {
	if !(c.Join == false) {
		c.Flags.Set(0)
	}
	if !(c.Leave == false) {
		c.Flags.Set(1)
	}
	if !(c.Invite == false) {
		c.Flags.Set(2)
	}
	if !(c.Ban == false) {
		c.Flags.Set(3)
	}
	if !(c.Unban == false) {
		c.Flags.Set(4)
	}
	if !(c.Kick == false) {
		c.Flags.Set(5)
	}
	if !(c.Unkick == false) {
		c.Flags.Set(6)
	}
	if !(c.Promote == false) {
		c.Flags.Set(7)
	}
	if !(c.Demote == false) {
		c.Flags.Set(8)
	}
	if !(c.Info == false) {
		c.Flags.Set(9)
	}
	if !(c.Settings == false) {
		c.Flags.Set(10)
	}
	if !(c.Pinned == false) {
		c.Flags.Set(11)
	}
	if !(c.Edit == false) {
		c.Flags.Set(12)
	}
	if !(c.Delete == false) {
		c.Flags.Set(13)
	}
	if !(c.GroupCall == false) {
		c.Flags.Set(14)
	}
	if !(c.Invites == false) {
		c.Flags.Set(15)
	}
	if !(c.Send == false) {
		c.Flags.Set(16)
	}
	if !(c.Forums == false) {
		c.Flags.Set(17)
	}
}

// Encode implements bin.Encoder.
func (c *ChannelAdminLogEventsFilter) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelAdminLogEventsFilter#ea107ae4 as nil")
	}
	b.PutID(ChannelAdminLogEventsFilterTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelAdminLogEventsFilter) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelAdminLogEventsFilter#ea107ae4 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channelAdminLogEventsFilter#ea107ae4: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelAdminLogEventsFilter) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelAdminLogEventsFilter#ea107ae4 to nil")
	}
	if err := b.ConsumeID(ChannelAdminLogEventsFilterTypeID); err != nil {
		return fmt.Errorf("unable to decode channelAdminLogEventsFilter#ea107ae4: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelAdminLogEventsFilter) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelAdminLogEventsFilter#ea107ae4 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channelAdminLogEventsFilter#ea107ae4: field flags: %w", err)
		}
	}
	c.Join = c.Flags.Has(0)
	c.Leave = c.Flags.Has(1)
	c.Invite = c.Flags.Has(2)
	c.Ban = c.Flags.Has(3)
	c.Unban = c.Flags.Has(4)
	c.Kick = c.Flags.Has(5)
	c.Unkick = c.Flags.Has(6)
	c.Promote = c.Flags.Has(7)
	c.Demote = c.Flags.Has(8)
	c.Info = c.Flags.Has(9)
	c.Settings = c.Flags.Has(10)
	c.Pinned = c.Flags.Has(11)
	c.Edit = c.Flags.Has(12)
	c.Delete = c.Flags.Has(13)
	c.GroupCall = c.Flags.Has(14)
	c.Invites = c.Flags.Has(15)
	c.Send = c.Flags.Has(16)
	c.Forums = c.Flags.Has(17)
	return nil
}

// SetJoin sets value of Join conditional field.
func (c *ChannelAdminLogEventsFilter) SetJoin(value bool) {
	if value {
		c.Flags.Set(0)
		c.Join = true
	} else {
		c.Flags.Unset(0)
		c.Join = false
	}
}

// GetJoin returns value of Join conditional field.
func (c *ChannelAdminLogEventsFilter) GetJoin() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(0)
}

// SetLeave sets value of Leave conditional field.
func (c *ChannelAdminLogEventsFilter) SetLeave(value bool) {
	if value {
		c.Flags.Set(1)
		c.Leave = true
	} else {
		c.Flags.Unset(1)
		c.Leave = false
	}
}

// GetLeave returns value of Leave conditional field.
func (c *ChannelAdminLogEventsFilter) GetLeave() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(1)
}

// SetInvite sets value of Invite conditional field.
func (c *ChannelAdminLogEventsFilter) SetInvite(value bool) {
	if value {
		c.Flags.Set(2)
		c.Invite = true
	} else {
		c.Flags.Unset(2)
		c.Invite = false
	}
}

// GetInvite returns value of Invite conditional field.
func (c *ChannelAdminLogEventsFilter) GetInvite() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(2)
}

// SetBan sets value of Ban conditional field.
func (c *ChannelAdminLogEventsFilter) SetBan(value bool) {
	if value {
		c.Flags.Set(3)
		c.Ban = true
	} else {
		c.Flags.Unset(3)
		c.Ban = false
	}
}

// GetBan returns value of Ban conditional field.
func (c *ChannelAdminLogEventsFilter) GetBan() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(3)
}

// SetUnban sets value of Unban conditional field.
func (c *ChannelAdminLogEventsFilter) SetUnban(value bool) {
	if value {
		c.Flags.Set(4)
		c.Unban = true
	} else {
		c.Flags.Unset(4)
		c.Unban = false
	}
}

// GetUnban returns value of Unban conditional field.
func (c *ChannelAdminLogEventsFilter) GetUnban() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(4)
}

// SetKick sets value of Kick conditional field.
func (c *ChannelAdminLogEventsFilter) SetKick(value bool) {
	if value {
		c.Flags.Set(5)
		c.Kick = true
	} else {
		c.Flags.Unset(5)
		c.Kick = false
	}
}

// GetKick returns value of Kick conditional field.
func (c *ChannelAdminLogEventsFilter) GetKick() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(5)
}

// SetUnkick sets value of Unkick conditional field.
func (c *ChannelAdminLogEventsFilter) SetUnkick(value bool) {
	if value {
		c.Flags.Set(6)
		c.Unkick = true
	} else {
		c.Flags.Unset(6)
		c.Unkick = false
	}
}

// GetUnkick returns value of Unkick conditional field.
func (c *ChannelAdminLogEventsFilter) GetUnkick() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(6)
}

// SetPromote sets value of Promote conditional field.
func (c *ChannelAdminLogEventsFilter) SetPromote(value bool) {
	if value {
		c.Flags.Set(7)
		c.Promote = true
	} else {
		c.Flags.Unset(7)
		c.Promote = false
	}
}

// GetPromote returns value of Promote conditional field.
func (c *ChannelAdminLogEventsFilter) GetPromote() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(7)
}

// SetDemote sets value of Demote conditional field.
func (c *ChannelAdminLogEventsFilter) SetDemote(value bool) {
	if value {
		c.Flags.Set(8)
		c.Demote = true
	} else {
		c.Flags.Unset(8)
		c.Demote = false
	}
}

// GetDemote returns value of Demote conditional field.
func (c *ChannelAdminLogEventsFilter) GetDemote() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(8)
}

// SetInfo sets value of Info conditional field.
func (c *ChannelAdminLogEventsFilter) SetInfo(value bool) {
	if value {
		c.Flags.Set(9)
		c.Info = true
	} else {
		c.Flags.Unset(9)
		c.Info = false
	}
}

// GetInfo returns value of Info conditional field.
func (c *ChannelAdminLogEventsFilter) GetInfo() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(9)
}

// SetSettings sets value of Settings conditional field.
func (c *ChannelAdminLogEventsFilter) SetSettings(value bool) {
	if value {
		c.Flags.Set(10)
		c.Settings = true
	} else {
		c.Flags.Unset(10)
		c.Settings = false
	}
}

// GetSettings returns value of Settings conditional field.
func (c *ChannelAdminLogEventsFilter) GetSettings() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(10)
}

// SetPinned sets value of Pinned conditional field.
func (c *ChannelAdminLogEventsFilter) SetPinned(value bool) {
	if value {
		c.Flags.Set(11)
		c.Pinned = true
	} else {
		c.Flags.Unset(11)
		c.Pinned = false
	}
}

// GetPinned returns value of Pinned conditional field.
func (c *ChannelAdminLogEventsFilter) GetPinned() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(11)
}

// SetEdit sets value of Edit conditional field.
func (c *ChannelAdminLogEventsFilter) SetEdit(value bool) {
	if value {
		c.Flags.Set(12)
		c.Edit = true
	} else {
		c.Flags.Unset(12)
		c.Edit = false
	}
}

// GetEdit returns value of Edit conditional field.
func (c *ChannelAdminLogEventsFilter) GetEdit() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(12)
}

// SetDelete sets value of Delete conditional field.
func (c *ChannelAdminLogEventsFilter) SetDelete(value bool) {
	if value {
		c.Flags.Set(13)
		c.Delete = true
	} else {
		c.Flags.Unset(13)
		c.Delete = false
	}
}

// GetDelete returns value of Delete conditional field.
func (c *ChannelAdminLogEventsFilter) GetDelete() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(13)
}

// SetGroupCall sets value of GroupCall conditional field.
func (c *ChannelAdminLogEventsFilter) SetGroupCall(value bool) {
	if value {
		c.Flags.Set(14)
		c.GroupCall = true
	} else {
		c.Flags.Unset(14)
		c.GroupCall = false
	}
}

// GetGroupCall returns value of GroupCall conditional field.
func (c *ChannelAdminLogEventsFilter) GetGroupCall() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(14)
}

// SetInvites sets value of Invites conditional field.
func (c *ChannelAdminLogEventsFilter) SetInvites(value bool) {
	if value {
		c.Flags.Set(15)
		c.Invites = true
	} else {
		c.Flags.Unset(15)
		c.Invites = false
	}
}

// GetInvites returns value of Invites conditional field.
func (c *ChannelAdminLogEventsFilter) GetInvites() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(15)
}

// SetSend sets value of Send conditional field.
func (c *ChannelAdminLogEventsFilter) SetSend(value bool) {
	if value {
		c.Flags.Set(16)
		c.Send = true
	} else {
		c.Flags.Unset(16)
		c.Send = false
	}
}

// GetSend returns value of Send conditional field.
func (c *ChannelAdminLogEventsFilter) GetSend() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(16)
}

// SetForums sets value of Forums conditional field.
func (c *ChannelAdminLogEventsFilter) SetForums(value bool) {
	if value {
		c.Flags.Set(17)
		c.Forums = true
	} else {
		c.Flags.Unset(17)
		c.Forums = false
	}
}

// GetForums returns value of Forums conditional field.
func (c *ChannelAdminLogEventsFilter) GetForums() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(17)
}
