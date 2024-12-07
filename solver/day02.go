package solver

import (
	"bufio"
	_ "embed"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/mcapell/aoc2024/utils/slices"
)

//go:embed inputs/day_02.txt
var puzzle02 string

func init() {
	register(1, &Day02{})
}

type Day02 struct {
	input [][]int
}

// Parses and loads the input
func (d *Day02) LoadInput() {
	d.input = d.parseInput(strings.NewReader(puzzle02))
}

// Solve the first problem, return the result
func (d *Day02) First() uint64 {
	var safeReports uint64

	for _, report := range d.input {
		if d.reportIsSafe(report) {
			safeReports++
		}
	}

	return safeReports
}

// Solve the second problem, return the result
func (d *Day02) Second() uint64 {
	var safeReports uint64

	for _, report := range d.input {
		if d.reportIsSafe(report) {
			safeReports++
			continue
		}

		// Create a new slice without one of the levels
		for i := 0; i < len(report); i++ {
			p1 := slices.Copy(report[:i])
			p2 := slices.Copy(report[i+1:])
			newReport := append(p1, p2...)
			if d.reportIsSafe(newReport) {
				safeReports++
				break
			}
		}
	}

	return safeReports
}

func (d *Day02) parseInput(input io.Reader) [][]int {
	var reports [][]int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		var report []int
		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				panic(err)
			}

			report = append(report, num)
		}

		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return reports
}

func (d *Day02) reportIsSafe(report []int) bool {
	// check trend
	if !sort.IntsAreSorted(report) && !sort.IntsAreSorted(ReverseNewSlice(report)) {
		return false
	}

	for i := 0; i < len(report)-1; i++ {
		// check difference between adjacent levels
		diff := int(math.Abs(float64(report[i+1] - report[i])))
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}
