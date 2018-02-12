package main

import (
	"fmt"
	"github.com/kenneth-fischer/Advent/files"
	"log"
	"strconv"
	"strings"
)

func main() {
	printResult(rowChecksum, "rowChecksum", 5, 1, 9, 5)
	printResult(rowChecksum, "rowChecksum", 7, 5, 3)
	printResult(rowChecksum, "rowChecksum", 2, 4, 6, 8)
	spreadsheet := [][]int{[]int{5, 1, 9, 5}, []int{7, 5, 3}, []int{2, 4, 6, 8}}
	fmt.Printf("Total checksum: %d\n", checksum(spreadsheet, rowChecksum))

	bigSpreadsheet, err := readSpreadsheet("checksum.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := checksum(bigSpreadsheet, rowChecksum)
	fmt.Printf("Big Checksum: %d\n", sum)
	fmt.Println()

	printResult(rowQuotient, "rowQuotient", 5, 9, 2, 8)
	printResult(rowQuotient, "rowQuotient", 9, 4, 7, 3)
	printResult(rowQuotient, "rowQuotient", 3, 8, 6, 5)
	spreadsheet = [][]int{[]int{5, 9, 2, 8}, []int{9, 4, 7, 3}, []int{3, 8, 6, 5}}
	fmt.Printf("Total checksum: %d\n", checksum(spreadsheet, rowQuotient))

	sum = checksum(bigSpreadsheet, rowQuotient)
	fmt.Printf("Big Checksum: %d\n", sum)

}

func readSpreadsheet(path string) ([][]int, error) {
	lines, err := files.ReadLines(path)
	if err != nil {
		return nil, err
	}

	rows := [][]int{}
	for _, line := range lines {
		row, err := readRow(line)
		if err != nil {
			return rows, err
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func readRow(line string) ([]int, error) {
	row := []int{}

	for _, str := range strings.Split(line, "\t") {
		value, err := strconv.Atoi(str)
		if err != nil {
			return row, err
		}
		row = append(row, value)
	}
	return row, nil
}

func printResult(f func(values ...int) int, fname string, values ...int) {
	args := ""
	for i := 0; i < 5; i++ {
		if i >= len(values) {
			break
		}
		if i == 0 {
			args = fmt.Sprintf("%d", values[i])
		} else {
			args = fmt.Sprintf("%s, %d", args, values[i])
		}
	}
	if len(values) > 5 {
		args = fmt.Sprintf("%s, ...", args)
	}
	fmt.Printf("%s(%s) = %d\n", fname, args, f(values...))
}

func checksum(grid [][]int, rowFunc func(values ...int) int) int {
	sum := 0
	for _, row := range grid {
		sum += rowFunc(row...)
	}
	return sum
}

func rowQuotient(values ...int) int {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i]%values[j] == 0 {
				return values[i] / values[j]
			} else if values[j]%values[i] == 0 {
				return values[j] / values[i]
			}
		}
	}
	return 0
}

func rowChecksum(values ...int) int {
	gotBounds := false
	max, min := 0, 0
	for _, value := range values {
		if !gotBounds {
			max, min = value, value
			gotBounds = true
		} else {
			if value < min {
				min = value
			}
			if value > max {
				max = value
			}
		}
	}
	return max - min
}
