package gojsonpath

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"

	"github.com/bserdar/gojsonpath/parser"
)

//go:generate antlr4 -Dlanguage=Go Jsonpath.g4 -o parser

type errorListener struct {
	antlr.DefaultErrorListener
	err error
}

type ErrSyntax string
type ErrInvalidExpression string

func (e ErrSyntax) Error() string            { return "Syntax error: " + string(e) }
func (e ErrInvalidExpression) Error() string { return "Invalid expression: " + string(e) }

func (lst *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if lst.err == nil {
		lst.err = ErrSyntax(fmt.Sprintf("line %d:%d %s ", line, column, msg))
	}
}

func Parse(input string) (Path, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("PANIC:", input)
			fmt.Println(r)
		}
	}()
	pr := getParser(input)
	pr.RemoveErrorListeners()
	errListener := errorListener{}
	pr.AddErrorListener(&errListener)
	c := pr.Path()
	if errListener.err != nil {
		return Path{}, fmt.Errorf("%w, input: %s", errListener.err, input)
	}
	// Build a matcher from AST
	return astPath(c.Expression())
}

// GetParser returns a parser that will parse the input string
func getParser(input string) *parser.JsonpathParser {
	lexer := parser.NewJsonpathLexer(antlr.NewInputStream(input))
	lexer.RemoveErrorListeners()
	errListener := errorListener{}
	lexer.AddErrorListener(&errListener)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJsonpathParser(stream)
	p.BuildParseTrees = true
	return p
}

func astPath(ctx parser.IExpressionContext) (path Path, err error) {
	// A Path can be a chain, recursive descent, member index, or selector
	switch expr := ctx.(type) {
	case *parser.SelectorExpressionContext:
		switch expr.Selector().GetText() {
		case "$":
			path.selectors = []selector{rootElementSelector{}}
			path.recursive = []bool{false}
		case "@":
			path.selectors = []selector{currentElementSelector{}}
			path.recursive = []bool{false}
		case "*":
			path.selectors = []selector{wildcardSelector{}}
			path.recursive = []bool{false}
		}
		return

	case *parser.ChainExpressionContext:
		sub := expr.AllExpression() // must be 2
		var left, right Path
		if left, err = astPath(sub[0]); err != nil {
			return
		}
		if right, err = astPath(sub[1]); err != nil {
			return
		}
		// Combine paths
		path.selectors = append(left.selectors, right.selectors...)
		path.recursive = append(left.recursive[:len(left.recursive)-1], false)
		path.recursive = append(path.recursive, right.recursive...)
		return

	case *parser.RecursiveDescentTermExpressionContext: // This is an expression that ends with ... It is equivalent to ..*
		if path, err = astPath(expr.Expression()); err != nil {
			return
		}
		path.selectors = append(path.selectors, wildcardSelector{})
		path.recursive = append(path.recursive[:len(path.recursive)-1], true)
		return

	case *parser.RecursiveDescentExpressionContext:
		sub := expr.AllExpression() // must be 2
		var left, right Path
		if left, err = astPath(sub[0]); err != nil {
			return
		}
		if right, err = astPath(sub[1]); err != nil {
			return
		}
		// Combine paths
		path.selectors = append(left.selectors, right.selectors...)
		path.recursive = append(left.recursive[:len(left.recursive)-1], true)
		path.recursive = append(path.recursive, right.recursive...)
		return

	case *parser.MemberIndexExpressionContext:
		var left, right Path
		if left, err = astPath(expr.Expression()); err != nil {
			return
		}
		indexExprs := expr.AllIndexExpression()
		if right, err = astBracketedSelector(indexExprs); err != nil {
			return
		}
		// Combine paths
		path.selectors = append(left.selectors, right.selectors...)
		path.recursive = append(left.recursive[:len(left.recursive)-1], false)
		path.recursive = append(path.recursive, right.recursive...)
		return

	case *parser.ArrayLiteralExpressionContext: // Paths of the form where member index is at the end(...[0]) appear like array literals
		list := expr.ArrayLiteral().(*parser.ArrayLiteralContext).ElementList().(*parser.ElementListContext).AllExpression()
		path, err = astBracketedSelectorExpressions(list)
		return

	case *parser.LiteralExpressionContext:
		var sel selector
		if sel, err = astKeySelectorFromLiteral(expr); err != nil {
			return
		}
		path.selectors = []selector{sel}
		path.recursive = []bool{false}
		return

	case *parser.IdentifierExpressionContext: // This is a key selector
		var sel selector
		if sel, err = astKeySelectorFromIdentifier(expr); err != nil {
			return
		}
		path.selectors = []selector{sel}
		path.recursive = []bool{false}
		return

	default:
		// Use the expression text as the key
		var e expression
		if e, err = astExpression(expr, true); err != nil {
			return
		}
		path.selectors = []selector{
			&keySelector{
				key:  expr.GetText(),
				expr: e,
			},
		}
		path.recursive = []bool{false}

		return
	}

	return
}

func astKeySelectorFromLiteral(ctx *parser.LiteralExpressionContext) (selector, error) {
	lit := ctx.Literal()
	if str := lit.(*parser.LiteralContext).StringLiteral(); str != nil {
		return &keySelector{
			key: unescapeString(str.GetText()),
		}, nil
	}
	expr, _ := astExpression(ctx, true)
	return &keySelector{
		key:  lit.GetText(),
		expr: expr,
	}, nil
}

func astKeySelectorFromIdentifier(ctx *parser.IdentifierExpressionContext) (selector, error) {
	return &keySelector{
		key: ctx.GetText(),
	}, nil
}

func astBracketedSelector(ctxs []parser.IIndexExpressionContext) (path Path, err error) {
	u := &unionSelector{
		parts: make([]selector, 0, len(ctxs)),
	}
	for _, c := range ctxs {
		ctx := c.(*parser.IndexExpressionContext)
		var sel selector
		if expr := ctx.Expression(); expr != nil {
			sel, err = astBracketedExpressionSelector(expr)
			if err != nil {
				return
			}
		} else if slice := ctx.Slice(); slice != nil {
			sel, err = astSlice(slice.(*parser.SliceContext))
			if err != nil {
				return
			}
		}
		if sel != nil {
			u.parts = append(u.parts, sel)
		}
	}
	if len(u.parts) == 1 {
		path.selectors = []selector{u.parts[0]}
	} else {
		path.selectors = []selector{u}
	}
	path.recursive = []bool{false}
	return
}

func astBracketedSelectorExpressions(ctxs []parser.IExpressionContext) (path Path, err error) {
	u := &unionSelector{
		parts: make([]selector, 0, len(ctxs)),
	}
	for _, c := range ctxs {
		var sel selector
		if sel, err = astBracketedExpressionSelector(c); err != nil {
			return
		}
		u.parts = append(u.parts, sel)
	}
	if len(u.parts) == 1 {
		path.selectors = []selector{u.parts[0]}
	} else {
		path.selectors = []selector{u}
	}
	path.recursive = []bool{false}
	return
}

// Expression within a bracket
func astBracketedExpressionSelector(ctx parser.IExpressionContext) (selector, error) {
	switch expr := ctx.(type) {
	case *parser.LiteralExpressionContext:
		return astKeySelectorFromLiteral(expr)
	case *parser.IdentifierExpressionContext:
		return astKeySelectorFromIdentifier(expr)
	case *parser.FilterExpressionContext:
		filter, err := astExpression(expr.Expression(), false)
		if err != nil {
			return nil, err
		}
		return &filterSelector{filter: filter}, nil
	case *parser.SelectorExpressionContext:
		switch expr.Selector().GetText() {
		case "$":
			return rootElementSelector{}, nil
		case "@":
			return currentElementSelector{}, nil
		case "*":
			return wildcardSelector{}, nil
		}
	}
	expr, err := astExpression(ctx, true)
	if err != nil {
		return nil, err
	}
	return &keySelector{key: ctx.GetText(), expr: expr}, nil
}

func astSlice(ctx *parser.SliceContext) (selector, error) {
	expressions := [3]expression{}
	index := 0
	for _, ch := range ctx.GetChildren() {
		if tok, ok := ch.(antlr.TerminalNode); ok {
			if tok.GetText() == ":" {
				index++
			}
		} else if expr, ok := ch.(parser.IExpressionContext); ok {
			var err error
			expressions[index], err = astExpression(expr, false)
			if err != nil {
				return nil, err
			}
		}
	}
	switch index {
	case 0:
		return &indexSelector{
			index: expressions[0],
		}, nil
	case 1:
		return &rangeSelector{
			start: expressions[0],
			end:   expressions[1],
		}, nil
	}
	if expressions[2] == nil {
		return &rangeSelector{
			start: exprValue{value: 0},
			end:   expressions[1],
		}, nil
	}
	return &sliceSelector{
		start: expressions[0],
		end:   expressions[1],
		step:  expressions[2],
	}, nil
}
