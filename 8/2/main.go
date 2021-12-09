package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	inputDigits  []*Digit
	outputDigits []*Digit
}

func NewEntry(input []string, output []string) *Entry {
	inputDigits, outputDigits := []*Digit{}, []*Digit{}

	for _, segments := range input {
		digit := &Digit{
			number:   -1,
			segments: segments,
		}
		digit.number = digit.easyGuess()
		inputDigits = append(inputDigits, digit)
	}

	for _, segments := range output {
		digit := &Digit{
			number:   -1,
			segments: segments,
		}
		digit.number = digit.easyGuess()
		outputDigits = append(outputDigits, digit)
	}

	entry := &Entry{
		inputDigits:  inputDigits,
		outputDigits: outputDigits,
	}

	return entry
}

func unsolvedFromDigits(digits []*Digit) []*Digit {
	unsolved := []*Digit{}
	for _, dig := range digits {
		if dig.number == -1 {
			unsolved = append(unsolved, dig)
		}
	}
	return unsolved
}

func (e *Entry) inputDigitsWithSegmentCount(count int) []*Digit {
	withSegs := []*Digit{}
	for _, dig := range e.inputDigits {
		if dig.segmentCount() == count {
			withSegs = append(withSegs, dig)
		}
	}
	return withSegs
}

func (e *Entry) knownDigit(num int) *Digit {
	for _, dig := range e.inputDigits {
		if dig.number == num {
			return dig
		}
	}
	return nil
}

func (e *Entry) solution() int {
	digits := ""
	for _, dig := range e.outputDigits {
		for _, xDig := range e.inputDigits {
			if xDig.containsSegments(dig.segments, true) {
				digits += strconv.Itoa(xDig.number)
				dig.number = xDig.number
				break
			}
		}
	}
	final, _ := strconv.Atoi(digits)
	return final
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

func (d *Digit) containsSegments(segs string, exact bool) bool {
	if len(segs) > len(d.segments) {
		return false
	}

	if exact && len(d.segments) != len(segs) {
		return false
	}

	for _, seg := range segs {
		contains := false
		for _, xSeg := range d.segments {
			if xSeg == seg {
				contains = true
			}
		}
		if !contains {
			return false
		}
	}

	return true
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
		entries = append(entries, NewEntry(strings.Fields(text[0]), strings.Fields(text[1])))
	}

	solutionSum := 0

	for _, entry := range entries {
		//
		// we already know 1,4,7,8

		// we should be able to deduce the 9 because it will have 6 segments
		// and contain all the segments from 1,4,7 combined
		for _, dig := range unsolvedFromDigits(entry.inputDigitsWithSegmentCount(6)) {
			if dig.containsSegments(entry.knownDigit(1).segments, false) &&
				dig.containsSegments(entry.knownDigit(4).segments, false) &&
				dig.containsSegments(entry.knownDigit(7).segments, false) {
				dig.number = 9
			}
		}

		// the two remaining digits with 6 segments are 0,6
		// 0 contains 1,7 so we should be able to deduce that, and solve the other
		for _, dig := range unsolvedFromDigits(entry.inputDigitsWithSegmentCount(6)) {
			if dig.containsSegments(entry.knownDigit(1).segments, false) &&
				dig.containsSegments(entry.knownDigit(7).segments, false) {
				dig.number = 0
			}
		}
		// the remaining 6 segment digit is 6
		unsolvedFromDigits(entry.inputDigitsWithSegmentCount(6))[0].number = 6

		// remaining digits to solve are 2,3,5, all 5 segments digits
		// 3 also contains 1,7 so we can figure that one out
		for _, dig := range unsolvedFromDigits(entry.inputDigitsWithSegmentCount(5)) {
			if dig.containsSegments(entry.knownDigit(1).segments, false) &&
				dig.containsSegments(entry.knownDigit(7).segments, false) {
				dig.number = 3
			}
		}
		// 5 is contained by 6 and 2 is not, so we can figure that out
		for _, dig := range unsolvedFromDigits(entry.inputDigitsWithSegmentCount(5)) {
			if entry.knownDigit(6).containsSegments(dig.segments, false) {
				dig.number = 5
			}
		}
		// the remaining is 2
		unsolvedFromDigits(entry.inputDigitsWithSegmentCount(5))[0].number = 2
		// done solving

		solutionSum += entry.solution()
	}

	fmt.Println(solutionSum)
}
