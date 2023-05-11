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

// PhoneCheckGroupCallRequest represents TL type `phone.checkGroupCall#b59cf977`.
// Check whether the group call Server Forwarding Unit is currently receiving the streams
// with the specified WebRTC source IDs.
// Returns an intersection of the source IDs specified in sources, and the source IDs
// currently being forwarded by the SFU.
//
// See https://core.telegram.org/method/phone.checkGroupCall for reference.
type PhoneCheckGroupCallRequest struct {
	// Group call
	Call InputGroupCall
	// Source IDs
	Sources []int
}

// PhoneCheckGroupCallRequestTypeID is TL type id of PhoneCheckGroupCallRequest.
const PhoneCheckGroupCallRequestTypeID = 0xb59cf977

// Ensuring interfaces in compile-time for PhoneCheckGroupCallRequest.
var (
	_ bin.Encoder     = &PhoneCheckGroupCallRequest{}
	_ bin.Decoder     = &PhoneCheckGroupCallRequest{}
	_ bin.BareEncoder = &PhoneCheckGroupCallRequest{}
	_ bin.BareDecoder = &PhoneCheckGroupCallRequest{}
)

func (c *PhoneCheckGroupCallRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Call.Zero()) {
		return false
	}
	if !(c.Sources == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *PhoneCheckGroupCallRequest) String() string {
	if c == nil {
		return "PhoneCheckGroupCallRequest(nil)"
	}
	type Alias PhoneCheckGroupCallRequest
	return fmt.Sprintf("PhoneCheckGroupCallRequest%+v", Alias(*c))
}

// FillFrom fills PhoneCheckGroupCallRequest from given interface.
func (c *PhoneCheckGroupCallRequest) FillFrom(from interface {
	GetCall() (value InputGroupCall)
	GetSources() (value []int)
}) {
	c.Call = from.GetCall()
	c.Sources = from.GetSources()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneCheckGroupCallRequest) TypeID() uint32 {
	return PhoneCheckGroupCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneCheckGroupCallRequest) TypeName() string {
	return "phone.checkGroupCall"
}

// TypeInfo returns info about TL type.
func (c *PhoneCheckGroupCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.checkGroupCall",
		ID:   PhoneCheckGroupCallRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
		{
			Name:       "Sources",
			SchemaName: "sources",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *PhoneCheckGroupCallRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode phone.checkGroupCall#b59cf977 as nil")
	}
	b.PutID(PhoneCheckGroupCallRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *PhoneCheckGroupCallRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode phone.checkGroupCall#b59cf977 as nil")
	}
	if err := c.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.checkGroupCall#b59cf977: field call: %w", err)
	}
	b.PutVectorHeader(len(c.Sources))
	for _, v := range c.Sources {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *PhoneCheckGroupCallRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode phone.checkGroupCall#b59cf977 to nil")
	}
	if err := b.ConsumeID(PhoneCheckGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.checkGroupCall#b59cf977: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *PhoneCheckGroupCallRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode phone.checkGroupCall#b59cf977 to nil")
	}
	{
		if err := c.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.checkGroupCall#b59cf977: field call: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.checkGroupCall#b59cf977: field sources: %w", err)
		}

		if headerLen > 0 {
			c.Sources = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode phone.checkGroupCall#b59cf977: field sources: %w", err)
			}
			c.Sources = append(c.Sources, value)
		}
	}
	return nil
}

// GetCall returns value of Call field.
func (c *PhoneCheckGroupCallRequest) GetCall() (value InputGroupCall) {
	if c == nil {
		return
	}
	return c.Call
}

// GetSources returns value of Sources field.
func (c *PhoneCheckGroupCallRequest) GetSources() (value []int) {
	if c == nil {
		return
	}
	return c.Sources
}

// PhoneCheckGroupCall invokes method phone.checkGroupCall#b59cf977 returning error if any.
// Check whether the group call Server Forwarding Unit is currently receiving the streams
// with the specified WebRTC source IDs.
// Returns an intersection of the source IDs specified in sources, and the source IDs
// currently being forwarded by the SFU.
//
// Possible errors:
//
//	400 GROUPCALL_JOIN_MISSING: You haven't joined this group call.
//
// See https://core.telegram.org/method/phone.checkGroupCall for reference.
func (c *Client) PhoneCheckGroupCall(ctx context.Context, request *PhoneCheckGroupCallRequest) ([]int, error) {
	var result IntVector

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []int(result.Elems), nil
}
