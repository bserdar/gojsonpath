package gojsonpath

type NodeType int

const (
	ValueNode NodeType = iota
	ObjectNode
	ArrayNode
)

// DocModel represents the underlying hierarchical document model.
type DocModel interface {
	// Return the root node
	Root() any
	// Return the type of the given node
	Type(any) NodeType

	// Works on object and array nodes
	Len(any) int
	// Lookup a key in an object node
	Key(any, string) (any, bool)
	// Keys of an object node
	Keys(any) []string
	// Lookup and item in an array node
	Elem(any, int) any
	// The input is a value node. Return the value of that node
	Value(any) any
}

// map[string]any model
type MapModel struct {
	Doc any
}

func (m MapModel) Root() any { return m.Doc }

func (MapModel) Type(in any) NodeType {
	if _, ok := in.(map[string]any); ok {
		return ObjectNode
	}
	if _, ok := in.([]any); ok {
		return ArrayNode
	}
	return ValueNode
}

func (MapModel) Len(in any) int {
	if arr, ok := in.([]any); ok {
		return len(arr)
	}
	if obj, ok := in.(map[string]any); ok {
		return len(obj)
	}
	if in == nil {
		return 0
	}
	return 1
}

func (MapModel) Key(in any, key string) (any, bool) {
	obj, ok := in.(map[string]any)
	if !ok {
		return nil, false
	}
	return obj[key], true
}

func (MapModel) Keys(in any) []string {
	obj, ok := in.(map[string]any)
	if !ok {
		return nil
	}
	out := make([]string, 0, len(obj))
	for k := range obj {
		out = append(out, k)
	}
	return out
}

func (MapModel) Elem(in any, index int) any {
	arr, ok := in.([]any)
	if !ok {
		return nil
	}
	return arr[index]
}

func (MapModel) Value(in any) any { return in }
