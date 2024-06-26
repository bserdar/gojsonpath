package gojsonpath

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestSimpleParser(t *testing.T) {
	type testCase struct {
		input    string
		path     string
		expected string
	}

	tests := []testCase{
		{
			input:    `[1,2,3]`,
			path:     `/1`,
			expected: `[2]`,
		},
		{
			input:    `{"a":[1,2,3]}`,
			path:     `/a/1`,
			expected: `[2]`,
		},
		{
			input:    `{"a":{"b":[1,2,3]}}`,
			path:     `/a/b/2`,
			expected: `[3]`,
		},
		{
			input:    `{"a":{"b":[1,2,3]}}`,
			path:     `/a/*/2`,
			expected: `[3]`,
		},
		{
			input:    `{"a":{"b":[1,2,3], "c":[2,3,4]}}`,
			path:     `/a/*/2`,
			expected: `[3,4]`,
		},
	}

	for _, tc := range tests {
		var in any
		err := json.Unmarshal([]byte(tc.input), &in)
		if err != nil {
			panic(fmt.Sprintf("%s: %s", tc.input, err))
		}

		path, err := Parse(tc.path)
		if err != nil {
			panic(fmt.Sprintf("%s: %s", tc.path, err))
		}
		var expected []any
		err = json.Unmarshal([]byte(tc.expected), &expected)
		if err != nil {
			panic(fmt.Sprintf("%s: %s", tc.expected, err))
		}

		result, err := Find(MapModel{Doc: in}, path)
		if err != nil {
			t.Errorf("%s: %s", tc.input, err)
			continue
		}
		for _, c := range expected {
			found := false
			for i := range result {
				if reflect.DeepEqual(result[i], c) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Not found %v--%v", expected, result)
			}
		}
	}
}
