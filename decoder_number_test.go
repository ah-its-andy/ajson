package ajson_test

import (
	"testing"

	"github.com/ah-its-andy/ajson"
)

func TestNumberDecoder_Decode_ValidNumber(t *testing.T) {
	decoder := ajson.NumberDecoder{}
	reader := ajson.NewReader("123")
	_, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestNumberDecoder_Decode_InvalidNumber(t *testing.T) {
	decoder := ajson.NumberDecoder{}
	reader := ajson.NewReader("123abc")
	_, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestNumberDecoder_Decode_EmptyString(t *testing.T) {
	decoder := ajson.NumberDecoder{}
	reader := ajson.NewReader("")
	_, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestNumberDecoder_Decode_EOF(t *testing.T) {
	decoder := ajson.NumberDecoder{}
	reader := ajson.NewReader("")
	reader.VisitNext() // Force EOF
	_, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestNumberDecoder_Decode_NumberWithExponential(t *testing.T) {
	decoder := ajson.NumberDecoder{}
	reader := ajson.NewReader("1.23e+10")
	_, err := decoder.Decode(nil, nil, reader, ajson.Options{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
