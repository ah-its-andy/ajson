package ajson_test

import (
	"testing"

	"github.com/ah-its-andy/ajson"
)

// Helper function to create a JSONDocument with a known structure
func createTestJSONDocument(t *testing.T) *ajson.JSONDocument {
	// Assuming there's a way to create a JSONDocument from a string for simplicity
	// This part is pseudo-code and needs to be replaced with actual implementation
	doc, err := ajson.Parse(`{"root": {"field1": "value1", "array": [1, 2, 3]}}`)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	return doc
}

func TestGetNode_RootNode(t *testing.T) {
	doc := createTestJSONDocument(t)
	node, err := doc.GetNode("$")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if node.Name() != "" {
		t.Errorf("Expected root node, got %s", node.Name())
	}
}

func TestGetNode_ExistingField(t *testing.T) {
	doc := createTestJSONDocument(t)
	node, err := doc.GetNode("$.root.field1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if node.Value() != "value1" {
		t.Errorf("Expected value 'value1', got '%s'", node.Value())
	}
}

func TestGetNode_NonExistingField(t *testing.T) {
	doc := createTestJSONDocument(t)
	node, err := doc.GetNode("$.root.nonExisting")
	if err != nil {
		t.Fatalf("Expected an error, got nil")
	}
	if node != nil {
		t.Errorf("Expected nil node, got %v", node)
	}
}

func TestGetNode_ExistingIndex(t *testing.T) {
	doc := createTestJSONDocument(t)
	node, err := doc.GetNode("$.root.array[1]")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if node.Value() != "2" {
		t.Errorf("Expected value '2', got '%s'", node.Value())
	}
}

func TestGetNode_NonExistingIndex(t *testing.T) {
	doc := createTestJSONDocument(t)
	node, err := doc.GetNode("$.root.array[5]")
	if err != nil {
		t.Fatalf("Expected an error, got nil")
	}
	if node != nil {
		t.Errorf("Expected nil node, got %v", node)
	}
}

func TestGetNode_InvalidPathSyntax(t *testing.T) {
	doc := createTestJSONDocument(t)
	_, err := doc.GetNode("$.root.[?@]")
	if err == nil {
		t.Errorf("Expected an error for invalid path syntax")
	}
}

func TestGetNode_EmptyPath(t *testing.T) {
	doc := createTestJSONDocument(t)
	node, err := doc.GetNode("")
	if err != nil {
		t.Errorf("Expected no error for empty path")
	}
	if node != nil {
		t.Errorf("Expected nil for empty path")
	}
}
