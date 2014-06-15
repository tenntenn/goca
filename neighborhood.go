package goca

import (
	"fmt"
)

var (
	NotSuportTopologyError = fmt.Errorf("NotSuportTopologyError")
)

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
	topology, ok := p.Toplogy.(SquareLattice)
	if !ok {
		return EmptyPattern, NotSuportTopologyError
	}

	sl := make(SquareLattice, len(topology))
	for i := 0; i < len(sl); i++ {
		sl[i] = 2*uint64(m) + 1
	}

	neighbors := NewPattern(sl)

	coordinates := topology.Coordinates(index)
	r := int64(m)
	c := make([]int64, len(sl))
	for i := -r; i < r; i++ {
		for j, _ := range coordinates {
			c[j] = int64(coordinates[j]) + i
		}
		neighbors.Set(p.GetByIndex(index), c...)
	}

	return neighbors, nil
}
