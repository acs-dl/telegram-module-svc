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

// PhoneSaveCallLogRequest represents TL type `phone.saveCallLog#41248786`.
// Save phone call debug information
//
// See https://core.telegram.org/method/phone.saveCallLog for reference.
type PhoneSaveCallLogRequest struct {
	// Phone call
	Peer InputPhoneCall
	// Logs
	File InputFileClass
}

// PhoneSaveCallLogRequestTypeID is TL type id of PhoneSaveCallLogRequest.
const PhoneSaveCallLogRequestTypeID = 0x41248786

// Ensuring interfaces in compile-time for PhoneSaveCallLogRequest.
var (
	_ bin.Encoder     = &PhoneSaveCallLogRequest{}
	_ bin.Decoder     = &PhoneSaveCallLogRequest{}
	_ bin.BareEncoder = &PhoneSaveCallLogRequest{}
	_ bin.BareDecoder = &PhoneSaveCallLogRequest{}
)

func (s *PhoneSaveCallLogRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer.Zero()) {
		return false
	}
	if !(s.File == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PhoneSaveCallLogRequest) String() string {
	if s == nil {
		return "PhoneSaveCallLogRequest(nil)"
	}
	type Alias PhoneSaveCallLogRequest
	return fmt.Sprintf("PhoneSaveCallLogRequest%+v", Alias(*s))
}

// FillFrom fills PhoneSaveCallLogRequest from given interface.
func (s *PhoneSaveCallLogRequest) FillFrom(from interface {
	GetPeer() (value InputPhoneCall)
	GetFile() (value InputFileClass)
}) {
	s.Peer = from.GetPeer()
	s.File = from.GetFile()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneSaveCallLogRequest) TypeID() uint32 {
	return PhoneSaveCallLogRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneSaveCallLogRequest) TypeName() string {
	return "phone.saveCallLog"
}

// TypeInfo returns info about TL type.
func (s *PhoneSaveCallLogRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.saveCallLog",
		ID:   PhoneSaveCallLogRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "File",
			SchemaName: "file",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *PhoneSaveCallLogRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode phone.saveCallLog#41248786 as nil")
	}
	b.PutID(PhoneSaveCallLogRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *PhoneSaveCallLogRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode phone.saveCallLog#41248786 as nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.saveCallLog#41248786: field peer: %w", err)
	}
	if s.File == nil {
		return fmt.Errorf("unable to encode phone.saveCallLog#41248786: field file is nil")
	}
	if err := s.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.saveCallLog#41248786: field file: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *PhoneSaveCallLogRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode phone.saveCallLog#41248786 to nil")
	}
	if err := b.ConsumeID(PhoneSaveCallLogRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.saveCallLog#41248786: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *PhoneSaveCallLogRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode phone.saveCallLog#41248786 to nil")
	}
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.saveCallLog#41248786: field peer: %w", err)
		}
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.saveCallLog#41248786: field file: %w", err)
		}
		s.File = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *PhoneSaveCallLogRequest) GetPeer() (value InputPhoneCall) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetFile returns value of File field.
func (s *PhoneSaveCallLogRequest) GetFile() (value InputFileClass) {
	if s == nil {
		return
	}
	return s.File
}

// PhoneSaveCallLog invokes method phone.saveCallLog#41248786 returning error if any.
// Save phone call debug information
//
// See https://core.telegram.org/method/phone.saveCallLog for reference.
func (c *Client) PhoneSaveCallLog(ctx context.Context, request *PhoneSaveCallLogRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
