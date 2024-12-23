package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type queueNode struct {
	score   int
	node    *node
	dir     rune
	visited int
}
type neighbor struct {
	node     *node
	distance int
}

type coords struct {
	x, y int
}
type node struct {
	coords     coords
	n, s, e, w *neighbor
}

type field struct {
	corners map[coords]*node
}

type visitedNode struct {
	node *node
	dir  rune
}

func main() {
	// fmt.Println(field, start, end)
	// prevNodes := make(map[visitedNode]map[visitedNode]bool)
	visited := make(map[visitedNode]int)
	visitedPath := make(map[visitedNode]string)
	// visited2 := make(map[visitedNode]map[string]bool)
	file, _ := os.Create("log.log")
	log.SetOutput(file)
	field, start, end := parse()
	// log.Println(len(field.corners))
	stack := []queueNode{queueNode{0, field.corners[start], '>', 0}}
	// winningPaths := map[string]bool{}
	larg := 0
	for len(stack) > 0 {
		// log.Println("-----------------------------------")
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// fmt.Println(len(queue))
		// if paths[cur.path] {
		// 	continue
		// }
		// if cur.node.coords == end && cur.score == 7036 {
		// if cur.node.coords == end && cur.score == 11048 {
		// 	continue
		// }
		if len(stack) > larg {
			larg = len(stack)
			fmt.Println(cur.score)
		}

		newQueueNodeL := turnLeft(&cur)
		newQueueNodeR := turnRight(&cur)
		newQueueNodeF := moveForward(&cur)

		visCheckL := visitedNode{node: newQueueNodeL.node, dir: newQueueNodeL.dir}
		visCheckR := visitedNode{node: newQueueNodeR.node, dir: newQueueNodeR.dir}
		visCheckF := visitedNode{node: newQueueNodeF.node, dir: newQueueNodeF.dir}

		if moveForward(&newQueueNodeL) != newQueueNodeL && visited[visCheckL] == 0 || (newQueueNodeL.score < visited[visCheckL] && newQueueNodeL.score > 0) {
			visited[visCheckL] = newQueueNodeL.score
			visitedPath[visCheckL] = visitedPath[visitedNode{node: cur.node, dir: cur.dir}] + "<"
			stack = append(stack, newQueueNodeL)
		}
		if moveForward(&newQueueNodeR) != newQueueNodeL && visited[visCheckR] == 0 || (newQueueNodeR.score < visited[visCheckR] && newQueueNodeR.score > 0) {
			visited[visCheckR] = newQueueNodeR.score
			visitedPath[visCheckR] = visitedPath[visitedNode{node: cur.node, dir: cur.dir}] + ">"
			stack = append(stack, newQueueNodeR)
		}
		if newQueueNodeF != cur && (visited[visCheckF] == 0 || newQueueNodeF.score < visited[visCheckF] && newQueueNodeF.score > 0) {
			visitedPath[visCheckF] = visitedPath[visitedNode{node: cur.node, dir: cur.dir}] + fmt.Sprintf("%d,%d-", cur.node.coords.x, cur.node.coords.y)
			visited[visCheckF] = newQueueNodeF.score
			stack = append(stack, newQueueNodeF)
		}
	}
	fmt.Println(visited[visitedNode{field.corners[end], '>'}])
	fmt.Println(visited[visitedNode{field.corners[end], '<'}])
	fmt.Println(visited[visitedNode{field.corners[end], '^'}])
	fmt.Println(visited[visitedNode{field.corners[end], 'v'}])

	queue := []coords{end}
	dirq := []rune{'>'}
	points := map[coords]bool{}
	for len(queue) > 0 {
		cur := field.corners[queue[0]]
		dir := dirq[0]
		queue = queue[1:]
		dirq = dirq[1:]

		connectedNodes := map[rune]coords{}
		connectedLs := map[rune]int{}
		if cur.coords == end {
			// visitedNode{field.corners[end], 'v'}
		}
		if cur.n != nil {
			connectedNodes['v'] = cur.n.node.coords
			connectedLs['v'] = cur.n.distance
		}
		if cur.s != nil {
			connectedNodes['^'] = cur.s.node.coords
			connectedLs['^'] = cur.s.distance
		}
		if cur.e != nil {
			connectedNodes['<'] = cur.e.node.coords
			connectedLs['<'] = cur.e.distance
		}
		if cur.w != nil {
			connectedNodes['>'] = cur.w.node.coords
			connectedLs['>'] = cur.w.distance
		}
		fmt.Println(cur, connectedLs)
		for cdir, c := range connectedNodes {
			// fmt.Println(visited[visitedNode{node: field.corners[c], dir: (dir)}], visited[visitedNode{node: cur, dir: dir}]-connectedLs[dir]-1)
			offset := 1000

			if cdir == dir || (cdir == '<' && dir == '>') || (cdir == '>' && dir == '<') || (cdir == 'v' && dir == '^') || (cdir == '^' && dir == 'v') {
				offset = 0
			}
			// offset := 0
			if visited[visitedNode{node: field.corners[c], dir: (cdir)}] == visited[visitedNode{node: cur, dir: dir}]-connectedLs[cdir]-1-offset || visited[visitedNode{node: field.corners[c], dir: (cdir)}] == visited[visitedNode{node: cur, dir: dir}]-connectedLs[cdir]-offset {
				points[c] = true
				queue = append(queue, c)
				dirq = append(dirq, cdir)
			}
		}
	}

	// optimalPath := []coords{}
	// optimalScore := []int{}
	// optimalString := visitedPath[visitedNode{field.corners[end], '>'}]
	// optimalCoordStrings := strings.Split(optimalString, "-")
	// for _, coordString := range optimalCoordStrings {
	// 	coordss := strings.Split(coordString, ",")
	// 	x, _ := strconv.Atoi(strings.Trim(coordss[0], "<v>^"))
	// 	y, _ := strconv.Atoi(strings.Trim(coordss[1], "<v>^"))
	// 	oPath := coords{x, y}
	// }

	fmt.Println(len(points))

	visitedEdges := map[[2]coords]bool{}
	for coord := range points {
		for coord1 := range points {
			if coord1 == coord {
				continue
			}
			visitedEdges[[2]coords{coord, coord1}] = true
		}
	}
	for edge := range visitedEdges {

		x, y := edge[0].x-edge[1].x, edge[0].y-edge[1].y

		dx, dy := 0, 0
		if x < 0 {
			dx = 1
		}
		if x > 0 {
			dx = -1
		}
		if y > 0 {
			dy = -1
		}
		if y < 0 {
			dy = 1
		}
		pointToAdd := edge[0]
		if !field.determineNeighbors(edge[0], edge[1]) {
			continue
		}
		if dx != 0 {

			for pointToAdd.x != edge[1].x {

				pointToAdd.x += dx
				points[pointToAdd] = true
				fmt.Println(pointToAdd, edge[0], edge[1], x, y, dx, dy)
			}
		} else if dy != 0 {

			for pointToAdd.y != edge[1].y {

				pointToAdd.y += dy
				points[pointToAdd] = true
				fmt.Println(pointToAdd, edge[0], edge[1], x, y, dx, dy)
			}
		}
		points[edge[0]] = true
		points[edge[1]] = true
	}
	fmt.Println(points)
	fmt.Println(len(points))

	o := ""
	for i := range 150 {
		for j := range 150 {
			if points[coords{i, j}] {
				o += "O"
			} else {
				o += "."
			}
		}
		o += "\n"
	}
	fi, _ := os.Create("out.check")
	fi.WriteString(o)
	fi.Close()

}

func turnLeft(cur *queueNode) queueNode {
	newQueueNode := *cur
	newQueueNode.score += 1000
	switch cur.dir {
	case '<':
		newQueueNode.dir = 'v'
	case '>':
		newQueueNode.dir = '^'
	case 'v':
		newQueueNode.dir = '>'
	case '^':
		newQueueNode.dir = '<'
	}
	return newQueueNode
}
func turnRight(cur *queueNode) queueNode {
	newQueueNode := *cur
	newQueueNode.score += 1000
	switch cur.dir {
	case '<':
		newQueueNode.dir = '^'
	case '>':
		newQueueNode.dir = 'v'
	case 'v':
		newQueueNode.dir = '<'
	case '^':
		newQueueNode.dir = '>'
	}
	return newQueueNode
}
func moveForward(cur *queueNode) queueNode {
	newQueueNode := *cur
	switch cur.dir {
	case '<':
		if newQueueNode.node.w != nil {
			newQueueNode.score += newQueueNode.node.w.distance + 1
			newQueueNode.node = newQueueNode.node.w.node
			newQueueNode.visited++
			return newQueueNode
		}
	case '^':
		if newQueueNode.node.n != nil {
			newQueueNode.score += newQueueNode.node.n.distance + 1
			newQueueNode.node = newQueueNode.node.n.node
			newQueueNode.visited++
		}
		return newQueueNode
	case 'v':
		if newQueueNode.node.s != nil {
			newQueueNode.score += newQueueNode.node.s.distance + 1
			newQueueNode.node = newQueueNode.node.s.node
			newQueueNode.visited++
		}
		return newQueueNode
	case '>':
		if newQueueNode.node.e != nil {
			newQueueNode.score += newQueueNode.node.e.distance + 1
			newQueueNode.node = newQueueNode.node.e.node
			newQueueNode.visited++
		}
		return newQueueNode
	}
	return *cur
}

func parse() (field, coords, coords) {
	f := field{
		make(map[coords]*node),
	}
	file, _ := os.Open("input.txt")
	// file, _ := os.Open("test.txt")
	// file, _ := os.Open("test2.txt")
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, strings.Trim(scanner.Text(), "\n"))
	}
	file.Close()
	var start, end coords

	for i, line := range lines {
		for j, char := range line {

			if isCorner(lines, i, j) {
				f.corners[coords{i, j}] = &node{coords: coords{i, j}}
			}
			if char == 'S' {
				start = coords{i, j}
			}
			if char == 'E' {
				end = coords{i, j}
			}
		}
	}

	for n := range f.corners {
		for dx := -1; dx < 2; dx++ {
			for dy := -1; dy < 2; dy++ {
				if dx*dy != 0 || (dx == 0 && dy == 0) {
					continue
				}
				cur := coords{n.x + dx, n.y + dy}
				dist := 1
				for f.corners[cur] == nil && lines[cur.x][cur.y] == '.' {
					cur.x += dx
					cur.y += dy
					dist += 1
				}
				if f.corners[cur] == nil || (cur.x == n.x && cur.y == n.y) {
					continue
				}
				dir := 'n'
				if dx == 1 {
					dir = 's'
				}
				if dy == 1 {
					dir = 'e'
				}
				if dy == -1 {
					dir = 'w'
				}
				switch dir {
				case 'n':
					log.Println("adding", cur, "as neighbor for", n)
					f.corners[n].n = &neighbor{f.corners[cur], dist}
				case 's':
					f.corners[n].s = &neighbor{f.corners[cur], dist}
				case 'e':
					f.corners[n].e = &neighbor{f.corners[cur], dist}
				case 'w':
					f.corners[n].w = &neighbor{f.corners[cur], dist}
				}
			}
		}

	}

	return f, start, end
}

func isCorner(lines []string, x, y int) bool {
	if lines[x][y] == '#' {
		return false
	}
	count := 0
	neighbors := map[coords]bool{}
	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			if dx*dy != 0 || (dx == 0 && dy == 0) {
				continue
			}

			if lines[x+dx][y+dy] == '.' {
				count++
				neighbors[coords{x + dx, y + dy}] = true
			}
		}
	}
	if count != 2 {
		return true
	}
	return !((neighbors[coords{x - 1, y}] && neighbors[coords{x + 1, y}]) || (neighbors[coords{x, y + 1}] && neighbors[coords{x, y - 1}]))

}

func (f *field) determineNeighbors(c1, c2 coords) bool {
	n1, n2 := f.corners[c1], f.corners[c2]
	if n1.e != nil && n1.e.node == n2 {
		return true
	}
	if n1.n != nil && n1.n.node == n2 {
		return true
	}
	if n1.s != nil && n1.s.node == n2 {
		return true
	}
	if n1.w != nil && n1.w.node == n2 {
		return true
	}
	return false
}
