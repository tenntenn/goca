package main

import (
	"../../goca"
	"fmt"
	"image/png"

	"os"
)

func main() {
	size := uint64(500)
	rule := goca.ECARule(90)
	ca := goca.NewECA(rule, size)
	//	writer := goca.Text1DWriter{os.Stdout}
	sim := goca.NewSimulator(ca)
	sim.Initilizer = goca.HandlerFunc(func(step int, ca goca.CA) error {
		ca.Pattern().Set(goca.State(1), int64(size/2))
		sim.Context["patterns"] = []goca.Pattern{ca.Pattern().Copy()}
		//writer.Write(ca.Pattern())
		return nil
	})
	sim.StepAfter = goca.HandlerFunc(func(step int, ca goca.CA) error {
		p := ca.Pattern()
		patterns, _ := sim.Context["patterns"].([]goca.Pattern)
		sim.Context["patterns"] = append(patterns, p.Copy())
		//		writer.Write(p)
		return nil
	})
	sim.Terminator = goca.HandlerFunc(func(step int, ca goca.CA) error {
		f, _ := os.Create(fmt.Sprintf("rule%d.png", rule))
		patterns, _ := sim.Context["patterns"].([]goca.Pattern)
		img := goca.SquareLattice1DImage{
			Patterns:     patterns,
			StateToColor: goca.MonochroicConverter{0, 1},
		}
		return png.Encode(f, img)
	})

	sim.Run(500)
}
