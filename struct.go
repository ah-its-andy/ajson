package ajson

import (
	"errors"
	"fmt"
	"strings"
)

type JSONNodeKind int

const (
	JSONNodeObject JSONNodeKind = iota
	JSONNodeArray
	JSONNodeString
	JSONNodeNumber
	JSONNodeBoolean
	JSONNodeNull
	JSONNodeRoot
)

type JSONNode interface {
	Kind() JSONNodeKind
	Name() string
	Value() string
	Depth() int

	Encode() (string, error)
	String() string
}

type JsonTokenizer interface {
	Token() *JSONToken
}

type JSONToken struct {
	NodeKind  JSONNodeKind
	NodeName  string
	NodeValue string

	ParentNode JSONNode
}

func (node *JSONToken) Kind() JSONNodeKind {
	return node.NodeKind
}

func (node *JSONToken) Name() string {
	return node.NodeName
}

func (node *JSONToken) Value() string {
	return node.NodeValue
}

func (node *JSONToken) Parent() JSONNode {
	return node.ParentNode
}

func (node *JSONToken) Depth() int {
	return node.ParentNode.Depth() + 1
}

func (node *JSONToken) Encode() (string, error) {
	return "", errors.New("not implemented")
}

func (node *JSONToken) String() string {
	v, err := node.Encode()
	if err != nil {
		return err.Error()
	}
	return v
}

type JSONConstant struct {
	JSONToken
}

func (node *JSONConstant) Kind() JSONNodeKind {
	return node.JSONToken.NodeKind
}

func (node *JSONConstant) Name() string {
	return node.JSONToken.NodeName
}
func (node *JSONConstant) Value() string {
	return node.JSONToken.NodeValue
}

func (node *JSONConstant) Token() *JSONToken {
	return &node.JSONToken
}

func (node *JSONConstant) Encode() (string, error) {
	builder := strings.Builder{}
	if node.ParentNode != nil && node.ParentNode.Kind() == JSONNodeObject {
		builder.WriteString("\"")
		builder.WriteString(node.NodeName)
		builder.WriteString("\" : ")
	}
	switch node.NodeKind {
	case JSONNodeBoolean, JSONNodeNumber:
		builder.WriteString(node.NodeValue)

	case JSONNodeNull:
		builder.WriteString("null")

	case JSONNodeString:
		builder.WriteString("\"")
		builder.WriteString(node.NodeValue)
		builder.WriteString("\"")

	case JSONNodeObject, JSONNodeArray:
		return "", errors.New("not implemented")
	default:
		return "", fmt.Errorf("unsupported type %v", node.NodeKind)
	}
	return builder.String(), nil
}

func (node *JSONConstant) String() string {
	v, err := node.Encode()
	if err != nil {
		return err.Error()
	}
	return v
}

type JSONBranch struct {
	JSONToken
	JSONSubNodes []JSONNode

	intent string
}

func (node *JSONBranch) Kind() JSONNodeKind {
	return node.JSONToken.NodeKind
}

func (node *JSONBranch) Name() string {
	return node.JSONToken.NodeName
}
func (node *JSONBranch) Value() string {
	return node.JSONToken.NodeValue
}

func (node *JSONBranch) SubNodes() []JSONNode {
	return node.JSONSubNodes
}

func (node *JSONBranch) Token() *JSONToken {
	return &node.JSONToken
}

func (node *JSONBranch) Encode() (string, error) {
	builder := strings.Builder{}
	if len(node.NodeName) > 0 {
		builder.WriteString("\"")
		builder.WriteString(node.NodeName)
		builder.WriteString("\" : ")
	}
	if node.NodeKind == JSONNodeObject {
		builder.WriteString("{\n")
	} else if node.NodeKind == JSONNodeArray {
		builder.WriteString("[\n")
	} else {
		return "", fmt.Errorf("invalid JSON node type %v", node.NodeKind)
	}
	for i, subNode := range node.JSONSubNodes {
		subNodeEncoded, err := subNode.Encode()
		if err != nil {
			return "", err
		}
		builder.WriteString(node.genIntent(subNode.Depth()))
		builder.WriteString(subNodeEncoded)
		if i != len(node.JSONSubNodes)-1 {
			builder.WriteString(",\n")
		}
	}
	if node.NodeKind == JSONNodeObject {
		builder.WriteString("}")
	} else if node.NodeKind == JSONNodeArray {
		builder.WriteString("]")
	}
	if len(node.NodeName) > 0 {
		builder.WriteString("\n")
	}
	return builder.String(), nil
}

func (node *JSONBranch) String() string {
	v, err := node.Encode()
	if err != nil {
		return err.Error()
	}
	return v
}

func (node *JSONBranch) genIntent(depth int) string {
	if node.intent == "" || len(node.intent) == 0 {
		return ""
	}

	intentBuilder := strings.Builder{}
	for i := 0; i < depth; i++ {
		intentBuilder.WriteString(node.intent)
	}
	return intentBuilder.String()
}
