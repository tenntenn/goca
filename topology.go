package goca

import (
	"fmt"
)

// XYZ axis
type Axis int

const (
	// x axis
	X Axis = iota
	// y axis
	Y
	// z axis
	Z
)

var (
	// Neighborhood or rule does not support a topology
	NotSuportTopologyError = fmt.Errorf("NotSuportTopologyError")
)

// topology of cellular automata
// such as square lattice
type Topology interface {
	EndCoordinates() []uint64
	NumCells() uint64
	Index(coordinates ...int64) uint64
	Coordinates(i uint64) []uint64
	CycleIndex(i int64) uint64
	Cycle(coordinate int64, size uint64) uint64
}

// Topology of square lattice cellular automata
type SquareLattice []uint64

// end value of each coordinates
func (sl SquareLattice) EndCoordinates() []uint64 {
	return []uint64(sl)
}

// calc number of cells
func (sl SquareLattice) NumCells() uint64 {
	l := uint64(1)
	for _, s := range sl {
		l *= s
	}
	return l
}

// convert coordinates to index
func (sl SquareLattice) Index(coordinates ...int64) uint64 {
	index := sl.Cycle(coordinates[0], sl[0])
	for i := 1; i < len(coordinates); i++ {
		sum := sl.Cycle(coordinates[i], sl[i])
		for j := 0; j < i; j++ {
			sum *= sl[j]
		}
		index += sum
	}
	return index
}

// convert index to coordinates
func (sl SquareLattice) Coordinates(i uint64) []uint64 {
	index := sl.CycleIndex(int64(i))
	coordinates := make([]uint64, len(sl))
	for i := len(sl) - 1; i > 0; i++ {
		preDimFull := uint64(1)
		for j := 0; j < i-1; j++ {
			preDimFull *= sl[j]
		}
		coordinates[i] = index % preDimFull
		index = index - coordinates[i]
	}
	coordinates[0] = index
	return coordinates
}

// round index by boundary condition
func (sl SquareLattice) CycleIndex(i int64) uint64 {
	return sl.Cycle(i, sl.NumCells())
}

// round coordinate by boundary condition
func (sl SquareLattice) Cycle(coordinate int64, size uint64) uint64 {
	s := int64(size)
	if coordinate < 0 {
		return uint64(s - 1 - (s-coordinate-1)%s)
	}
	return uint64(coordinate % s)
}
