package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid struct {
	cells  [][]bool
	width  int
	height int
}

func NewGrid(coords [][2]int, width int, height int) *Grid {
	rows := make([][]bool, height)
	for i := 0; i < height; i++ {
		rows[i] = make([]bool, width)
	}
	for _, coord := range coords {
		rows[coord[1]][coord[0]] = true
	}

	return &Grid{rows, width, height}
}

func (g *Grid) setCell(x int, y int, val bool) {
	g.cells[y][x] = val
}

func (g *Grid) cellAt(x int, y int) bool {
	return g.cells[y][x]
}

func (g *Grid) activeCellCount() int {
	count := 0
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			if g.cellAt(x, y) {
				count++
			}
		}
	}
	return count
}

func (g *Grid) printCells() {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			char := "."
			if cell := g.cellAt(x, y); cell {
				char = "#"
			}
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func (g *Grid) fold(plane string, which int) {
	start := which + 1
	switch plane {
	case "x":
		// vertical fold
		for x := start; x < g.width; x++ {
			for y := 0; y < g.height; y++ {
				// only really care to do anything if value is true
				if g.cellAt(x, y) {
					g.setCell(which-(x-which), y, true)
					g.setCell(x, y, false)
				}
			}
		}

		g.width = which
	case "y":
		// horizontal fold
		for y := start; y < g.height; y++ {
			for x := 0; x < g.width; x++ {
				// only really care to do anything if value is true
				if g.cellAt(x, y) {
					g.setCell(x, which-(y-which), true)
					g.setCell(x, y, false)
				}
			}
		}

		g.height = which
	}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	coords := [][2]int{}
	folds := [][2]interface{}{}
	xMax, yMax := 0, 0

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), ",") {
			pair := strings.Split(scanner.Text(), ",")
			x, _ := strconv.Atoi(pair[0])
			y, _ := strconv.Atoi(pair[1])
			coords = append(coords, [2]int{x, y})
			if x > xMax {
				xMax = x
			}
			if y > yMax {
				yMax = y
			}
		} else if strings.Contains(scanner.Text(), "=") {
			fold := strings.Split(scanner.Text(), "=")
			line, _ := strconv.Atoi(fold[1])
			folds = append(folds, [2]interface{}{string(fold[0][len(fold[0])-1]), line})
		}
	}

	grid := NewGrid(coords, xMax+1, yMax+1)

	////////////////////////////////////
	// challenge 1
	//
	// just do the first fold
	grid.fold(folds[0][0].(string), folds[0][1].(int))
	fmt.Println(grid.activeCellCount())

	////////////////////////////////////
	// challenge 2
	//
	grid = NewGrid(coords, xMax+1, yMax+1)
	for _, fold := range folds {
		grid.fold(fold[0].(string), fold[1].(int))
	}
	grid.printCells()
}
