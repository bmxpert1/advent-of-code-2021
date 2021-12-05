package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

type Vent struct {
	start Coords
	end   Coords
}

func (vent *Vent) isEasy() bool {
	return vent.start.x == vent.end.x || vent.start.y == vent.end.y
}

func (vent *Vent) getThroughCoords() (coords []Coords) {
	// add the start
	coords = append(coords, vent.start)

	if vent.isEasy() {
		// vertical,horizontal vents
		//
		if vent.start.x == vent.end.x {
			// vertical, y's are different
			var startY int
			var endY int
			if vent.start.y < vent.end.y {
				startY = vent.start.y
				endY = vent.end.y
			} else {
				startY = vent.end.y
				endY = vent.start.y
			}
			for y := startY + 1; y < endY; y++ {
				coords = append(coords, Coords{vent.start.x, y})
			}
		} else if vent.start.y == vent.end.y {
			// horizontal, x's are different
			var startX int
			var endX int
			if vent.start.x < vent.end.x {
				startX = vent.start.x
				endX = vent.end.x
			} else {
				startX = vent.end.x
				endX = vent.start.x
			}
			for x := startX + 1; x < endX; x++ {
				coords = append(coords, Coords{x, vent.start.y})
			}
		}
	} else {
		// diagonal vents
		//
		xFactor, yFactor := 1, 1

		if vent.start.x > vent.end.x {
			xFactor = -1
		}

		if vent.start.y > vent.end.y {
			yFactor = -1
		}

		diff := int(math.Abs(float64(vent.end.x) - float64(vent.start.x)))

		for i := 1; i < diff; i++ {
			coords = append(coords, Coords{vent.start.x + (i * xFactor), vent.start.y + (i * yFactor)})
		}
	}

	// add the end
	coords = append(coords, vent.end)
	return coords
}

func main() {
	// read vents from txt
	vents := []*Vent{}
	file, _ := os.Open("vents.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		vent := new(Vent)
		bothCoords := strings.Split(scanner.Text(), "->")
		for i, coords := range bothCoords {
			xy := strings.Split(coords, ",")
			var xyi [2]int
			for i, s := range xy {
				xyi[i], _ = strconv.Atoi(strings.Trim(s, " "))
			}
			if i == 0 {
				vent.start.x = xyi[0]
				vent.start.y = xyi[1]
			} else if i == 1 {
				vent.end.x = xyi[0]
				vent.end.y = xyi[1]
			}
		}
		vents = append(vents, vent)
	}
	// vents loaded

	////////////////////////////
	// challenge 1
	//
	// with this challenge we only consider vents where x1 == x2 or y1 == y2
	easyVents := []*Vent{}
	cells := map[[2]int]int{}

	// filters vents for easy vents
	for _, vent := range vents {
		if vent.isEasy() {
			easyVents = append(easyVents, vent)
		}
	}

	for _, vent := range easyVents {
		for _, coords := range vent.getThroughCoords() {
			cells[[2]int{coords.x, coords.y}]++
		}
	}

	// find cells with value > 1
	count := 0
	for _, cell := range cells {
		if cell > 1 {
			count++
		}
	}

	fmt.Println(count)

	////////////////////////////
	// challenge 2
	//
	cells = map[[2]int]int{}

	for _, vent := range vents {
		for _, coords := range vent.getThroughCoords() {
			cells[[2]int{coords.x, coords.y}]++
		}
	}

	// find cells with value > 1
	count = 0
	for _, cell := range cells {
		if cell > 1 {
			count++
		}
	}

	fmt.Println(count)
}
