package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Pair struct {
	orig  string
	value string
}

func run(template string, steps int, mods *map[string]string) string {
	pairs := []*Pair{}

	for i := 0; i < len(template)-1; i++ {
		pairs = append(pairs, &Pair{template[i : i+2], template[i : i+2]})
	}

	// targetCount := len(pairs)
	// for i := 0; i < steps-1; i++ {
	// 	targetCount += targetCount
	// }

	// fmt.Printf("target count: %v\n", targetCount)

	final := ""
	// currentCount := 0

	for i, pair := range pairs {
		fmt.Println(i)
		// final += pair.runPair(mods, targetCount, &currentCount, i == 0)
		fmt.Printf("running pair: %v\n", pair.orig)
		supplement := pair.runPair(mods, steps, 0, i == 0)
		fmt.Printf("finished pair: %v, %v\n", pair.orig, supplement)
		final += supplement
	}

	return final
}

// func (p *Pair) runPair(mods *map[string]string, targetCount int, currentCount *int, first bool) string {
func (p *Pair) runPair(mods *map[string]string, targetSteps int, currentStep int, first bool, aggString string) string {
	if currentStep >= targetSteps {
		return ""
	}

	fmt.Printf("running inner pair: %v, %v, step: %v\n", p.orig, first, currentStep)
	final := ""

	p.inject(mods)
	final = p.value

	for i, pair := range p.split() {
		supplement := pair.runPair(mods, targetSteps, currentStep+1, first && i == 0)
		final += supplement
		if p.orig == "NN" {
			fmt.Printf("*** %v [%v]\n", supplement, final)
		}
	}

	if !first && len(final) > 0 {
		final = final[1:]
	}

	fmt.Printf("finished inner pair: %v, %v\n", p.orig, final)

	return final
}

func (p *Pair) inject(mods *map[string]string) {
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

	fmt.Println(run(currentTemplate, 2, &mods))

	// for step := 0; step < 10; step++ {
	// 	// generate pairs to lookup
	// 	pairs := []string{}

	// 	for i := 0; i < len(currentTemplate)-1; i++ {
	// 		pairs = append(pairs, currentTemplate[i:i+2])
	// 	}

	// 	newTemplate := ""

	// 	for i, pair := range pairs {
	// 		if i == 0 {
	// 			newTemplate += string(pair[0])
	// 		}
	// 		newTemplate += mods[pair] + string(pair[1])
	// 	}

	// 	currentTemplate = newTemplate
	// }

	// countMap := map[string]int{}
	// countMin, countMax := math.MaxInt64, math.MinInt64

	// for _, ltr := range strings.Split(letters, "") {
	// 	countMap[ltr] = strings.Count(currentTemplate, ltr)
	// 	if countMap[ltr] < countMin && countMap[ltr] != 0 {
	// 		countMin = countMap[ltr]
	// 	}
	// 	if countMap[ltr] > countMax {
	// 		countMax = countMap[ltr]
	// 	}
	// }

	// fmt.Println(countMax - countMin)

	///////////////////////////////
	// challenge 2
	//
	// same thing but do it 40 times
	// currentTemplate = template

	// for step := 0; step < 40; step++ {
	// 	fmt.Printf("%v A\n", step)

	// 	// generate pairs to lookup
	// 	pairs := make([]string, len(currentTemplate)-1)

	// 	for i := 0; i < len(currentTemplate)-1; i++ {
	// 		pairs[i] = currentTemplate[i : i+2]
	// 	}

	// 	fmt.Printf("%v B, %v\n", step, len(pairs))

	// 	parts := make([]string, len(pairs))

	// 	for i, pair := range pairs {
	// 		supplement := ""
	// 		if i == 0 {
	// 			supplement += string(pair[0])
	// 		}
	// 		supplement += mods[pair] + string(pair[1])

	// 		parts[i] = supplement
	// 	}

	// 	fmt.Printf("%v C, %v\n", step, uintptr(len(pairs))*reflect.TypeOf(pairs).Elem().Size())

	// 	currentTemplate = strings.Join(parts, "")
	// }

	// countMap = map[string]int{}
	// countMin, countMax = math.MaxInt64, math.MinInt64

	// for _, ltr := range strings.Split(letters, "") {
	// 	countMap[ltr] = strings.Count(currentTemplate, ltr)
	// 	if countMap[ltr] < countMin && countMap[ltr] != 0 {
	// 		countMin = countMap[ltr]
	// 	}
	// 	if countMap[ltr] > countMax {
	// 		countMax = countMap[ltr]
	// 	}
	// }

	// fmt.Println(countMax - countMin)
}
