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

// MessagesPeerSettings represents TL type `messages.peerSettings#6880b94d`.
// Peer settings
//
// See https://core.telegram.org/constructor/messages.peerSettings for reference.
type MessagesPeerSettings struct {
	// Peer settings
	Settings PeerSettings
	// Mentioned chats
	Chats []ChatClass
	// Mentioned users
	Users []UserClass
}

// MessagesPeerSettingsTypeID is TL type id of MessagesPeerSettings.
const MessagesPeerSettingsTypeID = 0x6880b94d

// Ensuring interfaces in compile-time for MessagesPeerSettings.
var (
	_ bin.Encoder     = &MessagesPeerSettings{}
	_ bin.Decoder     = &MessagesPeerSettings{}
	_ bin.BareEncoder = &MessagesPeerSettings{}
	_ bin.BareDecoder = &MessagesPeerSettings{}
)

func (p *MessagesPeerSettings) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Settings.Zero()) {
		return false
	}
	if !(p.Chats == nil) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *MessagesPeerSettings) String() string {
	if p == nil {
		return "MessagesPeerSettings(nil)"
	}
	type Alias MessagesPeerSettings
	return fmt.Sprintf("MessagesPeerSettings%+v", Alias(*p))
}

// FillFrom fills MessagesPeerSettings from given interface.
func (p *MessagesPeerSettings) FillFrom(from interface {
	GetSettings() (value PeerSettings)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	p.Settings = from.GetSettings()
	p.Chats = from.GetChats()
	p.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesPeerSettings) TypeID() uint32 {
	return MessagesPeerSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesPeerSettings) TypeName() string {
	return "messages.peerSettings"
}

// TypeInfo returns info about TL type.
func (p *MessagesPeerSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.peerSettings",
		ID:   MessagesPeerSettingsTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *MessagesPeerSettings) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode messages.peerSettings#6880b94d as nil")
	}
	b.PutID(MessagesPeerSettingsTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *MessagesPeerSettings) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode messages.peerSettings#6880b94d as nil")
	}
	if err := p.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.peerSettings#6880b94d: field settings: %w", err)
	}
	b.PutVectorHeader(len(p.Chats))
	for idx, v := range p.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.peerSettings#6880b94d: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.peerSettings#6880b94d: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.peerSettings#6880b94d: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.peerSettings#6880b94d: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *MessagesPeerSettings) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode messages.peerSettings#6880b94d to nil")
	}
	if err := b.ConsumeID(MessagesPeerSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.peerSettings#6880b94d: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *MessagesPeerSettings) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode messages.peerSettings#6880b94d to nil")
	}
	{
		if err := p.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.peerSettings#6880b94d: field settings: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.peerSettings#6880b94d: field chats: %w", err)
		}

		if headerLen > 0 {
			p.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.peerSettings#6880b94d: field chats: %w", err)
			}
			p.Chats = append(p.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.peerSettings#6880b94d: field users: %w", err)
		}

		if headerLen > 0 {
			p.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.peerSettings#6880b94d: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// GetSettings returns value of Settings field.
func (p *MessagesPeerSettings) GetSettings() (value PeerSettings) {
	if p == nil {
		return
	}
	return p.Settings
}

// GetChats returns value of Chats field.
func (p *MessagesPeerSettings) GetChats() (value []ChatClass) {
	if p == nil {
		return
	}
	return p.Chats
}

// GetUsers returns value of Users field.
func (p *MessagesPeerSettings) GetUsers() (value []UserClass) {
	if p == nil {
		return
	}
	return p.Users
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (p *MessagesPeerSettings) MapChats() (value ChatClassArray) {
	return ChatClassArray(p.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (p *MessagesPeerSettings) MapUsers() (value UserClassArray) {
	return UserClassArray(p.Users)
}
