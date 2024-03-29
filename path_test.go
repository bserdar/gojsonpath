package gojsonpath

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
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

func TestEvaluator(t *testing.T) {
	data, err := os.ReadFile("testdata/regression_suite.yaml")
	if err != nil {
		panic(err)
	}
	var suite regressionSuite
	if err := yaml.Unmarshal(data, &suite); err != nil {
		t.Error(err)
		return
	}
	for _, query := range suite.Queries {
		if query.Consensus == "NOT_SUPPORTED" {
			continue
		}
		if query.Consensus == nil {
			continue
		}

		path, err := Parse(query.Selector)
		if err != nil {
			t.Error(err)
			continue
		}
		model := MapModel{
			Doc: query.Document,
		}
		t.Log(query.Selector)
		result, err := Find(model, path)
		if err != nil {
			t.Error(err)
		}
		t.Log(query.Selector, "ok")
		consensus, ok := query.Consensus.([]any)
		if !ok {
			t.Log("Not an array:", query.Selector)
			continue
		}
		for _, c := range consensus {
			found := false
			for i := range result {
				if reflect.DeepEqual(result[i], c) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("In %s: Not found %v--%v", query.Selector, c, result)
			}
		}
	}
}
