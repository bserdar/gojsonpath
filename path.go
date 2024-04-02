package gojsonpath

import (
	"errors"
	"strconv"
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
	match(*context) bool
	selectChildNodes(*context) (children []Element, canSelect bool)
}

func Find(doc DocModel, path Path) ([]any, error) {
	results := make([]any, 0)
	err := Search(doc, path, func(el []Element) {
		results = append(results, el[len(el)-1].Node)
	})
	return results, err
}

func Search(doc DocModel, path Path, capture func([]Element)) error {
	if len(path.selectors) == 0 {
		return nil
	}
	ctx := &context{
		doc:  doc,
		path: make([]Element, 0, 128),
	}
	el := Element{
		Node: doc.Root(),
	}
	el.Type = doc.Type(el.Node)
	ctx.path = append(ctx.path, el)
	return evaluateAt(ctx, path, 0, capture)
}

func evaluateAt(ctx *context, path Path, pathIx int, capture func([]Element)) error {
	if pathIx >= len(path.selectors) {
		// No selectors. Capture this element
		capture(ctx.path)
		return nil
	}
	if !path.selectors[pathIx].match(ctx) {
		return nil
	}
	// This document node satisfies the selector
	// Move to a child document node, and the next selector
	// Descend
	if !path.recursive[pathIx] {
		// Move to next selector
		pathIx++
		// Is there a selector?
		if pathIx >= len(path.selectors) {
			capture(ctx.path)
			return nil
		}

		// Can this selector select its children?
		childNodeElements, canSelect := path.selectors[pathIx].selectChildNodes(ctx)
		if canSelect {
			for _, childNode := range childNodeElements {
				ctx.path = append(ctx.path, childNode)
				if err := evaluateAt(ctx, path, pathIx, capture); err != nil {
					return err
				}
				ctx.path = ctx.path[:len(ctx.path)-1]
			}
			return nil
		} // if canSelect

		// Selector at pathIx Cannot select children, so we iterate all
		node := ctx.path[len(ctx.path)-1].Node
		switch ctx.doc.Type(node) {
		case ValueNode:
		case ObjectNode:
			for _, key := range ctx.doc.Keys(node) {
				value, _ := ctx.doc.Key(node, key)
				childElem := Element{
					Node: value,
					Type: ctx.doc.Type(value),
					Name: key,
				}
				ctx.path = append(ctx.path, childElem)
				if err := evaluateAt(ctx, path, pathIx, capture); err != nil {
					return err
				}
				ctx.path = ctx.path[:len(ctx.path)-1]
			}
		case ArrayNode:
			n := ctx.doc.Len(node)
			for i := 0; i < n; i++ {
				value := ctx.doc.Elem(node, i)
				index := i
				childElem := Element{
					Node:  value,
					Type:  ctx.doc.Type(value),
					Index: index,
				}
				ctx.path = append(ctx.path, childElem)
				if err := evaluateAt(ctx, path, pathIx, capture); err != nil {
					return err
				}
				ctx.path = ctx.path[:len(ctx.path)-1]
			}
		} // switch doc.Type(node)
		return nil
	} // if !path.recursive[pathIx]

	// Recursive descent
	// Evaluate the rest of the path for every node, recursively
	var descend func() error
	descend = func() error {
		node := ctx.path[len(ctx.path)-1].Node
		switch ctx.doc.Type(node) {
		case ValueNode:
			if err := evaluateAt(ctx, path, pathIx+1, capture); err != nil {
				return err
			}
		case ObjectNode:
			for _, key := range ctx.doc.Keys(node) {
				value, _ := ctx.doc.Key(node, key)
				childElem := Element{
					Node: value,
					Type: ctx.doc.Type(value),
					Name: key,
				}
				ctx.path = append(ctx.path, childElem)
				if err := evaluateAt(ctx, path, pathIx+1, capture); err != nil {
					return err
				}
				if err := descend(); err != nil {
					return err
				}
				ctx.path = ctx.path[:len(ctx.path)-1]
			}
		case ArrayNode:
			n := ctx.doc.Len(node)
			for i := 0; i < n; i++ {
				value := ctx.doc.Elem(node, i)
				index := i
				childElem := Element{
					Node:  value,
					Type:  ctx.doc.Type(value),
					Index: index,
				}
				ctx.path = append(ctx.path, childElem)
				if err := evaluateAt(ctx, path, pathIx+1, capture); err != nil {
					return err
				}
				if err := descend(); err != nil {
					return err
				}
				ctx.path = ctx.path[:len(ctx.path)-1]
			}
		} // switch doc.Type(node)
		return nil
	}
	return descend()
}

type rootElementSelector struct{}

func (rootElementSelector) match(ctx *context) bool {
	if len(ctx.path) == 1 {
		return true
	}
	return ctx.path[len(ctx.path)-1].Node == ctx.doc.Root()
}

func (rootElementSelector) selectChildNodes(*context) (children []Element, canSelect bool) {
	return nil, false
}

type currentElementSelector struct{}

func (currentElementSelector) match(ctx *context) bool {
	return len(ctx.path) != 0
}

func (currentElementSelector) selectChildNodes(*context) (children []Element, canSelect bool) {
	return nil, false
}

type unionSelector struct {
	parts []selector
}

func (sel *unionSelector) match(ctx *context) bool {
	for _, part := range sel.parts {
		if part.match(ctx) {
			return true
		}
	}
	return false
}

func (sel *unionSelector) selectChildNodes(*context) (children []Element, canSelect bool) {
	return nil, false
}

type wildcardSelector struct{}

func (wildcardSelector) match(*context) bool                         { return true }
func (wildcardSelector) selectChildNodes(*context) ([]Element, bool) { return nil, false }

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

func (sel *keySelector) intValue(ctx *context) *int {
	if len(sel.key) != 0 {
		i, err := strconv.Atoi(sel.key)
		if err == nil {
			return &i
		}
	}
	if sel.expr == nil {
		return nil
	}
	val, _ := sel.expr.evaluate(ctx)
	if v, ok := val.value.(int); ok {
		return &v
	}
	if s, ok := val.value.(string); ok {
		i, err := strconv.Atoi(s)
		if err == nil {
			return &i
		}
	}
	return nil
}

func (sel *keySelector) match(ctx *context) bool {
	if isObjectElement(ctx.path) {
		return ctx.path[len(ctx.path)-1].Name == sel.key
	}
	if isArrayElement(ctx.path) {
		if v := sel.intValue(ctx); v != nil {
			if *v < 0 {
				n := ctx.doc.Len(ctx.path[len(ctx.path)-2].Node)
				*v = normalizeIndex(*v, n)
				if *v < 0 {
					return false
				}
			}
			return ctx.path[len(ctx.path)-1].Index == *v
		}
	}
	return false
}

func (sel *keySelector) selectChildNodes(ctx *context) ([]Element, bool) {
	if isObject(ctx.path) {
		node, ok := ctx.doc.Key(ctx.path[len(ctx.path)-1].Node, sel.key)
		if !ok {
			return nil, true
		}
		return []Element{
			{
				Node: node,
				Type: ctx.doc.Type(node),
				Name: sel.key,
			},
		}, true
	}
	if isArray(ctx.path) {
		if v := sel.intValue(ctx); v != nil {
			n := ctx.doc.Len(ctx.path[len(ctx.path)-1].Node)
			index := normalizeIndex(*v, n)
			if index < 0 {
				return nil, true
			}
			if index >= n {
				return nil, true
			}
			node := ctx.doc.Elem(ctx.path[len(ctx.path)-1].Node, index)
			return []Element{
				{
					Node:  node,
					Type:  ctx.doc.Type(node),
					Index: index,
				},
			}, true
		}
	}
	return nil, false
}

type filterSelector struct {
	filter expression
}

func (sel *filterSelector) match(ctx *context) bool {
	v, err := sel.filter.evaluate(ctx)
	if err != nil {
		return false
	}
	return v.asBool()
}

func (sel *filterSelector) selectChildNodes(*context) (children []Element, canSelect bool) {
	return nil, false
}
