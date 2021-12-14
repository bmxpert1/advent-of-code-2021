package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Pair struct {
	orig  string
	value string
}

func run(template string, steps int, mods *map[string]string, countMap *map[string]int) {
	pairs := []*Pair{}

	// register initial template characters into countMap
	for _, c := range template {
		(*countMap)[string(c)]++
	}

	for i := 0; i < len(template)-1; i++ {
		pairs = append(pairs, &Pair{template[i : i+2], template[i : i+2]})
	}

	for _, pair := range pairs {
		pair.runPair(mods, countMap, steps, 0)
	}
}

func (p *Pair) runPair(mods *map[string]string, countMap *map[string]int, targetSteps int, currentStep int) {
	if currentStep >= targetSteps {
		return
	}

	p.inject(mods, countMap)

	for _, pair := range p.split() {
		pair.runPair(mods, countMap, targetSteps, currentStep+1)
	}
}

func (p *Pair) inject(mods *map[string]string, countMap *map[string]int) {
	(*countMap)[(*mods)[p.value]]++
	p.value = string(p.value[0]) + (*mods)[p.value] + string(p.value[1])
}

func (p *Pair) split() (pairs [2]*Pair) {
	for i := 0; i < len(p.value)-1; i++ {
		pairs[i] = &Pair{p.value[i : i+2], p.value[i : i+2]}
	}

	return pairs
}

func main() {
	file, _ := os.Open("example_input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var template string
	mods := map[string]string{}

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), " -> ") {
			pair := strings.Split(scanner.Text(), " -> ")
			couple, middle := pair[0], pair[1]
			mods[couple] = middle
		} else if len(scanner.Text()) > 0 {
			// starting template
			template = scanner.Text()
		}
	}

	////////////////////////////
	// challenge 1
	//
	currentTemplate := template
	countMap := map[string]int{}

	run(currentTemplate, 10, &mods, &countMap)

	countMin, countMax := math.MaxInt64, math.MinInt64

	for _, v := range countMap {
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

	///////////////////////////////
	// challenge 2
	//
	// same thing but do it 40 times
	currentTemplate = template
	countMap = map[string]int{}

	run(currentTemplate, 40, &mods, &countMap)

	countMin, countMax = math.MaxInt64, math.MinInt64

	for _, v := range countMap {
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
