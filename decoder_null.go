package ajson

type NullDecoder struct{}

func (decoder *NullDecoder) CanDecode(_ Decoder, n JSONNode, reader *Reader, options Options) bool {
	return reader.Peek() == 'n'
}

func (decoder *NullDecoder) Decode(_ Decoder, n JSONNode, reader *Reader, options Options) (JSONNode, error) {
	if err := reader.Visit('n'); err != nil {
		return nil, err
	}
	if err := reader.Visit('u'); err != nil {
		return nil, err
	}
	if err := reader.Visit('l'); err != nil {
		return nil, err
	}
	if err := reader.Visit('l'); err != nil {
		return nil, err
	}

	return &JSONConstant{
		JSONToken: JSONToken{
			NodeKind:   JSONNodeNull,
			ParentNode: n,
		},
	}, nil
}
