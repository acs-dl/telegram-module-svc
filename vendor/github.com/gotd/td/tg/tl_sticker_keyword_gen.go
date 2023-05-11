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

// StickerKeyword represents TL type `stickerKeyword#fcfeb29c`.
// Keywords for a certain sticker
//
// See https://core.telegram.org/constructor/stickerKeyword for reference.
type StickerKeyword struct {
	// Sticker ID
	DocumentID int64
	// Keywords
	Keyword []string
}

// StickerKeywordTypeID is TL type id of StickerKeyword.
const StickerKeywordTypeID = 0xfcfeb29c

// Ensuring interfaces in compile-time for StickerKeyword.
var (
	_ bin.Encoder     = &StickerKeyword{}
	_ bin.Decoder     = &StickerKeyword{}
	_ bin.BareEncoder = &StickerKeyword{}
	_ bin.BareDecoder = &StickerKeyword{}
)

func (s *StickerKeyword) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.DocumentID == 0) {
		return false
	}
	if !(s.Keyword == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StickerKeyword) String() string {
	if s == nil {
		return "StickerKeyword(nil)"
	}
	type Alias StickerKeyword
	return fmt.Sprintf("StickerKeyword%+v", Alias(*s))
}

// FillFrom fills StickerKeyword from given interface.
func (s *StickerKeyword) FillFrom(from interface {
	GetDocumentID() (value int64)
	GetKeyword() (value []string)
}) {
	s.DocumentID = from.GetDocumentID()
	s.Keyword = from.GetKeyword()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickerKeyword) TypeID() uint32 {
	return StickerKeywordTypeID
}

// TypeName returns name of type in TL schema.
func (*StickerKeyword) TypeName() string {
	return "stickerKeyword"
}

// TypeInfo returns info about TL type.
func (s *StickerKeyword) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickerKeyword",
		ID:   StickerKeywordTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "DocumentID",
			SchemaName: "document_id",
		},
		{
			Name:       "Keyword",
			SchemaName: "keyword",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StickerKeyword) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerKeyword#fcfeb29c as nil")
	}
	b.PutID(StickerKeywordTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StickerKeyword) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerKeyword#fcfeb29c as nil")
	}
	b.PutLong(s.DocumentID)
	b.PutVectorHeader(len(s.Keyword))
	for _, v := range s.Keyword {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StickerKeyword) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerKeyword#fcfeb29c to nil")
	}
	if err := b.ConsumeID(StickerKeywordTypeID); err != nil {
		return fmt.Errorf("unable to decode stickerKeyword#fcfeb29c: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StickerKeyword) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerKeyword#fcfeb29c to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode stickerKeyword#fcfeb29c: field document_id: %w", err)
		}
		s.DocumentID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stickerKeyword#fcfeb29c: field keyword: %w", err)
		}

		if headerLen > 0 {
			s.Keyword = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode stickerKeyword#fcfeb29c: field keyword: %w", err)
			}
			s.Keyword = append(s.Keyword, value)
		}
	}
	return nil
}

// GetDocumentID returns value of DocumentID field.
func (s *StickerKeyword) GetDocumentID() (value int64) {
	if s == nil {
		return
	}
	return s.DocumentID
}

// GetKeyword returns value of Keyword field.
func (s *StickerKeyword) GetKeyword() (value []string) {
	if s == nil {
		return
	}
	return s.Keyword
}
