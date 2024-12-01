package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mcapell/aoc2024/solver"
)

func main() {
	if len(os.Args) > 0 && os.Args[0] == "all" {
		for _, d := range solver.GetAll() {
			runTimed(d)
		}
		return
	}

	runTimed(solver.GetLast())
}

func runTimed(problem solver.Solver) {
	problem.LoadInput()
	t := time.Now()

	fmt.Printf("First result: %d\n", problem.First())
	first := time.Since(t)

	fmt.Printf("Second result: %d\n", problem.Second())
	second := time.Since(t)

	fmt.Printf("Run %T: First: %dμs, Second: %dμs\n", problem, first.Microseconds(), second.Microseconds())
}
