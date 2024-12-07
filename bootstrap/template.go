package main

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

type Template struct {
	day       int
	dayFormat string
}

const dayTemplate = `
package solver

import (
	"bufio"
	_ "embed"
	"io"
	"strings"
)

//go:embed inputs/day_XX.txt
var puzzleXX string

func init() {
	register(NN, &DayXX{})
}

type DayXX struct {
}

func (d *DayXX) LoadInput() {
	panic("not implemented")
}

func (d *DayXX) First() uint64 {
	panic("not implemented")
}

func (d *DayXX) Second() uint64 {
	panic("not implemented")
}

func (d *DayXX) parseInput(input io.Reader) []string {
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
}`

const testTemplate = `
package solver

import "testing"

func TestDayXX(t *testing.T) {
	var (
		d = &Day07{}
	)
	t.Run("first problem", func(t *testing.T) {
		var (
			expected uint64 = 1
			result          = d.First()
		)
		if expected != result {
			t.Errorf("expected %d. got=%d", expected, result)
		}
	})

	t.Run("second problem", func(t *testing.T) {
		var (
			expected uint64 = 1
			result          = d.Second()
		)
		if expected != result {
			t.Errorf("expected %d. got=%d", expected, result)
		}
	})
}
`

func newTemplate(day int) Template {
	dayFormat := fmt.Sprintf("%0.2d", day)
	return Template{
		day:       day,
		dayFormat: dayFormat,
	}
}

func (t *Template) solverPath(dir string) string {
	return filepath.Join(dir, fmt.Sprintf("day%s.go", t.dayFormat))
}

func (t *Template) testPath(dir string) string {
	return filepath.Join(dir, fmt.Sprintf("day%s_test.go", t.dayFormat))
}

func (t *Template) inputPath(dir string) string {
	return filepath.Join(dir, fmt.Sprintf("day_%s.txt", t.dayFormat))
}

func (t *Template) solverContent() string {
	s := strings.ReplaceAll(dayTemplate, "XX", t.dayFormat)
	return strings.ReplaceAll(s, "NN", fmt.Sprintf("%d", t.day-1))
}

func (t *Template) testContent() string {
	return strings.ReplaceAll(testTemplate, "XX", t.dayFormat)
}

func (t *Template) inputContent() string {
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", t.day)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(b)
}
