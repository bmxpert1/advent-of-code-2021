package main

import (
	"bufio"
	"os"
	"strings"
)

type Pair struct {
	value  string
	yields [2]string
	adds   string
}

func main() {
	pairMap := map[string]*Pair{}

	file, _ := os.Open("example_input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var template string

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), " -> ") {
			pair := strings.Split(scanner.Text(), " -> ")
			couple, middle := pair[0], pair[1]
			pairMap[couple] = &Pair{
				value:  couple,
				yields: [2]string{string(couple[0]) + middle, middle + string(couple[1])},
				adds:   middle,
			}
		} else if len(scanner.Text()) > 0 {
			// starting template
			template = scanner.Text()
		}
	}

}
