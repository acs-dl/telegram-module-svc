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

// BotsSetBotInfoRequest represents TL type `bots.setBotInfo#10cf3123`.
// Set our about text and description (bots only)
//
// See https://core.telegram.org/method/bots.setBotInfo for reference.
type BotsSetBotInfoRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Bot field of BotsSetBotInfoRequest.
	//
	// Use SetBot and GetBot helpers.
	Bot InputUserClass
	// Language code, if left empty update the fallback about text and description
	LangCode string
	// Name field of BotsSetBotInfoRequest.
	//
	// Use SetName and GetName helpers.
	Name string
	// New about text
	//
	// Use SetAbout and GetAbout helpers.
	About string
	// New description
	//
	// Use SetDescription and GetDescription helpers.
	Description string
}

// BotsSetBotInfoRequestTypeID is TL type id of BotsSetBotInfoRequest.
const BotsSetBotInfoRequestTypeID = 0x10cf3123

// Ensuring interfaces in compile-time for BotsSetBotInfoRequest.
var (
	_ bin.Encoder     = &BotsSetBotInfoRequest{}
	_ bin.Decoder     = &BotsSetBotInfoRequest{}
	_ bin.BareEncoder = &BotsSetBotInfoRequest{}
	_ bin.BareDecoder = &BotsSetBotInfoRequest{}
)

func (s *BotsSetBotInfoRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Bot == nil) {
		return false
	}
	if !(s.LangCode == "") {
		return false
	}
	if !(s.Name == "") {
		return false
	}
	if !(s.About == "") {
		return false
	}
	if !(s.Description == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *BotsSetBotInfoRequest) String() string {
	if s == nil {
		return "BotsSetBotInfoRequest(nil)"
	}
	type Alias BotsSetBotInfoRequest
	return fmt.Sprintf("BotsSetBotInfoRequest%+v", Alias(*s))
}

// FillFrom fills BotsSetBotInfoRequest from given interface.
func (s *BotsSetBotInfoRequest) FillFrom(from interface {
	GetBot() (value InputUserClass, ok bool)
	GetLangCode() (value string)
	GetName() (value string, ok bool)
	GetAbout() (value string, ok bool)
	GetDescription() (value string, ok bool)
}) {
	if val, ok := from.GetBot(); ok {
		s.Bot = val
	}

	s.LangCode = from.GetLangCode()
	if val, ok := from.GetName(); ok {
		s.Name = val
	}

	if val, ok := from.GetAbout(); ok {
		s.About = val
	}

	if val, ok := from.GetDescription(); ok {
		s.Description = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsSetBotInfoRequest) TypeID() uint32 {
	return BotsSetBotInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsSetBotInfoRequest) TypeName() string {
	return "bots.setBotInfo"
}

// TypeInfo returns info about TL type.
func (s *BotsSetBotInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.setBotInfo",
		ID:   BotsSetBotInfoRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
		{
			Name:       "Name",
			SchemaName: "name",
			Null:       !s.Flags.Has(3),
		},
		{
			Name:       "About",
			SchemaName: "about",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Description",
			SchemaName: "description",
			Null:       !s.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *BotsSetBotInfoRequest) SetFlags() {
	if !(s.Bot == nil) {
		s.Flags.Set(2)
	}
	if !(s.Name == "") {
		s.Flags.Set(3)
	}
	if !(s.About == "") {
		s.Flags.Set(0)
	}
	if !(s.Description == "") {
		s.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (s *BotsSetBotInfoRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode bots.setBotInfo#10cf3123 as nil")
	}
	b.PutID(BotsSetBotInfoRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *BotsSetBotInfoRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode bots.setBotInfo#10cf3123 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.setBotInfo#10cf3123: field flags: %w", err)
	}
	if s.Flags.Has(2) {
		if s.Bot == nil {
			return fmt.Errorf("unable to encode bots.setBotInfo#10cf3123: field bot is nil")
		}
		if err := s.Bot.Encode(b); err != nil {
			return fmt.Errorf("unable to encode bots.setBotInfo#10cf3123: field bot: %w", err)
		}
	}
	b.PutString(s.LangCode)
	if s.Flags.Has(3) {
		b.PutString(s.Name)
	}
	if s.Flags.Has(0) {
		b.PutString(s.About)
	}
	if s.Flags.Has(1) {
		b.PutString(s.Description)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *BotsSetBotInfoRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode bots.setBotInfo#10cf3123 to nil")
	}
	if err := b.ConsumeID(BotsSetBotInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.setBotInfo#10cf3123: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *BotsSetBotInfoRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode bots.setBotInfo#10cf3123 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode bots.setBotInfo#10cf3123: field flags: %w", err)
		}
	}
	if s.Flags.Has(2) {
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.setBotInfo#10cf3123: field bot: %w", err)
		}
		s.Bot = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.setBotInfo#10cf3123: field lang_code: %w", err)
		}
		s.LangCode = value
	}
	if s.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.setBotInfo#10cf3123: field name: %w", err)
		}
		s.Name = value
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.setBotInfo#10cf3123: field about: %w", err)
		}
		s.About = value
	}
	if s.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.setBotInfo#10cf3123: field description: %w", err)
		}
		s.Description = value
	}
	return nil
}

// SetBot sets value of Bot conditional field.
func (s *BotsSetBotInfoRequest) SetBot(value InputUserClass) {
	s.Flags.Set(2)
	s.Bot = value
}

// GetBot returns value of Bot conditional field and
// boolean which is true if field was set.
func (s *BotsSetBotInfoRequest) GetBot() (value InputUserClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.Bot, true
}

// GetLangCode returns value of LangCode field.
func (s *BotsSetBotInfoRequest) GetLangCode() (value string) {
	if s == nil {
		return
	}
	return s.LangCode
}

// SetName sets value of Name conditional field.
func (s *BotsSetBotInfoRequest) SetName(value string) {
	s.Flags.Set(3)
	s.Name = value
}

// GetName returns value of Name conditional field and
// boolean which is true if field was set.
func (s *BotsSetBotInfoRequest) GetName() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.Name, true
}

// SetAbout sets value of About conditional field.
func (s *BotsSetBotInfoRequest) SetAbout(value string) {
	s.Flags.Set(0)
	s.About = value
}

// GetAbout returns value of About conditional field and
// boolean which is true if field was set.
func (s *BotsSetBotInfoRequest) GetAbout() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.About, true
}

// SetDescription sets value of Description conditional field.
func (s *BotsSetBotInfoRequest) SetDescription(value string) {
	s.Flags.Set(1)
	s.Description = value
}

// GetDescription returns value of Description conditional field and
// boolean which is true if field was set.
func (s *BotsSetBotInfoRequest) GetDescription() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.Description, true
}

// BotsSetBotInfo invokes method bots.setBotInfo#10cf3123 returning error if any.
// Set our about text and description (bots only)
//
// Possible errors:
//
//	400 USER_BOT_REQUIRED: This method can only be called by a bot.
//
// See https://core.telegram.org/method/bots.setBotInfo for reference.
// Can be used by bots.
func (c *Client) BotsSetBotInfo(ctx context.Context, request *BotsSetBotInfoRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
