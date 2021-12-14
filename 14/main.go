package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

	for step := 0; step < 10; step++ {
		// generate pairs to lookup
		pairs := []string{}

		for i := 0; i < len(currentTemplate)-1; i++ {
			pairs = append(pairs, currentTemplate[i:i+2])
		}

		newTemplate := ""

		for i, pair := range pairs {
			if i == 0 {
				newTemplate += string(pair[0])
			}
			newTemplate += mods[pair] + string(pair[1])
		}

		currentTemplate = newTemplate
	}

	countMap := map[string]int{}
	countMin, countMax := math.MaxInt64, math.MinInt64

	for _, ltr := range strings.Split(letters, "") {
		countMap[ltr] = strings.Count(currentTemplate, ltr)
		if countMap[ltr] < countMin && countMap[ltr] != 0 {
			countMin = countMap[ltr]
		}
		if countMap[ltr] > countMax {
			countMax = countMap[ltr]
		}
	}

	fmt.Println(countMax - countMin)

	///////////////////////////////
	// challenge 2
	//
	// same thing but do it 40 times
	currentTemplate = template

	for step := 0; step < 40; step++ {
		fmt.Println(step)

		// generate pairs to lookup
		pairs := make([]string, len(currentTemplate)-1)

		for i := 0; i < len(currentTemplate)-1; i++ {
			pairs[i] = currentTemplate[i : i+2]
		}

		newTemplate := ""

		for i, pair := range pairs {
			if i == 0 {
				newTemplate += string(pair[0])
			}
			newTemplate += mods[pair] + string(pair[1])
		}

		currentTemplate = newTemplate
	}

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
