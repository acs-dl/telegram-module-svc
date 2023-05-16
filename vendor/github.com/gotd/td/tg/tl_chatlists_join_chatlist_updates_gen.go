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

// ChatlistsJoinChatlistUpdatesRequest represents TL type `chatlists.joinChatlistUpdates#e089f8f5`.
//
// See https://core.telegram.org/method/chatlists.joinChatlistUpdates for reference.
type ChatlistsJoinChatlistUpdatesRequest struct {
	// Chatlist field of ChatlistsJoinChatlistUpdatesRequest.
	Chatlist InputChatlistDialogFilter
	// Peers field of ChatlistsJoinChatlistUpdatesRequest.
	Peers []InputPeerClass
}

// ChatlistsJoinChatlistUpdatesRequestTypeID is TL type id of ChatlistsJoinChatlistUpdatesRequest.
const ChatlistsJoinChatlistUpdatesRequestTypeID = 0xe089f8f5

// Ensuring interfaces in compile-time for ChatlistsJoinChatlistUpdatesRequest.
var (
	_ bin.Encoder     = &ChatlistsJoinChatlistUpdatesRequest{}
	_ bin.Decoder     = &ChatlistsJoinChatlistUpdatesRequest{}
	_ bin.BareEncoder = &ChatlistsJoinChatlistUpdatesRequest{}
	_ bin.BareDecoder = &ChatlistsJoinChatlistUpdatesRequest{}
)

func (j *ChatlistsJoinChatlistUpdatesRequest) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.Chatlist.Zero()) {
		return false
	}
	if !(j.Peers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *ChatlistsJoinChatlistUpdatesRequest) String() string {
	if j == nil {
		return "ChatlistsJoinChatlistUpdatesRequest(nil)"
	}
	type Alias ChatlistsJoinChatlistUpdatesRequest
	return fmt.Sprintf("ChatlistsJoinChatlistUpdatesRequest%+v", Alias(*j))
}

// FillFrom fills ChatlistsJoinChatlistUpdatesRequest from given interface.
func (j *ChatlistsJoinChatlistUpdatesRequest) FillFrom(from interface {
	GetChatlist() (value InputChatlistDialogFilter)
	GetPeers() (value []InputPeerClass)
}) {
	j.Chatlist = from.GetChatlist()
	j.Peers = from.GetPeers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatlistsJoinChatlistUpdatesRequest) TypeID() uint32 {
	return ChatlistsJoinChatlistUpdatesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatlistsJoinChatlistUpdatesRequest) TypeName() string {
	return "chatlists.joinChatlistUpdates"
}

// TypeInfo returns info about TL type.
func (j *ChatlistsJoinChatlistUpdatesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatlists.joinChatlistUpdates",
		ID:   ChatlistsJoinChatlistUpdatesRequestTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Chatlist",
			SchemaName: "chatlist",
		},
		{
			Name:       "Peers",
			SchemaName: "peers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (j *ChatlistsJoinChatlistUpdatesRequest) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode chatlists.joinChatlistUpdates#e089f8f5 as nil")
	}
	b.PutID(ChatlistsJoinChatlistUpdatesRequestTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *ChatlistsJoinChatlistUpdatesRequest) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode chatlists.joinChatlistUpdates#e089f8f5 as nil")
	}
	if err := j.Chatlist.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatlists.joinChatlistUpdates#e089f8f5: field chatlist: %w", err)
	}
	b.PutVectorHeader(len(j.Peers))
	for idx, v := range j.Peers {
		if v == nil {
			return fmt.Errorf("unable to encode chatlists.joinChatlistUpdates#e089f8f5: field peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatlists.joinChatlistUpdates#e089f8f5: field peers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (j *ChatlistsJoinChatlistUpdatesRequest) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode chatlists.joinChatlistUpdates#e089f8f5 to nil")
	}
	if err := b.ConsumeID(ChatlistsJoinChatlistUpdatesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode chatlists.joinChatlistUpdates#e089f8f5: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *ChatlistsJoinChatlistUpdatesRequest) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode chatlists.joinChatlistUpdates#e089f8f5 to nil")
	}
	{
		if err := j.Chatlist.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatlists.joinChatlistUpdates#e089f8f5: field chatlist: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.joinChatlistUpdates#e089f8f5: field peers: %w", err)
		}

		if headerLen > 0 {
			j.Peers = make([]InputPeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatlists.joinChatlistUpdates#e089f8f5: field peers: %w", err)
			}
			j.Peers = append(j.Peers, value)
		}
	}
	return nil
}

// GetChatlist returns value of Chatlist field.
func (j *ChatlistsJoinChatlistUpdatesRequest) GetChatlist() (value InputChatlistDialogFilter) {
	if j == nil {
		return
	}
	return j.Chatlist
}

// GetPeers returns value of Peers field.
func (j *ChatlistsJoinChatlistUpdatesRequest) GetPeers() (value []InputPeerClass) {
	if j == nil {
		return
	}
	return j.Peers
}

// MapPeers returns field Peers wrapped in InputPeerClassArray helper.
func (j *ChatlistsJoinChatlistUpdatesRequest) MapPeers() (value InputPeerClassArray) {
	return InputPeerClassArray(j.Peers)
}

// ChatlistsJoinChatlistUpdates invokes method chatlists.joinChatlistUpdates#e089f8f5 returning error if any.
//
// See https://core.telegram.org/method/chatlists.joinChatlistUpdates for reference.
func (c *Client) ChatlistsJoinChatlistUpdates(ctx context.Context, request *ChatlistsJoinChatlistUpdatesRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
