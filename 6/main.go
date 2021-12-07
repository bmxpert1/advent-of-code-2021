package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// try as an array of uint8
func sim(days int) int64 {
	// read initial ages from json
	var ages []uint8
	content, _ := ioutil.ReadFile("example_fish.json")
	json.Unmarshal([]byte(content), &ages)

	var fishCount int64

	for i := 0; i < days; i++ {
		fishCount = 0
		newAges := []uint8{}
		for _, age := range ages {
			if age == 0 {
				// reset, spawn
				age = 6
				newAges = append(newAges, 8) // the new fish
				fishCount++
			} else {
				age--
			}
			newAges = append(newAges, age)
			fishCount++
		}
		ages = newAges
		fmt.Printf("%v:%v\n", i, fishCount)
	}

	return fishCount
}

// try it as an array of pointers
func sim2(days int) int64 {
	// load the possible values into a pointer map to save on memory
	numPtrs := map[uint8]*uint8{}
	for i := 0; i < 9; i++ {
		ptr := new(uint8)
		*ptr = uint8(i)
		numPtrs[uint8(i)] = ptr
	}

	// read initial ages from json
	var startingAges []uint8
	content, _ := ioutil.ReadFile("example_fish.json")
	json.Unmarshal([]byte(content), &startingAges)

	var fishCount int64

	ages := []*uint8{}
	for _, age := range startingAges {
		ages = append(ages, numPtrs[age])
	}

	for i := 0; i < days; i++ {
		fishCount = 0
		newAges := []*uint8{}
		for _, age := range ages {
			if *age == 0 {
				// reset, spawn
				age = numPtrs[6]
				newAges = append(newAges, numPtrs[8]) // the new fish
				fishCount++
			} else {
				age = numPtrs[*age-1]
			}
			newAges = append(newAges, age)
			fishCount++
		}
		ages = nil
		ages = newAges
		fmt.Printf("%v:%v\n", i, fishCount)
	}

	return int64(fishCount)
}

// try it as a string
func sim3(days int) int64 {
	// read initial ages from txt
	agesInitial, _ := ioutil.ReadFile("example_fish.txt")
	ages := string(agesInitial)

	var fishCount int64

	for i := 0; i < days; i++ {
		fishCount = 0
		newAges := ""
		for _, age := range strings.Split(ages, ",") {
			agei, _ := strconv.Atoi(age)
			if agei == 0 {
				// reset, spawn
				agei = 6
				newAges += "8,"
				fishCount++
			} else {
				agei--
			}
			newAges += strconv.Itoa(agei) + ","
			fishCount++
		}
		ages = strings.TrimSuffix(newAges, ",")
		fmt.Printf("%v:%v\n", i, fishCount)
	}

	return fishCount
}

func main() {
	/////////////////////////////
	// challenge 1
	//
	// run sim for 80 days
	fmt.Println(sim(80))

	/////////////////////////////
	// challenge 2
	//
	// run sim for 256 days
	fmt.Println(sim(256))
}
