package ajson

type JSONDocument struct {
	root JSONNode
}

func (doc *JSONDocument) Kind() JSONNodeKind {
	return JSONNodeRoot
}
func (doc *JSONDocument) Name() string {
	return "{ROOT}"
}
func (doc *JSONDocument) Value() string {
	return ""
}

func (doc *JSONDocument) Encode() (string, error) {
	return doc.root.Encode()
}

func (doc *JSONDocument) String() string {
	return doc.root.String()
}

func (doc *JSONDocument) Depth() int {
	return 0
}

func (doc *JSONDocument) Parent() JSONNode {
	return nil
}
