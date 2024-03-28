package gojsonpath

import (
	"encoding/json"
	"testing"
)

func parseAndFind(in string, doc string) ([]any, error) {
	path, err := Parse(in)
	if err != nil {
		return nil, err
	}
	var v any
	if err := json.Unmarshal([]byte(doc), &v); err != nil {
		return nil, err
	}
	model := MapModel{
		Doc: v,
	}
	return Find(model, path)
}

func TestPathEval(t *testing.T) {
	result, err := parseAndFind(`$.*`, `{"a":1,"b":2,"c": {"d":3}}`)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
