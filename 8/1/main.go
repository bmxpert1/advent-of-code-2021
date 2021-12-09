package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	input  []string
	output []string
}

func (e *Entry) outputDigits() (digits []*Digit) {
	for _, seg := range e.output {
		digits = append(digits, &Digit{segments: seg})
	}
	return digits
}

type Digit struct {
	number   int
	segments string
}

func (d *Digit) segmentCount() int {
	return len(d.segments)
}

// this just tries to determine the number based on segment count
func (d *Digit) easyGuess() int {
	poss := segmentCountDigitMap()[d.segmentCount()]
	if len(poss) == 1 {
		return poss[0].number
	}
	return -1
}

func digits() map[int]*Digit {
	return map[int]*Digit{
		0: {
			number:   0,
			segments: "abcefg", // 6 (0,6,9)
			// contains: [1, 7]
		},
		1: {
			number:   1,
			segments: "cf", // 2 (1)*
			// contains: []
		},
		2: {
			number:   2,
			segments: "acdeg", // 5 (2,3,5)
			// contains: []
		},
		3: {
			number:   3,
			segments: "acdfg", // 5 (2,3,5)
			// contains: [1, 7]
		},
		4: {
			number:   4,
			segments: "bcdf", // 4 (4)*
			// contains: []
		},
		5: {
			number:   5,
			segments: "abdfg", // 5 (2,3,5)
			// contains: []
		},
		6: {
			number:   6,
			segments: "abdefg", // 6 (0,6,9)
			// contains: [5]
		},
		7: {
			number:   7,
			segments: "acf", // 3 (7)*
			// contains: [1]
		},
		8: {
			number:   8,
			segments: "abcdefg", // 7 (8)*
			// contains: [0,1,2,3,4,5,6,7,9]
		},
		9: {
			number:   9,
			segments: "abcdfg", // 6 (0,6,9)
			// contains: [1, 3, 4, 5, 7]
		},
	}
}

func segmentCountDigitMap() map[int][]*Digit {
	result := map[int][]*Digit{}
	for _, dig := range digits() {
		result[dig.segmentCount()] = append(result[dig.segmentCount()], dig)
	}
	return result
}

func main() {
	// read entries from txt
	entries := []*Entry{}
	file, _ := os.Open("../entries.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " | ")
		entry := &Entry{
			input:  strings.Fields(text[0]),
			output: strings.Fields(text[1]),
		}
		entries = append(entries, entry)
	}

	// count 1,4,7,8
	count := 0
	for _, entry := range entries {
		for _, digit := range entry.outputDigits() {
			switch digit.easyGuess() {
			case 1, 4, 7, 8:
				count++
			}
		}
	}

	fmt.Println(count)
}
