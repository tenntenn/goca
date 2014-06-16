package goca

type Simulator struct {
	CA         CA
	Initilizer Handler
	StepBefore Handler
	StepAfter  Handler
	Terminator Handler
	Context    map[string]interface{}
}

func NewSimulator(ca CA) *Simulator {
	return &Simulator{
		CA:      ca,
		Context: make(map[string]interface{}),
	}
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

	if s.Terminator != nil {
		s.Terminator.Handle(step, s.CA)
	}
}
