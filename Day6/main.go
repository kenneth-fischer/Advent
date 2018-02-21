package main

import (
	"fmt"
	"github.com/kenneth-fischer/Advent/files"
	"log"
	"strconv"
	"strings"
)

func main() {
	redistribute(true, 0, 2, 7, 0)
	fmt.Println()
	findCycle(true, 0, 2, 7, 0)
	fmt.Println()
	banks := loadBanks("banks.txt")
	findCycle(false, banks...)

}

func findCycle(debug bool, banksIn ...int) {
	configs := [][]int{banksIn}
	iterations := 0
	prev := -1

	for {
		banksOut := redistribute(false, banksIn...)
		iterations++
		if debug { 
			fmt.Printf("%3d. %v -> %v\n", iterations, banksIn, banksOut)
		}
		prev = indexOf(banksOut, configs) 
		if prev >= 0 {
			break
		}
		configs = append(configs, banksOut)
		banksIn = banksOut
	}
	fmt.Printf("Cycle found after %d iterations. %d iterations in loop\n", iterations, iterations - prev)
}

func indexOf(config []int, configs [][]int) int {
	for i, candidate := range configs {
		if matches(candidate, config) {
			return i
		}
	}
	return -1
}

func matches(banks1, banks2 []int) bool {
	if len(banks1) != len(banks2) {
		return false
	}
	for i := 0; i < len(banks1); i++ {
		if banks1[i] != banks2[i] {
			return false
		}
	}
	return true
}


func redistribute(debug bool, banksIn ...int) []int {
	banksOut := make([]int, len(banksIn))
	copy(banksOut, banksIn)
	chosen := 0
	for i := 1; i < len(banksOut); i++ {
		if banksOut[i] > banksOut[chosen] {
			chosen = i
		}
	}
	toBeMoved := banksOut[chosen]
	if debug { fmt.Println(banksOut) }
	banksOut[chosen] = 0
	if debug { fmt.Println(banksOut) }

	for i := 0; i < toBeMoved; i++ {
		chosen = (chosen + 1) % len(banksOut)
		banksOut[chosen]++
		if debug { fmt.Println(banksOut) }
	}
	if debug { fmt.Println(banksOut) }
	return banksOut
}

func loadBanks(path string) []int {
	ints := []int{}
	numbers, err := files.ReadString(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, number := range strings.Split(numbers, "\t") {
		myInt, _ := strconv.Atoi(number)
		ints = append(ints, myInt)
	}
	return ints
}

