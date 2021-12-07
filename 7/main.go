package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
)

func main() {
	var crabs []int
	content, _ := ioutil.ReadFile("input.json")
	json.Unmarshal([]byte(content), &crabs)

	// get min/max crab positions
	minCrabPos, maxCrabPos := math.MaxInt32, -math.MaxInt32
	distinctCrabPos := []int{}

	for _, pos := range crabs {
		if pos < minCrabPos {
			minCrabPos = pos
		}
		if pos > maxCrabPos {
			maxCrabPos = pos
		}
	}

	for p := minCrabPos; p <= maxCrabPos; p++ {
		distinctCrabPos = append(distinctCrabPos, p)
	}

	//////////////////////////////
	// challenge 1
	//
	// loop over all crabs to generate the total distance (fuel) to target distinct position
	// map distinct pos -> fuel usage
	fuelMap := map[int]int64{}
	for _, pos := range crabs {
		for _, tgt := range distinctCrabPos {
			fuelMap[tgt] += int64(math.Abs(float64(pos) - float64(tgt)))
		}
	}

	min := int64(math.MaxInt64)
	for _, cnt := range fuelMap {
		if cnt < min {
			min = cnt
		}
	}

	fmt.Println(min)

	/////////////////////////////////
	// challenge 2
	//
	fuelMap = map[int]int64{}
	for _, pos := range crabs {
		for _, tgt := range distinctCrabPos {
			dist := int64(math.Abs(float64(pos) - float64(tgt)))
			sum := int64(0)
			if dist > 0 {
				for i := dist; i >= 1; i-- {
					sum += i
				}
			}
			// fmt.Printf("from %v to %v: %v fuel\n", pos, tgt, sum)
			fuelMap[tgt] += sum
		}
	}

	min = int64(math.MaxInt64)
	for _, cnt := range fuelMap {
		if cnt < min {
			min = cnt
		}
	}

	fmt.Println(min)
}
