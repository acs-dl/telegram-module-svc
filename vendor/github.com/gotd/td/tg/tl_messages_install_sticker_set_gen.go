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

// MessagesInstallStickerSetRequest represents TL type `messages.installStickerSet#c78fe460`.
// Install a stickerset
//
// See https://core.telegram.org/method/messages.installStickerSet for reference.
type MessagesInstallStickerSetRequest struct {
	// Stickerset to install
	Stickerset InputStickerSetClass
	// Whether to archive stickerset
	Archived bool
}

// MessagesInstallStickerSetRequestTypeID is TL type id of MessagesInstallStickerSetRequest.
const MessagesInstallStickerSetRequestTypeID = 0xc78fe460

// Ensuring interfaces in compile-time for MessagesInstallStickerSetRequest.
var (
	_ bin.Encoder     = &MessagesInstallStickerSetRequest{}
	_ bin.Decoder     = &MessagesInstallStickerSetRequest{}
	_ bin.BareEncoder = &MessagesInstallStickerSetRequest{}
	_ bin.BareDecoder = &MessagesInstallStickerSetRequest{}
)

func (i *MessagesInstallStickerSetRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Stickerset == nil) {
		return false
	}
	if !(i.Archived == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *MessagesInstallStickerSetRequest) String() string {
	if i == nil {
		return "MessagesInstallStickerSetRequest(nil)"
	}
	type Alias MessagesInstallStickerSetRequest
	return fmt.Sprintf("MessagesInstallStickerSetRequest%+v", Alias(*i))
}

// FillFrom fills MessagesInstallStickerSetRequest from given interface.
func (i *MessagesInstallStickerSetRequest) FillFrom(from interface {
	GetStickerset() (value InputStickerSetClass)
	GetArchived() (value bool)
}) {
	i.Stickerset = from.GetStickerset()
	i.Archived = from.GetArchived()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesInstallStickerSetRequest) TypeID() uint32 {
	return MessagesInstallStickerSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesInstallStickerSetRequest) TypeName() string {
	return "messages.installStickerSet"
}

// TypeInfo returns info about TL type.
func (i *MessagesInstallStickerSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.installStickerSet",
		ID:   MessagesInstallStickerSetRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Stickerset",
			SchemaName: "stickerset",
		},
		{
			Name:       "Archived",
			SchemaName: "archived",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *MessagesInstallStickerSetRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode messages.installStickerSet#c78fe460 as nil")
	}
	b.PutID(MessagesInstallStickerSetRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *MessagesInstallStickerSetRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode messages.installStickerSet#c78fe460 as nil")
	}
	if i.Stickerset == nil {
		return fmt.Errorf("unable to encode messages.installStickerSet#c78fe460: field stickerset is nil")
	}
	if err := i.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.installStickerSet#c78fe460: field stickerset: %w", err)
	}
	b.PutBool(i.Archived)
	return nil
}

// Decode implements bin.Decoder.
func (i *MessagesInstallStickerSetRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode messages.installStickerSet#c78fe460 to nil")
	}
	if err := b.ConsumeID(MessagesInstallStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.installStickerSet#c78fe460: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *MessagesInstallStickerSetRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode messages.installStickerSet#c78fe460 to nil")
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.installStickerSet#c78fe460: field stickerset: %w", err)
		}
		i.Stickerset = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.installStickerSet#c78fe460: field archived: %w", err)
		}
		i.Archived = value
	}
	return nil
}

// GetStickerset returns value of Stickerset field.
func (i *MessagesInstallStickerSetRequest) GetStickerset() (value InputStickerSetClass) {
	if i == nil {
		return
	}
	return i.Stickerset
}

// GetArchived returns value of Archived field.
func (i *MessagesInstallStickerSetRequest) GetArchived() (value bool) {
	if i == nil {
		return
	}
	return i.Archived
}

// MessagesInstallStickerSet invokes method messages.installStickerSet#c78fe460 returning error if any.
// Install a stickerset
//
// Possible errors:
//
//	400 STICKERSET_INVALID: The provided sticker set is invalid.
//
// See https://core.telegram.org/method/messages.installStickerSet for reference.
func (c *Client) MessagesInstallStickerSet(ctx context.Context, request *MessagesInstallStickerSetRequest) (MessagesStickerSetInstallResultClass, error) {
	var result MessagesStickerSetInstallResultBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.StickerSetInstallResult, nil
}
