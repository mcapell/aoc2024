package solver

var solvers [24]Solver

type Solver interface {
	// Parses and loads the input
	LoadInput()
	// Solve the first problem, return the result
	First() uint64
	// Solve the second problem, return the result
	Second() uint64
}

// Get all registered problems
func GetAll() []Solver {
	var r []Solver
	for _, s := range solvers {
		if s != nil {
			r = append(r, s)
		}
	}
	return r
}

// Get the last registered problem
func GetLast() Solver {
	for i := 23; i >= 0; i-- {
		if p := solvers[i]; p != nil {
			return p
		}
	}

	panic("no solvers loaded")
}

func register(i int, p Solver) {
	solvers[i] = p
}
