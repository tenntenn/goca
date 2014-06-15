package main

import (
	"../../goca"

	"os"
)

func main() {
	size := uint64(30)
	rule := goca.ECARule(90)
	ca := goca.NewECA(rule, size)
	writer := goca.Text1DWriter{os.Stdout}
	sim := &goca.Simulator{
		CA: ca,
		Initilizer: goca.HandlerFunc(func(step int, ca goca.CA) error {
			ca.Pattern().Set(goca.State(1), int64(size/2))
			writer.Write(ca.Pattern())
			return nil
		}),
		StepAfter: &goca.WriterHandler{writer},
	}

	sim.Run(20)
}
