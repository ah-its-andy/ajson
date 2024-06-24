package ajson_test

import (
	"testing"

	"github.com/ah-its-andy/ajson"
)

func TestDecode(t *testing.T) {
	caseStr := `{
		"name": "root",
		"value": {
			"key1": "value1",
			"key2": {
				"subkey1": "subvalue1",
				"subkey2": ["subelement1", "subelement2"]
			},
			"key3": ["element1", "element2"]
		}
	}`
	decoder := ajson.NewJSONDecoder(
		&ajson.ObjectDecoder{
			FieldDecoder: &ajson.StringDecoder{},
		},
		&ajson.ArrayDecoder{},
		&ajson.StringDecoder{},
		&ajson.NumberDecoder{},
		&ajson.BooleanDecoder{},
		&ajson.NullDecoder{})

	doc := &ajson.JSONDocument{}
	root, err := decoder.Decode(nil, doc, ajson.NewReader(caseStr), *ajson.DefaultOptions)
	if err != nil {
		t.Error(err)
	}
	rootJSN, _ := root.Encode()
	t.Log(rootJSN)
}
