package ajson

type ObjectDecoder struct {
	FieldDecoder Decoder
}

func (decoder *ObjectDecoder) CanDecode(_ Decoder, n JSONNode, reader *Reader, options Options) bool {
	return reader.Peek() == '{'
}

func (decoder *ObjectDecoder) Decode(d Decoder, n JSONNode, reader *Reader, options Options) (JSONNode, error) {
	if err := reader.Visit('{'); err != nil {
		return nil, err
	}
	reader.SkipWhitespace()

	node := &JSONBranch{
		JSONToken: JSONToken{
			NodeKind:   JSONNodeObject,
			ParentNode: n,
		},
		JSONSubNodes: []JSONNode{},
		intent:       options.IntentChar,
	}

	for reader.Peek() != '}' {
		field, err := decoder.FieldDecoder.Decode(d, node, reader, options)
		if err != nil {
			return nil, err
		}

		reader.SkipWhitespace()
		if err := reader.Visit(':'); err != nil {
			return nil, err
		}
		reader.SkipWhitespace()

		value, err := d.Decode(decoder, node, reader, options)
		if err != nil {
			return nil, err
		}
		tokenizer, _ := value.(JsonTokenizer)

		tokenizer.Token().NodeName = field.Value()
		node.JSONSubNodes = append(node.JSONSubNodes, value)

		reader.SkipWhitespace()

		if reader.Peek() == '}' {
			break
		}

		if err := reader.Visit(','); err != nil {
			return nil, err
		}
		reader.SkipWhitespace()
	}

	if err := reader.Visit('}'); err != nil {
		return nil, err
	}

	return node, nil
}
