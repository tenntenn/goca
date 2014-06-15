package goca

type Simulator struct {
	CA         CA
	Initilizer Handler
	StepBefore Handler
	StepAfter  Handler
	Terminator Handler
}

type Handler interface {
	Handle(step int, ca CA) error
}

type HandlerFunc func(step int, ca CA) error

func (f HandlerFunc) Handle(step int, ca CA) error {
	return f(step, ca)
}

func (s *Simulator) Run(step int) {
	if s.Initilizer != nil {
		s.Initilizer.Handle(0, s.CA)
	}
	for i := 1; i <= step; i++ {
		if s.StepBefore != nil {
			s.StepBefore.Handle(i, s.CA)
		}
		s.CA.Transit()
		if s.StepAfter != nil {
			s.StepAfter.Handle(i, s.CA)
		}
	}
}
