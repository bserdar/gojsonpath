package gojsonpath

var predefinedFunctions = map[string]func([]any) (any, error){
	"length": lengthFunc,
}

func lengthFunc(in []any) (any, error) {
	if len(in) == 0 {
		return 0, nil
	}
	if x, ok := in[0].([]any); ok {
		return len(x), nil
	}
	return 0, nil
}
