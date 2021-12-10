package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Table struct {
	data          string
	width         int
	length        int
	sprawlChecked [][2]int
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

func (t *Table) hasSprawlChecked(x int, y int) bool {
	for _, coords := range t.sprawlChecked {
		if coords[0] == x && coords[1] == y {
			return true
		}
	}

	return false
}

func (t *Table) basinSize(x int, y int) int {
	t.sprawlChecked = [][2]int{{x, y}}
	size := 1
	// sprawl up
	size += t.sprawl(x, y-1, "up")
	// sprawl right
	size += t.sprawl(x+1, y, "right")
	// sprawl down
	size += t.sprawl(x, y+1, "down")
	// sprawl left
	size += t.sprawl(x-1, y, "left")
	return size
}

func (t *Table) sprawl(x int, y int, dir string) int {
	if t.hasSprawlChecked(x, y) {
		return 0
	}
	t.sprawlChecked = append(t.sprawlChecked, [2]int{x, y})

	cell := t.cellAt(x, y)
	if cell >= 9 {
		return 0
	}
	size := 1
	if dir != "down" {
		// up
		size += t.sprawl(x, y-1, "up")
	}
	if dir != "right" {
		// left
		size += t.sprawl(x-1, y, "left")
	}
	if dir != "left" {
		// right
		size += t.sprawl(x+1, y, "right")
	}
	if dir != "up" {
		// down
		size += t.sprawl(x, y+1, "down")
	}
	return size
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

	lows := [][2]int{}

	///////////////////////////////////////////
	// challenge 1
	//
	riskSum := 0
	for x := 0; x < width; x++ {
		for y := 0; y < length; y++ {
			if table.isLow(x, y) {
				lows = append(lows, [2]int{x, y})
				riskSum += 1 + table.cellAt(x, y) // 1 plus its height
			}
		}
	}

	fmt.Println(riskSum)

	////////////////////////////////////////////
	// challenge 2
	//
	basinSizes := []int{}
	for _, low := range lows {
		basinSizes = append(basinSizes, table.basinSize(low[0], low[1]))
	}
	sort.Ints(basinSizes)

	// multiply top 3 biggest basins
	fmt.Println(basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3])
}
