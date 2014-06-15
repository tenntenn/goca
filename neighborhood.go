package goca

// getting neighborhood algorithm
type Neighborhood interface {
	// get neighbors pattern with index i
	Get(p Pattern, i uint64) (Pattern, error)
}

// Moore neighborhood
// http://en.wikipedia.org/wiki/Moore_neighborhood
// MooreNeighborhood only support for SquareLattice.
type MooreNeighborhood uint64

// get neighbors of SquareLattice pattern.
// if topology of given pattern is not SquareLattice,
// NotSuportTopologyError occured.
func (m MooreNeighborhood) Get(p Pattern, index uint64) (Pattern, error) {
	topology, ok := p.Topology.(SquareLattice)
	if !ok {
		return EmptyPattern, NotSuportTopologyError
	}

	sl := make(SquareLattice, len(topology))
	length := 2*int64(m) + 1
	for i := 0; i < len(sl); i++ {
		sl[i] = uint64(length)
	}

	neighbors := NewPattern(sl)

	coordinates := topology.Coordinates(index)
	r := int64(m)
	from := make([]int64, len(sl))
	to := make([]int64, len(sl))
	for i := int64(0); i < length; i++ {
		for j, _ := range coordinates {
			from[j] = int64(coordinates[j]) + (i - r)
			to[j] = i
		}
		neighbors.Set(p.Get(from...), to...)
	}

	return neighbors, nil
}
