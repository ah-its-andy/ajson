package ajson

import "errors"

type StringDecoder struct{}

func (decoder *StringDecoder) CanDecode(_ Decoder, n JSONNode, reader *Reader, options Options) bool {
	return reader.Peek() == '"'
}
func (decoder *StringDecoder) Decode(_ Decoder, n JSONNode, reader *Reader, options Options) (JSONNode, error) {
	if err := reader.Visit('"'); err != nil {
		return nil, err
	}

	value := []byte{}
	for !reader.IsEOF() && reader.Peek() != '"' {
		value = append(value, reader.Peek())
		reader.VisitNext()
	}

	if reader.IsEOF() {
		return nil, errors.New("unexpected end of input")
	}

	if err := reader.Visit('"'); err != nil {
		return nil, err
	}

	return &JSONConstant{
		JSONToken: JSONToken{
			NodeKind:   JSONNodeString,
			NodeValue:  string(value),
			ParentNode: n,
		},
	}, nil
}
