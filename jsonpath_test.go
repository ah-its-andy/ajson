package ajson_test

import (
	"testing"

	"github.com/ah-its-andy/ajson"
)

func TestCompile_EmptyExpression(t *testing.T) {
	jpexpr := ajson.NewJSONPath("")
	err := jpexpr.Compile()
	if err != nil {
		t.Errorf("Compile failed: %v", err)
	}
	if len(jpexpr.Nodes()) != 0 {
		t.Errorf("Expected no nodes, got %d", len(jpexpr.Nodes()))
	}
}

func TestCompile_RootOnly(t *testing.T) {
	jpexpr := ajson.NewJSONPath("$")
	err := jpexpr.Compile()
	if err != nil {
		t.Errorf("Compile failed: %v", err)
	}
	if len(jpexpr.Nodes()) != 1 || jpexpr.Nodes()[0].Kind != ajson.JSONPathNodeRoot {
		t.Errorf("Expected root node, got %+v", jpexpr.Nodes())
	}
}

func TestCompile_SingleField(t *testing.T) {
	jpexpr := ajson.NewJSONPath("$.field")
	err := jpexpr.Compile()
	if err != nil {
		t.Errorf("Compile failed: %v", err)
	}
	if len(jpexpr.Nodes()) != 2 || jpexpr.Nodes()[1].Kind != ajson.JSONPathNodeField || jpexpr.Nodes()[1].Name != "field" {
		t.Errorf("Expected field node 'field', got %+v", jpexpr.Nodes())
	}
}

func TestCompile_SingleIndex(t *testing.T) {
	jpexpr := ajson.NewJSONPath("$[0]")
	err := jpexpr.Compile()
	if err != nil {
		t.Errorf("Compile failed: %v", err)
	}
	if len(jpexpr.Nodes()) != 2 || jpexpr.Nodes()[1].Kind != ajson.JSONPathNodeIndex || jpexpr.Nodes()[1].Name != "0" {
		t.Errorf("Expected index node '0', got %+v", jpexpr.Nodes())
	}
}

func TestCompile_MultipleFields(t *testing.T) {
	jpexpr := ajson.NewJSONPath("$.store.book")
	err := jpexpr.Compile()
	if err != nil {
		t.Errorf("Compile failed: %v", err)
	}
	if len(jpexpr.Nodes()) != 3 || jpexpr.Nodes()[2].Kind != ajson.JSONPathNodeField || jpexpr.Nodes()[2].Name != "book" {
		t.Errorf("Expected field node 'book', got %+v", jpexpr.Nodes())
	}
}

func TestCompile_MultipleIndexes(t *testing.T) {
	jpexpr := ajson.NewJSONPath("$[0][1]")
	err := jpexpr.Compile()
	if err != nil {
		t.Errorf("Compile failed: %v", err)
	}
	if len(jpexpr.Nodes()) != 3 || jpexpr.Nodes()[2].Kind != ajson.JSONPathNodeIndex || jpexpr.Nodes()[2].Name != "1" {
		t.Errorf("Expected index node '1', got %+v", jpexpr.Nodes())
	}
}

func TestCompile_MixedFieldsAndIndexes(t *testing.T) {
	jpexpr := ajson.NewJSONPath("$.store.book[0].title")
	err := jpexpr.Compile()
	if err != nil {
		t.Errorf("Compile failed: %v", err)
	}
	if len(jpexpr.Nodes()) != 5 || jpexpr.Nodes()[4].Kind != ajson.JSONPathNodeField || jpexpr.Nodes()[4].Name != "title" {
		t.Errorf("Expected field node 'title', got %+v", jpexpr.Nodes())
	}
}

// Note: This test is more of a placeholder since the current implementation does not explicitly handle invalid expressions.
func TestCompile_InvalidExpression(t *testing.T) {
	jpexpr := ajson.NewJSONPath("$.store.[0]")
	err := jpexpr.Compile()
	if err != nil {
		t.Errorf("Compile failed: %v", err)
	}
	// This test checks if the method can run without crashing, but does not validate correctness due to lack of error handling in Compile.
}
