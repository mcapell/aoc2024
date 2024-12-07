package solver

import (
	"bufio"
	_ "embed"
	"io"
	"strings"

	"github.com/mcapell/aoc2024/utils/slices"
)

type PositionWithDirection struct {
	X         int
	Y         int
	Direction int
}

//go:embed inputs/day_06.txt
var puzzle06 string

func init() {
	register(5, &Day06{})
}

type Day06 struct {
	input [][]byte
}

func (d *Day06) LoadInput() {
	d.input = d.parseInput(strings.NewReader(puzzle06))
}

func (d *Day06) First() int {
	pos := d.findStartingPosition()

	input := slices.DeepCopy2d(d.input)
	d.run(input, pos[0], pos[1], nil)

	var count int
	for _, row := range input {
		for _, v := range row {
			if v == 'X' {
				count++
			}
		}
	}
	return count
}

func (d *Day06) Second() int {
	pos := d.findStartingPosition()

	input := slices.DeepCopy2d(d.input)

	obstacles := map[int]PositionWithDirection{}
	d.run(input, pos[0], pos[1], obstacles)

	var numLoops int
	for _, obstacle := range obstacles {
		i := slices.DeepCopy2d(d.input)

		// Can't put an obstacle on guard's position
		if obstacle.Y == pos[0] && obstacle.X == pos[1] {
			continue
		}

		// Place a new obstacle
		i[obstacle.Y][obstacle.X] = '#'

		// Check if the new input has a loop
		if d.run(i, pos[0], pos[1], nil) {
			numLoops++
		}
	}

	return numLoops
}

func (d *Day06) parseInput(input io.Reader) [][]byte {
	var rows [][]byte
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		var row []byte
		for _, c := range line {
			row = append(row, byte(c))
		}
		rows = append(rows, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return rows
}

func (d *Day06) findStartingPosition() []int {
	for row, rowValues := range d.input {
		for col := range rowValues {
			if d.input[row][col] == '^' {
				return []int{row, col}
			}
		}
	}

	panic("starting position not found")
}

func (d *Day06) run(input [][]byte, row, col int, obstacles map[int]PositionWithDirection) bool {
	var (
		up    int = 0
		right int = 1
		down  int = 2
		left  int = 3

		direction        = up
		visitedPositions = map[PositionWithDirection]bool{}
	)

	for {
		input[row][col] = 'X'
		pos := PositionWithDirection{
			X:         col,
			Y:         row,
			Direction: direction,
		}
		if _, exist := visitedPositions[pos]; exist {
			return true
		}

		visitedPositions[pos] = true

		for d.peekNext(input, row, col, direction) == '#' {
			direction = (direction + 1) % 4
		}

		switch direction {
		case up:
			row--
		case right:
			col++
		case down:
			row++
		case left:
			col--
		}

		if row < 0 || row == len(input) || col < 0 || col == len(input[0]) {
			break
		}

		if obstacles != nil {
			obstacles[row*1000+col] = PositionWithDirection{
				X: col,
				Y: row,
			}
		}
	}

	return false
}

func (d *Day06) peekNext(input [][]byte, row, col, direction int) byte {
	var (
		up    int = 0
		right int = 1
		down  int = 2
		left  int = 3

		next byte
	)

	switch direction {
	case up:
		if row-1 < 0 {
			break
		}
		next = input[row-1][col]
	case right:
		if col+1 == len(input[0]) {
			break
		}
		next = input[row][col+1]
	case down:
		if row+1 == len(input) {
			break
		}
		next = input[row+1][col]
	case left:
		if col-1 < 0 {
			break
		}
		next = input[row][col-1]
	}

	return next
}
