package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	printCapcha("1122", nextChar)
	printCapcha("1111", nextChar)
	printCapcha("91212129", nextChar)
	printCapcha("55512125", nextChar)
	sequence, _ := ReadString("input1")
	printCapcha(sequence, nextChar)

	fmt.Println()

	printCapcha("1212", halfwayRound)
	printCapcha("1221", halfwayRound)
	printCapcha("123425", halfwayRound)
	printCapcha("123123", halfwayRound)
	printCapcha("12131415", halfwayRound)
	sequence, _ = ReadString("input1")
	printCapcha(sequence, halfwayRound)
}

func ReadString(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func printCapcha(sequence string, next func(string, int) string) {
	caption := sequence[0:]
	if len(caption) > 10 {
		caption = caption[0:10] + "..."
	}
	fmt.Printf("Capcha(%q) = %d\n", caption, GetCapcha(sequence, next))
}

func GetCapcha(sequence string, next func(string, int) string) int {
	sum := 0
	for i := 0; i < len(sequence); i++ {
		digit := char(sequence, i)
		nextDigit := next(sequence, i)
		if digit == nextDigit {
			nextValue, _ := strconv.ParseInt(string(digit), 10, 64)
			sum += int(nextValue)
		}
	}
	return sum
}

func char(sequence string, index int) string {
	return string(sequence[index])
}

func nextChar(sequence string, index int) string {
	if index+1 == len(sequence) {
		return string(sequence[0])
	}
	return string(sequence[index+1])
}

func halfwayRound(sequence string, index int) string {
	halfOfLength := len(sequence) / 2
	next := (index + halfOfLength) % len(sequence)
	return string(sequence[next])
}
