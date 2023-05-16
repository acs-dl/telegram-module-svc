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

// Poll represents TL type `poll#86e18161`.
// Poll
//
// See https://core.telegram.org/constructor/poll for reference.
type Poll struct {
	// ID of the poll
	ID int64
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the poll is closed and doesn't accept any more answers
	Closed bool
	// Whether cast votes are publicly visible to all users (non-anonymous poll)
	PublicVoters bool
	// Whether multiple options can be chosen as answer
	MultipleChoice bool
	// Whether this is a quiz (with wrong and correct answers, results shown in the return
	// type)
	Quiz bool
	// The question of the poll
	Question string
	// The possible answers, vote using messages.sendVote¹.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.sendVote
	Answers []PollAnswer
	// Amount of time in seconds the poll will be active after creation, 5-600. Can't be used
	// together with close_date.
	//
	// Use SetClosePeriod and GetClosePeriod helpers.
	ClosePeriod int
	// Point in time (Unix timestamp) when the poll will be automatically closed. Must be at
	// least 5 and no more than 600 seconds in the future; can't be used together with
	// close_period.
	//
	// Use SetCloseDate and GetCloseDate helpers.
	CloseDate int
}

// PollTypeID is TL type id of Poll.
const PollTypeID = 0x86e18161

// Ensuring interfaces in compile-time for Poll.
var (
	_ bin.Encoder     = &Poll{}
	_ bin.Decoder     = &Poll{}
	_ bin.BareEncoder = &Poll{}
	_ bin.BareDecoder = &Poll{}
)

func (p *Poll) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ID == 0) {
		return false
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.Closed == false) {
		return false
	}
	if !(p.PublicVoters == false) {
		return false
	}
	if !(p.MultipleChoice == false) {
		return false
	}
	if !(p.Quiz == false) {
		return false
	}
	if !(p.Question == "") {
		return false
	}
	if !(p.Answers == nil) {
		return false
	}
	if !(p.ClosePeriod == 0) {
		return false
	}
	if !(p.CloseDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *Poll) String() string {
	if p == nil {
		return "Poll(nil)"
	}
	type Alias Poll
	return fmt.Sprintf("Poll%+v", Alias(*p))
}

// FillFrom fills Poll from given interface.
func (p *Poll) FillFrom(from interface {
	GetID() (value int64)
	GetClosed() (value bool)
	GetPublicVoters() (value bool)
	GetMultipleChoice() (value bool)
	GetQuiz() (value bool)
	GetQuestion() (value string)
	GetAnswers() (value []PollAnswer)
	GetClosePeriod() (value int, ok bool)
	GetCloseDate() (value int, ok bool)
}) {
	p.ID = from.GetID()
	p.Closed = from.GetClosed()
	p.PublicVoters = from.GetPublicVoters()
	p.MultipleChoice = from.GetMultipleChoice()
	p.Quiz = from.GetQuiz()
	p.Question = from.GetQuestion()
	p.Answers = from.GetAnswers()
	if val, ok := from.GetClosePeriod(); ok {
		p.ClosePeriod = val
	}

	if val, ok := from.GetCloseDate(); ok {
		p.CloseDate = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Poll) TypeID() uint32 {
	return PollTypeID
}

// TypeName returns name of type in TL schema.
func (*Poll) TypeName() string {
	return "poll"
}

// TypeInfo returns info about TL type.
func (p *Poll) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "poll",
		ID:   PollTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Closed",
			SchemaName: "closed",
			Null:       !p.Flags.Has(0),
		},
		{
			Name:       "PublicVoters",
			SchemaName: "public_voters",
			Null:       !p.Flags.Has(1),
		},
		{
			Name:       "MultipleChoice",
			SchemaName: "multiple_choice",
			Null:       !p.Flags.Has(2),
		},
		{
			Name:       "Quiz",
			SchemaName: "quiz",
			Null:       !p.Flags.Has(3),
		},
		{
			Name:       "Question",
			SchemaName: "question",
		},
		{
			Name:       "Answers",
			SchemaName: "answers",
		},
		{
			Name:       "ClosePeriod",
			SchemaName: "close_period",
			Null:       !p.Flags.Has(4),
		},
		{
			Name:       "CloseDate",
			SchemaName: "close_date",
			Null:       !p.Flags.Has(5),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (p *Poll) SetFlags() {
	if !(p.Closed == false) {
		p.Flags.Set(0)
	}
	if !(p.PublicVoters == false) {
		p.Flags.Set(1)
	}
	if !(p.MultipleChoice == false) {
		p.Flags.Set(2)
	}
	if !(p.Quiz == false) {
		p.Flags.Set(3)
	}
	if !(p.ClosePeriod == 0) {
		p.Flags.Set(4)
	}
	if !(p.CloseDate == 0) {
		p.Flags.Set(5)
	}
}

// Encode implements bin.Encoder.
func (p *Poll) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode poll#86e18161 as nil")
	}
	b.PutID(PollTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *Poll) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode poll#86e18161 as nil")
	}
	p.SetFlags()
	b.PutLong(p.ID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode poll#86e18161: field flags: %w", err)
	}
	b.PutString(p.Question)
	b.PutVectorHeader(len(p.Answers))
	for idx, v := range p.Answers {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode poll#86e18161: field answers element with index %d: %w", idx, err)
		}
	}
	if p.Flags.Has(4) {
		b.PutInt(p.ClosePeriod)
	}
	if p.Flags.Has(5) {
		b.PutInt(p.CloseDate)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *Poll) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode poll#86e18161 to nil")
	}
	if err := b.ConsumeID(PollTypeID); err != nil {
		return fmt.Errorf("unable to decode poll#86e18161: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *Poll) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode poll#86e18161 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode poll#86e18161: field id: %w", err)
		}
		p.ID = value
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode poll#86e18161: field flags: %w", err)
		}
	}
	p.Closed = p.Flags.Has(0)
	p.PublicVoters = p.Flags.Has(1)
	p.MultipleChoice = p.Flags.Has(2)
	p.Quiz = p.Flags.Has(3)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode poll#86e18161: field question: %w", err)
		}
		p.Question = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode poll#86e18161: field answers: %w", err)
		}

		if headerLen > 0 {
			p.Answers = make([]PollAnswer, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PollAnswer
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode poll#86e18161: field answers: %w", err)
			}
			p.Answers = append(p.Answers, value)
		}
	}
	if p.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode poll#86e18161: field close_period: %w", err)
		}
		p.ClosePeriod = value
	}
	if p.Flags.Has(5) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode poll#86e18161: field close_date: %w", err)
		}
		p.CloseDate = value
	}
	return nil
}

// GetID returns value of ID field.
func (p *Poll) GetID() (value int64) {
	if p == nil {
		return
	}
	return p.ID
}

// SetClosed sets value of Closed conditional field.
func (p *Poll) SetClosed(value bool) {
	if value {
		p.Flags.Set(0)
		p.Closed = true
	} else {
		p.Flags.Unset(0)
		p.Closed = false
	}
}

// GetClosed returns value of Closed conditional field.
func (p *Poll) GetClosed() (value bool) {
	if p == nil {
		return
	}
	return p.Flags.Has(0)
}

// SetPublicVoters sets value of PublicVoters conditional field.
func (p *Poll) SetPublicVoters(value bool) {
	if value {
		p.Flags.Set(1)
		p.PublicVoters = true
	} else {
		p.Flags.Unset(1)
		p.PublicVoters = false
	}
}

// GetPublicVoters returns value of PublicVoters conditional field.
func (p *Poll) GetPublicVoters() (value bool) {
	if p == nil {
		return
	}
	return p.Flags.Has(1)
}

// SetMultipleChoice sets value of MultipleChoice conditional field.
func (p *Poll) SetMultipleChoice(value bool) {
	if value {
		p.Flags.Set(2)
		p.MultipleChoice = true
	} else {
		p.Flags.Unset(2)
		p.MultipleChoice = false
	}
}

// GetMultipleChoice returns value of MultipleChoice conditional field.
func (p *Poll) GetMultipleChoice() (value bool) {
	if p == nil {
		return
	}
	return p.Flags.Has(2)
}

// SetQuiz sets value of Quiz conditional field.
func (p *Poll) SetQuiz(value bool) {
	if value {
		p.Flags.Set(3)
		p.Quiz = true
	} else {
		p.Flags.Unset(3)
		p.Quiz = false
	}
}

// GetQuiz returns value of Quiz conditional field.
func (p *Poll) GetQuiz() (value bool) {
	if p == nil {
		return
	}
	return p.Flags.Has(3)
}

// GetQuestion returns value of Question field.
func (p *Poll) GetQuestion() (value string) {
	if p == nil {
		return
	}
	return p.Question
}

// GetAnswers returns value of Answers field.
func (p *Poll) GetAnswers() (value []PollAnswer) {
	if p == nil {
		return
	}
	return p.Answers
}

// SetClosePeriod sets value of ClosePeriod conditional field.
func (p *Poll) SetClosePeriod(value int) {
	p.Flags.Set(4)
	p.ClosePeriod = value
}

// GetClosePeriod returns value of ClosePeriod conditional field and
// boolean which is true if field was set.
func (p *Poll) GetClosePeriod() (value int, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(4) {
		return value, false
	}
	return p.ClosePeriod, true
}

// SetCloseDate sets value of CloseDate conditional field.
func (p *Poll) SetCloseDate(value int) {
	p.Flags.Set(5)
	p.CloseDate = value
}

// GetCloseDate returns value of CloseDate conditional field and
// boolean which is true if field was set.
func (p *Poll) GetCloseDate() (value int, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(5) {
		return value, false
	}
	return p.CloseDate, true
}
