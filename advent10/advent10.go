package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type node struct {
	left, right, up, down *node
	x, y, value           int
}

func main() {
	startTime := time.Now()
	coordsMap := parse()
	sum := 0
	for c, n := range coordsMap {
		if n.value == 0 {
			sum += getValidPaths(c[0], c[1], coordsMap, false)
		}
	}
	fmt.Println(sum)
	sum = 0
	for c, n := range coordsMap {
		if n.value == 0 {
			sum += getValidPaths(c[0], c[1], coordsMap, true)
		}
	}
	fmt.Println(sum)
	fmt.Println(time.Since(startTime))
}

func parse() map[[2]int]*node {

	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")

	parseMap := map[[2]int]*node{}

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		for j, char := range line {
			value, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			n := &node{x: i, y: j, value: value}
			parseMap[[2]int{i, j}] = n
		}
	}

	for coords, n := range parseMap {
		x, y := coords[0], coords[1]
		var up, left, down, right [2]int = [2]int{x - 1, y}, [2]int{x, y - 1}, [2]int{x + 1, y}, [2]int{x, y + 1}
		if parseMap[up] != nil {
			n.up = parseMap[up]
		}
		if parseMap[down] != nil {
			n.down = parseMap[down]
		}
		if parseMap[left] != nil {
			n.left = parseMap[left]
		}
		if parseMap[right] != nil {
			n.right = parseMap[right]
		}
	}
	return parseMap
}

func getValidPaths(x, y int, parseMap map[[2]int]*node, part2 bool) int {
	count := 0
	start := parseMap[[2]int{x, y}]

	queue, visited := make([]node, 0), make(map[node]bool)

	queue = append(queue, *start)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.value == 9 && !visited[cur] {
			count++
			if !part2 {
				visited[cur] = true
			}
			continue
		}
		visited[cur] = true
		if cur.up != nil && cur.up.value == cur.value+1 && !visited[*cur.up] {
			queue = append(queue, *cur.up)
		}
		if cur.down != nil && cur.down.value == cur.value+1 && !visited[*cur.down] {
			queue = append(queue, *cur.down)
		}
		if cur.left != nil && cur.left.value == cur.value+1 && !visited[*cur.left] {
			queue = append(queue, *cur.left)
		}
		if cur.right != nil && cur.right.value == cur.value+1 && !visited[*cur.right] {
			queue = append(queue, *cur.right)
		}
	}
	return count
}
