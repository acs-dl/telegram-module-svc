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

// MessagesSentEncryptedMessage represents TL type `messages.sentEncryptedMessage#560f8935`.
// Message without file attachments sent to an encrypted file.
//
// See https://core.telegram.org/constructor/messages.sentEncryptedMessage for reference.
type MessagesSentEncryptedMessage struct {
	// Date of sending
	Date int
}

// MessagesSentEncryptedMessageTypeID is TL type id of MessagesSentEncryptedMessage.
const MessagesSentEncryptedMessageTypeID = 0x560f8935

// construct implements constructor of MessagesSentEncryptedMessageClass.
func (s MessagesSentEncryptedMessage) construct() MessagesSentEncryptedMessageClass { return &s }

// Ensuring interfaces in compile-time for MessagesSentEncryptedMessage.
var (
	_ bin.Encoder     = &MessagesSentEncryptedMessage{}
	_ bin.Decoder     = &MessagesSentEncryptedMessage{}
	_ bin.BareEncoder = &MessagesSentEncryptedMessage{}
	_ bin.BareDecoder = &MessagesSentEncryptedMessage{}

	_ MessagesSentEncryptedMessageClass = &MessagesSentEncryptedMessage{}
)

func (s *MessagesSentEncryptedMessage) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSentEncryptedMessage) String() string {
	if s == nil {
		return "MessagesSentEncryptedMessage(nil)"
	}
	type Alias MessagesSentEncryptedMessage
	return fmt.Sprintf("MessagesSentEncryptedMessage%+v", Alias(*s))
}

// FillFrom fills MessagesSentEncryptedMessage from given interface.
func (s *MessagesSentEncryptedMessage) FillFrom(from interface {
	GetDate() (value int)
}) {
	s.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSentEncryptedMessage) TypeID() uint32 {
	return MessagesSentEncryptedMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSentEncryptedMessage) TypeName() string {
	return "messages.sentEncryptedMessage"
}

// TypeInfo returns info about TL type.
func (s *MessagesSentEncryptedMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sentEncryptedMessage",
		ID:   MessagesSentEncryptedMessageTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSentEncryptedMessage) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sentEncryptedMessage#560f8935 as nil")
	}
	b.PutID(MessagesSentEncryptedMessageTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSentEncryptedMessage) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sentEncryptedMessage#560f8935 as nil")
	}
	b.PutInt(s.Date)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSentEncryptedMessage) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sentEncryptedMessage#560f8935 to nil")
	}
	if err := b.ConsumeID(MessagesSentEncryptedMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sentEncryptedMessage#560f8935: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSentEncryptedMessage) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sentEncryptedMessage#560f8935 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sentEncryptedMessage#560f8935: field date: %w", err)
		}
		s.Date = value
	}
	return nil
}

// GetDate returns value of Date field.
func (s *MessagesSentEncryptedMessage) GetDate() (value int) {
	if s == nil {
		return
	}
	return s.Date
}

// MessagesSentEncryptedFile represents TL type `messages.sentEncryptedFile#9493ff32`.
// Message with a file enclosure sent to a protected chat
//
// See https://core.telegram.org/constructor/messages.sentEncryptedFile for reference.
type MessagesSentEncryptedFile struct {
	// Sending date
	Date int
	// Attached file
	File EncryptedFileClass
}

// MessagesSentEncryptedFileTypeID is TL type id of MessagesSentEncryptedFile.
const MessagesSentEncryptedFileTypeID = 0x9493ff32

// construct implements constructor of MessagesSentEncryptedMessageClass.
func (s MessagesSentEncryptedFile) construct() MessagesSentEncryptedMessageClass { return &s }

// Ensuring interfaces in compile-time for MessagesSentEncryptedFile.
var (
	_ bin.Encoder     = &MessagesSentEncryptedFile{}
	_ bin.Decoder     = &MessagesSentEncryptedFile{}
	_ bin.BareEncoder = &MessagesSentEncryptedFile{}
	_ bin.BareDecoder = &MessagesSentEncryptedFile{}

	_ MessagesSentEncryptedMessageClass = &MessagesSentEncryptedFile{}
)

func (s *MessagesSentEncryptedFile) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Date == 0) {
		return false
	}
	if !(s.File == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSentEncryptedFile) String() string {
	if s == nil {
		return "MessagesSentEncryptedFile(nil)"
	}
	type Alias MessagesSentEncryptedFile
	return fmt.Sprintf("MessagesSentEncryptedFile%+v", Alias(*s))
}

// FillFrom fills MessagesSentEncryptedFile from given interface.
func (s *MessagesSentEncryptedFile) FillFrom(from interface {
	GetDate() (value int)
	GetFile() (value EncryptedFileClass)
}) {
	s.Date = from.GetDate()
	s.File = from.GetFile()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSentEncryptedFile) TypeID() uint32 {
	return MessagesSentEncryptedFileTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSentEncryptedFile) TypeName() string {
	return "messages.sentEncryptedFile"
}

// TypeInfo returns info about TL type.
func (s *MessagesSentEncryptedFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sentEncryptedFile",
		ID:   MessagesSentEncryptedFileTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "File",
			SchemaName: "file",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSentEncryptedFile) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sentEncryptedFile#9493ff32 as nil")
	}
	b.PutID(MessagesSentEncryptedFileTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSentEncryptedFile) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sentEncryptedFile#9493ff32 as nil")
	}
	b.PutInt(s.Date)
	if s.File == nil {
		return fmt.Errorf("unable to encode messages.sentEncryptedFile#9493ff32: field file is nil")
	}
	if err := s.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sentEncryptedFile#9493ff32: field file: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSentEncryptedFile) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sentEncryptedFile#9493ff32 to nil")
	}
	if err := b.ConsumeID(MessagesSentEncryptedFileTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sentEncryptedFile#9493ff32: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSentEncryptedFile) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sentEncryptedFile#9493ff32 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sentEncryptedFile#9493ff32: field date: %w", err)
		}
		s.Date = value
	}
	{
		value, err := DecodeEncryptedFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sentEncryptedFile#9493ff32: field file: %w", err)
		}
		s.File = value
	}
	return nil
}

// GetDate returns value of Date field.
func (s *MessagesSentEncryptedFile) GetDate() (value int) {
	if s == nil {
		return
	}
	return s.Date
}

// GetFile returns value of File field.
func (s *MessagesSentEncryptedFile) GetFile() (value EncryptedFileClass) {
	if s == nil {
		return
	}
	return s.File
}

// MessagesSentEncryptedMessageClassName is schema name of MessagesSentEncryptedMessageClass.
const MessagesSentEncryptedMessageClassName = "messages.SentEncryptedMessage"

// MessagesSentEncryptedMessageClass represents messages.SentEncryptedMessage generic type.
//
// See https://core.telegram.org/type/messages.SentEncryptedMessage for reference.
//
// Example:
//
//	g, err := tg.DecodeMessagesSentEncryptedMessage(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.MessagesSentEncryptedMessage: // messages.sentEncryptedMessage#560f8935
//	case *tg.MessagesSentEncryptedFile: // messages.sentEncryptedFile#9493ff32
//	default: panic(v)
//	}
type MessagesSentEncryptedMessageClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesSentEncryptedMessageClass

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

	// Date of sending
	GetDate() (value int)
}

// DecodeMessagesSentEncryptedMessage implements binary de-serialization for MessagesSentEncryptedMessageClass.
func DecodeMessagesSentEncryptedMessage(buf *bin.Buffer) (MessagesSentEncryptedMessageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesSentEncryptedMessageTypeID:
		// Decoding messages.sentEncryptedMessage#560f8935.
		v := MessagesSentEncryptedMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesSentEncryptedMessageClass: %w", err)
		}
		return &v, nil
	case MessagesSentEncryptedFileTypeID:
		// Decoding messages.sentEncryptedFile#9493ff32.
		v := MessagesSentEncryptedFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesSentEncryptedMessageClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesSentEncryptedMessageClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesSentEncryptedMessage boxes the MessagesSentEncryptedMessageClass providing a helper.
type MessagesSentEncryptedMessageBox struct {
	SentEncryptedMessage MessagesSentEncryptedMessageClass
}

// Decode implements bin.Decoder for MessagesSentEncryptedMessageBox.
func (b *MessagesSentEncryptedMessageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesSentEncryptedMessageBox to nil")
	}
	v, err := DecodeMessagesSentEncryptedMessage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SentEncryptedMessage = v
	return nil
}

// Encode implements bin.Encode for MessagesSentEncryptedMessageBox.
func (b *MessagesSentEncryptedMessageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SentEncryptedMessage == nil {
		return fmt.Errorf("unable to encode MessagesSentEncryptedMessageClass as nil")
	}
	return b.SentEncryptedMessage.Encode(buf)
}