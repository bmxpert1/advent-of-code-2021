package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func sum(nums []int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func main() {
	var data []int
	content, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		fmt.Println(err)
	}

	// challenge 1
	count := 0

	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			count++
		}
	}

	fmt.Println(count)

	// challenge 2
	count = 0
	var cons = [][]int{}

	for i := 0; i < len(data)-2; i++ {
		cons = append(cons, data[i:i+3])
	}

	for i := 1; i < len(cons); i++ {
		if sum(cons[i]) > sum(cons[i-1]) {
			count++
		}
	}

	fmt.Println(count)
}
