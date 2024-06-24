package ajson

import (
	"errors"
)

type JSONParser struct {
	// decoders []Decoder
	options Options
}

func NewJSONParser(opts ...Options) *JSONParser {
	var options Options
	if len(opts) == 0 {
		options = opts[0]
	} else {
		options = *DefaultOptions
	}
	return &JSONParser{
		options: options,
	}
}

func (p *JSONParser) Parse(data string) (*JSONDocument, error) {
	r := &Reader{
		data:    data,
		dataLen: len(data),
	}

	decoder := NewJSONDecoder(p.options.Decoders...)

	doc := &JSONDocument{}

	r.SkipWhitespace()

	node, err := decoder.Decode(nil, doc, r, p.options)
	if err != nil {
		return nil, err
	}

	r.SkipWhitespace()

	if !r.IsEOF() {
		return nil, errors.New(ErrUnexpectedTrailingToken)
	}
	doc.root = node
	return doc, nil
}
