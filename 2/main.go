package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	// read in json data
	var data [][]interface{}
	content, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		fmt.Println(err)
	}

	// challenge 1
	pos := []int{0, 0}

	for _, adj := range data {
		amt := int(adj[1].(float64))
		switch adj[0].(string) {
		case "forward":
			pos[0] += amt
		case "down":
			pos[1] += amt
		case "up":
			pos[1] -= amt
		}
	}

	fmt.Println(pos[0] * pos[1])

	// challenge 2
	pos[0], pos[1] = 0, 0
	aim := 0

	for _, adj := range data {
		amt := int(adj[1].(float64))
		switch adj[0].(string) {
		case "forward":
			pos[0] += amt
			pos[1] += aim * amt
		case "down":
			aim += amt
		case "up":
			aim -= amt
		}
	}

	fmt.Println(pos[0] * pos[1])
}
