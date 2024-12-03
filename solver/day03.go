package solver

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

//go:embed inputs/day_03.txt
var puzzle03 string
var instructionRe *regexp.Regexp

type Instruction interface {
	GetType() string
}

type MulInstruction struct {
	Left  int
	Right int
}

func (m *MulInstruction) GetType() string { return "mul" }

type DontInstruction struct{}

func (m *DontInstruction) GetType() string { return "dont" }

type DoInstruction struct{}

func (m *DoInstruction) GetType() string { return "do" }

func init() {
	register(2, &Day03{})
	instructionRe = regexp.MustCompile(`(mul\(\d+,\d+\))|(don\'t\(\))|(do\(\))`)
}

type Day03 struct {
	input []string
}

// Parses and loads the input
func (d *Day03) LoadInput() {
	d.input = d.parseInput(strings.NewReader(puzzle03))
}

// Solve the first problem, return the result
func (d *Day03) First() int {
	valid := d.getValidInstructions(d.input)
	var result int
	for _, instruction := range valid {
		if ins, ok := instruction.(*MulInstruction); ok {
			result += ins.Left * ins.Right
		}
	}

	return result
}

// Solve the second problem, return the result
func (d *Day03) Second() (_ int) {
	valid := d.getValidInstructions(d.input)
	var result int
	var enable = true
	for _, instruction := range valid {
		if ins, ok := instruction.(*MulInstruction); ok && enable {
			result += ins.Left * ins.Right
		}

		if _, ok := instruction.(*DoInstruction); ok {
			enable = true
		}

		if _, ok := instruction.(*DontInstruction); ok {
			enable = false
		}

	}

	return result
}

func (d *Day03) parseInput(input io.Reader) []string {
	var memory []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		memory = append(memory, strings.TrimSpace(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return memory
}

func (d *Day03) getValidInstructions(input []string) []Instruction {
	var instructions []Instruction
	for _, page := range input {
		allInstructions := instructionRe.FindAllString(page, -1)
		for _, in := range allInstructions {
			if strings.HasPrefix(in, "mul(") {
				s := strings.TrimPrefix(strings.TrimSuffix(in, ")"), "mul(")
				parts := strings.Split(s, ",")

				left, _ := strconv.Atoi(parts[0])
				right, _ := strconv.Atoi(parts[1])
				instructions = append(instructions, &MulInstruction{Left: left, Right: right})
			} else if strings.HasPrefix(in, "do(") {
				instructions = append(instructions, &DoInstruction{})
			} else if strings.HasPrefix(in, "don't(") {
				instructions = append(instructions, &DontInstruction{})
			} else {
				panic(fmt.Sprintf("unexpected instruction: %s", in))
			}
		}
	}

	return instructions
}
