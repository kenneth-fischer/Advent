package main

import (
	"fmt"
	"github.com/kenneth-fischer/Advent/files"
	"log"
	"strconv"
)

var (
	pc           int
	steps        int
	instructions []int
)

func main() {
	run(true, increment, 0, 3, 0, 1, -3)
	jumps := loadInstructions("jumps.txt")
	run(false, increment, jumps...)
	fmt.Println()
	run(true, condIncrement, 0, 3, 0, 1, -3)
	jumps = loadInstructions("jumps.txt")
	run(false, condIncrement, jumps...)
}

func loadInstructions(path string) []int {
	ints := []int{}
	lines, err := files.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		myInt, _ := strconv.Atoi(line)
		ints = append(ints, myInt)
	}
	return ints
}

func run(debug bool, nextInstr func(int) int, jumps ...int) {
	pc = 0
	steps = 0
	instructions = jumps
	for done := isDone(); !done; done = isDone() {
		if debug {
			printState()
		}
		step(nextInstr)
	}
	if debug {
		printState()
	}
	fmt.Printf("Finished in %d steps\n", steps)
}

func printState() {
	fmt.Printf("%2d.", steps)
	for i := 0; i < len(instructions); i++ {
		if i == pc {
			fmt.Printf(" (%d)", instructions[i])
		} else {
			fmt.Printf("  %d ", instructions[i])
		}
	}
	fmt.Println()
}

func step(nextInstr func(int) int) {
	jump := instructions[pc]
	instructions[pc] = nextInstr(instructions[pc])
	steps += 1
	pc += jump
}

func increment(jump int) int {
	return jump + 1
}

func condIncrement(jump int) int {
	if jump >= 3 {
		return jump - 1
	}
	return jump + 1
}

func isDone() bool {
	return pc >= len(instructions)
}
