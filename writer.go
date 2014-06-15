package goca

import (
	"fmt"
	"io"
)

type PatternWriter interface {
	Write(p Pattern) error
}

type WriterHandler struct {
	PatternWriter
}

func (wh *WriterHandler) Handle(step int, ca CA) error {
	return wh.Write(ca.Pattern())
}

type Text1DWriter struct {
	Writer io.Writer
}

func (w Text1DWriter) Write(p Pattern) error {
	topology, ok := p.Topology.(SquareLattice)
	if !ok || len(topology) != 1 {
		return NotSuportTopologyError
	}

	end := topology.EndCoordinates()
	for x := int64(0); x < int64(end[X]); x++ {
		s := p.Get(x)
		fmt.Fprint(w.Writer, s)
	}
	fmt.Fprintln(w.Writer)
	return nil
}

type Text2DWriter struct {
	Writer io.Writer
}

func (w Text2DWriter) Write(p Pattern) error {
	topology, ok := p.Topology.(SquareLattice)
	if !ok || len(topology) != 2 {
		return NotSuportTopologyError
	}

	end := topology.EndCoordinates()
	for y := int64(0); y < int64(end[Y]); y++ {
		for x := int64(0); x < int64(end[X]); x++ {
			s := p.Get(x, y)
			fmt.Fprint(w.Writer, s)
		}
		fmt.Fprintln(w.Writer)
	}
	return nil
}
