package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	day := flag.Int("day", 0, "Day of Advent of Code")
	flag.Parse()

	if *day == 0 {
		fmt.Println("Missing day argument")
		os.Exit(1)
	}

	tmpl := newTemplate(*day)

	writeFile(tmpl.solverPath("solver/"), tmpl.solverContent())
	writeFile(tmpl.testPath("solver/"), tmpl.testContent())
	writeFile(tmpl.inputPath("solver/inputs/"), tmpl.inputContent())
}

func writeFile(filePath, content string) {
	_, err := os.Stat(filePath)
	if err == nil {
		panic("the file already exist")
	}

	fmt.Printf("creating file %s\n", filePath)
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}
