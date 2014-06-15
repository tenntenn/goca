package goca

import (
	"fmt"
	"io"
	"os"
)

type Simulator struct {
	CA         CA
	Initilizer Initilizer
	Logger     ProgressLogger
	Writer     PatternWriter
}

func NewSimulator(ca CA, ini Initilizer, w PatternWriter) *Simulator {
	return &Simulator{
		CA:         ca,
		Initilizer: ini,
		Logger:     &SimpleLogger{os.Stdout},
		Writer:     w,
	}
}

type ProgressLogger interface {
	Progress(step int, ca CA)
}

type SimpleLogger struct {
	Writer io.Writer
}

func (sl *SimpleLogger) Progress(step int, ca CA) {
	fmt.Fprintln(sl.Writer, "Step:", step)
}

type Initilizer interface {
	Initilize(ca CA)
}

func (s *Simulator) Run(step int) {
	s.Initilizer.Initilize(s.CA)
	for i := 1; i <= step; i++ {
		s.Logger.Progress(i, s.CA)
		s.CA.Transit()
		s.Writer.Write(s.CA.Pattern())
	}
}
