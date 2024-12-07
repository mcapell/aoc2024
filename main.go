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

	fmt.Printf("\tFirst solved in %d%s, Second in %d%s\n",
		getScaledValue(first.Microseconds()),
		getUnit(first.Microseconds()),
		getScaledValue(second.Microseconds()),
		getUnit(second.Microseconds()))
}

func getUnit(micros int64) string {
	var (
		units = []string{"μs", "ms", "s"}
		value int
	)

	for value = 0; micros > 1000 && value < 3; value++ {
		micros /= 1000
	}

	return units[value]
}

func getScaledValue(micros int64) int64 {
	for i := 0; micros > 1000 && i < 3; i++ {
		micros /= 1000
	}

	return micros
}
