package solver

import (
	"bufio"
	_ "embed"
	"io"
	"sort"
	"strconv"
	"strings"
)

//go:embed inputs/day_01.txt
var input string

func init() {
	register(0, &Day01{})
}

type Day01 struct {
	input [][]int
}

func (d *Day01) LoadInput() {
	first, second := parseInput(strings.NewReader(input))
	sort.Ints(first)
	sort.Ints(second)

	d.input = [][]int{first, second}
}

func (d *Day01) First() int {
	first, second := d.input[0], d.input[1]

	distance := 0
	for i := 0; i < len(first); i++ {
		distance += getDistance(first[i], second[i])
	}

	return distance
}

func (d *Day01) Second() int {
	first, second := d.input[0], d.input[1]

	// count how many times each value appears on list 2
	counts := make(map[int]int)
	for _, v := range second {
		if _, exist := counts[v]; !exist {
			counts[v] = 0
		}

		counts[v] += 1
	}

	// Get distances
	distance := 0
	for i := 0; i < len(first); i++ {
		distance += first[i] * counts[first[i]]
	}

	return distance
}

func parseInput(input io.Reader) ([]int, []int) {
	var (
		first  []int
		second []int
	)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "  ")
		if len(parts) == 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err1 == nil && err2 == nil {
				first = append(first, num1)
				second = append(second, num2)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return first, second
}

func getDistance(i, j int) int {
	if i < j {
		i, j = j, i
	}

	return i - j
}
