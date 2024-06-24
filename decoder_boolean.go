package ajson

import "errors"

type BooleanDecoder struct{}

func (decoder *BooleanDecoder) CanDecode(_ Decoder, n JSONNode, reader *Reader, options Options) bool {
	return reader.Peek() == 't' || reader.Peek() == 'f'
}

func (decoder *BooleanDecoder) Decode(_ Decoder, n JSONNode, reader *Reader, options Options) (JSONNode, error) {
	if reader.VisitIfNext('t') {
		if err := reader.Visit('r'); err != nil {
			return nil, err
		}
		if err := reader.Visit('u'); err != nil {
			return nil, err
		}
		if err := reader.Visit('e'); err != nil {
			return nil, err
		}

		if err := reader.PeekEndChar(); err != nil {
			return nil, err
		}

		return &JSONConstant{
			JSONToken: JSONToken{
				NodeKind:   JSONNodeBoolean,
				NodeValue:  "true",
				ParentNode: n,
			},
		}, nil
	}

	if reader.VisitIfNext('f') {
		if err := reader.Visit('a'); err != nil {
			return nil, err
		}
		if err := reader.Visit('l'); err != nil {
			return nil, err
		}
		if err := reader.Visit('s'); err != nil {
			return nil, err
		}
		if err := reader.Visit('e'); err != nil {
			return nil, err
		}

		return &JSONConstant{
			JSONToken: JSONToken{
				NodeKind:  JSONNodeBoolean,
				NodeValue: "false",
			},
		}, nil
	}

	return nil, errors.New("invalid boolean value")
}
