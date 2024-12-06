package solver

import (
	"bufio"
	_ "embed"
	"io"
	"strconv"
	"strings"

	"github.com/mcapell/aoc2024/utils/slices"
)

//go:embed inputs/day_05.txt
var puzzle05 string

func init() {
	register(4, &Day05{})
}

type Day05 struct {
	rules map[int][]int
	pages [][]int
}

func (d *Day05) LoadInput() {
	rules, pages := d.parseInput(strings.NewReader(puzzle05))
	d.rules = rules
	d.pages = pages
}

func (d *Day05) First() int {
	var result int
	for _, page := range d.pages {
		if d.isCorrect(page) {
			result += page[len(page)/2]
		}
	}

	return result
}

func (d *Day05) Second() int {
	var result int
	for _, page := range d.pages {
		if d.isCorrect(page) {
			continue
		}

		for !d.isCorrect(page) {
			for k, previous := range d.rules {
				for _, v := range previous {
					posK := slices.IndexOf(page, k)
					posV := slices.IndexOf(page, v)
					if posK != -1 && posV != -1 && posK > posV {
						page[posK], page[posV] = page[posV], page[posK]
					}
				}
			}
		}

		result += page[len(page)/2]
	}
	return result
}

func (d *Day05) parseInput(input io.Reader) (map[int][]int, [][]int) {
	var (
		rules = map[int][]int{}
		pages = [][]int{}
	)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue

		}
		ruleParts := strings.Split(line, "|")
		if len(ruleParts) == 2 {
			k, _ := strconv.Atoi(ruleParts[0])
			v, _ := strconv.Atoi(ruleParts[1])
			if _, ok := rules[k]; !ok {
				rules[k] = []int{}
			}

			rules[k] = append(rules[k], v)
		} else if len(ruleParts) == 1 {
			var page []int
			p := strings.Split(line, ",")
			for _, s := range p {
				n, _ := strconv.Atoi(s)
				page = append(page, n)
			}

			pages = append(pages, page)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return rules, pages
}

func (d *Day05) isCorrect(page []int) bool {
	seen := map[int]bool{}
	isCorrect := true
	for _, p := range page {
		if r, ok := d.rules[p]; ok {
			for _, prev := range r {
				if _, found := seen[prev]; found {
					isCorrect = false
					break
				}
			}
		}
		if !isCorrect {
			break
		}
		seen[p] = true
	}

	return isCorrect
}
