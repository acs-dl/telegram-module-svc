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

// PhoneRequestCallRequest represents TL type `phone.requestCall#42ff96ed`.
// Start a telegram phone call
//
// See https://core.telegram.org/method/phone.requestCall for reference.
type PhoneRequestCallRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to start a video call
	Video bool
	// Destination of the phone call
	UserID InputUserClass
	// Random ID to avoid resending the same object
	RandomID int
	// Parameter for E2E encryption key exchange »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/end-to-end/voice-calls
	GAHash []byte
	// Phone call settings
	Protocol PhoneCallProtocol
}

// PhoneRequestCallRequestTypeID is TL type id of PhoneRequestCallRequest.
const PhoneRequestCallRequestTypeID = 0x42ff96ed

// Ensuring interfaces in compile-time for PhoneRequestCallRequest.
var (
	_ bin.Encoder     = &PhoneRequestCallRequest{}
	_ bin.Decoder     = &PhoneRequestCallRequest{}
	_ bin.BareEncoder = &PhoneRequestCallRequest{}
	_ bin.BareDecoder = &PhoneRequestCallRequest{}
)

func (r *PhoneRequestCallRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Video == false) {
		return false
	}
	if !(r.UserID == nil) {
		return false
	}
	if !(r.RandomID == 0) {
		return false
	}
	if !(r.GAHash == nil) {
		return false
	}
	if !(r.Protocol.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *PhoneRequestCallRequest) String() string {
	if r == nil {
		return "PhoneRequestCallRequest(nil)"
	}
	type Alias PhoneRequestCallRequest
	return fmt.Sprintf("PhoneRequestCallRequest%+v", Alias(*r))
}

// FillFrom fills PhoneRequestCallRequest from given interface.
func (r *PhoneRequestCallRequest) FillFrom(from interface {
	GetVideo() (value bool)
	GetUserID() (value InputUserClass)
	GetRandomID() (value int)
	GetGAHash() (value []byte)
	GetProtocol() (value PhoneCallProtocol)
}) {
	r.Video = from.GetVideo()
	r.UserID = from.GetUserID()
	r.RandomID = from.GetRandomID()
	r.GAHash = from.GetGAHash()
	r.Protocol = from.GetProtocol()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneRequestCallRequest) TypeID() uint32 {
	return PhoneRequestCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneRequestCallRequest) TypeName() string {
	return "phone.requestCall"
}

// TypeInfo returns info about TL type.
func (r *PhoneRequestCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.requestCall",
		ID:   PhoneRequestCallRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Video",
			SchemaName: "video",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "GAHash",
			SchemaName: "g_a_hash",
		},
		{
			Name:       "Protocol",
			SchemaName: "protocol",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *PhoneRequestCallRequest) SetFlags() {
	if !(r.Video == false) {
		r.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (r *PhoneRequestCallRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode phone.requestCall#42ff96ed as nil")
	}
	b.PutID(PhoneRequestCallRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *PhoneRequestCallRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode phone.requestCall#42ff96ed as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.requestCall#42ff96ed: field flags: %w", err)
	}
	if r.UserID == nil {
		return fmt.Errorf("unable to encode phone.requestCall#42ff96ed: field user_id is nil")
	}
	if err := r.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.requestCall#42ff96ed: field user_id: %w", err)
	}
	b.PutInt(r.RandomID)
	b.PutBytes(r.GAHash)
	if err := r.Protocol.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.requestCall#42ff96ed: field protocol: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *PhoneRequestCallRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode phone.requestCall#42ff96ed to nil")
	}
	if err := b.ConsumeID(PhoneRequestCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.requestCall#42ff96ed: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *PhoneRequestCallRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode phone.requestCall#42ff96ed to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.requestCall#42ff96ed: field flags: %w", err)
		}
	}
	r.Video = r.Flags.Has(0)
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.requestCall#42ff96ed: field user_id: %w", err)
		}
		r.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.requestCall#42ff96ed: field random_id: %w", err)
		}
		r.RandomID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phone.requestCall#42ff96ed: field g_a_hash: %w", err)
		}
		r.GAHash = value
	}
	{
		if err := r.Protocol.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.requestCall#42ff96ed: field protocol: %w", err)
		}
	}
	return nil
}

// SetVideo sets value of Video conditional field.
func (r *PhoneRequestCallRequest) SetVideo(value bool) {
	if value {
		r.Flags.Set(0)
		r.Video = true
	} else {
		r.Flags.Unset(0)
		r.Video = false
	}
}

// GetVideo returns value of Video conditional field.
func (r *PhoneRequestCallRequest) GetVideo() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(0)
}

// GetUserID returns value of UserID field.
func (r *PhoneRequestCallRequest) GetUserID() (value InputUserClass) {
	if r == nil {
		return
	}
	return r.UserID
}

// GetRandomID returns value of RandomID field.
func (r *PhoneRequestCallRequest) GetRandomID() (value int) {
	if r == nil {
		return
	}
	return r.RandomID
}

// GetGAHash returns value of GAHash field.
func (r *PhoneRequestCallRequest) GetGAHash() (value []byte) {
	if r == nil {
		return
	}
	return r.GAHash
}

// GetProtocol returns value of Protocol field.
func (r *PhoneRequestCallRequest) GetProtocol() (value PhoneCallProtocol) {
	if r == nil {
		return
	}
	return r.Protocol
}

// PhoneRequestCall invokes method phone.requestCall#42ff96ed returning error if any.
// Start a telegram phone call
//
// Possible errors:
//
//	400 CALL_PROTOCOL_FLAGS_INVALID: Call protocol flags invalid.
//	400 INPUT_USER_DEACTIVATED: The specified user was deleted.
//	400 PARTICIPANT_VERSION_OUTDATED: The other participant does not use an up to date telegram client with support for calls.
//	400 USER_ID_INVALID: The provided user ID is invalid.
//	403 USER_IS_BLOCKED: You were blocked by this user.
//	403 USER_PRIVACY_RESTRICTED: The user's privacy settings do not allow you to do this.
//
// See https://core.telegram.org/method/phone.requestCall for reference.
func (c *Client) PhoneRequestCall(ctx context.Context, request *PhoneRequestCallRequest) (*PhonePhoneCall, error) {
	var result PhonePhoneCall

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
