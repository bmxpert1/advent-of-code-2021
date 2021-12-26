package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
)

var (
	xmin, xmax, ymin, ymax int
)

func testVelocity(xVel int, yVel int) (success bool, pos [][2]int) {
	pos = append(pos, [2]int{0, 0})
	currx, curry := pos[0][0], pos[0][1]

	for curry >= ymin && !insideBounds(currx, curry) {
		currx += xVel
		curry += yVel
		pos = append(pos, [2]int{currx, curry})

		// adjust velocities
		if xVel > 0 {
			xVel--
		} else if xVel < 0 {
			xVel++
		}
		yVel--
	}

	return insideBounds(currx, curry), pos
}

func insideBounds(x int, y int) bool {
	return x >= xmin && x <= xmax && y >= ymin && y <= ymax
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	r, _ := regexp.Compile(`x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)`)
	bounds := r.FindSubmatch(input)
	xmins, xmaxs, ymins, ymaxs := string(bounds[1]), string(bounds[2]), string(bounds[3]), string(bounds[4])
	xmin, _ = strconv.Atoi(xmins)
	xmax, _ = strconv.Atoi(xmaxs)
	ymin, _ = strconv.Atoi(ymins)
	ymax, _ = strconv.Atoi(ymaxs)

	maxy := math.MinInt64
	successMap := map[[2]int]bool{}

	for x := -500; x < 500; x++ {
		for y := -500; y < 500; y++ {
			if success, pos := testVelocity(x, y); success {
				successMap[[2]int{x, y}] = true

				for _, p := range pos {
					if p[1] > maxy {
						maxy = p[1]
					}
				}
			}
		}
	}

	// challenge 1
	fmt.Println(maxy)

	// challenge 2
	fmt.Println(len(successMap))
}
