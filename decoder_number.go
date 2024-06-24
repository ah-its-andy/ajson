package ajson

import (
	"errors"
)

type NumberDecoder struct{}

func (decoder *NumberDecoder) CanDecode(_ Decoder, n JSONNode, reader *Reader, options Options) bool {
	return decoder.isNumberChar(reader.Peek())
}

func (decoder *NumberDecoder) Decode(_ Decoder, n JSONNode, reader *Reader, options Options) (JSONNode, error) {
	value := []byte{}
	for !reader.IsEOF() && decoder.isNumberChar(reader.Peek()) {
		value = append(value, reader.Peek())
		reader.VisitNext()
	}

	if reader.IsEOF() {
		return nil, errors.New("unexpected end of input")
	}

	return &JSONConstant{
		JSONToken: JSONToken{
			NodeKind:   JSONNodeNumber,
			NodeValue:  string(value),
			ParentNode: n,
		},
	}, nil
}

func (decoder *NumberDecoder) isNumberChar(ch byte) bool {
	return (ch >= '0' && ch <= '9') || ch == '-' || ch == '+' || ch == '.' || ch == 'e' || ch == 'E'
}
