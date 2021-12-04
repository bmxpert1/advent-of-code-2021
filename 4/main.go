package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Option struct {
	value    int
	selected bool
}

type Board struct {
	options  [5][5]Option
	complete bool
}

func sum(nums []int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func (b *Board) checkPick(pick int) {
	if b.complete {
		return
	}

	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if b.options[y][x].value == pick {
				b.options[y][x].selected = true
			}
		}
	}
}

func (b *Board) checkWinner() bool {
	// check rows
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !b.options[y][x].selected {
				break
			}
			if x == 4 {
				return true
			}
		}
	}

	// check columns
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !b.options[y][x].selected {
				break
			}
			if y == 4 {
				return true
			}
		}
	}

	return false
}

func (b *Board) unmarked() (nums []int) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !b.options[y][x].selected {
				nums = append(nums, b.options[y][x].value)
			}
		}
	}

	return nums
}

func main() {
	// load picks from json
	var picks []int
	content, _ := ioutil.ReadFile("picks.json")
	json.Unmarshal([]byte(content), &picks)

	// read boards from txt
	boards := []*Board{}
	file, _ := os.Open("boards.txt")
	scanner := bufio.NewScanner(file)

	// var genBoard Board
	genBoard := new(Board)
	y := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			boards = append(boards, genBoard)
			genBoard = new(Board)
			y = 0
		} else {
			for x, opt := range strings.Fields(scanner.Text()) {
				optInt, _ := strconv.Atoi(opt)
				genBoard.options[y][x].value = optInt
			}
			y++
		}
	}
	// append last board
	boards = append(boards, genBoard)
	//
	// boards loaded
	/////////////////////////////

	/////////////////////////////
	// challenge 1
	//
	var firstWinner *Board
	var lastWinner *Board
	var firstWinningPick int
	var lastWinningPick int

	for i, pick := range picks {
		for _, b := range boards {
			if b.complete {
				continue
			}

			b.checkPick(pick)

			if i >= 4 {
				// start checking for winners
				if b.checkWinner() {
					b.complete = true
					if firstWinner == nil {
						firstWinner = b
						firstWinningPick = pick
					}
					lastWinner = b
					lastWinningPick = pick
				}
			}
		}
	}

	fmt.Println(sum(firstWinner.unmarked()) * firstWinningPick)

	/////////////////////////////
	// challenge 2
	//
	fmt.Println(sum(lastWinner.unmarked()) * lastWinningPick)
}
