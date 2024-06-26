package gojsonpath

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/bserdar/gojsonpath/parser"
)

type context struct {
	doc  DocModel
	path DocPath
}

type expression interface {
	evaluate(*context) (exprValue, error)
}

type exprValue struct {
	value any
}

// If the expression value is a node, then we keep the nodeValue as
// its value
type nodeValue struct {
	node any
	doc  DocModel
}

func (c exprValue) evaluate(*context) (exprValue, error) { return c, nil }

// Return an int, float, bool, or string as int or float64. If not convertable, returns nil
func valueAsNumber(value any) any {
	i, ok := value.(int)
	if ok {
		return i
	}
	f, ok := value.(float64)
	if ok {
		return f
	}
	b, ok := value.(bool)
	if ok && b {
		return 1
	}
	s, ok := value.(string)
	if ok {
		i, err := strconv.Atoi(s)
		if err == nil {
			return i
		}
		f, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return f
		}
	}
	if n, ok := value.(nodeValue); ok {
		return valueAsNumber(n.doc.Value(n.node))
	}
	return nil
}

// Returns an int, float, bool, or numeric string as int
func valueAsInt(value any) (int, bool) {
	v := valueAsNumber(value)
	if v == nil {
		return 0, false
	}
	if i, ok := v.(int); ok {
		return i, true
	}
	return int(v.(float64)), true
}

func (c exprValue) String() string {
	v := c.getValue()
	s, ok := v.(string)
	if ok {
		return s
	}
	return fmt.Sprint(v)
}

func (c exprValue) asBool() bool {
	value := c.getValue()
	if b, ok := value.(bool); ok {
		return b
	}
	if i, ok := value.(int); ok {
		return i != 0
	}
	if f, ok := value.(float64); ok {
		return f != 0
	}
	if s, ok := value.(string); ok {
		return len(s) != 0
	}
	return c.value != nil
}

func (c exprValue) getValue() any {
	n, ok := c.value.(nodeValue)
	if ok {
		return n.doc.Value(n.node)
	}
	return c.value
}

type unaryMinusExpression struct {
	expr expression
}

func (u unaryMinusExpression) evaluate(ctx *context) (exprValue, error) {
	c, err := u.expr.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	num := valueAsNumber(c)
	if num == nil {
		return exprValue{}, nil
	}
	if ivalue, ok := num.(int); ok {
		return exprValue{value: -ivalue}, nil
	}
	if fvalue, ok := num.(float64); ok {
		return exprValue{value: -fvalue}, nil
	}
	return exprValue{}, nil
}

type filterExpression struct {
	expr expression
}

func (expr filterExpression) evaluate(ctx *context) (exprValue, error) {
	v, err := expr.expr.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	return exprValue{value: v.asBool()}, nil
}

type notExpression struct {
	expr expression
}

func (expr notExpression) evaluate(ctx *context) (exprValue, error) {
	v, err := expr.expr.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	return exprValue{value: !v.asBool()}, nil
}

type equalityExpression struct {
	left  expression
	right expression
}

func (expr equalityExpression) evaluate(ctx *context) (exprValue, error) {
	left, err := expr.left.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	right, err := expr.right.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	leftValue := left.getValue()
	rightValue := right.getValue()

	if i, ok := leftValue.(int); ok {
		if j, ok := rightValue.(int); ok {
			return exprValue{value: i == j}, nil
		}
		if f, ok := rightValue.(float64); ok {
			return exprValue{value: i == int(f)}, nil
		}
		return exprValue{value: false}, nil
	}
	if i, ok := leftValue.(float64); ok {
		if j, ok := rightValue.(float64); ok {
			return exprValue{value: i == j}, nil
		}
		if j, ok := rightValue.(int); ok {
			return exprValue{value: i == float64(j)}, nil
		}
	}
	if l, ok := leftValue.(string); ok {
		if r, ok := rightValue.(string); ok {
			return exprValue{value: l == r}, nil
		}
	}
	if b, ok := leftValue.(bool); ok {
		if r, ok := rightValue.(bool); ok {
			return exprValue{value: b == r}, nil
		}
	}
	return exprValue{value: false}, nil
}

type equality3Expression struct {
	left  expression
	right expression
}

func (expr equality3Expression) evaluate(ctx *context) (exprValue, error) {
	left, err := expr.left.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	right, err := expr.right.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	leftValue := left.getValue()
	rightValue := right.getValue()
	if i, ok := leftValue.(int); ok {
		if j, ok := rightValue.(int); ok {
			return exprValue{value: i == j}, nil
		}
		return exprValue{value: false}, nil
	}
	if i, ok := leftValue.(float64); ok {
		if j, ok := rightValue.(float64); ok {
			return exprValue{value: i == j}, nil
		}
	}
	if l, ok := leftValue.(string); ok {
		if r, ok := rightValue.(string); ok {
			return exprValue{value: l == r}, nil
		}
	}
	if b, ok := leftValue.(bool); ok {
		if r, ok := rightValue.(bool); ok {
			return exprValue{value: b == r}, nil
		}
	}
	return exprValue{value: false}, nil
}

type pathExpression struct {
	path Path
}

func (expr pathExpression) evaluate(ctx *context) (exprValue, error) {
	result := make([]any, 0)
	err := evaluateAt(ctx, expr.path, 0, func(item DocPath) {
		result = append(result, item.Last().Node)
	})
	if err != nil {
		return exprValue{}, err
	}
	if len(result) == 1 {
		return exprValue{value: nodeValue{
			node: result[0],
			doc:  ctx.doc,
		}}, nil
	}
	for i := range result {
		result[i] = nodeValue{
			node: result[i],
			doc:  ctx.doc,
		}
	}
	return exprValue{value: result}, nil
}

type landExpression struct {
	left, right expression
}

func (expr landExpression) evaluate(ctx *context) (exprValue, error) {
	v, err := expr.left.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	if !v.asBool() {
		return exprValue{value: false}, nil
	}
	v, err = expr.right.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	return exprValue{value: v.asBool()}, nil
}

type lorExpression struct {
	left, right expression
}

func (expr lorExpression) evaluate(ctx *context) (exprValue, error) {
	v, err := expr.left.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	if v.asBool() {
		return exprValue{value: true}, nil
	}
	v, err = expr.right.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	return exprValue{value: v.asBool()}, nil
}

// returns comparison result (-1,0,1), whether or not comparison is valid, and error
func compareExpr(ctx *context, left, right expression) (int, bool, error) {
	l, err := left.evaluate(ctx)
	if err != nil {
		return 0, false, err
	}
	r, err := right.evaluate(ctx)
	if err != nil {
		return 0, false, err
	}
	lvalue := l.getValue()
	rvalue := r.getValue()

	if lvalue == nil && rvalue == nil {
		return 0, true, nil
	}
	if lvalue == nil && rvalue != nil {
		return -1, true, nil
	}
	if lvalue != nil && rvalue == nil {
		return 1, true, nil
	}

	if i, ok := lvalue.(int); ok {
		j, ok := valueAsInt(rvalue)
		if !ok {
			return 0, false, nil
		}
		if i < j {
			return -1, true, nil
		}
		if i > j {
			return 1, true, nil
		}
		return 0, true, nil
	}

	if i, ok := lvalue.(float64); ok {
		n := valueAsNumber(rvalue)
		if n == nil {
			return 0, false, nil
		}
		var j float64
		if f, ok := n.(float64); ok {
			j = f
		} else if f, ok := n.(int); ok {
			j = float64(f)
		}
		if i < j {
			return -1, true, nil
		}
		if i > j {
			return 1, true, nil
		}
		return 0, true, nil
	}

	if i, ok := lvalue.(string); ok {
		j, ok := rvalue.(string)
		if !ok {
			j = fmt.Sprint(rvalue)
		}
		if i < j {
			return -1, true, nil
		}
		if i > j {
			return 1, true, nil
		}
		return 0, true, nil
	}
	return 0, false, nil
}

type relationalExpression struct {
	op func(*context) (exprValue, error)
}

func (rel relationalExpression) evaluate(ctx *context) (exprValue, error) {
	return rel.op(ctx)
}

type mulExpression struct {
	left, right expression
}

func (mul mulExpression) evaluate(ctx *context) (exprValue, error) {
	l, err := mul.left.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	r, err := mul.right.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	nl := valueAsNumber(l.value)
	nr := valueAsNumber(r.value)
	if nl == nil || nr == nil {
		return exprValue{}, nil
	}
	if fl, ok := nl.(float64); ok {
		if fr, ok := nr.(float64); ok {
			return exprValue{value: fl * fr}, nil
		}
		return exprValue{value: fl * float64(nr.(int))}, nil
	}
	i := nl.(int)
	if ir, ok := nr.(int); ok {
		return exprValue{value: i * ir}, nil
	}
	return exprValue{value: float64(i) * nr.(float64)}, nil
}

type divExpression struct {
	left, right expression
}

func (div divExpression) evaluate(ctx *context) (exprValue, error) {
	l, err := div.left.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	r, err := div.right.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	nl := valueAsNumber(l.value)
	nr := valueAsNumber(r.value)
	if nl == nil || nr == nil {
		return exprValue{}, nil
	}
	if fl, ok := nl.(float64); ok {
		if fr, ok := nr.(float64); ok {
			if fr == 0 {
				return exprValue{}, nil
			}
			return exprValue{value: fl / fr}, nil
		}
		i := nr.(int)
		if i == 0 {
			return exprValue{}, nil
		}
		return exprValue{value: fl / float64(i)}, nil
	}
	i := nl.(int)
	if ir, ok := nr.(int); ok {
		if ir == 0 {
			return exprValue{}, nil
		}
		return exprValue{value: i / ir}, nil
	}
	fr := nr.(float64)
	if fr == 0 {
		return exprValue{}, nil
	}
	return exprValue{value: float64(i) / fr}, nil
}

type modExpression struct {
	left, right expression
}

func (mod modExpression) evaluate(ctx *context) (exprValue, error) {
	l, err := mod.left.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	r, err := mod.right.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	nl := valueAsNumber(l.value)
	nr := valueAsNumber(r.value)
	if nl == nil || nr == nil {
		return exprValue{}, nil
	}
	if fl, ok := nl.(float64); ok {
		if fr, ok := nr.(float64); ok {
			if fr == 0 {
				return exprValue{}, nil
			}
			return exprValue{value: math.Mod(fl, fr)}, nil
		}
		i := nr.(int)
		if i == 0 {
			return exprValue{}, nil
		}
		return exprValue{value: math.Mod(fl, float64(i))}, nil
	}
	i := nl.(int)
	if ir, ok := nr.(int); ok {
		if ir == 0 {
			return exprValue{}, nil
		}
		return exprValue{value: i % ir}, nil
	}
	fr := nr.(float64)
	if fr == 0 {
		return exprValue{}, nil
	}
	return exprValue{value: math.Mod(float64(i), fr)}, nil
}

type powExpression struct {
	left, right expression
}

func (pow powExpression) evaluate(ctx *context) (exprValue, error) {
	l, err := pow.left.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	r, err := pow.right.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	nl := valueAsNumber(l.value)
	nr := valueAsNumber(r.value)
	if nl == nil || nr == nil {
		return exprValue{}, nil
	}
	if fl, ok := nl.(float64); ok {
		if fr, ok := nr.(float64); ok {
			return exprValue{value: math.Pow(fl, fr)}, nil
		}
		i := nr.(int)
		return exprValue{value: math.Pow(fl, float64(i))}, nil
	}
	i := nl.(int)
	if ir, ok := nr.(int); ok {
		return exprValue{value: math.Pow(float64(i), float64(ir))}, nil
	}
	fr := nr.(float64)
	return exprValue{value: math.Pow(float64(i), fr)}, nil
}

type addExpression struct {
	left, right expression
}

func (expr addExpression) evaluate(ctx *context) (exprValue, error) {
	l, err := expr.left.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	r, err := expr.right.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	nl := valueAsNumber(l.value)
	nr := valueAsNumber(r.value)
	if nl == nil || nr == nil {
		return exprValue{}, nil
	}
	if fl, ok := nl.(float64); ok {
		if fr, ok := nr.(float64); ok {
			return exprValue{value: fl + fr}, nil
		}
		i := nr.(int)
		return exprValue{value: fl + float64(i)}, nil
	}
	i := nl.(int)
	if ir, ok := nr.(int); ok {
		return exprValue{value: i + ir}, nil
	}
	fr := nr.(float64)
	return exprValue{value: float64(i) + fr}, nil
}

type subExpression struct {
	left, right expression
}

func (expr subExpression) evaluate(ctx *context) (exprValue, error) {
	l, err := expr.left.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	r, err := expr.right.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	nl := valueAsNumber(l.value)
	nr := valueAsNumber(r.value)
	if nl == nil || nr == nil {
		return exprValue{}, nil
	}
	if fl, ok := nl.(float64); ok {
		if fr, ok := nr.(float64); ok {
			return exprValue{value: fl - fr}, nil
		}
		i := nr.(int)
		return exprValue{value: fl - float64(i)}, nil
	}
	i := nl.(int)
	if ir, ok := nr.(int); ok {
		return exprValue{value: i - ir}, nil
	}
	fr := nr.(float64)
	return exprValue{value: float64(i) - fr}, nil
}

type arrayExpression struct {
	arr []expression
}

func (expr arrayExpression) evaluate(ctx *context) (exprValue, error) {
	ret := make([]any, 0, len(expr.arr))
	for _, x := range expr.arr {
		v, err := x.evaluate(ctx)
		if err != nil {
			return exprValue{}, err
		}
		ret = append(ret, v.getValue())
	}
	return exprValue{value: ret}, nil
}

type inExpression struct {
	left  expression
	right expression
}

func (expr inExpression) evaluate(ctx *context) (exprValue, error) {
	l, err := expr.left.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	r, err := expr.right.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	if r.value == nil {
		return exprValue{}, err
	}
	arr, ok := r.value.([]any)
	if !ok {
		arr = []any{r.value}
	}
	for _, x := range arr {
		var v any
		if n, ok := x.(nodeValue); ok {
			v = n.doc.Value(n.node)
		} else {
			v = x
		}
		if l == v {
			return exprValue{value: true}, nil
		}
	}
	return exprValue{value: false}, nil
}

// Determines if the expression is a path expression
func isPathExpression(ctx parser.IExpressionContext) bool {
	switch expr := ctx.(type) {
	case *parser.SelectorExpressionContext:
		return true
	case *parser.ChainExpressionContext, *parser.MemberIndexExpressionContext:
		text := expr.GetText()
		if strings.HasPrefix(text, "@") || strings.HasPrefix(text, "$") {
			return true
		}
	case *parser.RecursiveDescentExpressionContext, *parser.RecursiveDescentTermExpressionContext, *parser.RecursiveDescentMemberIndexExpressionContext:
		return true
	}
	return false
}

type methodCallExpression struct {
	receiver expression
	name     string
	args     []expression
}

func (expr methodCallExpression) evaluate(ctx *context) (exprValue, error) {
	receiver, err := expr.receiver.evaluate(ctx)
	if err != nil {
		return exprValue{}, err
	}
	funcCall := predefinedFunctions[expr.name]
	if funcCall == nil {
		return exprValue{}, fmt.Errorf("Function not found: %s", expr.name)
	}
	args := make([]any, 0, len(expr.args)+1)
	args = append(args, receiver.getValue())
	for _, x := range expr.args {
		e, err := x.evaluate(ctx)
		if err != nil {
			return exprValue{}, err
		}
		args = append(args, e.getValue())
	}
	val, err := funcCall(args)
	if err != nil {
		return exprValue{}, err
	}
	return exprValue{value: val}, nil
}

type functionCallExpression struct {
	name string
	args []expression
}

func (expr functionCallExpression) evaluate(ctx *context) (exprValue, error) {
	funcCall := predefinedFunctions[expr.name]
	if funcCall == nil {
		return exprValue{}, fmt.Errorf("Function not found: %s", expr.name)
	}
	args := make([]any, 0, len(expr.args))
	for _, x := range expr.args {
		e, err := x.evaluate(ctx)
		if err != nil {
			return exprValue{}, err
		}
		args = append(args, e.getValue())
	}
	val, err := funcCall(args)
	if err != nil {
		return exprValue{}, err
	}
	return exprValue{value: val}, nil
}

// inPath tells if we are parsing a path. If we are parsing a path
// component, strings recognized as expressions are returned as
// string literals
func astExpression(ctx parser.IExpressionContext, inPath bool) (expression, error) {
	if inPath {
		switch expr := ctx.(type) {
		case *parser.LiteralExpressionContext:
			return astLiteral(expr.Literal().(*parser.LiteralContext)), nil
		case *parser.UnaryPlusExpressionContext, *parser.UnaryMinusExpressionContext, *parser.AdditiveExpressionContext:
			return exprValue{value: expr.GetText()}, nil
		case *parser.FilterExpressionContext:
			x, err := astExpression(expr.Expression(), false)
			if err != nil {
				return nil, err
			}
			return filterExpression{
				expr: x,
			}, nil
		case *parser.ArgumentsExpressionContext:
			return astArgumentExpression(expr)
		}
	} else {
		switch expr := ctx.(type) {

		case *parser.LiteralExpressionContext:
			return astLiteral(expr.Literal().(*parser.LiteralContext)), nil
		case *parser.UnaryPlusExpressionContext:
			return astExpression(expr.Expression(), inPath)
		case *parser.UnaryMinusExpressionContext:
			v, err := astExpression(expr.Expression(), inPath)
			if err != nil {
				return nil, err
			}
			if constant, ok := v.(exprValue); ok {
				if ivalue, ok := constant.value.(int); ok {
					return exprValue{value: -ivalue}, nil
				}
				if fvalue, ok := constant.value.(float64); ok {
					return exprValue{value: -fvalue}, nil
				}
			}
			return unaryMinusExpression{
				expr: v,
			}, nil

		case *parser.EqualityExpressionContext:
			var expressions [2]expression
			var op string
			index := 0
			for _, ch := range expr.GetChildren() {
				if tok, ok := ch.(antlr.TerminalNode); ok {
					t := tok.GetText()
					switch t {
					case "==", "===", "!=", "!==":
						op = t
					}
				} else if expr, ok := ch.(parser.IExpressionContext); ok {
					var err error
					expressions[index], err = astExpression(expr, inPath)
					if err != nil {
						return nil, err
					}
					index++
				}
			}
			switch op {
			case "==":
				return equalityExpression{
					left:  expressions[0],
					right: expressions[1],
				}, nil
			case "===":
				return equality3Expression{
					left:  expressions[0],
					right: expressions[1],
				}, nil
			case "!=":
				return notExpression{
					expr: equalityExpression{
						left:  expressions[0],
						right: expressions[1],
					},
				}, nil
			case "!==":
				return notExpression{
					expr: equality3Expression{
						left:  expressions[0],
						right: expressions[1],
					},
				}, nil
			}

		case *parser.LogicalAndExpressionContext:
			e := expr.AllExpression()
			left, err := astExpression(e[0], inPath)
			if err != nil {
				return nil, err
			}
			right, err := astExpression(e[1], inPath)
			if err != nil {
				return nil, err
			}
			return landExpression{
				left:  left,
				right: right,
			}, nil

		case *parser.LogicalOrExpressionContext:
			e := expr.AllExpression()
			left, err := astExpression(e[0], inPath)
			if err != nil {
				return nil, err
			}
			right, err := astExpression(e[1], inPath)
			if err != nil {
				return nil, err
			}
			return lorExpression{
				left:  left,
				right: right,
			}, nil

		case *parser.ParenthesizedExpressionContext:
			return astExpression(expr.Expression(), inPath)

		case *parser.RelationalExpressionContext:
			var expressions [2]expression
			var op string
			index := 0
			for _, ch := range expr.GetChildren() {
				if tok, ok := ch.(antlr.TerminalNode); ok {
					t := tok.GetText()
					switch t {
					case "<=", "<", ">=", ">":
						op = t
					}
				} else if expr, ok := ch.(parser.IExpressionContext); ok {
					var err error
					expressions[index], err = astExpression(expr, inPath)
					if err != nil {
						return nil, err
					}
					index++
				}
			}
			switch op {
			case "<":
				return relationalExpression{
					op: func(ctx *context) (exprValue, error) {
						res, valid, err := compareExpr(ctx, expressions[0], expressions[1])
						return exprValue{value: res < 0 && valid}, err
					},
				}, nil
			case "<=":
				return relationalExpression{
					op: func(ctx *context) (exprValue, error) {
						res, valid, err := compareExpr(ctx, expressions[0], expressions[1])
						return exprValue{value: res <= 0 && valid}, err
					},
				}, nil
			case ">":
				return relationalExpression{
					op: func(ctx *context) (exprValue, error) {
						res, valid, err := compareExpr(ctx, expressions[0], expressions[1])
						return exprValue{value: res > 0 && valid}, err
					},
				}, nil
			case ">=":
				return relationalExpression{
					op: func(ctx *context) (exprValue, error) {
						res, valid, err := compareExpr(ctx, expressions[0], expressions[1])
						return exprValue{value: res >= 0 && valid}, err
					},
				}, nil
			}

		case *parser.NotExpressionContext:
			e, err := astExpression(expr.Expression(), inPath)
			if err != nil {
				return nil, err
			}
			return notExpression{
				expr: e,
			}, nil

		case *parser.MultiplicativeExpressionContext:
			var expressions [2]expression
			var op string
			index := 0
			for _, ch := range expr.GetChildren() {
				if tok, ok := ch.(antlr.TerminalNode); ok {
					t := tok.GetText()
					switch t {
					case "*", "/", "%":
						op = t
					}
				} else if expr, ok := ch.(parser.IExpressionContext); ok {
					var err error
					expressions[index], err = astExpression(expr, inPath)
					if err != nil {
						return nil, err
					}
					index++
				}
			}
			switch op {
			case "*":
				return mulExpression{
					left:  expressions[0],
					right: expressions[1],
				}, nil
			case "/":
				return divExpression{
					left:  expressions[0],
					right: expressions[1],
				}, nil
			case "%":
				return modExpression{
					left:  expressions[0],
					right: expressions[1],
				}, nil
			}

		case *parser.PowerExpressionContext:
			expressions := expr.AllExpression()
			left, err := astExpression(expressions[0], inPath)
			if err != nil {
				return nil, err
			}
			right, err := astExpression(expressions[1], inPath)
			if err != nil {
				return nil, err
			}
			return powExpression{
				left:  left,
				right: right,
			}, nil

		case *parser.AdditiveExpressionContext:
			var expressions [2]expression
			var op string
			index := 0
			for _, ch := range expr.GetChildren() {
				if tok, ok := ch.(antlr.TerminalNode); ok {
					t := tok.GetText()
					switch t {
					case "+", "-":
						op = t
					}
				} else if expr, ok := ch.(parser.IExpressionContext); ok {
					var err error
					expressions[index], err = astExpression(expr, inPath)
					if err != nil {
						return nil, err
					}
					index++
				}
			}
			switch op {
			case "+":
				return addExpression{
					left:  expressions[0],
					right: expressions[1],
				}, nil
			case "-":
				return subExpression{
					left:  expressions[0],
					right: expressions[1],
				}, nil
			}

		case *parser.InExpressionContext:
			slice := expr.AllExpression()
			left, err := astExpression(slice[0], false)
			if err != nil {
				return nil, err
			}
			right, err := astExpression(slice[1], false)
			if err != nil {
				return nil, err
			}
			return inExpression{
				left:  left,
				right: right,
			}, nil

		case *parser.ArrayLiteralExpressionContext:
			arr, err := astElementList(expr.ArrayLiteral().(*parser.ArrayLiteralContext).ElementList().(*parser.ElementListContext))
			if err != nil {
				return nil, err
			}
			return arrayExpression{
				arr: arr,
			}, nil

		case *parser.ArgumentsExpressionContext:
			return astArgumentExpression(expr)

			// 	case *parser.PathExpressionContext:
			// 	case *parser.IdentifierExpressionContext:
			// 	case *parser.ObjectLiteralExpressionContext:
			// 	case *parser.ChainExpressionContext:
			// 	case *parser.MemberIndexExpressionContext:
			// 	case *parser.ArgumentsExpressionContext:
			// 	case *parser.BitNotExpressionContext:
			// 	case *parser.BitAndExpressionContext:
			// 	case *parser.BitXOrExpressionContext:
			// 	case *parser.BitOrExpressionContext:
			// 	case *parser.TernaryExpressionContext:
		}

		// Check if this is a path
		// Path starts with @ or $
		if ch := ctx.GetChild(0); ch != nil {
			parse := func(c interface{ GetText() string }) expression {
				tok := c.GetText()
				if strings.HasPrefix(tok, "@") || strings.HasPrefix(tok, "$") {
					if p, err := astPath(ctx); err == nil {
						return pathExpression{path: p}
					}
				}
				return nil
			}
			switch c := ch.(type) {
			case *parser.SelectorExpressionContext:
				if e := parse(c); e != nil {
					return e, nil
				}
			case *parser.SelectorContext:
				if e := parse(c); e != nil {
					return e, nil
				}
			case *parser.ChainExpressionContext:
				if e := parse(c); e != nil {
					return e, nil
				}
			case *parser.MemberIndexExpressionContext:
				if e := parse(c); e != nil {
					return e, nil
				}
			}
		}
	}

	return nil, ErrInvalidExpression(ctx.GetText())
}

func astLiteral(ctx *parser.LiteralContext) exprValue {
	if ctx.NullLiteral() != nil {
		return exprValue{}
	}
	if l := ctx.BooleanLiteral(); l != nil {
		return exprValue{value: l.GetText() == "true"}
	}
	if l := ctx.StringLiteral(); l != nil {
		return exprValue{value: unescapeString(l.GetText())}
	}
	return exprValue{value: numericLiteral(ctx.NumericLiteral().(*parser.NumericLiteralContext))}
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

func astElementList(ctx interface {
	AllExpression() []parser.IExpressionContext
}) ([]expression, error) {
	list := ctx.AllExpression()
	arr := make([]expression, 0, len(list))
	for _, e := range list {
		elem, err := astExpression(e, false)
		if err != nil {
			return nil, err
		}
		arr = append(arr, elem)
	}
	return arr, nil
}

func astArgumentExpression(ctx *parser.ArgumentsExpressionContext) (expression, error) {
	args, err := astElementList(ctx.Arguments().(*parser.ArgumentsContext))
	if err != nil {
		return nil, err
	}
	e := ctx.Expression()
	var name string
	var fCall bool
	var receiver expression
	if isPathExpression(e) {
		path, err := astPath(e)
		if err != nil {
			return nil, err
		}
		// Last component of path is method
		if len(path.selectors) < 2 {
			return nil, ErrInvalidAST
		}
		if key, ok := path.selectors[len(path.selectors)-1].(*keySelector); ok {
			name = key.key
			fCall = false
			path.selectors = path.selectors[:len(path.selectors)-1]
			path.recursive = path.recursive[:len(path.recursive)-1]
			receiver = &pathExpression{
				path: path,
			}
		}
	} else {
		// Function call?
		if id, ok := e.(*parser.IdentifierExpressionContext); ok {
			name = id.GetText()
			fCall = true
		} else {
			// Method call without a path?
			if chain, ok := e.(*parser.ChainExpressionContext); ok {
				exprs := chain.AllExpression()
				name = exprs[1].GetText()
				var err error
				receiver, err = astExpression(exprs[0], false)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	if len(name) == 0 {
		return nil, ErrInvalidAST
	}
	if fCall {
		return functionCallExpression{
			name: name,
			args: args,
		}, nil
	}
	return methodCallExpression{
		name:     name,
		args:     args,
		receiver: receiver,
	}, nil
}
