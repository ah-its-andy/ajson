package ajson

type JSONPathNodeKind int

const (
	JSONPathNodeRoot JSONPathNodeKind = iota
	JSONPathNodeField
	JSONPathNodeIndex
)

type JSONPathNode struct {
	Kind JSONPathNodeKind
	Name string
	Next *JSONPathNode
}

type JSONPathExpression struct {
	expr  string
	nodes []*JSONPathNode
}

func NewJSONPath(jsonpath string) *JSONPathExpression {
	return &JSONPathExpression{
		expr: jsonpath,
	}
}

func (jpexpr *JSONPathExpression) Nodes() []*JSONPathNode {
	return jpexpr.nodes
}

func (jpexpr *JSONPathExpression) Compile() error {
	// Split the expression into individual nodes
	nodes := make([]*JSONPathNode, 0)
	if len(jpexpr.expr) == 0 {
		return nil
	}

	exprLen := len(jpexpr.expr)
	currentIndex := 0
	for currentIndex < exprLen {
		currentChar := jpexpr.expr[currentIndex]
		if currentChar == '$' {
			node := &JSONPathNode{
				Kind: JSONPathNodeRoot,
				Name: "$",
			}
			nodes = append(nodes, node)
			currentIndex++
			continue
		} else if currentChar == '[' {
			node := &JSONPathNode{
				Kind: JSONPathNodeIndex,
			}
			nameChars := make([]byte, 0)
			currentIndex++
			for currentIndex < exprLen && jpexpr.expr[currentIndex] != ']' {
				nameChars = append(nameChars, jpexpr.expr[currentIndex])
				currentIndex++
			}
			currentIndex++
			node.Name = string(nameChars)
			nodes = append(nodes, node)
			continue
		} else if currentChar == '.' {
			node := &JSONPathNode{
				Kind: JSONPathNodeField,
			}
			nameChars := make([]byte, 0)
			currentIndex++
			for currentIndex < exprLen && jpexpr.expr[currentIndex] != '[' && jpexpr.expr[currentIndex] != '.' {
				nameChars = append(nameChars, jpexpr.expr[currentIndex])
				currentIndex++
			}
			node.Name = string(nameChars)
			nodes = append(nodes, node)
			continue
		} else {
			currentIndex++
		}
	}

	jpexpr.nodes = nodes

	return nil
}
