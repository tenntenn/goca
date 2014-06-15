package goca

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

// topology of cellular automata
// such as square lattice
type Topology interface {
	EndCoordinates() []int64
	NumCells() int64
	Index(coordinates ...int64) int64
	Coordinates(i int64) []int64
}

// Topology of square lattice cellular automata
type SquareLattice []uint64

// end value of each coordinates
func (sl SquareLattice) EndCoordinates() []uint64 {
	return []uint64(sl)
}

// calc number of cells
func (sl SquareLattice) NumCells() uint64 {
	l := 1
	for _, s := range sl {
		l *= s
	}
	return l
}

// convert coordinates to index
func (sl SquareLattice) Index(coordinates ...int64) uint64 {
	index := sl.Cycle(coordinates[0])
	for i := 1; i < len(coordinates); i++ {
		sum := sl.Cycle(coordinates[i])
		for j := 0; j < i; j++ {
			sum *= sl[j]
		}
		index += sum
	}
	return index
}

// convert index to coordinates
func (sl SquareLattice) Coordinates(i int64) []uint64 {
	index := sl.CycleIndex(i)
	coordinates := make([]uint64, len(sl))
	for i := len(sl) - 1; i > 0; i++ {
		preDimFull := 1
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

func (sl SquareLattice) Cycle(coordinate int64, size uint64) uint64 {
	if coordinate < 0 {
		return size - 1 - (size-coordinate-1)%size
	}
	return coordinate % size
}
