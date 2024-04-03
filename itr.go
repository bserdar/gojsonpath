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

type sliceIterator struct {
	items []Segment
	seg   Segment
}

func (s *sliceIterator) next() bool {
	if len(s.items) == 0 {
		return false
	}
	s.seg = s.items[0]
	s.items = s.items[1:]
	return true
}

func (s *sliceIterator) item(seg *Segment) { *seg = s.seg }
