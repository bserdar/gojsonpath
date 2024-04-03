package gojsonpath

import (
	"errors"
	"strconv"
)

var (
	ErrNoCurrentElement = errors.New("No current element")
	ErrInvalidAST       = errors.New("Invalid AST")
)

// A Path is a compiled JSONPath object.
type Path struct {
	selectors []selector
	recursive []bool
}

// Append paths one after the other
func Append(p ...Path) Path {
	ret := Path{}
	for _, x := range p {
		ret.selectors = append(ret.selectors, x.selectors...)
		ret.recursive = append(ret.recursive, x.recursive...)
	}
	return ret
}

// Common path definitions
var (
	SelectRootPath = Path{
		selectors: []selector{rootElementSelector{}},
		recursive: []bool{false},
	}
	SelectCurrentPath = Path{
		selectors: []selector{currentElementSelector{}},
		recursive: []bool{false},
	}
	WildcardPath = Path{
		selectors: []selector{wildcardSelector{}},
		recursive: []bool{false},
	}
)

// Returns a new path that is a copy of p but recursively descends
func RecursiveDescent(p Path) Path {
	ret := Path{
		selectors: p.selectors,
		recursive: make([]bool, len(p.recursive)),
	}
	copy(ret.recursive, p.recursive)
	ret.recursive[len(ret.recursive)-1] = true
	return ret
}

// Returns a path that selects a key
func KeySelectorPath(key string) Path {
	return Path{
		selectors: []selector{&keySelector{key: key}},
		recursive: []bool{false},
	}
}

// Returns a path that selects an index
func IndexSelectorPath(index int) Path {
	return Path{
		selectors: []selector{&keySelector{expr: exprValue{value: index}}},
		recursive: []bool{false},
	}
}

type selector interface {
	match(*context) bool
	selectChildNodes(*context) (children iterator, canSelect bool)
}

// Finds all document elements addressed by the path
func Find(doc DocModel, path Path) ([]any, error) {
	results := make([]any, 0)
	err := Search(doc, path, func(docpath DocPath) {
		results = append(results, docpath.Last().Node)
	})
	return results, err
}

// Search iterates all document nodes depth-first, and calls `capture`
// for those document nodes that `path` matches.
func Search(doc DocModel, path Path, capture func(DocPath)) error {
	if len(path.selectors) == 0 {
		return nil
	}
	ctx := &context{
		doc: doc,
		path: DocPath{
			P: make([]Segment, 0, 128),
		},
	}
	el := Segment{
		Node: doc.Root(),
	}
	el.Type = doc.Type(el.Node)
	ctx.path.P = append(ctx.path.P, el)
	return evaluateAt(ctx, path, 0, capture)
}

func evaluateAt(ctx *context, path Path, pathIx int, capture func(DocPath)) error {
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
			for childNodeElements.next() {
				var childNode Segment
				childNodeElements.item(&childNode)
				ctx.path = ctx.path.Push(childNode)
				if err := evaluateAt(ctx, path, pathIx, capture); err != nil {
					return err
				}
				ctx.path = ctx.path.Pop()
			}
			return nil
		} // if canSelect

		// Selector at pathIx Cannot select children, so we iterate all
		node := ctx.path.Last().Node
		switch ctx.doc.Type(node) {
		case ValueNode:
		case ObjectNode:
			for _, key := range ctx.doc.Keys(node) {
				value, _ := ctx.doc.Key(node, key)
				ctx.path = ctx.path.Key(key, value, ctx.doc.Type(value))
				if err := evaluateAt(ctx, path, pathIx, capture); err != nil {
					return err
				}
				ctx.path = ctx.path.Pop()
			}
		case ArrayNode:
			n := ctx.doc.Len(node)
			for i := 0; i < n; i++ {
				value := ctx.doc.Elem(node, i)
				index := i
				ctx.path = ctx.path.Index(index, value, ctx.doc.Type(value))
				if err := evaluateAt(ctx, path, pathIx, capture); err != nil {
					return err
				}
				ctx.path = ctx.path.Pop()
			}
		} // switch doc.Type(node)
		return nil
	} // if !path.recursive[pathIx]

	// Recursive descent
	// Evaluate the rest of the path for every node, recursively
	var descend func() error
	descend = func() error {
		node := ctx.path.Last().Node
		switch ctx.doc.Type(node) {
		case ValueNode:
			if err := evaluateAt(ctx, path, pathIx+1, capture); err != nil {
				return err
			}
		case ObjectNode:
			for _, key := range ctx.doc.Keys(node) {
				value, _ := ctx.doc.Key(node, key)
				ctx.path = ctx.path.Key(key, value, ctx.doc.Type(value))
				if err := evaluateAt(ctx, path, pathIx+1, capture); err != nil {
					return err
				}
				if err := descend(); err != nil {
					return err
				}
				ctx.path = ctx.path.Pop()
			}
		case ArrayNode:
			n := ctx.doc.Len(node)
			for i := 0; i < n; i++ {
				value := ctx.doc.Elem(node, i)
				index := i
				ctx.path = ctx.path.Index(index, value, ctx.doc.Type(value))
				if err := evaluateAt(ctx, path, pathIx+1, capture); err != nil {
					return err
				}
				if err := descend(); err != nil {
					return err
				}
				ctx.path = ctx.path.Pop()
			}
		} // switch doc.Type(node)
		return nil
	}
	return descend()
}

type rootElementSelector struct{}

func (rootElementSelector) match(ctx *context) bool {
	if ctx.path.Len() == 1 {
		return true
	}
	return ctx.path.Last().Node == ctx.doc.Root()
}

func (rootElementSelector) selectChildNodes(*context) (children iterator, canSelect bool) {
	return nil, false
}

type currentElementSelector struct{}

func (currentElementSelector) match(ctx *context) bool {
	return ctx.path.Len() != 0
}

func (currentElementSelector) selectChildNodes(*context) (children iterator, canSelect bool) {
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

func (sel *unionSelector) selectChildNodes(*context) (children iterator, canSelect bool) {
	return nil, false
}

type wildcardSelector struct{}

func (wildcardSelector) match(*context) bool                        { return true }
func (wildcardSelector) selectChildNodes(*context) (iterator, bool) { return nil, false }

type keySelector struct {
	key  string
	expr expression
}

// Determines if last element of el is an object element
func isObjectElement(p DocPath) bool {
	if p.Len() < 2 {
		return false
	}
	return p.P[len(p.P)-2].Type == ObjectNode
}

// Determines if last element of el is an object
func isObject(p DocPath) bool {
	if p.Len() < 1 {
		return false
	}
	return p.Last().Type == ObjectNode
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
		return ctx.path.Last().Name == sel.key
	}
	if isArrayElement(ctx.path) {
		if v := sel.intValue(ctx); v != nil {
			if *v < 0 {
				n := ctx.doc.Len(ctx.path.P[len(ctx.path.P)-2].Node)
				*v = normalizeIndex(*v, n)
				if *v < 0 {
					return false
				}
			}
			return ctx.path.Last().Index == *v
		}
	}
	return false
}

func (sel *keySelector) selectChildNodes(ctx *context) (iterator, bool) {
	if isObject(ctx.path) {
		node, ok := ctx.doc.Key(ctx.path.Last().Node, sel.key)
		if !ok {
			return emptyIterator{}, true
		}
		return &singleIterator{
			seg: Segment{
				Node: node,
				Type: ctx.doc.Type(node),
				Name: sel.key,
			},
		}, true
	}
	if isArray(ctx.path) {
		if v := sel.intValue(ctx); v != nil {
			n := ctx.doc.Len(ctx.path.Last().Node)
			index := normalizeIndex(*v, n)
			if index < 0 {
				return emptyIterator{}, true
			}
			if index >= n {
				return emptyIterator{}, true
			}
			node := ctx.doc.Elem(ctx.path.Last().Node, index)
			return &singleIterator{
				seg: Segment{
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

func (sel *filterSelector) selectChildNodes(*context) (children iterator, canSelect bool) {
	return nil, false
}
