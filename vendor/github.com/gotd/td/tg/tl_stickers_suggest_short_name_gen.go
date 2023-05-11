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

// StickersSuggestShortNameRequest represents TL type `stickers.suggestShortName#4dafc503`.
// Suggests a short name for a given stickerpack name
//
// See https://core.telegram.org/method/stickers.suggestShortName for reference.
type StickersSuggestShortNameRequest struct {
	// Sticker pack name
	Title string
}

// StickersSuggestShortNameRequestTypeID is TL type id of StickersSuggestShortNameRequest.
const StickersSuggestShortNameRequestTypeID = 0x4dafc503

// Ensuring interfaces in compile-time for StickersSuggestShortNameRequest.
var (
	_ bin.Encoder     = &StickersSuggestShortNameRequest{}
	_ bin.Decoder     = &StickersSuggestShortNameRequest{}
	_ bin.BareEncoder = &StickersSuggestShortNameRequest{}
	_ bin.BareDecoder = &StickersSuggestShortNameRequest{}
)

func (s *StickersSuggestShortNameRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StickersSuggestShortNameRequest) String() string {
	if s == nil {
		return "StickersSuggestShortNameRequest(nil)"
	}
	type Alias StickersSuggestShortNameRequest
	return fmt.Sprintf("StickersSuggestShortNameRequest%+v", Alias(*s))
}

// FillFrom fills StickersSuggestShortNameRequest from given interface.
func (s *StickersSuggestShortNameRequest) FillFrom(from interface {
	GetTitle() (value string)
}) {
	s.Title = from.GetTitle()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickersSuggestShortNameRequest) TypeID() uint32 {
	return StickersSuggestShortNameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StickersSuggestShortNameRequest) TypeName() string {
	return "stickers.suggestShortName"
}

// TypeInfo returns info about TL type.
func (s *StickersSuggestShortNameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickers.suggestShortName",
		ID:   StickersSuggestShortNameRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StickersSuggestShortNameRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickers.suggestShortName#4dafc503 as nil")
	}
	b.PutID(StickersSuggestShortNameRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StickersSuggestShortNameRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickers.suggestShortName#4dafc503 as nil")
	}
	b.PutString(s.Title)
	return nil
}

// Decode implements bin.Decoder.
func (s *StickersSuggestShortNameRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickers.suggestShortName#4dafc503 to nil")
	}
	if err := b.ConsumeID(StickersSuggestShortNameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.suggestShortName#4dafc503: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StickersSuggestShortNameRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickers.suggestShortName#4dafc503 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickers.suggestShortName#4dafc503: field title: %w", err)
		}
		s.Title = value
	}
	return nil
}

// GetTitle returns value of Title field.
func (s *StickersSuggestShortNameRequest) GetTitle() (value string) {
	if s == nil {
		return
	}
	return s.Title
}

// StickersSuggestShortName invokes method stickers.suggestShortName#4dafc503 returning error if any.
// Suggests a short name for a given stickerpack name
//
// Possible errors:
//
//	400 TITLE_INVALID: The specified stickerpack title is invalid.
//
// See https://core.telegram.org/method/stickers.suggestShortName for reference.
func (c *Client) StickersSuggestShortName(ctx context.Context, title string) (*StickersSuggestedShortName, error) {
	var result StickersSuggestedShortName

	request := &StickersSuggestShortNameRequest{
		Title: title,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
