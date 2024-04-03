package gojsonpath

// DocPath is a list of nodes of the document
type DocPath struct {
	P []Segment
}

// A Segment is a node in a document path
type Segment struct {
	Node any
	Type NodeType
	// if previous element is an object, name is set
	Name string
	// If previous element is an array, name is set
	Index int
}

// Add a key-value to the path. The value is of type valueType. Returns the new path
func (path DocPath) Key(key string, value any, valueType NodeType) DocPath {
	path.P = append(path.P, Segment{
		Node: value,
		Name: key,
		Type: valueType,
	})
	return path
}

// Add an array index-value to the path. The value is of type valueType. Returns the new path
func (path DocPath) Index(index int, value any, valueType NodeType) DocPath {
	path.P = append(path.P, Segment{
		Node:  value,
		Index: index,
		Type:  valueType,
	})
	return path
}

func (path DocPath) Push(seg Segment) DocPath {
	path.P = append(path.P, seg)
	return path
}

// Returns a new path by removing the last element from the path
func (path DocPath) Pop() DocPath {
	path.P = path.P[:len(path.P)-1]
	return path
}

// Last returns the last element in the path
func (path DocPath) Last() Segment {
	return path.P[len(path.P)-1]
}

func (path DocPath) Len() int { return len(path.P) }
