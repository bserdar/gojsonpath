package gojsonpath

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"

	"github.com/bserdar/gojsonpath/parser"
)

// ParseSimplePath creates a compiled JSON path from a simplified
// path expression of the form
//
//	/ key1 / key2 / index / * / key3
//
// where keys and indexes address object keys and array indexes, and *
// addresses any matching node. This form does not support recursive
// descent. The path must be an absolute path (i.e. start with '/'.)
// Keys can be given as is, or as string literals (i.e. quoted.)
func ParseSimplePath(input string) (Path, error) {
	pr := getParser(input)
	pr.RemoveErrorListeners()
	errListener := errorListener{}
	pr.AddErrorListener(&errListener)
	if errListener.err != nil {
		return Path{}, fmt.Errorf("%w, input: %s", errListener.err, input)
	}
	// Build a matcher
	p, err := astSimplePath(pr.SimplePath().(*parser.SimplePathContext).SimplePathExpr())
	if err != nil {
		return p, err
	}
	return Append(SelectRootPath, p), nil
}

func astSimplePath(ctx parser.ISimplePathExprContext) (path Path, err error) {
	if lit := ctx.Literal(); lit != nil {
		v := astLiteral(lit.(*parser.LiteralContext))
		if lit.StringLiteral() != nil {
			path.selectors = []selector{
				&keySelector{
					key: v.value.(string),
				},
			}
			path.recursive = []bool{false}
			return
		}
		path.selectors = []selector{
			&keySelector{
				key: v.String(),
			},
		}
		path.recursive = []bool{false}
		return
	}

	if id := ctx.Identifier(); id != nil {
		path.selectors = []selector{
			&keySelector{
				key: id.GetText(),
			},
		}
		path.recursive = []bool{false}
		return
	}

	if paths := ctx.AllSimplePathExpr(); len(paths) > 0 {
		// There are two components
		var p1, p2 Path
		p1, err = astSimplePath(paths[0])
		if err != nil {
			return
		}
		p2, err = astSimplePath(paths[1])
		if err != nil {
			return
		}
		path = Append(p1, p2)
		return
	}

	for _, ch := range ctx.GetChildren() {
		if tok, ok := ch.(antlr.TerminalNode); ok {
			t := tok.GetText()
			if t == "/" {
				path = SelectRootPath
			}
			if t == "*" {
				path = WildcardPath
			}
		}
	}
	return
}
