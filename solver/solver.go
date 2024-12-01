package solver

var solvers [24]Solver

type Solver interface {
	// Parses and loads the input
	LoadInput()
	// Solve the first problem, return the result
	First() int
	// Solve the second problem, return the result
	Second() int
}

// Get all registered problems
func GetAll() []Solver {
	return solvers[:]
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
