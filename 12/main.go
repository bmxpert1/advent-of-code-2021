package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Node struct {
	name        string
	connections []*Connection
	visitCount  int
}

func (n *Node) isBig() bool {
	return strings.ToUpper(n.name) == n.name
}

func (n *Node) isSmall() bool {
	return !n.isBig()
}

func (n *Node) addConnection(targetNode *Node, addReverse bool) {
	n.connections = append(n.connections, &Connection{n, targetNode})
	if addReverse {
		targetNode.addConnection(n, false)
	}
}

type Connection struct {
	sourceNode *Node
	targetNode *Node
}

func main() {
	// read input from txt
	paths := []string{}
	file, _ := os.Open("simple_example_input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		paths = append(paths, scanner.Text())
	}

	nodeMap := map[string]*Node{}

	for _, path := range paths {
		parts := strings.Split(path, "-")
		start, end := parts[0], parts[1]

		if nodeMap[start] == nil {
			nodeMap[start] = &Node{name: start}
		}

		if nodeMap[end] == nil {
			nodeMap[end] = &Node{name: end}
		}

		nodeMap[start].addConnection(nodeMap[end], true)
	}

	spew.Dump(nodeMap["b"])
}
