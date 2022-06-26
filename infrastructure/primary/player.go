package main

import (
	"bufio"
	"fmt"
	"os"
	"scrabble/usecases"
)

func main() {
	words, lettersAvailable := parseInputs()

	bestWord := usecases.FindBestWordUsecase(words, lettersAvailable)

	fmt.Println(bestWord)
}

func parseInputs() ([]string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	words := parseWords(scanner)
	lettersAvailable := parseLettersAvailable(scanner)

	return words, lettersAvailable
}

func parseWords(scanner *bufio.Scanner) []string {
	var wordCount int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &wordCount)

	words := make([]string, wordCount)
	for i := 0; i < wordCount; i++ {
		scanner.Scan()
		words[i] = scanner.Text()
	}
	scanner.Scan()

	return words
}

func parseLettersAvailable(scanner *bufio.Scanner) string {
	lettersAvailable := scanner.Text()
	return lettersAvailable
}
