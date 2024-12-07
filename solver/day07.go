package solver

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
)

//go:embed inputs/day_07.txt
var puzzle07 string

func init() {
	register(6, &Day07{})
}

type Day07 struct {
	input [][]int
}

func (d *Day07) LoadInput() {
	d.input = d.parseInput(strings.NewReader(puzzle07))
}

func (d *Day07) First() uint64 {
	var result uint64
	for _, equation := range d.input {
		res, nums := equation[0], equation[1:]
		if d.isValidFirst(res, nums) {
			result += uint64(res)
		}
	}
	return result
}

func (d *Day07) Second() uint64 {
	var result uint64
	for _, equation := range d.input {
		res, nums := equation[0], equation[1:]
		if d.isValidSecond(res, nums) {
			result += uint64(res)
		}
	}
	return result
}

func (d *Day07) parseInput(input io.Reader) [][]int {
	var rows [][]int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, ": ")

		var row []int
		res, _ := strconv.Atoi(parts[0])
		row = append(row, res)

		for _, value := range strings.Split(parts[1], " ") {
			n, _ := strconv.Atoi(value)
			row = append(row, n)
		}

		rows = append(rows, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return rows
}

func (d *Day07) isValidFirst(result int, nums []int) bool {
	if len(nums) == 1 {
		return result == nums[0]
	}

	return d.isValidFirst(result, append([]int{nums[0] + nums[1]}, nums[2:]...)) ||
		d.isValidFirst(result, append([]int{nums[0] * nums[1]}, nums[2:]...))
}

func (d *Day07) isValidSecond(result int, nums []int) bool {
	if len(nums) == 1 {
		return result == nums[0]
	}

	concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", nums[0], nums[1]))

	return d.isValidSecond(result, append([]int{nums[0] + nums[1]}, nums[2:]...)) ||
		d.isValidSecond(result, append([]int{nums[0] * nums[1]}, nums[2:]...)) ||
		d.isValidSecond(result, append([]int{concat}, nums[2:]...))
}
