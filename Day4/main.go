package main

import (
	"fmt"
	"github.com/kenneth-fischer/Advent/files"
	"log"
	"strings"
)

func main() {
	checkIfValid("aa bb cc dd ee", noDupeWords)
	checkIfValid("aa bb cc dd ee aa", noDupeWords)
	checkIfValid("aa bb cc dd ee aaa", noDupeWords)
	err, total, valid := countValidPassPhrases("passPhrases.txt", noDupeWords)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d of %d pass phrases contain no duplicate words\n", valid, total)
	fmt.Println()

	checkIfValid("abcde fghij", noAnagrams)
	checkIfValid("abcde xyz ecdab", noAnagrams)
	checkIfValid("a ab abc abd abf abj", noAnagrams)
	checkIfValid("iiii oiii ooii oooi oooo", noAnagrams)
	checkIfValid("oiii ioii iioi iiio", noAnagrams)
	err, total, valid = countValidPassPhrases("passPhrases.txt", noAnagrams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d of %d pass phrases contain no anagrams\n", valid, total)
}

func checkIfValid(phrase string, validate func(string) error) {
	if err := validate(phrase); err != nil {
		fmt.Printf("Invalid: %q, %v\n", phrase, err)
	} else {
		fmt.Printf("Valid:   %q\n", phrase)
	}
}

func countValidPassPhrases(path string, validate func(string) error) (err error, total int, valid int) {
	lines, err := files.ReadLines("passPhrases.txt")
	if err != nil {
		return
	}
	total = len(lines)
	for _, line := range lines {
		if validateErr := validate(line); validateErr != nil {
			//fmt.Println(validateErr)
		} else {
			valid++
		}
	}
	return
}

func noDupeWords(text string) error {
	words := strings.Split(text, " ")
	occurs := map[string]int{}

	for _, word := range words {
		if _, exists := occurs[word]; exists {
			return fmt.Errorf("phrase contains more than one %q", word)
		}
		occurs[word] = 1
	}
	return nil
}

func noAnagrams(text string) error {
	words := strings.Split(text, " ")

	for i := 0; i < len(words)-1; i++ {
		firstWord := words[i]

		for j := i + 1; j < len(words); j++ {
			secondWord := words[j]
			if areAnagrams(firstWord, secondWord) {
				return fmt.Errorf("phrase contains anagrams %q and %q", firstWord, secondWord)
			}
		}
	}
	return nil
}

func areAnagrams(word1, word2 string) bool {
	lettersInWord1 := getLetterOccurences(word1)
	lettersInWord2 := getLetterOccurences(word2)

	if len(lettersInWord1) != len(lettersInWord2) {
		return false
	}

	for letter, word1Occurs := range lettersInWord1 {
		word2Occurs, found := lettersInWord2[letter]
		if !found {
			return false
		}
		if word1Occurs != word2Occurs {
			return false
		}
	}
	return true

}

func getLetterOccurences(word string) map[string]int {
	letters := map[string]int{}

	for _, ch := range word {
		letter := string(ch)
		if occurs, exists := letters[letter]; exists {
			letters[letter] = occurs + 1
		} else {
			letters[letter] = 1
		}
	}
	return letters
}
