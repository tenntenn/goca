package goca

// type of transition rule of cellular automaton
type Rule interface {
	Apply(p Pattern) State
}

type RuleFunc func(p Pattern) State

func (f RuleFunc) Apply(p Pattern) State {
	return f(p)
}

// type of a cellular automata
type CA interface {
	Pattern() Pattern
	Transit() error
	Rule() Rule
	Topology() Topology
	Neighborhood() Neighborhood
}

// simple implmentation of CA
type SimpleCA struct {
	current      Pattern
	next         Pattern
	rule         Rule
	topology     Topology
	neighborhood Neighborhood
}

// create simple cellular automata
func NewCA(r Rule, t Topology, n Neighborhood) CA {
	return &SimpleCA{
		current:      NewPattern(t),
		next:         NewPattern(t),
		rule:         r,
		topology:     t,
		neighborhood: n,
	}
}

// return current pattern
func (ca *SimpleCA) Pattern() Pattern {
	return ca.current
}

// transition next states
func (ca *SimpleCA) Transit() error {
	for i := uint64(0); i < ca.topology.NumCells(); i++ {
		neighbors, err := ca.neighborhood.Get(ca.current, i)
		if err != nil {
			return err
		}
		ca.next.SetAtIndex(ca.rule.Apply(neighbors), i)
	}
	ca.current, ca.next = ca.next, ca.current
	return nil
}

// return transition rule of ca
func (ca *SimpleCA) Rule() Rule {
	return ca.rule
}

// return topology of ca
func (ca *SimpleCA) Topology() Topology {
	return ca.topology
}

// return neighborhood getting algorithm
func (ca *SimpleCA) Neighborhood() Neighborhood {
	return ca.neighborhood
}
