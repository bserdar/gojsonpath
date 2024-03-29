package gojsonpath

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bserdar/gojsonpath/parser"
)

type expression interface {
	evaluate() constantValue
}

type constantValue struct {
	value any
}

func (c constantValue) evaluate() constantValue { return c }

func (c constantValue) asInt() int {
	i, ok := c.value.(int)
	if ok {
		return i
	}
	f, ok := c.value.(float64)
	if ok {
		return int(f)
	}
	b, ok := c.value.(bool)
	if ok && b {
		return 1
	}
	return 0
}

func (c constantValue) asString() string {
	s, ok := c.value.(string)
	if ok {
		return s
	}
	return fmt.Sprint(c.value)
}

func astSingleExpression(ctx parser.ISingleExpressionContext) (expression, error) {
	switch expr := ctx.(type) {
	case *parser.LiteralExpressionContext:
		return astLiteral(expr.Literal().(*parser.LiteralContext)), nil

	case *parser.PathExpressionContext:
	case *parser.IdentifierExpressionContext:
	case *parser.ArrayLiteralExpressionContext:
	case *parser.ObjectLiteralExpressionContext:
	case *parser.ChainExpressionContext:
	case *parser.MemberIndexExpressionContext:
	case *parser.ArgumentsExpressionContext:
	case *parser.UnaryPlusExpressionContext:
	case *parser.UnaryMinusExpressionContext:
	case *parser.BitNotExpressionContext:
	case *parser.NotExpressionContext:
	case *parser.PowerExpressionContext:
	case *parser.MultiplicativeExpressionContext:
	case *parser.AdditiveExpressionContext:
	case *parser.CoalesceExpressionContext:
	case *parser.RelationalExpressionContext:
	case *parser.InExpressionContext:
	case *parser.EqualityExpressionContext:
	case *parser.BitAndExpressionContext:
	case *parser.BitXOrExpressionContext:
	case *parser.BitOrExpressionContext:
	case *parser.LogicalAndExpressionContext:
	case *parser.LogicalOrExpressionContext:
	case *parser.TernaryExpressionContext:
	case *parser.ParenthesizedExpressionContext:
	case *parser.FilterExpressionContext:
	}
	return nil, nil
}

func astLiteral(ctx *parser.LiteralContext) constantValue {
	if ctx.NullLiteral() != nil {
		return constantValue{}
	}
	if l := ctx.BooleanLiteral(); l != nil {
		return constantValue{value: l.GetText() == "true"}
	}
	if l := ctx.StringLiteral(); l != nil {
		return constantValue{value: unescapeString(l.GetText())}
	}
	return constantValue{value: numericLiteral(ctx.NumericLiteral().(*parser.NumericLiteralContext))}
}

func numericLiteral(ctx *parser.NumericLiteralContext) any {
	if lit := ctx.DecimalLiteral(); lit != nil {
		number := lit.GetText()
		if strings.IndexRune(number, 'e') != -1 || strings.IndexRune(number, '.') != -1 {
			v, _ := strconv.ParseFloat(number, 64)
			return float64(v)
		}
		v, _ := strconv.ParseInt(number, 10, 64)
		return int(v)
	}

	if lit := ctx.HexIntegerLiteral(); lit != nil {
		v, _ := strconv.ParseInt(lit.GetText()[2:], 16, 64)
		return int(v)
	}

	if lit := ctx.OctalIntegerLiteral2(); lit != nil {
		v, _ := strconv.ParseInt(lit.GetText()[2:], 8, 64)
		return int(v)
	}
	lit := ctx.BinaryIntegerLiteral().GetText()[2:]
	v, _ := strconv.ParseInt(strings.Replace(lit, "_", "", -1), 2, 64)
	return int(v)
}

var singleEscapeRunes = map[rune]rune{
	'\'': '\'',
	'"':  '"',
	'b':  '\b',
	'f':  '\f',
	'n':  '\n',
	'r':  '\r',
	'v':  '\v',
}

func unescapeString(in string) string {
	out := make([]rune, 0, len(in))
	esc := false
	hex := false
	uni := false
	var buf []rune
	for _, r := range in[1 : len(in)-1] {
		if esc {
			if hex {
				buf = append(buf, r)
				if len(buf) == 2 {
					hex = false
					esc = false
					v, _ := strconv.ParseInt(string(buf), 16, 8)
					out = append(out, rune(v))
					buf = nil
				}
			} else if uni {
				buf = append(buf, r)
				if len(buf) > 1 && buf[0] == '{' {
					if r == '}' {
						v, _ := strconv.ParseInt(string(buf[1:]), 16, 32)
						out = append(out, rune(v))
						buf = nil
						uni = false
						esc = false
					} else {
						out = append(out, r)
					}
				} else if len(buf) == 4 {
					v, _ := strconv.ParseInt(string(buf), 16, 32)
					out = append(out, rune(v))
					buf = nil
					uni = false
					esc = false
				}
			} else if r == 'x' {
				hex = true
				buf = make([]rune, 0, 4)
			} else if r == 'u' {
				uni = true
				buf = make([]rune, 0, 4)
			} else {
				if x, ok := singleEscapeRunes[r]; ok {
					out = append(out, x)
				}
				esc = false
			}
		} else if r == '\\' {
			esc = true
		} else {
			out = append(out, r)
		}
	}
	return string(out)
}
