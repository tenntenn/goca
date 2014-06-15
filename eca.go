package goca

type ECARule uint32

func (r ECARule) Apply(p Pattern) State {
	current := (p.Get(0) << 2) | (p.Get(1) << 1) | p.Get(2)
	if int64(r)&(1<<uint64(current)) != 0 {
		return State(1)
	}
	return State(0)
}

func NewECA(rule ECARule, size uint64) CA {
	return NewCA(
		rule,
		SquareLattice([]uint64{size}),
		MooreNeighborhood(1),
	)
}
