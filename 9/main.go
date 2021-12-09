package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Table struct {
	data   string
	width  int
	length int
}

func (t *Table) cellAt(x int, y int) int {
	// handle going outside of the square (table)
	if x < 0 || x > t.width-1 || y < 0 || y > t.length-1 {
		return math.MaxInt64
	}

	num, _ := strconv.Atoi(string(t.data[y*t.width+x]))
	return num
}

func (t *Table) isLow(x int, y int) bool {
	cell := t.cellAt(x, y)

	// check above
	return cell < t.cellAt(x, y-1) &&
		// check right
		cell < t.cellAt(x+1, y) &&
		// check below
		cell < t.cellAt(x, y+1) &&
		// check left
		cell < t.cellAt(x-1, y)
}

func main() {
	var table *Table
	// read input from txt
	content, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(content), "\n")
	width := len(lines[0])
	length := len(lines)

	table = &Table{
		data:   strings.Replace(string(content), "\n", "", -1),
		width:  width,
		length: length,
	}

	///////////////////////////////////////////
	// challenge 1
	//
	riskSum := 0
	for x := 0; x < width; x++ {
		for y := 0; y < length; y++ {
			if table.isLow(x, y) {
				riskSum += 1 + table.cellAt(x, y) // 1 plus its height
			}
		}
	}

	fmt.Println(riskSum)
}
