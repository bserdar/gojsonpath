package gojsonpath

type iterator interface {
	next() bool
	item(*Segment)
}

type emptyIterator struct{}

func (emptyIterator) next() bool    { return false }
func (emptyIterator) item(*Segment) {}

type singleIterator struct {
	seg  Segment
	done bool
}

func (i *singleIterator) next() bool {
	ret := !i.done
	i.done = true
	return ret
}

func (i *singleIterator) item(s *Segment) { *s = i.seg }

type rangeIterator struct {
	doc            DocModel
	node           any
	from, to, step int
	at             int
	state          int
}

func (s *rangeIterator) next() bool {
	switch s.state {
	case 0: // Not started
		s.at = s.from
		if s.step > 0 {
			if s.at >= s.to {
				s.state = 2
				return false
			}
			s.state = 1
			return true
		}
		if s.at <= s.to {
			s.state = 2
			return false
		}
		s.state = 1
		return true

	case 1: // iterating
		s.at += s.step
		if s.step > 0 {
			if s.at >= s.to {
				s.state = 2
				return false
			}
			return true
		}
		if s.at <= s.to {
			s.state = 2
			return false
		}
		return true
	}
	return false
}

func (s *rangeIterator) item(seg *Segment) {
	seg.Node = s.doc.Elem(s.node, s.at)
	seg.Type = s.doc.Type(seg.Node)
	seg.Index = s.at
}
