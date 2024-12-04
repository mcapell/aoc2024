package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mcapell/aoc2024/solver"
)

func main() {
	var runAll bool
	flag.BoolVar(&runAll, "all", false, "Run all problems")
	flag.Parse()

	if runAll {
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
	fmt.Printf("Running %T\n", problem)

	fmt.Printf("\tFirst result: %d\n", problem.First())
	first := time.Since(t)

	fmt.Printf("\tSecond result: %d\n", problem.Second())
	second := time.Since(t)

	fmt.Printf("\tFirst solved in %dμs, Second in %dμs\n", first.Microseconds(), second.Microseconds())
}
