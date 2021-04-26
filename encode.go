package uelevelclip

import (
	"fmt"
	"io"
)

var DefaultEncodeOpt = &EncodeOpt{
	Indent: "   ",
}

type EncodeOpt struct {
	Indent string
}

type Encoder struct {
	w      io.Writer
	opt    *EncodeOpt
	indent int
}

func NewEncoder(w io.Writer, opt *EncodeOpt) *Encoder {
	if opt == nil {
		opt = DefaultEncodeOpt
	}
	return &Encoder{w: w, opt: opt}
}

func (e *Encoder) Encode(b *Block) error {
	return b.encode(e)
}

func (e *Encoder) writeIndent() error {
	for i := 0; i < e.indent; i++ {
		_, err := fmt.Fprint(e.w, e.indent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Block) encode(e *Encoder) error {
	if err := e.writeIndent(); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(e.w, "Begin %v", b.Class); err != nil {
		return err
	}

	if b.Option != nil && len(b.Option) > 0 {
		for k, v := range b.Option {
			if _, err := fmt.Fprintf(e.w, " %v=%v", k, v); err != nil {
				return err
			}
		}
	}

	if _, err := fmt.Fprintln(e.w); err != nil {
		return err
	}
	e.indent++
	for _, c := range b.Children {
		c.encode(e)
	}
	e.indent--
	if err := e.writeIndent(); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(e.w, "End %v\n", b.Class); err != nil {
		return err
	}
	return nil
}

func (b *Line) encode(e *Encoder) error {
	if err := e.writeIndent(); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(e.w, string(*b)); err != nil {
		return err
	}
	return nil
}
