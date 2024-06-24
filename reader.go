package ajson

import (
	"errors"
)

type Reader struct {
	data         string
	dataLen      int
	currentIndex int
}

func NewReader(data string) *Reader {
	return &Reader{
		data:         data,
		dataLen:      len(data),
		currentIndex: 0,
	}
}

func (r *Reader) IsEOF() bool {
	return r.currentIndex >= r.dataLen
}

func (r *Reader) Peek() byte {
	if r.IsEOF() {
		return 0
	}

	return r.data[r.currentIndex]
}

func (r *Reader) VisitOrEOF(expected byte) error {
	if r.IsEOF() {
		return nil
	}
	if r.data[r.currentIndex] != expected {
		return errors.New(ErrUnexpectedEndOfInput)
	}
	r.VisitNext()
	return nil
}

func (r *Reader) Visit(expected byte) error {
	if r.IsEOF() || r.data[r.currentIndex] != expected {
		return errors.New(ErrUnexpectedEndOfInput)
	}

	r.VisitNext()

	return nil
}

func (r *Reader) VisitIfNext(expected byte) bool {
	if r.Peek() == expected {
		r.VisitNext()
		return true
	}

	return false
}

func (r *Reader) PeekEndChar() error {
	if r.IsEOF() {
		return nil
	}
	if r.data[r.currentIndex] == ',' {
		return nil
	}
	if r.data[r.currentIndex] == ']' {
		return nil
	}
	if r.data[r.currentIndex] == '}' {
		return nil
	}
	return errors.New(ErrUnexpectedEndOfInput)
}

func (r *Reader) VisitNext() {
	r.currentIndex++
}

func (r *Reader) IsWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func (r *Reader) SkipWhitespace() {
	for !r.IsEOF() && r.IsWhitespace(r.Peek()) {
		r.VisitNext()
	}
}
