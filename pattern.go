package goca

// state of a cell of celluar automata
type State int64

// pattern of celluar automata
type Pattern struct {
	States  []State
	Toplogy Topology
}

var EmptyPattern = Pattern{nil, nil}

// create pattern from toplogy
func NewPattern(toplogy Topology) Pattern {
	return Pattern{
		states:  make([]State, toplogy.Length()),
		toplogy: toplogy,
	}
}

// get state with given coordinates c
func (p Pattern) Get(c ...int64) State {
	return p.GetByIndex(p.toplogy.Index(coordinates))
}

// get state with given index i
func (p Pattern) GetByIndex(i uint64) State {
	index := p.toplogy.CycleIndex(i)
	return p.states[index]
}

// set state with given coordinates c
func (p Pattern) Set(s State, c ...int64) {
	p.SetByIndex(state, p.toplogy.Index(coordinates))
}

// set state with given index i
func (p Pattern) SetAtIndex(s State, i uint64) {
	index := p.toplogy.CycleIndex(i)
	p.states[index] = state
}
