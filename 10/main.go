package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	openingRunes = "([{<"
	closingRunes = ")]}>"
)

type Token struct {
	prev *Token
	next *Token
	char rune
	open bool
}

func NewToken(prev *Token, char rune) *Token {
	return &Token{
		prev: prev,
		char: char,
		open: true,
	}
}

func isOpeningToken(token rune) bool {
	return strings.ContainsRune(openingRunes, token)
}

func isClosingToken(token rune) bool {
	return strings.ContainsRune(closingRunes, token)
}

func oppositeToken(token rune) rune {
	if isOpeningToken(token) {
		return rune(closingRunes[strings.IndexRune(openingRunes, token)])
	} else {
		return rune(openingRunes[strings.IndexRune(closingRunes, token)])
	}
}

func parseLine(line string) (valid bool, invalidToken rune, tail *Token) {
	var curr *Token

	for i, t := range line {
		if i == 0 {
			if isOpeningToken(t) {
				curr = NewToken(nil, t)
				continue
			}

			return false, t, curr
		}

		if isOpeningToken(t) {
			if curr == nil {
				curr = NewToken(nil, t)
				continue
			}

			curr.next = NewToken(curr, t)
			curr = curr.next
		} else if isClosingToken(t) {
			opp := oppositeToken(t)
			if curr.char == opp {
				curr.open = false
				curr = curr.prev
			} else {
				return false, t, curr
			}
		} else {
			return false, t, curr
		}
	}

	return true, ' ', curr
}

func runesToComplete(tail *Token) []rune {
	runes := []rune{}

	for tail != nil {
		runes = append(runes, oppositeToken(tail.char))
		tail = tail.prev
	}

	return runes
}

func main() {
	// read input from txt
	lines := []string{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	///////////////////////
	// challenge 1
	//
	// invalids := []rune{}
	score := 0
	for _, line := range lines {
		if valid, invalidToken, _ := parseLine(line); !valid {
			// invalids = append(invalids, invalidToken)
			score += scores[invalidToken]
		}
	}

	fmt.Println(score)

	////////////////////////
	// challenge 2
	//

	scores = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	incScores := []int{}
	for _, line := range lines {
		if valid, _, tail := parseLine(line); valid {
			score := 0

			for _, t := range runesToComplete(tail) {
				score = (score * 5) + scores[t]
			}

			incScores = append(incScores, score)
		}
	}
	sort.Ints(incScores)
	fmt.Println(incScores[int(float64(len(incScores)/2))]) // median
}
