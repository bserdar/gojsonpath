package gojsonpath

// Used for [index]
type indexSelector struct {
	index expression
}

// Determines if last element of el is an array element
func isArrayElement(p DocPath) bool {
	if p.Len() < 2 {
		return false
	}
	return p.P[len(p.P)-2].Type == ArrayNode
}

// Determines if last element of el is an array
func isArray(p DocPath) bool {
	if p.Len() < 1 {
		return false
	}
	return p.Last().Type == ArrayNode
}

func normalizeIndex(index any, len int) int {
	i, ok := valueAsInt(index)
	if !ok {
		return -1
	}
	if i >= 0 {
		return i
	}
	if -len <= i {
		return i + len
	}
	return 0
}

func (sel *indexSelector) match(ctx *context) bool {
	if !isArrayElement(ctx.path) {
		return false
	}
	ivalue, err := sel.index.evaluate(ctx)
	if err != nil {
		return false
	}
	index, ok := valueAsInt(ivalue.value)
	if !ok {
		return false
	}
	if index < 0 {
		n := ctx.doc.Len(ctx.path.Last().Node)
		index := normalizeIndex(index, n)
		if index < 0 {
			return false
		}
	}
	return ctx.path.Last().Index == index
}

func (sel *indexSelector) selectChildNodes(ctx *context) (children iterator, canSelect bool) {
	if !isArray(ctx.path) {
		return emptyIterator{}, true
	}
	n := ctx.doc.Len(ctx.path.Last().Node)
	ivalue, err := sel.index.evaluate(ctx)
	if err != nil {
		return nil, false
	}
	index := normalizeIndex(ivalue.value, n)
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

// Used for [from:to], [:to], or [from:]
type rangeSelector struct {
	start expression
	end   expression
}

func getRange(ctx *context, start, end expression, len int) (from, to int) {
	var s int
	if start != nil {
		ivalue, _ := start.evaluate(ctx)
		s = normalizeIndex(ivalue.value, len)
	}
	if s < 0 {
		return -1, -1
	}
	var e int
	if end != nil {
		ivalue, _ := end.evaluate(ctx)
		e = normalizeIndex(ivalue.value, len)
	} else {
		e = len
	}
	if e < 0 {
		return -1, -1
	}
	return s, e
}

func (sel *rangeSelector) match(ctx *context) bool {
	if !isArrayElement(ctx.path) {
		return false
	}
	n := ctx.doc.Len(ctx.path.P[len(ctx.path.P)-2].Node)
	start, end := getRange(ctx, sel.start, sel.end, n)
	index := ctx.path.Last().Index
	return index >= start && index < end
}

func (sel *rangeSelector) selectChildNodes(ctx *context) (children iterator, canSelect bool) {
	if !isArray(ctx.path) {
		return emptyIterator{}, true
	}
	n := ctx.doc.Len(ctx.path.Last().Node)
	start, end := getRange(ctx, sel.start, sel.end, n)
	if start < 0 || start >= end {
		return emptyIterator{}, true
	}
	ret := make([]Segment, 0)
	for i := start; i < end && i < n; i++ {
		node := ctx.doc.Elem(ctx.path.Last().Node, i)
		ret = append(ret, Segment{
			Node:  node,
			Type:  ctx.doc.Type(node),
			Index: i,
		})
	}
	return &sliceIterator{
		items: ret,
	}, true
}

type sliceSelector struct {
	start expression
	end   expression
	step  expression
}

func (sel *sliceSelector) match(ctx *context) bool {
	if !isArrayElement(ctx.path) {
		return false
	}
	n := ctx.doc.Len(ctx.path.P[len(ctx.path.P)-2].Node)
	start, end := getRange(ctx, sel.start, sel.end, n)
	if start < 0 {
		return false
	}
	index := ctx.path.Last().Index
	if index < start || index >= end {
		return false
	}
	var step int
	if start < end {
		step = 1
	} else {
		step = -1
	}
	if sel.step != nil {
		ivalue, err := sel.step.evaluate(ctx)
		if err != nil {
			return false
		}
		var ok bool
		step, ok = valueAsInt(ivalue.value)
		if !ok {
			return false
		}
	}
	if step == 0 {
		return false
	}
	if start < end {
		if step < 0 {
			return false
		}
		return (index-start)%step == 0
	}
	if step > 0 {
		return false
	}
	return (start-index)%(-step) == 0
}

func (sel *sliceSelector) selectChildNodes(ctx *context) (children iterator, canSelect bool) {
	if !isArray(ctx.path) {
		return emptyIterator{}, true
	}
	n := ctx.doc.Len(ctx.path.Last().Node)
	start, end := getRange(ctx, sel.start, sel.end, n)
	if start < 0 {
		return emptyIterator{}, true
	}
	var step int
	if start < end {
		step = 1
	} else {
		step = -1
	}
	if sel.step != nil {
		ivalue, err := sel.step.evaluate(ctx)
		if err != nil {
			return nil, false
		}
		var ok bool
		step, ok = valueAsInt(ivalue.value)
		if !ok {
			return nil, false
		}
	}
	if step == 0 {
		return emptyIterator{}, true
	}
	var ret []Segment
	if start < end {
		if step < 0 {
			return emptyIterator{}, true
		}
		ret = make([]Segment, 0)
		for i := start; i < end && i < n; i += step {
			node := ctx.doc.Elem(ctx.path.Last().Node, i)
			ret = append(ret, Segment{
				Node:  node,
				Type:  ctx.doc.Type(node),
				Index: i,
			})
		}
		return &sliceIterator{
			items: ret,
		}, true
	}
	if step > 0 {
		return emptyIterator{}, true
	}
	if start >= n {
		return emptyIterator{}, true
	}
	ret = make([]Segment, 0)
	for i := start; i > end; i += step {
		node := ctx.doc.Elem(ctx.path.Last().Node, i)
		ret = append(ret, Segment{
			Node:  node,
			Type:  ctx.doc.Type(node),
			Index: i,
		})
	}
	return &sliceIterator{
		items: ret,
	}, true
}
