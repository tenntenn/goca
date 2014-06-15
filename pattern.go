package goca

// state of a cell of celluar automata
type State int64

// pattern of celluar automata
type Pattern struct {
	States  []State
	Toplogy Topology
}

// Empty pattern
var EmptyPattern = Pattern{nil, nil}

// create pattern from toplogy
func NewPattern(t Topology) Pattern {
	return Pattern{
		States:  make([]State, t.NumCells()),
		Toplogy: t,
	}
}

// get state with given coordinates c
func (p Pattern) Get(c ...int64) State {
	return p.GetByIndex(p.Toplogy.Index(c...))
}

// get state with given index i
func (p Pattern) GetByIndex(i uint64) State {
	index := p.Toplogy.CycleIndex(int64(i))
	return p.States[index]
}

// set state with given coordinates c
func (p Pattern) Set(s State, c ...int64) {
	p.SetAtIndex(s, p.Toplogy.Index(c...))
}

// set state with given index i
func (p Pattern) SetAtIndex(s State, i uint64) {
	index := p.Toplogy.CycleIndex(int64(i))
	p.States[index] = s
}
