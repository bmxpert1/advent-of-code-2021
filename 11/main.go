package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type OctoTable struct {
	octopi [][]*Octo
	// originalOctopi [][]*Octo
	width   int
	length  int
	flashes int
}

func NewOctoTable(octopi [][]*Octo) *OctoTable {
	// octoCopy := make([][]*Octo, len(octopi))
	// for i := range octopi {
	// 	octoCopy[i] = make([]*Octo, len(octopi[i]))
	// 	copy(octoCopy[i], octopi[i])
	// }
	tbl := &OctoTable{octopi, len(octopi[0]), len(octopi), 0}
	for y, octoRow := range tbl.octopi {
		for x, octo := range octoRow {
			octo.table = tbl
			octo.x = x
			octo.y = y
		}
	}
	return tbl
}

// func (t *OctoTable) reset() {
// 	octoCopy := make([][]*Octo, len(t.originalOctopi))
// 	for i := range t.originalOctopi {
// 		octoCopy[i] = make([]*Octo, len(t.originalOctopi[i]))
// 		copy(octoCopy[i], t.originalOctopi[i])
// 	}

// 	t.octopi = octoCopy
// 	t.flashes = 0

// 	// spew.Dump(t)
// }

func (t *OctoTable) octoAt(x int, y int) *Octo {
	if y < 0 || x < 0 || y > t.length-1 || x > t.width-1 {
		return nil
	}

	return t.octopi[y][x]
}

func (t *OctoTable) step() (flashes int) {
	for y := 0; y < t.length; y++ {
		for x := 0; x < t.width; x++ {
			octo := t.octoAt(x, y)
			octo.flashed = false
			octo.energize(false)
		}
	}

	for y := 0; y < t.length; y++ {
		for x := 0; x < t.width; x++ {
			if octo := t.octoAt(x, y); octo.energy > 9 {
				octo.flash()
			}
		}
	}

	flashes = 0
	for y := 0; y < t.length; y++ {
		for x := 0; x < t.width; x++ {
			if octo := t.octoAt(x, y); octo.flashed {
				flashes++
			}
		}
	}

	return flashes
}

type Octo struct {
	table   *OctoTable
	energy  int
	flashed bool
	x       int
	y       int
}

func (o *Octo) energize(allowFlash bool) {
	if !o.flashed {
		o.energy += 1

		if o.energy > 9 && allowFlash {
			o.flash()
		}
	}
}

func (o *Octo) flash() {
	o.flashed = true
	o.energy = 0

	// energize adjacent including diagonal
	//
	// TL
	if octo := o.table.octoAt(o.x-1, o.y-1); octo != nil {
		octo.energize(true)
	}
	// T
	if octo := o.table.octoAt(o.x, o.y-1); octo != nil {
		octo.energize(true)
	}
	// TR
	if octo := o.table.octoAt(o.x+1, o.y-1); octo != nil {
		octo.energize(true)
	}
	// R
	if octo := o.table.octoAt(o.x+1, o.y); octo != nil {
		octo.energize(true)
	}
	// BR
	if octo := o.table.octoAt(o.x+1, o.y+1); octo != nil {
		octo.energize(true)
	}
	// B
	if octo := o.table.octoAt(o.x, o.y+1); octo != nil {
		octo.energize(true)
	}
	// BL
	if octo := o.table.octoAt(o.x-1, o.y+1); octo != nil {
		octo.energize(true)
	}
	// L
	if octo := o.table.octoAt(o.x-1, o.y); octo != nil {
		octo.energize(true)
	}

	o.table.flashes++
}

func main() {
	// read input from txt
	octos := [][]*Octo{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		octoLine := []*Octo{}

		for _, nrg := range scanner.Text() {
			energy, _ := strconv.Atoi(string(nrg))
			octoLine = append(octoLine, &Octo{
				energy:  energy,
				flashed: false,
			})
		}

		octos = append(octos, octoLine)
	}

	octoTable := NewOctoTable(octos)

	//////////////////////////
	// challenge 1
	//
	steps := 100

	for i := 0; i < steps; i++ {
		octoTable.step()
	}

	fmt.Println(octoTable.flashes)

	//////////////////////////
	// challenge 2
	//
	octos = [][]*Octo{}
	file, _ = os.Open("input.txt")
	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		octoLine := []*Octo{}

		for _, nrg := range scanner.Text() {
			energy, _ := strconv.Atoi(string(nrg))
			octoLine = append(octoLine, &Octo{
				energy:  energy,
				flashed: false,
			})
		}

		octos = append(octos, octoLine)
	}

	octoTable = NewOctoTable(octos)
	step := 1
	for flashes := octoTable.step(); flashes != octoTable.width*octoTable.length; flashes = octoTable.step() {
		step++
	}
	fmt.Println(step)
}
