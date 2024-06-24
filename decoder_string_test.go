package ajson_test

import (
	"strings"
	"testing"

	"github.com/ah-its-andy/ajson"
)

func TestDecodeString_ValidString(t *testing.T) {
	reader := ajson.NewReader("\"hello\"")
	decoder := ajson.StringDecoder{}
	node, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if node.Value() != "hello" {
		t.Errorf("Expected 'hello', got '%s'", node.Value())
	}
}

func TestDecodeString_EmptyString(t *testing.T) {
	reader := ajson.NewReader("\"\"")
	decoder := ajson.StringDecoder{}
	node, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if node.Value() != "" {
		t.Errorf("Expected empty string, got '%s'", node.Value())
	}
}

func TestDecodeString_UnexpectedEOF(t *testing.T) {
	reader := ajson.NewReader("\"hello")
	decoder := ajson.StringDecoder{}
	_, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err == nil || err.Error() != ajson.ErrUnexpectedEndOfInput {
		t.Fatalf("Expected 'unexpected end of input' error, got %v", err)
	}
}

func TestDecodeString_SpecialCharacters(t *testing.T) {
	reader := ajson.NewReader("\"hello\\nworld\"")
	decoder := ajson.StringDecoder{}
	node, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	// This test might fail due to the current implementation not handling escape sequences
	if !strings.EqualFold(node.Value(), "hello\\nworld") {
		t.Errorf("Expected 'hello\\nworld', got '%s'", node.Value())
	}
}

func TestDecodeString_WithQuotes(t *testing.T) {
	reader := ajson.NewReader("\"hello \\\"world\\\"\"")
	decoder := ajson.StringDecoder{}
	node, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	// This test might fail due to the current implementation not handling escape sequences
	if node.Value() != "hello \"world\"" {
		t.Errorf("Expected 'hello \"world\"', got '%s'", node.Value())
	}
}
