package ajson_test

import (
	"testing"

	"github.com/ah-its-andy/ajson"
)

func TestBooleanDecoder_Decode_ValidTrue(t *testing.T) {
	decoder := ajson.BooleanDecoder{}
	reader := ajson.NewReader("true")
	node, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if node == nil || node.String() != "true" {
		t.Errorf("Expected true, got %v", node)
	}
}

func TestBooleanDecoder_Decode_ValidFalse(t *testing.T) {
	decoder := ajson.BooleanDecoder{}
	reader := ajson.NewReader("false")
	node, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if node == nil || node.String() != "false" {
		t.Errorf("Expected false, got %v", node)
	}
}

func TestBooleanDecoder_Decode_InvalidBoolean(t *testing.T) {
	decoder := ajson.BooleanDecoder{}
	reader := ajson.NewReader("tru")
	_, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestBooleanDecoder_Decode_InvalidBoolean2(t *testing.T) {
	decoder := ajson.BooleanDecoder{}
	reader := ajson.NewReader("true123")
	_, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestBooleanDecoder_Decode_EmptyString(t *testing.T) {
	decoder := ajson.BooleanDecoder{}
	reader := ajson.NewReader("")
	_, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}
