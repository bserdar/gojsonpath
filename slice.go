package gojsonpath

// Used for [index]
type indexSelector struct {
	index expression
}

// Determines if last element of el is an array element
func isArrayElement(el []Element) bool {
	if len(el) < 2 {
		return false
	}
	return el[len(el)-2].Type == ArrayNode
}

// Determines if last element of el is an array
func isArray(el []Element) bool {
	if len(el) < 1 {
		return false
	}
	return el[len(el)-1].Type == ArrayNode
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

func (sel *indexSelector) match(doc DocModel, el []Element) bool {
	if !isArrayElement(el) {
		return false
	}
	ivalue, err := sel.index.evaluate(doc, el)
	if err != nil {
		return false
	}
	index, ok := valueAsInt(ivalue.value)
	if !ok {
		return false
	}
	if index < 0 {
		n := doc.Len(el[len(el)-1].Node)
		index := normalizeIndex(index, n)
		if index < 0 {
			return false
		}
	}
	return el[len(el)-1].Index == index
}

func (sel *indexSelector) selectChildNodes(doc DocModel, el []Element) (children []Element, canSelect bool) {
	if !isArray(el) {
		return nil, true
	}
	n := doc.Len(el[len(el)-1].Node)
	ivalue, err := sel.index.evaluate(doc, el)
	if err != nil {
		return nil, false
	}
	index := normalizeIndex(ivalue.value, n)
	if index < 0 {
		return nil, true
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

// Used for [from:to], [:to], or [from:]
type rangeSelector struct {
	start expression
	end   expression
}

func getRange(doc DocModel, el []Element, start, end expression, len int) (from, to int) {
	var s int
	if start != nil {
		ivalue, _ := start.evaluate(doc, el)
		s = normalizeIndex(ivalue.value, len)
	}
	if s < 0 {
		return -1, -1
	}
	var e int
	if end != nil {
		ivalue, _ := end.evaluate(doc, el)
		e = normalizeIndex(ivalue.value, len)
	} else {
		e = len
	}
	if e < 0 {
		return -1, -1
	}
	return s, e
}

func (sel *rangeSelector) match(doc DocModel, el []Element) bool {
	if !isArrayElement(el) {
		return false
	}
	n := doc.Len(el[len(el)-2].Node)
	start, end := getRange(doc, el, sel.start, sel.end, n)
	index := el[len(el)-1].Index
	return index >= start && index < end
}

func (sel *rangeSelector) selectChildNodes(doc DocModel, el []Element) (children []Element, canSelect bool) {
	if !isArray(el) {
		return nil, true
	}
	n := doc.Len(el[len(el)-1].Node)
	start, end := getRange(doc, el, sel.start, sel.end, n)
	if start < 0 || start >= end {
		return nil, true
	}
	ret := make([]Element, 0)
	for i := start; i < end && i < n; i++ {
		node := doc.Elem(el[len(el)-1].Node, i)
		ret = append(ret, Element{
			Node:  node,
			Type:  doc.Type(node),
			Index: i,
		})
	}
	return ret, true
}

type sliceSelector struct {
	start expression
	end   expression
	step  expression
}

func (sel *sliceSelector) match(doc DocModel, el []Element) bool {
	if !isArrayElement(el) {
		return false
	}
	n := doc.Len(el[len(el)-2].Node)
	start, end := getRange(doc, el, sel.start, sel.end, n)
	if start < 0 {
		return false
	}
	index := el[len(el)-1].Index
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
		ivalue, err := sel.step.evaluate(doc, el)
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

func (sel *sliceSelector) selectChildNodes(doc DocModel, el []Element) (children []Element, canSelect bool) {
	if !isArray(el) {
		return nil, true
	}
	n := doc.Len(el[len(el)-1].Node)
	start, end := getRange(doc, el, sel.start, sel.end, n)
	if start < 0 {
		return nil, true
	}
	var step int
	if start < end {
		step = 1
	} else {
		step = -1
	}
	if sel.step != nil {
		ivalue, err := sel.step.evaluate(doc, el)
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
		return nil, true
	}
	var ret []Element
	if start < end {
		if step < 0 {
			return nil, true
		}
		ret = make([]Element, 0)
		for i := start; i < end && i < n; i += step {
			node := doc.Elem(el[len(el)-1].Node, i)
			ret = append(ret, Element{
				Node:  node,
				Type:  doc.Type(node),
				Index: i,
			})
		}
		return ret, true
	}
	if step > 0 {
		return nil, true
	}
	if start >= n {
		return nil, true
	}
	ret = make([]Element, 0)
	for i := start; i > end; i += step {
		node := doc.Elem(el[len(el)-1].Node, i)
		ret = append(ret, Element{
			Node:  node,
			Type:  doc.Type(node),
			Index: i,
		})
	}
	return ret, true
}
