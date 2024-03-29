package gojsonpath

import (
	"errors"
)

var (
	ErrNoCurrentElement = errors.New("No current element")
	ErrInvalidAST       = errors.New("Invalid AST")
)

// A Path addresses a set of locations on a hierarchical object
type Path struct {
	selectors []selector
	recursive []bool
}

type Element struct {
	Node any
	Type NodeType
	// if previous element is an object, name is set
	Name string
	// If previous element is an array, name is set
	Index int
}

type selector interface {
	match(DocModel, []Element) bool
	selectChildNodes(DocModel, []Element) (children []Element, canSelect bool)
}

func Find(doc DocModel, path Path) ([]any, error) {
	if len(path.selectors) == 0 {
		return nil, nil
	}
	currentPath := make([]Element, 0, 128)
	el := Element{
		Node: doc.Root(),
	}
	el.Type = doc.Type(el.Node)
	currentPath = append(currentPath, el)
	results := make([]any, 0)
	err := evaluateAt(doc, path, 0, currentPath, func(el Element) {
		results = append(results, el.Node)
	})
	return results, err
}

func evaluateAt(doc DocModel, path Path, pathIx int, currentPath []Element, capture func(Element)) error {
	if pathIx >= len(path.selectors) {
		// No selectors. Capture this element
		capture(currentPath[len(currentPath)-1])
		return nil
	}
	if !path.selectors[pathIx].match(doc, currentPath) {
		return nil
	}
	// Cannot descend, complete
	if pathIx >= len(path.recursive) {
		capture(currentPath[len(currentPath)-1])
		return nil
	}

	// Descend
	if !path.recursive[pathIx] {
		// Single descend
		// Is there a selector?
		if pathIx+1 < len(path.selectors) {
			// There is a selector.
			nextSelector := path.selectors[pathIx+1]
			childNodeElements, canSelect := nextSelector.selectChildNodes(doc, currentPath)
			if canSelect {
				for _, childNode := range childNodeElements {
					currentPath = append(currentPath, childNode)
					if err := evaluateAt(doc, path, pathIx+2, currentPath, capture); err != nil {
						return err
					}
					currentPath = currentPath[:len(currentPath)-1]
				}
				return nil
			} // if canSelect
			// Cannot select. Iterate all children
			node := currentPath[len(currentPath)-1].Node
			switch doc.Type(node) {
			case ValueNode:
				// No children, return
			case ObjectNode:
				for _, key := range doc.Keys(node) {
					value, _ := doc.Key(node, key)
					childElem := Element{
						Node: value,
						Type: doc.Type(value),
						Name: key,
					}
					currentPath = append(currentPath, childElem)
					if err := evaluateAt(doc, path, pathIx+1, currentPath, capture); err != nil {
						return err
					}
					currentPath = currentPath[:len(currentPath)-1]
				}
			case ArrayNode:
				n := doc.Len(node)
				for i := 0; i < n; i++ {
					value := doc.Elem(node, i)
					index := i
					childElem := Element{
						Node:  value,
						Type:  doc.Type(value),
						Index: index,
					}
					currentPath = append(currentPath, childElem)
					if err := evaluateAt(doc, path, pathIx+1, currentPath, capture); err != nil {
						return err
					}
					currentPath = currentPath[:len(currentPath)-1]
				}
			} // switch doc.Type(node)
			return nil
		} // pathIx+1 < len(path.selectors)
		// There is no more selectors. This is not supported
		return nil
	} // if !path.recursive[pathIx]

	// Recursive descent
	// Evaluate the rest of the path for every node, recursively
	var descend func() error
	descend = func() error {
		node := currentPath[len(currentPath)-1].Node
		switch doc.Type(node) {
		case ValueNode:
			if err := evaluateAt(doc, path, pathIx+1, currentPath, capture); err != nil {
				return err
			}
		case ObjectNode:
			for _, key := range doc.Keys(node) {
				value, _ := doc.Key(node, key)
				childElem := Element{
					Node: value,
					Type: doc.Type(value),
					Name: key,
				}
				currentPath = append(currentPath, childElem)
				if err := evaluateAt(doc, path, pathIx+1, currentPath, capture); err != nil {
					return err
				}
				if err := descend(); err != nil {
					return err
				}
				currentPath = currentPath[:len(currentPath)-1]
			}
		case ArrayNode:
			n := doc.Len(node)
			for i := 0; i < n; i++ {
				value := doc.Elem(node, i)
				index := i
				childElem := Element{
					Node:  value,
					Type:  doc.Type(value),
					Index: index,
				}
				currentPath = append(currentPath, childElem)
				if err := evaluateAt(doc, path, pathIx+1, currentPath, capture); err != nil {
					return err
				}
				if err := descend(); err != nil {
					return err
				}
				currentPath = currentPath[:len(currentPath)-1]
			}
		} // switch doc.Type(node)
		return nil
	}
	return descend()
}

type rootElementSelector struct{}

func (rootElementSelector) match(doc DocModel, el []Element) bool {
	if len(el) == 1 {
		return true
	}
	return el[len(el)-1].Node == doc.Root()
}

func (rootElementSelector) selectChildNodes(DocModel, []Element) (children []Element, canSelect bool) {
	return nil, false
}

type currentElementSelector struct{}

func (currentElementSelector) match(doc DocModel, el []Element) bool {
	return len(el) != 0
}

func (currentElementSelector) selectChildNodes(DocModel, []Element) (children []Element, canSelect bool) {
	return nil, false
}

type unionSelector struct {
	parts []selector
}

func (sel *unionSelector) match(doc DocModel, el []Element) bool {
	for _, part := range sel.parts {
		if part.match(doc, el) {
			return true
		}
	}
	return false
}

func (sel *unionSelector) selectChildNodes(DocModel, []Element) (children []Element, canSelect bool) {
	return nil, false
}

type wildcardSelector struct{}

func (wildcardSelector) match(DocModel, []Element) bool                         { return true }
func (wildcardSelector) selectChildNodes(DocModel, []Element) ([]Element, bool) { return nil, false }

type keySelector struct {
	key  string
	expr expression
}

// Determines if last element of el is an object element
func isObjectElement(el []Element) bool {
	if len(el) < 2 {
		return false
	}
	return el[len(el)-2].Type == ObjectNode
}

// Determines if last element of el is an object
func isObject(el []Element) bool {
	if len(el) < 1 {
		return false
	}
	return el[len(el)-1].Type == ObjectNode
}

func (sel *keySelector) intValue() *int {
	if sel.expr == nil {
		return nil
	}
	val := sel.expr.evaluate()
	if v, ok := val.value.(int); ok {
		return &v
	}
	return nil
}

func (sel *keySelector) match(doc DocModel, el []Element) bool {
	if isObjectElement(el) {
		return el[len(el)-1].Name == sel.key
	}
	if isArrayElement(el) {
		if v := sel.intValue(); v != nil {
			return el[len(el)-1].Index == *v
		}
	}
	return false
}

func (sel *keySelector) selectChildNodes(doc DocModel, el []Element) ([]Element, bool) {
	if isObject(el) {
		node, ok := doc.Key(el[len(el)-1].Node, sel.key)
		if !ok {
			return nil, true
		}
		return []Element{
			{
				Node: node,
				Type: doc.Type(node),
				Name: sel.key,
			},
		}, true
	}
	if isArray(el) {
		if v := sel.intValue(); v != nil {
			n := doc.Len(el[len(el)-1].Node)
			index := *v
			if index < 0 {
				index = n + index
				if index < 0 {
					return nil, true
				}
			}
			if index >= n {
				return nil, true
			}
			node := doc.Elem(el[len(el)-1].Node, index)
			return []Element{
				{
					Node:  node,
					Type:  doc.Type(node),
					Index: index,
				},
			}, true
		}
	}
	return nil, true
}
