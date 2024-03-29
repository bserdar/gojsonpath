package gojsonpath

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

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
	pr := getParser(input)
	pr.RemoveErrorListeners()
	errListener := errorListener{}
	pr.AddErrorListener(&errListener)
	c := pr.Path()
	if errListener.err != nil {
		return Path{}, fmt.Errorf("%w, input: %s", errListener.err, input)
	}
	// Build a matcher from AST
	return astPath(c.(*parser.PathContext))
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

func astPath(ctx *parser.PathContext) (path Path, err error) {
	path.selectors = make([]selector, 1, 32)
	path.recursive = make([]bool, 0, 32)
	if ctx.Root_selector() != nil {
		path.selectors[0] = &rootElementSelector{}
	} else {
		path.selectors[0] = &currentElementSelector{}
	}
	for _, el := range ctx.AllPath_element() {
		elCtx := el.(*parser.Path_elementContext)
		if elCtx.Recursive_descent() != nil {
			path.recursive = append(path.recursive, true)
		} else {
			path.recursive = append(path.recursive, false)
		}

		if bracketed := elCtx.Bracketed_selector(); bracketed != nil {
			if err = astBracketedSelector(bracketed.(*parser.Bracketed_selectorContext), &path); err != nil {
				return
			}
		} else if selector := elCtx.Selector(); selector != nil {
			if err = astSelector(selector.(*parser.SelectorContext), &path); err != nil {
				return
			}
		}
	}
	return
}

func astBracketedSelector(ctx *parser.Bracketed_selectorContext, path *Path) error {
	return astUnion(ctx.Union().(*parser.UnionContext), path)
}

func astUnion(ctx *parser.UnionContext, path *Path) error {
	parts := ctx.AllUnionPart()
	selector := &unionSelector{
		parts: make([]selector, 0, len(parts)),
	}
	for _, part := range parts {
		p, err := astUnionPart(part.(*parser.UnionPartContext))
		if err != nil {
			return err
		}
		selector.parts = append(selector.parts, p)
	}
	if len(selector.parts) == 1 {
		path.selectors = append(path.selectors, selector.parts[0])
	} else {
		path.selectors = append(path.selectors, selector)
	}
	return nil
}

func astUnionPart(ctx *parser.UnionPartContext) (selector, error) {
	if sel := ctx.Selector(); sel != nil {
		return astNestedSelector(sel.(*parser.SelectorContext))
	}
	return astSlice(ctx.Slice().(*parser.SliceContext))
}

func astSelector(ctx *parser.SelectorContext, path *Path) error {
	sel, err := astNestedSelector(ctx)
	if err != nil {
		return err
	}
	path.selectors = append(path.selectors, sel)
	return nil
}

func astNestedSelector(ctx *parser.SelectorContext) (selector, error) {
	if expr := ctx.SingleExpression(); expr != nil {
		switch e := expr.(type) {
		case *parser.LiteralExpressionContext:
			lit := e.Literal().(*parser.LiteralContext)
			if l := lit.StringLiteral(); l != nil {
				return &keySelector{key: unescapeString(l.GetText())}, nil
			}
			return &keySelector{
				key:  lit.GetText(),
				expr: astLiteral(lit),
			}, nil

		case *parser.IdentifierExpressionContext,
			*parser.ArrayLiteralExpressionContext,
			*parser.ObjectLiteralExpressionContext,
			*parser.ChainExpressionContext,
			*parser.MemberIndexExpressionContext,
			*parser.ArgumentsExpressionContext,
			*parser.UnaryPlusExpressionContext,
			*parser.UnaryMinusExpressionContext,
			*parser.BitNotExpressionContext,
			*parser.NotExpressionContext,
			*parser.PowerExpressionContext,
			*parser.MultiplicativeExpressionContext,
			*parser.AdditiveExpressionContext,
			*parser.CoalesceExpressionContext,
			*parser.RelationalExpressionContext,
			*parser.InExpressionContext,
			*parser.EqualityExpressionContext,
			*parser.BitAndExpressionContext,
			*parser.BitXOrExpressionContext,
			*parser.BitOrExpressionContext,
			*parser.LogicalAndExpressionContext,
			*parser.LogicalOrExpressionContext,
			*parser.TernaryExpressionContext,
			*parser.ParenthesizedExpressionContext:
			return &keySelector{key: expr.GetText()}, nil

		case *parser.PathExpressionContext:
		case *parser.FilterExpressionContext:

		}
		return nil, ErrSyntax(fmt.Sprintf("Unrecognized selector: %s", expr.GetText()))
	}
	for _, ch := range ctx.GetChildren() {
		if tok, ok := ch.(antlr.TerminalNode); ok {
			if tok.GetText() == "*" {
				return &wildcardSelector{}, nil
			}
		}
	}
	return nil, ErrInvalidAST
}

func astSlice(ctx *parser.SliceContext) (selector, error) {
	expressions := [3]expression{}
	index := 0
	for _, ch := range ctx.GetChildren() {
		if tok, ok := ch.(antlr.TerminalNode); ok {
			if tok.GetText() == ":" {
				index++
			}
		} else if expr, ok := ch.(parser.ISingleExpressionContext); ok {
			var err error
			expressions[index], err = astSingleExpression(expr)
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
			start: constantValue{value: 0},
			end:   expressions[1],
		}, nil
	}
	return &sliceSelector{
		start: expressions[0],
		end:   expressions[1],
		step:  expressions[2],
	}, nil
}
