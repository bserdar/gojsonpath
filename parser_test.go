package gojsonpath

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

type regressionSuite struct {
	Queries []regressionQuery `yaml:"queries"`
}

type regressionQuery struct {
	Id        string `yaml:"id"`
	Selector  string `yaml:"selector"`
	Document  any    `yaml:"document"`
	Consensus any    `yaml:"consensus"`
}

func TestParser(t *testing.T) {
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
		_, err := Parse(query.Selector)
		if err != nil {
			fmt.Println("Input", query.Selector)
			fmt.Println("Error", err)
		}
	}
}

func TestMakeKey(t *testing.T) {
	testCases := [][2]string{
		{`abc`, `abc`},
		{`1abc`, `"1abc"`},
		{`a-bc`, `"a-bc"`},
		{`a"bc`, `'a"bc'`},
		{`a"'bc`, `"a\"'bc"`},
	}

	for _, x := range testCases {
		if y := MakeKeyString(x[0]); y != x[1] {
			t.Errorf("Input: %s, expected: %s, output: %s", x[0], x[1], y)
		}
	}
}
