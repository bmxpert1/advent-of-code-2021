package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func findPaths(nodeMap map[string]*Node, specialNode *Node) (paths []string) {
	start, end := nodeMap["start"], nodeMap["end"]

	for _, node := range nodeMap {
		node.visitCount = 0
	}

	return start.pathsToTarget(end, specialNode)
}

func (n *Node) pathsToTarget(targetNode *Node, specialNode *Node) (paths []string) {
	n.visitCount++

	for _, conn := range n.connections {
		path := n.name
		if conn.targetNode.isSmall() &&
			((conn.targetNode != specialNode && conn.targetNode.visitCount > 0) ||
				(conn.targetNode == specialNode && conn.targetNode.visitCount > 1)) {
			// we can no longer visit this node
			continue
		} else {
			if conn.targetNode == targetNode {
				paths = append(paths, path+","+targetNode.name)
			} else {
				for _, p := range conn.targetNode.pathsToTarget(targetNode, specialNode) {
					paths = append(paths, path+","+p)
				}
			}
		}
	}
	n.visitCount--

	return paths
}

type Connection struct {
	sourceNode *Node
	targetNode *Node
}

func main() {
	// read input from txt
	connections := []string{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		connections = append(connections, scanner.Text())
	}

	nodeMap := map[string]*Node{}

	for _, conn := range connections {
		parts := strings.Split(conn, "-")
		start, end := parts[0], parts[1]

		if nodeMap[start] == nil {
			nodeMap[start] = &Node{name: start}
		}

		if nodeMap[end] == nil {
			nodeMap[end] = &Node{name: end}
		}

		nodeMap[start].addConnection(nodeMap[end], true)
	}

	//////////////////////////////////
	// challenge 1
	//
	fmt.Println(len(findPaths(nodeMap, nil)))

	//////////////////////////////////
	// challenge 2
	//
	paths := []string{}
	smalls := []*Node{}
	for _, node := range nodeMap {
		if node.isSmall() && node.name != "start" && node.name != "end" {
			smalls = append(smalls, node)
		}
	}
	for _, small := range smalls {
		paths = append(paths, findPaths(nodeMap, small)...)
	}
	uniqPaths := map[string]string{}
	for _, path := range paths {
		uniqPaths[path] = path
	}
	paths = []string{}
	for path := range uniqPaths {
		paths = append(paths, path)
	}

	fmt.Println(len(paths))
}
