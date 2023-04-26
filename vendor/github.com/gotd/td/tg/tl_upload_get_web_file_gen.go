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

// UploadGetWebFileRequest represents TL type `upload.getWebFile#24e6818d`.
//
// See https://core.telegram.org/method/upload.getWebFile for reference.
type UploadGetWebFileRequest struct {
	// Location field of UploadGetWebFileRequest.
	Location InputWebFileLocationClass
	// Offset field of UploadGetWebFileRequest.
	Offset int
	// Limit field of UploadGetWebFileRequest.
	Limit int
}

// UploadGetWebFileRequestTypeID is TL type id of UploadGetWebFileRequest.
const UploadGetWebFileRequestTypeID = 0x24e6818d

// Ensuring interfaces in compile-time for UploadGetWebFileRequest.
var (
	_ bin.Encoder     = &UploadGetWebFileRequest{}
	_ bin.Decoder     = &UploadGetWebFileRequest{}
	_ bin.BareEncoder = &UploadGetWebFileRequest{}
	_ bin.BareDecoder = &UploadGetWebFileRequest{}
)

func (g *UploadGetWebFileRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Location == nil) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *UploadGetWebFileRequest) String() string {
	if g == nil {
		return "UploadGetWebFileRequest(nil)"
	}
	type Alias UploadGetWebFileRequest
	return fmt.Sprintf("UploadGetWebFileRequest%+v", Alias(*g))
}

// FillFrom fills UploadGetWebFileRequest from given interface.
func (g *UploadGetWebFileRequest) FillFrom(from interface {
	GetLocation() (value InputWebFileLocationClass)
	GetOffset() (value int)
	GetLimit() (value int)
}) {
	g.Location = from.GetLocation()
	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadGetWebFileRequest) TypeID() uint32 {
	return UploadGetWebFileRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadGetWebFileRequest) TypeName() string {
	return "upload.getWebFile"
}

// TypeInfo returns info about TL type.
func (g *UploadGetWebFileRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upload.getWebFile",
		ID:   UploadGetWebFileRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Location",
			SchemaName: "location",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *UploadGetWebFileRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode upload.getWebFile#24e6818d as nil")
	}
	b.PutID(UploadGetWebFileRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *UploadGetWebFileRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode upload.getWebFile#24e6818d as nil")
	}
	if g.Location == nil {
		return fmt.Errorf("unable to encode upload.getWebFile#24e6818d: field location is nil")
	}
	if err := g.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode upload.getWebFile#24e6818d: field location: %w", err)
	}
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *UploadGetWebFileRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode upload.getWebFile#24e6818d to nil")
	}
	if err := b.ConsumeID(UploadGetWebFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.getWebFile#24e6818d: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *UploadGetWebFileRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode upload.getWebFile#24e6818d to nil")
	}
	{
		value, err := DecodeInputWebFileLocation(b)
		if err != nil {
			return fmt.Errorf("unable to decode upload.getWebFile#24e6818d: field location: %w", err)
		}
		g.Location = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getWebFile#24e6818d: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getWebFile#24e6818d: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetLocation returns value of Location field.
func (g *UploadGetWebFileRequest) GetLocation() (value InputWebFileLocationClass) {
	if g == nil {
		return
	}
	return g.Location
}

// GetOffset returns value of Offset field.
func (g *UploadGetWebFileRequest) GetOffset() (value int) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *UploadGetWebFileRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// UploadGetWebFile invokes method upload.getWebFile#24e6818d returning error if any.
//
// See https://core.telegram.org/method/upload.getWebFile for reference.
func (c *Client) UploadGetWebFile(ctx context.Context, request *UploadGetWebFileRequest) (*UploadWebFile, error) {
	var result UploadWebFile

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
