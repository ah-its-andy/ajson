package ajson_test

import (
	"testing"

	"github.com/ah-its-andy/ajson"
)

func TestObjectDecoder_Decode_ValidJSONObject(t *testing.T) {
	decoder := ajson.ObjectDecoder{FieldDecoder: &ajson.StringDecoder{}}
	reader := ajson.NewReader(`{"key": "value"}`)
	node, err := decoder.Decode(GetTestDecoder(), nil, reader, ajson.Options{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if node == nil {
		t.Fatalf("Expected a node, got nil")
	}
}

func TestObjectDecoder_Decode_InvalidJSONObject(t *testing.T) {
	decoder := ajson.ObjectDecoder{FieldDecoder: &ajson.StringDecoder{}}
	reader := ajson.NewReader(`{"key": "value"`)
	_, err := decoder.Decode(GetTestDecoder(), nil, reader, ajson.Options{})
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}
}

func TestObjectDecoder_Decode_EmptyJSONObject(t *testing.T) {
	decoder := ajson.ObjectDecoder{FieldDecoder: &ajson.StringDecoder{}}
	reader := ajson.NewReader(`{}`)
	node, err := decoder.Decode(GetTestDecoder(), nil, reader, ajson.Options{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if node == nil {
		t.Fatalf("Expected a node, got nil")
	}
}

func TestObjectDecoder_Decode_NestedJSONObject(t *testing.T) {
	decoder := ajson.ObjectDecoder{FieldDecoder: &ajson.StringDecoder{}}
	reader := ajson.NewReader(`{"nested": {"key": "value"}}`)
	node, err := decoder.Decode(GetTestDecoder(), nil, reader, ajson.Options{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if node == nil {
		t.Fatalf("Expected a node, got nil")
	}
}

func TestObjectDecoder_Decode_JSONObjectWithVariousTypes(t *testing.T) {
	decoder := ajson.ObjectDecoder{FieldDecoder: &ajson.StringDecoder{}}
	reader := ajson.NewReader(`{"string": "value", "number": 123, "boolean": true}`)
	node, err := decoder.Decode(GetTestDecoder(), nil, reader, ajson.Options{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if node == nil {
		t.Fatalf("Expected a node, got nil")
	}
}
