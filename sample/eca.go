package main

import (
	"../../goca"

	"os"
)

func main() {
	size := uint64(30)
	rule := goca.ECARule(90)
	ca := goca.NewECA(rule, size)
	sim := goca.NewSimulator(ca, goca.InitFunc(func(ca goca.CA) {
		ca.Pattern().Set(goca.State(1), int64(size/2))
	}), &goca.Text1DWriter{os.Stdout})

	sim.Run(20)
}
