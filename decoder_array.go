package ajson

import "fmt"

type ArrayDecoder struct{}

func (decoder *ArrayDecoder) CanDecode(d Decoder, n JSONNode, reader *Reader, options Options) bool {
	return reader.Peek() == '['
}

func (decoder *ArrayDecoder) Decode(d Decoder, n JSONNode, reader *Reader, options Options) (JSONNode, error) {
	if err := reader.Visit('['); err != nil {
		return nil, err
	}

	reader.SkipWhitespace()

	node := &JSONBranch{
		JSONToken: JSONToken{
			NodeKind:   JSONNodeArray,
			ParentNode: n,
		},
		JSONSubNodes: []JSONNode{},
		intent:       options.IntentChar,
	}

	index := 0
	for reader.Peek() != ']' {
		value, err := d.Decode(decoder, node, reader, options)
		if err != nil {
			return nil, err
		}
		tokenizer, _ := value.(JsonTokenizer)
		tokenizer.Token().NodeName = fmt.Sprintf("$%d", index)
		node.JSONSubNodes = append(node.JSONSubNodes, value)
		index++

		reader.SkipWhitespace()

		if reader.Peek() == ']' {
			break
		}

		if err := reader.Visit(','); err != nil {
			return nil, err
		}
		reader.SkipWhitespace()
	}

	if err := reader.Visit(']'); err != nil {
		return nil, err
	}

	return node, nil
}
