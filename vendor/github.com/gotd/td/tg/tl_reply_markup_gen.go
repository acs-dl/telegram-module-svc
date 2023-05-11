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

// ReplyKeyboardHide represents TL type `replyKeyboardHide#a03e5b85`.
// Hide sent bot keyboard
//
// See https://core.telegram.org/constructor/replyKeyboardHide for reference.
type ReplyKeyboardHide struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Use this flag if you want to remove the keyboard for specific users only. Targets: 1)
	// users that are @mentioned in the text of the Message object; 2) if the bot's message
	// is a reply (has reply_to_message_id), sender of the original message.Example: A user
	// votes in a poll, bot returns confirmation message in reply to the vote and removes the
	// keyboard for that user, while still showing the keyboard with poll options to users
	// who haven't voted yet
	Selective bool
}

// ReplyKeyboardHideTypeID is TL type id of ReplyKeyboardHide.
const ReplyKeyboardHideTypeID = 0xa03e5b85

// construct implements constructor of ReplyMarkupClass.
func (r ReplyKeyboardHide) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyKeyboardHide.
var (
	_ bin.Encoder     = &ReplyKeyboardHide{}
	_ bin.Decoder     = &ReplyKeyboardHide{}
	_ bin.BareEncoder = &ReplyKeyboardHide{}
	_ bin.BareDecoder = &ReplyKeyboardHide{}

	_ ReplyMarkupClass = &ReplyKeyboardHide{}
)

func (r *ReplyKeyboardHide) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Selective == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyKeyboardHide) String() string {
	if r == nil {
		return "ReplyKeyboardHide(nil)"
	}
	type Alias ReplyKeyboardHide
	return fmt.Sprintf("ReplyKeyboardHide%+v", Alias(*r))
}

// FillFrom fills ReplyKeyboardHide from given interface.
func (r *ReplyKeyboardHide) FillFrom(from interface {
	GetSelective() (value bool)
}) {
	r.Selective = from.GetSelective()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplyKeyboardHide) TypeID() uint32 {
	return ReplyKeyboardHideTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplyKeyboardHide) TypeName() string {
	return "replyKeyboardHide"
}

// TypeInfo returns info about TL type.
func (r *ReplyKeyboardHide) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replyKeyboardHide",
		ID:   ReplyKeyboardHideTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Selective",
			SchemaName: "selective",
			Null:       !r.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *ReplyKeyboardHide) SetFlags() {
	if !(r.Selective == false) {
		r.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (r *ReplyKeyboardHide) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyKeyboardHide#a03e5b85 as nil")
	}
	b.PutID(ReplyKeyboardHideTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReplyKeyboardHide) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyKeyboardHide#a03e5b85 as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode replyKeyboardHide#a03e5b85: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReplyKeyboardHide) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyKeyboardHide#a03e5b85 to nil")
	}
	if err := b.ConsumeID(ReplyKeyboardHideTypeID); err != nil {
		return fmt.Errorf("unable to decode replyKeyboardHide#a03e5b85: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReplyKeyboardHide) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyKeyboardHide#a03e5b85 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode replyKeyboardHide#a03e5b85: field flags: %w", err)
		}
	}
	r.Selective = r.Flags.Has(2)
	return nil
}

// SetSelective sets value of Selective conditional field.
func (r *ReplyKeyboardHide) SetSelective(value bool) {
	if value {
		r.Flags.Set(2)
		r.Selective = true
	} else {
		r.Flags.Unset(2)
		r.Selective = false
	}
}

// GetSelective returns value of Selective conditional field.
func (r *ReplyKeyboardHide) GetSelective() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(2)
}

// ReplyKeyboardForceReply represents TL type `replyKeyboardForceReply#86b40b08`.
// Force the user to send a reply
//
// See https://core.telegram.org/constructor/replyKeyboardForceReply for reference.
type ReplyKeyboardForceReply struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Requests clients to hide the keyboard as soon as it's been used. The keyboard will
	// still be available, but clients will automatically display the usual letter-keyboard
	// in the chat – the user can press a special button in the input field to see the
	// custom keyboard again.
	SingleUse bool
	// Use this parameter if you want to show the keyboard to specific users only. Targets:
	// 1) users that are @mentioned in the text of the Message object; 2) if the bot's
	// message is a reply (has reply_to_message_id), sender of the original message. Example:
	// A user requests to change the bot's language, bot replies to the request with a
	// keyboard to select the new language. Other users in the group don't see the keyboard.
	Selective bool
	// The placeholder to be shown in the input field when the keyboard is active; 1-64
	// characters.
	//
	// Use SetPlaceholder and GetPlaceholder helpers.
	Placeholder string
}

// ReplyKeyboardForceReplyTypeID is TL type id of ReplyKeyboardForceReply.
const ReplyKeyboardForceReplyTypeID = 0x86b40b08

// construct implements constructor of ReplyMarkupClass.
func (r ReplyKeyboardForceReply) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyKeyboardForceReply.
var (
	_ bin.Encoder     = &ReplyKeyboardForceReply{}
	_ bin.Decoder     = &ReplyKeyboardForceReply{}
	_ bin.BareEncoder = &ReplyKeyboardForceReply{}
	_ bin.BareDecoder = &ReplyKeyboardForceReply{}

	_ ReplyMarkupClass = &ReplyKeyboardForceReply{}
)

func (r *ReplyKeyboardForceReply) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.SingleUse == false) {
		return false
	}
	if !(r.Selective == false) {
		return false
	}
	if !(r.Placeholder == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyKeyboardForceReply) String() string {
	if r == nil {
		return "ReplyKeyboardForceReply(nil)"
	}
	type Alias ReplyKeyboardForceReply
	return fmt.Sprintf("ReplyKeyboardForceReply%+v", Alias(*r))
}

// FillFrom fills ReplyKeyboardForceReply from given interface.
func (r *ReplyKeyboardForceReply) FillFrom(from interface {
	GetSingleUse() (value bool)
	GetSelective() (value bool)
	GetPlaceholder() (value string, ok bool)
}) {
	r.SingleUse = from.GetSingleUse()
	r.Selective = from.GetSelective()
	if val, ok := from.GetPlaceholder(); ok {
		r.Placeholder = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplyKeyboardForceReply) TypeID() uint32 {
	return ReplyKeyboardForceReplyTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplyKeyboardForceReply) TypeName() string {
	return "replyKeyboardForceReply"
}

// TypeInfo returns info about TL type.
func (r *ReplyKeyboardForceReply) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replyKeyboardForceReply",
		ID:   ReplyKeyboardForceReplyTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SingleUse",
			SchemaName: "single_use",
			Null:       !r.Flags.Has(1),
		},
		{
			Name:       "Selective",
			SchemaName: "selective",
			Null:       !r.Flags.Has(2),
		},
		{
			Name:       "Placeholder",
			SchemaName: "placeholder",
			Null:       !r.Flags.Has(3),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *ReplyKeyboardForceReply) SetFlags() {
	if !(r.SingleUse == false) {
		r.Flags.Set(1)
	}
	if !(r.Selective == false) {
		r.Flags.Set(2)
	}
	if !(r.Placeholder == "") {
		r.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (r *ReplyKeyboardForceReply) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyKeyboardForceReply#86b40b08 as nil")
	}
	b.PutID(ReplyKeyboardForceReplyTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReplyKeyboardForceReply) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyKeyboardForceReply#86b40b08 as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode replyKeyboardForceReply#86b40b08: field flags: %w", err)
	}
	if r.Flags.Has(3) {
		b.PutString(r.Placeholder)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReplyKeyboardForceReply) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyKeyboardForceReply#86b40b08 to nil")
	}
	if err := b.ConsumeID(ReplyKeyboardForceReplyTypeID); err != nil {
		return fmt.Errorf("unable to decode replyKeyboardForceReply#86b40b08: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReplyKeyboardForceReply) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyKeyboardForceReply#86b40b08 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode replyKeyboardForceReply#86b40b08: field flags: %w", err)
		}
	}
	r.SingleUse = r.Flags.Has(1)
	r.Selective = r.Flags.Has(2)
	if r.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode replyKeyboardForceReply#86b40b08: field placeholder: %w", err)
		}
		r.Placeholder = value
	}
	return nil
}

// SetSingleUse sets value of SingleUse conditional field.
func (r *ReplyKeyboardForceReply) SetSingleUse(value bool) {
	if value {
		r.Flags.Set(1)
		r.SingleUse = true
	} else {
		r.Flags.Unset(1)
		r.SingleUse = false
	}
}

// GetSingleUse returns value of SingleUse conditional field.
func (r *ReplyKeyboardForceReply) GetSingleUse() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(1)
}

// SetSelective sets value of Selective conditional field.
func (r *ReplyKeyboardForceReply) SetSelective(value bool) {
	if value {
		r.Flags.Set(2)
		r.Selective = true
	} else {
		r.Flags.Unset(2)
		r.Selective = false
	}
}

// GetSelective returns value of Selective conditional field.
func (r *ReplyKeyboardForceReply) GetSelective() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(2)
}

// SetPlaceholder sets value of Placeholder conditional field.
func (r *ReplyKeyboardForceReply) SetPlaceholder(value string) {
	r.Flags.Set(3)
	r.Placeholder = value
}

// GetPlaceholder returns value of Placeholder conditional field and
// boolean which is true if field was set.
func (r *ReplyKeyboardForceReply) GetPlaceholder() (value string, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(3) {
		return value, false
	}
	return r.Placeholder, true
}

// ReplyKeyboardMarkup represents TL type `replyKeyboardMarkup#85dd99d1`.
// Bot keyboard
//
// See https://core.telegram.org/constructor/replyKeyboardMarkup for reference.
type ReplyKeyboardMarkup struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Requests clients to resize the keyboard vertically for optimal fit (e.g., make the
	// keyboard smaller if there are just two rows of buttons). If not set, the custom
	// keyboard is always of the same height as the app's standard keyboard.
	Resize bool
	// Requests clients to hide the keyboard as soon as it's been used. The keyboard will
	// still be available, but clients will automatically display the usual letter-keyboard
	// in the chat – the user can press a special button in the input field to see the
	// custom keyboard again.
	SingleUse bool
	// Use this parameter if you want to show the keyboard to specific users only. Targets:
	// 1) users that are @mentioned in the text of the Message object; 2) if the bot's
	// message is a reply (has reply_to_message_id), sender of the original message.Example:
	// A user requests to change the bot's language, bot replies to the request with a
	// keyboard to select the new language. Other users in the group don't see the keyboard.
	Selective bool
	// Persistent field of ReplyKeyboardMarkup.
	Persistent bool
	// Button row
	Rows []KeyboardButtonRow
	// The placeholder to be shown in the input field when the keyboard is active; 1-64
	// characters.
	//
	// Use SetPlaceholder and GetPlaceholder helpers.
	Placeholder string
}

// ReplyKeyboardMarkupTypeID is TL type id of ReplyKeyboardMarkup.
const ReplyKeyboardMarkupTypeID = 0x85dd99d1

// construct implements constructor of ReplyMarkupClass.
func (r ReplyKeyboardMarkup) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyKeyboardMarkup.
var (
	_ bin.Encoder     = &ReplyKeyboardMarkup{}
	_ bin.Decoder     = &ReplyKeyboardMarkup{}
	_ bin.BareEncoder = &ReplyKeyboardMarkup{}
	_ bin.BareDecoder = &ReplyKeyboardMarkup{}

	_ ReplyMarkupClass = &ReplyKeyboardMarkup{}
)

func (r *ReplyKeyboardMarkup) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Resize == false) {
		return false
	}
	if !(r.SingleUse == false) {
		return false
	}
	if !(r.Selective == false) {
		return false
	}
	if !(r.Persistent == false) {
		return false
	}
	if !(r.Rows == nil) {
		return false
	}
	if !(r.Placeholder == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyKeyboardMarkup) String() string {
	if r == nil {
		return "ReplyKeyboardMarkup(nil)"
	}
	type Alias ReplyKeyboardMarkup
	return fmt.Sprintf("ReplyKeyboardMarkup%+v", Alias(*r))
}

// FillFrom fills ReplyKeyboardMarkup from given interface.
func (r *ReplyKeyboardMarkup) FillFrom(from interface {
	GetResize() (value bool)
	GetSingleUse() (value bool)
	GetSelective() (value bool)
	GetPersistent() (value bool)
	GetRows() (value []KeyboardButtonRow)
	GetPlaceholder() (value string, ok bool)
}) {
	r.Resize = from.GetResize()
	r.SingleUse = from.GetSingleUse()
	r.Selective = from.GetSelective()
	r.Persistent = from.GetPersistent()
	r.Rows = from.GetRows()
	if val, ok := from.GetPlaceholder(); ok {
		r.Placeholder = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplyKeyboardMarkup) TypeID() uint32 {
	return ReplyKeyboardMarkupTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplyKeyboardMarkup) TypeName() string {
	return "replyKeyboardMarkup"
}

// TypeInfo returns info about TL type.
func (r *ReplyKeyboardMarkup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replyKeyboardMarkup",
		ID:   ReplyKeyboardMarkupTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Resize",
			SchemaName: "resize",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "SingleUse",
			SchemaName: "single_use",
			Null:       !r.Flags.Has(1),
		},
		{
			Name:       "Selective",
			SchemaName: "selective",
			Null:       !r.Flags.Has(2),
		},
		{
			Name:       "Persistent",
			SchemaName: "persistent",
			Null:       !r.Flags.Has(4),
		},
		{
			Name:       "Rows",
			SchemaName: "rows",
		},
		{
			Name:       "Placeholder",
			SchemaName: "placeholder",
			Null:       !r.Flags.Has(3),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *ReplyKeyboardMarkup) SetFlags() {
	if !(r.Resize == false) {
		r.Flags.Set(0)
	}
	if !(r.SingleUse == false) {
		r.Flags.Set(1)
	}
	if !(r.Selective == false) {
		r.Flags.Set(2)
	}
	if !(r.Persistent == false) {
		r.Flags.Set(4)
	}
	if !(r.Placeholder == "") {
		r.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (r *ReplyKeyboardMarkup) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyKeyboardMarkup#85dd99d1 as nil")
	}
	b.PutID(ReplyKeyboardMarkupTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReplyKeyboardMarkup) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyKeyboardMarkup#85dd99d1 as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode replyKeyboardMarkup#85dd99d1: field flags: %w", err)
	}
	b.PutVectorHeader(len(r.Rows))
	for idx, v := range r.Rows {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode replyKeyboardMarkup#85dd99d1: field rows element with index %d: %w", idx, err)
		}
	}
	if r.Flags.Has(3) {
		b.PutString(r.Placeholder)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReplyKeyboardMarkup) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyKeyboardMarkup#85dd99d1 to nil")
	}
	if err := b.ConsumeID(ReplyKeyboardMarkupTypeID); err != nil {
		return fmt.Errorf("unable to decode replyKeyboardMarkup#85dd99d1: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReplyKeyboardMarkup) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyKeyboardMarkup#85dd99d1 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode replyKeyboardMarkup#85dd99d1: field flags: %w", err)
		}
	}
	r.Resize = r.Flags.Has(0)
	r.SingleUse = r.Flags.Has(1)
	r.Selective = r.Flags.Has(2)
	r.Persistent = r.Flags.Has(4)
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode replyKeyboardMarkup#85dd99d1: field rows: %w", err)
		}

		if headerLen > 0 {
			r.Rows = make([]KeyboardButtonRow, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value KeyboardButtonRow
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode replyKeyboardMarkup#85dd99d1: field rows: %w", err)
			}
			r.Rows = append(r.Rows, value)
		}
	}
	if r.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode replyKeyboardMarkup#85dd99d1: field placeholder: %w", err)
		}
		r.Placeholder = value
	}
	return nil
}

// SetResize sets value of Resize conditional field.
func (r *ReplyKeyboardMarkup) SetResize(value bool) {
	if value {
		r.Flags.Set(0)
		r.Resize = true
	} else {
		r.Flags.Unset(0)
		r.Resize = false
	}
}

// GetResize returns value of Resize conditional field.
func (r *ReplyKeyboardMarkup) GetResize() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(0)
}

// SetSingleUse sets value of SingleUse conditional field.
func (r *ReplyKeyboardMarkup) SetSingleUse(value bool) {
	if value {
		r.Flags.Set(1)
		r.SingleUse = true
	} else {
		r.Flags.Unset(1)
		r.SingleUse = false
	}
}

// GetSingleUse returns value of SingleUse conditional field.
func (r *ReplyKeyboardMarkup) GetSingleUse() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(1)
}

// SetSelective sets value of Selective conditional field.
func (r *ReplyKeyboardMarkup) SetSelective(value bool) {
	if value {
		r.Flags.Set(2)
		r.Selective = true
	} else {
		r.Flags.Unset(2)
		r.Selective = false
	}
}

// GetSelective returns value of Selective conditional field.
func (r *ReplyKeyboardMarkup) GetSelective() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(2)
}

// SetPersistent sets value of Persistent conditional field.
func (r *ReplyKeyboardMarkup) SetPersistent(value bool) {
	if value {
		r.Flags.Set(4)
		r.Persistent = true
	} else {
		r.Flags.Unset(4)
		r.Persistent = false
	}
}

// GetPersistent returns value of Persistent conditional field.
func (r *ReplyKeyboardMarkup) GetPersistent() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(4)
}

// GetRows returns value of Rows field.
func (r *ReplyKeyboardMarkup) GetRows() (value []KeyboardButtonRow) {
	if r == nil {
		return
	}
	return r.Rows
}

// SetPlaceholder sets value of Placeholder conditional field.
func (r *ReplyKeyboardMarkup) SetPlaceholder(value string) {
	r.Flags.Set(3)
	r.Placeholder = value
}

// GetPlaceholder returns value of Placeholder conditional field and
// boolean which is true if field was set.
func (r *ReplyKeyboardMarkup) GetPlaceholder() (value string, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(3) {
		return value, false
	}
	return r.Placeholder, true
}

// ReplyInlineMarkup represents TL type `replyInlineMarkup#48a30254`.
// Bot or inline keyboard
//
// See https://core.telegram.org/constructor/replyInlineMarkup for reference.
type ReplyInlineMarkup struct {
	// Bot or inline keyboard rows
	Rows []KeyboardButtonRow
}

// ReplyInlineMarkupTypeID is TL type id of ReplyInlineMarkup.
const ReplyInlineMarkupTypeID = 0x48a30254

// construct implements constructor of ReplyMarkupClass.
func (r ReplyInlineMarkup) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyInlineMarkup.
var (
	_ bin.Encoder     = &ReplyInlineMarkup{}
	_ bin.Decoder     = &ReplyInlineMarkup{}
	_ bin.BareEncoder = &ReplyInlineMarkup{}
	_ bin.BareDecoder = &ReplyInlineMarkup{}

	_ ReplyMarkupClass = &ReplyInlineMarkup{}
)

func (r *ReplyInlineMarkup) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Rows == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyInlineMarkup) String() string {
	if r == nil {
		return "ReplyInlineMarkup(nil)"
	}
	type Alias ReplyInlineMarkup
	return fmt.Sprintf("ReplyInlineMarkup%+v", Alias(*r))
}

// FillFrom fills ReplyInlineMarkup from given interface.
func (r *ReplyInlineMarkup) FillFrom(from interface {
	GetRows() (value []KeyboardButtonRow)
}) {
	r.Rows = from.GetRows()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplyInlineMarkup) TypeID() uint32 {
	return ReplyInlineMarkupTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplyInlineMarkup) TypeName() string {
	return "replyInlineMarkup"
}

// TypeInfo returns info about TL type.
func (r *ReplyInlineMarkup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replyInlineMarkup",
		ID:   ReplyInlineMarkupTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Rows",
			SchemaName: "rows",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReplyInlineMarkup) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyInlineMarkup#48a30254 as nil")
	}
	b.PutID(ReplyInlineMarkupTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReplyInlineMarkup) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyInlineMarkup#48a30254 as nil")
	}
	b.PutVectorHeader(len(r.Rows))
	for idx, v := range r.Rows {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode replyInlineMarkup#48a30254: field rows element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReplyInlineMarkup) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyInlineMarkup#48a30254 to nil")
	}
	if err := b.ConsumeID(ReplyInlineMarkupTypeID); err != nil {
		return fmt.Errorf("unable to decode replyInlineMarkup#48a30254: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReplyInlineMarkup) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyInlineMarkup#48a30254 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode replyInlineMarkup#48a30254: field rows: %w", err)
		}

		if headerLen > 0 {
			r.Rows = make([]KeyboardButtonRow, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value KeyboardButtonRow
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode replyInlineMarkup#48a30254: field rows: %w", err)
			}
			r.Rows = append(r.Rows, value)
		}
	}
	return nil
}

// GetRows returns value of Rows field.
func (r *ReplyInlineMarkup) GetRows() (value []KeyboardButtonRow) {
	if r == nil {
		return
	}
	return r.Rows
}

// ReplyMarkupClassName is schema name of ReplyMarkupClass.
const ReplyMarkupClassName = "ReplyMarkup"

// ReplyMarkupClass represents ReplyMarkup generic type.
//
// See https://core.telegram.org/type/ReplyMarkup for reference.
//
// Example:
//
//	g, err := tg.DecodeReplyMarkup(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.ReplyKeyboardHide: // replyKeyboardHide#a03e5b85
//	case *tg.ReplyKeyboardForceReply: // replyKeyboardForceReply#86b40b08
//	case *tg.ReplyKeyboardMarkup: // replyKeyboardMarkup#85dd99d1
//	case *tg.ReplyInlineMarkup: // replyInlineMarkup#48a30254
//	default: panic(v)
//	}
type ReplyMarkupClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ReplyMarkupClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeReplyMarkup implements binary de-serialization for ReplyMarkupClass.
func DecodeReplyMarkup(buf *bin.Buffer) (ReplyMarkupClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ReplyKeyboardHideTypeID:
		// Decoding replyKeyboardHide#a03e5b85.
		v := ReplyKeyboardHide{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case ReplyKeyboardForceReplyTypeID:
		// Decoding replyKeyboardForceReply#86b40b08.
		v := ReplyKeyboardForceReply{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case ReplyKeyboardMarkupTypeID:
		// Decoding replyKeyboardMarkup#85dd99d1.
		v := ReplyKeyboardMarkup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case ReplyInlineMarkupTypeID:
		// Decoding replyInlineMarkup#48a30254.
		v := ReplyInlineMarkup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", bin.NewUnexpectedID(id))
	}
}

// ReplyMarkup boxes the ReplyMarkupClass providing a helper.
type ReplyMarkupBox struct {
	ReplyMarkup ReplyMarkupClass
}

// Decode implements bin.Decoder for ReplyMarkupBox.
func (b *ReplyMarkupBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ReplyMarkupBox to nil")
	}
	v, err := DecodeReplyMarkup(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ReplyMarkup = v
	return nil
}

// Encode implements bin.Encode for ReplyMarkupBox.
func (b *ReplyMarkupBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode ReplyMarkupClass as nil")
	}
	return b.ReplyMarkup.Encode(buf)
}