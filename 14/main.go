package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Pair struct {
	value        string
	yields       [2]string
	adds         string
	nextCount    int
	currentCount int
}

func main() {
	pairMap := map[string]*Pair{}

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var template string

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), " -> ") {
			pair := strings.Split(scanner.Text(), " -> ")
			couple, middle := pair[0], pair[1]
			pairMap[couple] = &Pair{
				value:  couple,
				yields: [2]string{string(couple[0]) + middle, middle + string(couple[1])},
				adds:   middle,
			}
		} else if len(scanner.Text()) > 0 {
			// starting template
			template = scanner.Text()
		}
	}

	var totals map[string]int

	runSteps := func(steps int) {
		totals = map[string]int{}

		// load initial character totals
		for _, c := range template {
			totals[string(c)]++
		}

		// reset pairMap counts
		for _, pair := range pairMap {
			pair.nextCount = 0
			pair.currentCount = 0
		}

		// register initial next loop counts
		for i := 0; i < len(template)-1; i++ {
			pairMap[template[i:i+2]].nextCount++
		}

		for i := 0; i < steps; i++ {
			for _, pair := range pairMap {
				pair.currentCount = pair.nextCount
				pair.nextCount = 0
			}

			for _, pair := range pairMap {
				if pair.currentCount > 0 {
					totals[pair.adds] += pair.currentCount

					for _, y := range pair.yields {
						pairMap[y].nextCount += pair.currentCount
					}
				}
			}
		}
	}

	///////////////////////
	// challenge 1
	//
	// 10 steps
	runSteps(10)

	countMin, countMax := math.MaxInt64, math.MinInt64
	for _, v := range totals {
		if v == 0 {
			continue
		}

		if v < countMin {
			countMin = v
		}
		if v > countMax {
			countMax = v
		}
	}

	fmt.Println(countMax - countMin)

	///////////////////////
	// challenge 1
	//
	// 40 steps
	runSteps(40)

	countMin, countMax = math.MaxInt64, math.MinInt64
	for _, v := range totals {
		if v == 0 {
			continue
		}

		if v < countMin {
			countMin = v
		}
		if v > countMax {
			countMax = v
		}
	}

	fmt.Println(countMax - countMin)
}
