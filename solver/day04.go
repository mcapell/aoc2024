package solver

import (
	"bufio"
	_ "embed"
	"io"
	"strings"
)

//go:embed inputs/day_04.txt
var puzzle04 string

func init() {
	register(4, &Day04{})
}

type Day04 struct {
	input []string
}

// Parses and loads the input
func (d *Day04) LoadInput() {
	d.input = d.parseInput(strings.NewReader(puzzle04))
}

// Solve the first problem, return the result
func (d *Day04) First() uint64 {
	words := d.findAllWords()
	count := 0
	for _, word := range words {
		if word == "XMAS" || word == "SAMX" {
			count++
		}
	}
	return uint64(count)
}

// Solve the second problem, return the result
func (d *Day04) Second() uint64 {
	return d.countXmasWords()
}

func (d *Day04) parseInput(input io.Reader) []string {
	var rows []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		rows = append(rows, strings.TrimSpace(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return rows
}

func (d *Day04) findAllWords() []string {
	var words []string
	// Find horizontal words
	for r := 0; r < len(d.input); r++ {
		for c := 0; c < len(d.input[r])-3; c++ {
			words = append(words, string([]byte{d.input[r][c], d.input[r][c+1], d.input[r][c+2], d.input[r][c+3]}))
		}
	}

	// Find vertical words
	for r := 0; r < len(d.input)-3; r++ {
		for c := 0; c < len(d.input[r]); c++ {
			words = append(words, string([]byte{d.input[r][c], d.input[r+1][c], d.input[r+2][c], d.input[r+3][c]}))
		}
	}

	// Find diagonal / words
	for r := 3; r < len(d.input); r++ {
		for c := 0; c < len(d.input[r])-3; c++ {
			words = append(words, string([]byte{d.input[r][c], d.input[r-1][c+1], d.input[r-2][c+2], d.input[r-3][c+3]}))
		}
	}

	// Find diagonal \ words
	for r := 0; r < len(d.input)-3; r++ {
		for c := 0; c < len(d.input[r])-3; c++ {
			words = append(words, string([]byte{d.input[r][c], d.input[r+1][c+1], d.input[r+2][c+2], d.input[r+3][c+3]}))
		}
	}

	return words
}

func (d *Day04) countXmasWords() uint64 {
	var count uint64

	// For each 3x3 matrix, check if X-MAS word exist
	for r := 0; r < len(d.input)-2; r++ {
		for c := 0; c < len(d.input[r])-2; c++ {
			// Ensure A is at the center
			if d.input[r+1][c+1] != 'A' {
				continue
			}

			// Check if M and S are at the correct place
			option1 := d.input[r][c] == 'M' && d.input[r+2][c+2] == 'S' && d.input[r][c+2] == 'S' && d.input[r+2][c] == 'M'
			option2 := d.input[r][c] == 'S' && d.input[r+2][c+2] == 'M' && d.input[r][c+2] == 'M' && d.input[r+2][c] == 'S'
			option3 := d.input[r][c] == 'M' && d.input[r+2][c+2] == 'S' && d.input[r][c+2] == 'M' && d.input[r+2][c] == 'S'
			option4 := d.input[r][c] == 'S' && d.input[r+2][c+2] == 'M' && d.input[r][c+2] == 'S' && d.input[r+2][c] == 'M'
			if option1 || option2 || option3 || option4 {
				count++
			}
		}
	}

	return count
}
