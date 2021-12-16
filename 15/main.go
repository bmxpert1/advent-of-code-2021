package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/fatih/color"
)

var distAdj = map[int][][][2]int{
	2: {
		{{0, 1}},
		{{1, 0}},
	},
	3: {
		{{0, 1}, {0, 2}},
		{{0, 1}, {1, 1}},
		{{1, 0}, {1, 1}},
		{{1, 0}, {2, 0}},
	},
	4: {
		{{0, 1}, {0, 2}, {0, 3}},
		{{0, 1}, {0, 2}, {1, 2}},
		{{0, 1}, {1, 1}, {1, 2}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{1, 0}, {1, 1}, {2, 1}},
		{{1, 0}, {2, 0}, {2, 1}},
		{{1, 0}, {2, 0}, {3, 0}},
	},
}

func sum(nums []int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

type Grid struct {
	cells    [][]int
	width    int
	height   int
	entered  map[[2]int]bool
	position [2]int
}

func NewGrid(cells [][]int) *Grid {
	width := len(cells[0])
	height := len(cells)

	return &Grid{cells, width, height, map[[2]int]bool{{0, 0}: true}, [2]int{0, 0}}
}

func (g *Grid) cellAt(x int, y int) int {
	// handle going outside of the square (table)
	if x < 0 || x > g.width-1 || y < 0 || y > g.height-1 {
		return -1
	}

	return g.cells[y][x]
}

func (g *Grid) print() {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			c := color.New(color.FgWhite)
			if g.entered[[2]int{x, y}] {
				c = c.Add(color.FgRed)
			}
			c.Printf("%d", g.cellAt(x, y))
		}
		fmt.Println()
	}
	fmt.Println()
}

// decide the next cell to move to and move
func (g *Grid) move(target [2]int) {
	nextPos := [][2]int{}
	x, y := g.position[0], g.position[1]

	// check right
	if cv := g.cellAt(x+1, y); cv != -1 {
		nextPos = append(nextPos, [2]int{x + 1, y})
	}

	// check down
	if cv := g.cellAt(x, y+1); cv != -1 {
		nextPos = append(nextPos, [2]int{x, y + 1})
	}

	var bestPos [2]int
	bestPosSum := math.MaxInt64
	allOpts := map[[2]int][][]int{}

	for _, pos := range nextPos {
		allOpts[pos] = g.genOptions(pos, 3)
	}

	for pos, opts := range allOpts {
		for _, opt := range opts {
			sum := sum(opt)
			if sum < bestPosSum {
				bestPosSum = sum
				bestPos = pos
			}
		}
	}

	g.position = bestPos
	g.entered[g.position] = true

	if g.position != target {
		g.move(target)
	}
}

func (g *Grid) genOptions(from [2]int, maxDistance int) (opts [][]int) {
	if maxDistance == 1 {
		return [][]int{{g.cellAt(from[0], from[1])}}
	}

	for _, adjs := range distAdj[maxDistance] {
		opt := []int{}
		// add ourself
		opt = append(opt, g.cellAt(from[0], from[1]))

		for _, adj := range adjs {
			if c := g.cellAt(from[0]+adj[0], from[1]+adj[1]); c != -1 {
				opt = append(opt, c)
			}
		}

		// make sure we [maxDistance] count
		if len(opt) == maxDistance {
			opts = append(opts, opt)
		}
	}

	if len(opts) == 0 {
		return g.genOptions(from, maxDistance-1)
	} else {
		return opts
	}
}

func (g *Grid) sumEntered() int {
	sum := 0
	for pos, _ := range g.entered {
		if pos != [2]int{0, 0} {
			sum += g.cellAt(pos[0], pos[1])
		}
	}
	return sum
}

func main() {
	cells := [][]int{}

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		newRow := []int{}

		for _, cell := range scanner.Text() {
			celli, _ := strconv.Atoi(string(cell))
			newRow = append(newRow, celli)
		}

		cells = append(cells, newRow)
	}

	grid := NewGrid(cells)

	grid.move([2]int{grid.width - 1, grid.height - 1})

	grid.print()

	fmt.Println(grid.sumEntered())
}
