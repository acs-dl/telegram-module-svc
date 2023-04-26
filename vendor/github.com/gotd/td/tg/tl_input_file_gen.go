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

// InputFile represents TL type `inputFile#f52ff27f`.
// Defines a file saved in parts using the method upload.saveFilePart¹.
//
// Links:
//  1. https://core.telegram.org/method/upload.saveFilePart
//
// See https://core.telegram.org/constructor/inputFile for reference.
type InputFile struct {
	// Random file identifier created by the client
	ID int64
	// Number of parts saved
	Parts int
	// Full name of the file
	Name string
	// In case the file's md5-hash¹ was passed, contents of the file will be checked prior
	// to use
	//
	// Links:
	//  1) https://en.wikipedia.org/wiki/MD5#MD5_hashes
	MD5Checksum string
}

// InputFileTypeID is TL type id of InputFile.
const InputFileTypeID = 0xf52ff27f

// construct implements constructor of InputFileClass.
func (i InputFile) construct() InputFileClass { return &i }

// Ensuring interfaces in compile-time for InputFile.
var (
	_ bin.Encoder     = &InputFile{}
	_ bin.Decoder     = &InputFile{}
	_ bin.BareEncoder = &InputFile{}
	_ bin.BareDecoder = &InputFile{}

	_ InputFileClass = &InputFile{}
)

func (i *InputFile) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.Parts == 0) {
		return false
	}
	if !(i.Name == "") {
		return false
	}
	if !(i.MD5Checksum == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputFile) String() string {
	if i == nil {
		return "InputFile(nil)"
	}
	type Alias InputFile
	return fmt.Sprintf("InputFile%+v", Alias(*i))
}

// FillFrom fills InputFile from given interface.
func (i *InputFile) FillFrom(from interface {
	GetID() (value int64)
	GetParts() (value int)
	GetName() (value string)
	GetMD5Checksum() (value string)
}) {
	i.ID = from.GetID()
	i.Parts = from.GetParts()
	i.Name = from.GetName()
	i.MD5Checksum = from.GetMD5Checksum()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputFile) TypeID() uint32 {
	return InputFileTypeID
}

// TypeName returns name of type in TL schema.
func (*InputFile) TypeName() string {
	return "inputFile"
}

// TypeInfo returns info about TL type.
func (i *InputFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputFile",
		ID:   InputFileTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Parts",
			SchemaName: "parts",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "MD5Checksum",
			SchemaName: "md5_checksum",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputFile) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFile#f52ff27f as nil")
	}
	b.PutID(InputFileTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputFile) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFile#f52ff27f as nil")
	}
	b.PutLong(i.ID)
	b.PutInt(i.Parts)
	b.PutString(i.Name)
	b.PutString(i.MD5Checksum)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputFile) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFile#f52ff27f to nil")
	}
	if err := b.ConsumeID(InputFileTypeID); err != nil {
		return fmt.Errorf("unable to decode inputFile#f52ff27f: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputFile) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFile#f52ff27f to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputFile#f52ff27f: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputFile#f52ff27f: field parts: %w", err)
		}
		i.Parts = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputFile#f52ff27f: field name: %w", err)
		}
		i.Name = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputFile#f52ff27f: field md5_checksum: %w", err)
		}
		i.MD5Checksum = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputFile) GetID() (value int64) {
	if i == nil {
		return
	}
	return i.ID
}

// GetParts returns value of Parts field.
func (i *InputFile) GetParts() (value int) {
	if i == nil {
		return
	}
	return i.Parts
}

// GetName returns value of Name field.
func (i *InputFile) GetName() (value string) {
	if i == nil {
		return
	}
	return i.Name
}

// GetMD5Checksum returns value of MD5Checksum field.
func (i *InputFile) GetMD5Checksum() (value string) {
	if i == nil {
		return
	}
	return i.MD5Checksum
}

// InputFileBig represents TL type `inputFileBig#fa4f0bb5`.
// Assigns a big file (over 10 MB in size), saved in part using the method upload
// saveBigFilePart¹.
//
// Links:
//  1. https://core.telegram.org/method/upload.saveBigFilePart
//
// See https://core.telegram.org/constructor/inputFileBig for reference.
type InputFileBig struct {
	// Random file id, created by the client
	ID int64
	// Number of parts saved
	Parts int
	// Full file name
	Name string
}

// InputFileBigTypeID is TL type id of InputFileBig.
const InputFileBigTypeID = 0xfa4f0bb5

// construct implements constructor of InputFileClass.
func (i InputFileBig) construct() InputFileClass { return &i }

// Ensuring interfaces in compile-time for InputFileBig.
var (
	_ bin.Encoder     = &InputFileBig{}
	_ bin.Decoder     = &InputFileBig{}
	_ bin.BareEncoder = &InputFileBig{}
	_ bin.BareDecoder = &InputFileBig{}

	_ InputFileClass = &InputFileBig{}
)

func (i *InputFileBig) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.Parts == 0) {
		return false
	}
	if !(i.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputFileBig) String() string {
	if i == nil {
		return "InputFileBig(nil)"
	}
	type Alias InputFileBig
	return fmt.Sprintf("InputFileBig%+v", Alias(*i))
}

// FillFrom fills InputFileBig from given interface.
func (i *InputFileBig) FillFrom(from interface {
	GetID() (value int64)
	GetParts() (value int)
	GetName() (value string)
}) {
	i.ID = from.GetID()
	i.Parts = from.GetParts()
	i.Name = from.GetName()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputFileBig) TypeID() uint32 {
	return InputFileBigTypeID
}

// TypeName returns name of type in TL schema.
func (*InputFileBig) TypeName() string {
	return "inputFileBig"
}

// TypeInfo returns info about TL type.
func (i *InputFileBig) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputFileBig",
		ID:   InputFileBigTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Parts",
			SchemaName: "parts",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputFileBig) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileBig#fa4f0bb5 as nil")
	}
	b.PutID(InputFileBigTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputFileBig) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileBig#fa4f0bb5 as nil")
	}
	b.PutLong(i.ID)
	b.PutInt(i.Parts)
	b.PutString(i.Name)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputFileBig) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileBig#fa4f0bb5 to nil")
	}
	if err := b.ConsumeID(InputFileBigTypeID); err != nil {
		return fmt.Errorf("unable to decode inputFileBig#fa4f0bb5: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputFileBig) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileBig#fa4f0bb5 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileBig#fa4f0bb5: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileBig#fa4f0bb5: field parts: %w", err)
		}
		i.Parts = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileBig#fa4f0bb5: field name: %w", err)
		}
		i.Name = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputFileBig) GetID() (value int64) {
	if i == nil {
		return
	}
	return i.ID
}

// GetParts returns value of Parts field.
func (i *InputFileBig) GetParts() (value int) {
	if i == nil {
		return
	}
	return i.Parts
}

// GetName returns value of Name field.
func (i *InputFileBig) GetName() (value string) {
	if i == nil {
		return
	}
	return i.Name
}

// InputFileClassName is schema name of InputFileClass.
const InputFileClassName = "InputFile"

// InputFileClass represents InputFile generic type.
//
// See https://core.telegram.org/type/InputFile for reference.
//
// Example:
//
//	g, err := tg.DecodeInputFile(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.InputFile: // inputFile#f52ff27f
//	case *tg.InputFileBig: // inputFileBig#fa4f0bb5
//	default: panic(v)
//	}
type InputFileClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputFileClass

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

	// Random file identifier created by the client
	GetID() (value int64)

	// Number of parts saved
	GetParts() (value int)

	// Full name of the file
	GetName() (value string)
}

// DecodeInputFile implements binary de-serialization for InputFileClass.
func DecodeInputFile(buf *bin.Buffer) (InputFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputFileTypeID:
		// Decoding inputFile#f52ff27f.
		v := InputFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileClass: %w", err)
		}
		return &v, nil
	case InputFileBigTypeID:
		// Decoding inputFileBig#fa4f0bb5.
		v := InputFileBig{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputFile boxes the InputFileClass providing a helper.
type InputFileBox struct {
	InputFile InputFileClass
}

// Decode implements bin.Decoder for InputFileBox.
func (b *InputFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputFileBox to nil")
	}
	v, err := DecodeInputFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputFile = v
	return nil
}

// Encode implements bin.Encode for InputFileBox.
func (b *InputFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputFile == nil {
		return fmt.Errorf("unable to encode InputFileClass as nil")
	}
	return b.InputFile.Encode(buf)
}
