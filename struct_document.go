package ajson

import "strconv"

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

func (doc *JSONDocument) GetNode(path string) (JSONNode, error) {
	if len(path) == 0 {
		return nil, nil
	}
	jsonPath := NewJSONPath(path)
	if err := jsonPath.Compile(); err != nil {
		return nil, err
	}
	node := doc.GetNodeByJSONPath(jsonPath)
	return node, nil
}

func (doc *JSONDocument) GetNodeByJSONPath(expr *JSONPathExpression) JSONNode {
	currentNode := doc.root
	for _, pathNode := range expr.Nodes() {
		if pathNode.Kind == JSONPathNodeRoot {
			continue
		}
		if pathNode.Kind == JSONPathNodeField {
			if branchNode, ok := currentNode.(*JSONBranch); !ok {
				return nil
			} else {
				hits := false
				for _, subNode := range branchNode.SubNodes() {
					if subNode.Name() == pathNode.Name {
						currentNode = subNode
						hits = true
						break
					}
				}
				if !hits {
					return nil
				}
			}
		} else if pathNode.Kind == JSONPathNodeIndex {
			if branchNode, ok := currentNode.(*JSONBranch); !ok {
				return nil
			} else {
				hits := false
				for i, subNode := range branchNode.SubNodes() {
					index := strconv.FormatInt(int64(i), 10)
					if index == pathNode.Name {
						currentNode = subNode
						hits = true
						break
					}
				}
				if !hits {
					return nil
				}
			}
		}
	}
	return currentNode
}
