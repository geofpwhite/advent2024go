package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type node struct {
	label     string
	neighbors map[string]bool
}

type nodes map[string]*node

func parse() nodes {
	nodes := make(nodes)
	file, _ := os.Open("input.txt")
	// file, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		connectionString := scanner.Text()
		computers := strings.Split(connectionString, "-")
		if nodes[computers[0]] == nil {
			nodes[computers[0]] = &node{label: computers[0], neighbors: make(map[string]bool)}
		}
		if nodes[computers[1]] == nil {
			nodes[computers[1]] = &node{label: computers[1], neighbors: make(map[string]bool)}
		}
		nodes[computers[0]].neighbors[computers[1]] = true
		nodes[computers[1]].neighbors[computers[0]] = true
	}
	return nodes
}

func main() {
	nodes := parse()
	fmt.Println(nodes)
	triplets := map[[3]string]bool{}
	numberOfTriangles := map[string]int{}
	for nodeLabel, node := range nodes { //get all the pairs of neighbors and check if they are neighbors
		possibleTripletsWithNode := [][2]string{}
		for label := range node.neighbors {
			for label2 := range node.neighbors {
				if label == label2 {
					continue
				}
				possibleTripletsWithNode = append(possibleTripletsWithNode, [2]string{label, label2})
			}
		}
		for _, neighborsToCheck := range possibleTripletsWithNode {
			if nodes[neighborsToCheck[0]].neighbors[neighborsToCheck[1]] {
				x := []string{nodeLabel, neighborsToCheck[0], neighborsToCheck[1]}
				slices.Sort(x)
				// triplets = append(triplets, [3]string{x[0], x[1], x[2]})
				triplets[[3]string{x[0], x[1], x[2]}] = true
			}
		}
	}
	fmt.Println(triplets)
	sum := 0
	for triple := range triplets {
		if len(triple) < 3 {
			continue
		}
		for _, label := range triple {
			if label == "" {
				break
			}
			numberOfTriangles[label]++

			// if label[0] == 't' {
			// 	sum++
			// 	break
			// }
		}
	}
	m := 0
	for _, value := range numberOfTriangles {
		if value > m {
			m = value
		}
	}
	LANgroup := []string{}
	for label, value := range numberOfTriangles {
		if value == m {
			LANgroup = append(LANgroup, label)
		}
	}
	fmt.Println(sum)
	fmt.Println(len(nodes))
	fmt.Println(numberOfTriangles)
	fmt.Println(LANgroup)
	slices.Sort(LANgroup)
	fmt.Println(LANgroup)
}

func (ns *nodes) newLargestCores(curLargestCores [][]string, newRelation string) [][]string {
	computers := strings.Split(newRelation, "-")

	if (*ns)[computers[0]] == nil {
		(*ns)[computers[0]] = &node{label: computers[0], neighbors: make(map[string]bool)}
	}
	if (*ns)[computers[1]] == nil {
		(*ns)[computers[1]] = &node{label: computers[1], neighbors: make(map[string]bool)}
	}
	(*ns)[computers[0]].neighbors[computers[1]] = true
	(*ns)[computers[1]].neighbors[computers[0]] = true
	// newLargestCores := [][]string{}
	return [][]string{}
}
