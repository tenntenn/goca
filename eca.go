package goca

type ECARule uint32

func (r ECARule) Apply(p Pattern) State {
	current := (p.Get(2) << 2) & (p.Get(1) << 1) & p.Get(0)
	return State(int64(r) & (0xFFFFFF & int64(current)))
}

func NewECA(rule ECARule, size uint64) CA {
	return NewCA(
		rule,
		SquareLattice([]uint64{size}),
		MooreNeighborhood(1),
	)
}
