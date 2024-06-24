package ajson_test

import (
	"testing"

	"github.com/ah-its-andy/ajson"
)

func TestArrayDecoder_Decode_EmptyArray(t *testing.T) {
	decoder := ajson.ArrayDecoder{}
	reader := ajson.NewReader("[]")
	result, err := decoder.Decode(GetTestDecoder(), nil, reader, ajson.Options{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(result.(*ajson.JSONBranch).JSONSubNodes) != 0 {
		t.Errorf("Expected empty array, got %v", result)
	}
}

func TestArrayDecoder_Decode_SimpleArray(t *testing.T) {
	decoder := ajson.ArrayDecoder{}
	reader := ajson.NewReader("[1, 2, 3]")
	_, err := decoder.Decode(GetTestDecoder(), nil, reader, ajson.Options{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// Additional checks for the decoded values can be added here
}

func TestArrayDecoder_Decode_MixedArray(t *testing.T) {
	decoder := ajson.ArrayDecoder{}
	reader := ajson.NewReader("[1, \"two\", 3]")
	_, err := decoder.Decode(GetTestDecoder(), nil, reader, ajson.Options{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// Additional checks for the decoded values can be added here
}

func TestArrayDecoder_Decode_InvalidFormat(t *testing.T) {
	decoder := ajson.ArrayDecoder{}
	reader := ajson.NewReader("[1, 2, 3")
	_, err := decoder.Decode(GetTestDecoder(), nil, reader, ajson.Options{})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestArrayDecoder_Decode_NestedArray(t *testing.T) {
	decoder := ajson.ArrayDecoder{}
	reader := ajson.NewReader("[[1, 2], [3, 4]]")
	_, err := decoder.Decode(GetTestDecoder(), nil, reader, ajson.Options{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// Additional checks for the decoded values can be added here
}
