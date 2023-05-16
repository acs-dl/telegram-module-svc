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

// InvokeWithTakeoutRequest represents TL type `invokeWithTakeout#aca9fd2e`.
// Invoke a method within a takeout session
//
// See https://core.telegram.org/constructor/invokeWithTakeout for reference.
type InvokeWithTakeoutRequest struct {
	// Takeout session ID
	TakeoutID int64
	// Query
	Query bin.Object
}

// InvokeWithTakeoutRequestTypeID is TL type id of InvokeWithTakeoutRequest.
const InvokeWithTakeoutRequestTypeID = 0xaca9fd2e

// Ensuring interfaces in compile-time for InvokeWithTakeoutRequest.
var (
	_ bin.Encoder     = &InvokeWithTakeoutRequest{}
	_ bin.Decoder     = &InvokeWithTakeoutRequest{}
	_ bin.BareEncoder = &InvokeWithTakeoutRequest{}
	_ bin.BareDecoder = &InvokeWithTakeoutRequest{}
)

func (i *InvokeWithTakeoutRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.TakeoutID == 0) {
		return false
	}
	if !(i.Query == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InvokeWithTakeoutRequest) String() string {
	if i == nil {
		return "InvokeWithTakeoutRequest(nil)"
	}
	type Alias InvokeWithTakeoutRequest
	return fmt.Sprintf("InvokeWithTakeoutRequest%+v", Alias(*i))
}

// FillFrom fills InvokeWithTakeoutRequest from given interface.
func (i *InvokeWithTakeoutRequest) FillFrom(from interface {
	GetTakeoutID() (value int64)
	GetQuery() (value bin.Object)
}) {
	i.TakeoutID = from.GetTakeoutID()
	i.Query = from.GetQuery()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InvokeWithTakeoutRequest) TypeID() uint32 {
	return InvokeWithTakeoutRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*InvokeWithTakeoutRequest) TypeName() string {
	return "invokeWithTakeout"
}

// TypeInfo returns info about TL type.
func (i *InvokeWithTakeoutRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "invokeWithTakeout",
		ID:   InvokeWithTakeoutRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TakeoutID",
			SchemaName: "takeout_id",
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InvokeWithTakeoutRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeWithTakeout#aca9fd2e as nil")
	}
	b.PutID(InvokeWithTakeoutRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InvokeWithTakeoutRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeWithTakeout#aca9fd2e as nil")
	}
	b.PutLong(i.TakeoutID)
	if err := i.Query.Encode(b); err != nil {
		return fmt.Errorf("unable to encode invokeWithTakeout#aca9fd2e: field query: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InvokeWithTakeoutRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeWithTakeout#aca9fd2e to nil")
	}
	if err := b.ConsumeID(InvokeWithTakeoutRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode invokeWithTakeout#aca9fd2e: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InvokeWithTakeoutRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeWithTakeout#aca9fd2e to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode invokeWithTakeout#aca9fd2e: field takeout_id: %w", err)
		}
		i.TakeoutID = value
	}
	{
		if err := i.Query.Decode(b); err != nil {
			return fmt.Errorf("unable to decode invokeWithTakeout#aca9fd2e: field query: %w", err)
		}
	}
	return nil
}

// GetTakeoutID returns value of TakeoutID field.
func (i *InvokeWithTakeoutRequest) GetTakeoutID() (value int64) {
	if i == nil {
		return
	}
	return i.TakeoutID
}

// GetQuery returns value of Query field.
func (i *InvokeWithTakeoutRequest) GetQuery() (value bin.Object) {
	if i == nil {
		return
	}
	return i.Query
}
