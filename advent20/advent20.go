package main

import (
	"bufio"
	"fmt"
	"os"
)

type coords struct{ x, y int }

type node struct {
	u, d, l, r *neighbor
	coords     coords
}
type neighbor struct {
	node     *node
	distance int
}
type queueNode struct {
	coords coords
	score  int
}

type nodes map[coords]*node

func main() {
	//part1()
	part2()
}

func part1() {
	nodes, start, end := parse()
	timeToBeat := findPath(nodes, start, end)
	fmt.Println(timeToBeat)
	file, _ := os.Open("input.txt")
	// file, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sum := 0
	k := 0
	times := map[int]int{}
	for i := 1; i < len(lines)-1; i++ {

		for j := 1; j < len(lines[i])-1; j++ {
			k++
			if lines[i][j] == '#' {
				lines[i] = lines[i][:j] + "." + lines[i][j+1:]
				// fmt.Println(lines[i])
				nodes, start, end := parse2(lines)
				checkTime := findPath(nodes, start, end)
				// fmt.Println(checkTime, timeToBeat)
				if checkTime != timeToBeat && checkTime != 0 {
					times[timeToBeat-checkTime]++
				}
				if checkTime <= timeToBeat-100 && checkTime > 0 {
					fmt.Println(timeToBeat, checkTime)
					sum++
				}
				lines[i] = lines[i][:j] + "#" + lines[i][j+1:]
			}
		}
	}
	fmt.Println(sum)
	s2 := 0
	for time, num := range times {
		if time >= 100 {
			s2 += num
		}
	}
	// fmt.Println(times)
	fmt.Println(s2)
	//get optimal path without cheats
	// for _, n := range nodes {
	// 	fmt.Println(n)
	// }
}

func part2() {
	nodes, start, end := parse()
	visited := findPath2(nodes, start, end)
	fmt.Println(visited)
	sum := 0
	for coords, score := range visited {
		for coords2, score2 := range visited {
			if score <= score2 {
				continue
			}
			distance := max(coords.x-coords2.x, coords2.x-coords.x) + max(coords.y-coords2.y, coords2.y-coords.y)
			if distance > 20 {
				continue
			}
			if score2+distance <= score-100 {
				sum++
			}
		}
	}
	fmt.Println(sum)
}

func findPath(nodes nodes, start coords, end coords) int {
	queue := []queueNode{queueNode{start, 0}}
	visited := map[coords]int{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if nodes[cur.coords] == nil {
			return 0
		}
		if nodes[cur.coords].u != nil && (visited[nodes[cur.coords].u.node.coords] == 0 || visited[nodes[cur.coords].u.node.coords] > cur.score+nodes[cur.coords].u.distance) {
			visited[nodes[cur.coords].u.node.coords] = cur.score + nodes[cur.coords].u.distance
			queue = append(queue, queueNode{nodes[cur.coords].u.node.coords, cur.score + nodes[cur.coords].u.distance})
		}
		if nodes[cur.coords].d != nil && (visited[nodes[cur.coords].d.node.coords] == 0 || visited[nodes[cur.coords].d.node.coords] > cur.score+nodes[cur.coords].d.distance) {
			visited[nodes[cur.coords].d.node.coords] = cur.score + nodes[cur.coords].d.distance
			queue = append(queue, queueNode{nodes[cur.coords].d.node.coords, cur.score + nodes[cur.coords].d.distance})
		}
		if nodes[cur.coords].l != nil && (visited[nodes[cur.coords].l.node.coords] == 0 || visited[nodes[cur.coords].l.node.coords] > cur.score+nodes[cur.coords].l.distance) {
			visited[nodes[cur.coords].l.node.coords] = cur.score + nodes[cur.coords].l.distance
			queue = append(queue, queueNode{nodes[cur.coords].l.node.coords, cur.score + nodes[cur.coords].l.distance})
		}
		if nodes[cur.coords].r != nil && (visited[nodes[cur.coords].r.node.coords] == 0 || visited[nodes[cur.coords].r.node.coords] > cur.score+nodes[cur.coords].r.distance) {
			visited[nodes[cur.coords].r.node.coords] = cur.score + nodes[cur.coords].r.distance
			queue = append(queue, queueNode{nodes[cur.coords].r.node.coords, cur.score + nodes[cur.coords].r.distance})
		}
	}
	return visited[end]
}
func findPath2(nodes nodes, start coords, end coords) map[coords]int {
	queue := []queueNode{queueNode{start, 0}}
	visited := map[coords]int{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if nodes[cur.coords] == nil {
			return nil
		}
		if nodes[cur.coords].u != nil && (visited[nodes[cur.coords].u.node.coords] == 0 || visited[nodes[cur.coords].u.node.coords] > cur.score+nodes[cur.coords].u.distance) {
			visited[nodes[cur.coords].u.node.coords] = cur.score + nodes[cur.coords].u.distance
			inBetweenCur := cur.coords
			inBetweenScore := cur.score + 1
			inBetweenCur.x--
			for inBetweenCur != nodes[cur.coords].u.node.coords {
				visited[inBetweenCur] = inBetweenScore
				inBetweenCur.x--
				inBetweenScore++
			}
			queue = append(queue, queueNode{nodes[cur.coords].u.node.coords, cur.score + nodes[cur.coords].u.distance})
		}
		if nodes[cur.coords].d != nil && (visited[nodes[cur.coords].d.node.coords] == 0 || visited[nodes[cur.coords].d.node.coords] > cur.score+nodes[cur.coords].d.distance) {
			visited[nodes[cur.coords].d.node.coords] = cur.score + nodes[cur.coords].d.distance
			inBetweenCur := cur.coords
			inBetweenScore := cur.score + 1
			inBetweenCur.x++
			for inBetweenCur != nodes[cur.coords].d.node.coords {
				visited[inBetweenCur] = inBetweenScore
				inBetweenCur.x++
				inBetweenScore++
			}
			queue = append(queue, queueNode{nodes[cur.coords].d.node.coords, cur.score + nodes[cur.coords].d.distance})
		}
		if nodes[cur.coords].l != nil && (visited[nodes[cur.coords].l.node.coords] == 0 || visited[nodes[cur.coords].l.node.coords] > cur.score+nodes[cur.coords].l.distance) {
			visited[nodes[cur.coords].l.node.coords] = cur.score + nodes[cur.coords].l.distance
			inBetweenCur := cur.coords
			inBetweenScore := cur.score + 1
			inBetweenCur.y--
			for inBetweenCur != nodes[cur.coords].l.node.coords {
				visited[inBetweenCur] = inBetweenScore
				inBetweenCur.y--
				inBetweenScore++
			}
			queue = append(queue, queueNode{nodes[cur.coords].l.node.coords, cur.score + nodes[cur.coords].l.distance})
		}
		if nodes[cur.coords].r != nil && (visited[nodes[cur.coords].r.node.coords] == 0 || visited[nodes[cur.coords].r.node.coords] > cur.score+nodes[cur.coords].r.distance) {
			visited[nodes[cur.coords].r.node.coords] = cur.score + nodes[cur.coords].r.distance
			inBetweenCur := cur.coords
			inBetweenScore := cur.score + 1
			inBetweenCur.y++
			for inBetweenCur != nodes[cur.coords].r.node.coords {
				visited[inBetweenCur] = inBetweenScore
				inBetweenCur.y++
				inBetweenScore++
			}
			queue = append(queue, queueNode{nodes[cur.coords].r.node.coords, cur.score + nodes[cur.coords].r.distance})
		}
	}
	return visited

}

func parse() (nodes, coords, coords) {
	ns := make(nodes)
	start, end := coords{}, coords{}
	file, _ := os.Open("input.txt")
	// file, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	for i, line := range lines {
		for j, char := range line {
			if char == 'S' {
				start = coords{i, j}
			}
			if char == 'E' {
				end = coords{i, j}
			}
			if char != '#' {
				if isCorner(lines, i, j) {
					ns[coords{i, j}] = &node{coords: coords{i, j}}
				}
			}
		}
	}
	for _, n := range ns {
		for dx := -1; dx < 2; dx++ {
			for dy := -1; dy < 2; dy++ {
				if dx*dy != 0 || (dx == 0 && dy == 0) {
					continue
				}
				cur := coords{n.coords.x + dx, n.coords.y + dy}
				dist := 1
				for ns[cur] == nil && lines[cur.x][cur.y] == '.' {
					cur.x += dx
					cur.y += dy
					dist += 1
				}
				if ns[cur] == nil || (cur.x == n.coords.x && cur.y == n.coords.y) {
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
					ns[n.coords].u = &neighbor{ns[cur], dist}
				case 's':
					ns[n.coords].d = &neighbor{ns[cur], dist}
				case 'e':
					ns[n.coords].r = &neighbor{ns[cur], dist}
				case 'w':
					ns[n.coords].l = &neighbor{ns[cur], dist}
				}
			}
		}

	}
	return ns, start, end
}

func isCorner(lines []string, x, y int) bool {
	if lines[x][y] == 'E' {
		return true
	}
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
func parse2(lines []string) (nodes, coords, coords) {
	ns := make(nodes)
	start, end := coords{}, coords{}
	// file, _ := os.Open("input.txt")
	for i, line := range lines {
		for j, char := range line {
			if char == 'S' {
				start = coords{i, j}
			}
			if char == 'E' {
				end = coords{i, j}
			}
			if char != '#' {
				if isCorner(lines, i, j) {
					ns[coords{i, j}] = &node{coords: coords{i, j}}
				}
			}
		}
	}
	for _, n := range ns {
		for dx := -1; dx < 2; dx++ {
			for dy := -1; dy < 2; dy++ {
				if dx*dy != 0 || (dx == 0 && dy == 0) {
					continue
				}
				cur := coords{n.coords.x + dx, n.coords.y + dy}
				dist := 1
				for ns[cur] == nil && lines[cur.x][cur.y] == '.' {
					cur.x += dx
					cur.y += dy
					dist += 1
				}
				if ns[cur] == nil || (cur.x == n.coords.x && cur.y == n.coords.y) {
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
					ns[n.coords].u = &neighbor{ns[cur], dist}
				case 's':
					ns[n.coords].d = &neighbor{ns[cur], dist}
				case 'e':
					ns[n.coords].r = &neighbor{ns[cur], dist}
				case 'w':
					ns[n.coords].l = &neighbor{ns[cur], dist}
				}
			}
		}

	}
	return ns, start, end
}
