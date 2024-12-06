package main

import (
	"fmt"
	"os"
	"strings"
)

type direction int

const (
	UP direction = iota
	DOWN
	LEFT
	RIGHT
)

type guard struct {
	dir  direction
	x, y int
}

func advent() {

	blockages, maxX, maxY, g := parse()
	visited := make(map[[2]int]bool)
	turns := make([][2]int, 0)
	steps := 0
	sum := 0
	for !g.moveOne(visited, maxX, maxY, blockages, &turns) {
		steps++
	}
	fmt.Println(len(visited))
	for x := range maxX {
		for y := range maxY {
			blockages, maxX, maxY, g = parse()
			if blockages[[2]int{x, y}] || (g.x == x && g.y == y) {
				continue
			}
			visited2 := make(map[guard]bool)
			blockages[[2]int{x, y}] = true
			// fmt.Println(x, y)
			// stop, cycle := g.moveOnePart2(visited2, maxX, maxY, blockages, &turns)
			stop, cycle := false, false
			for !stop && !cycle {
				stop, cycle = g.moveOnePart2(visited2, maxX, maxY, blockages, &turns)
				if cycle {
					sum++
					// fmt.Println(x, y)
					break
				}
			}
		}

	}
	fmt.Println(sum)
}

func (g *guard) turnRight() {
	switch g.dir {
	case UP:
		g.dir = RIGHT
	case RIGHT:
		g.dir = DOWN
	case DOWN:
		g.dir = LEFT
	case LEFT:
		g.dir = UP
	}
}

func (g *guard) moveOne(visited map[[2]int]bool, maxX, maxY int, blockages map[[2]int]bool, turns *[][2]int) bool { //returns true if we are finished
	var squareToCheck [2]int
	switch g.dir {
	case UP:
		squareToCheck = [2]int{g.x - 1, g.y}
	case DOWN:
		squareToCheck = [2]int{g.x + 1, g.y}
	case LEFT:
		squareToCheck = [2]int{g.x, g.y - 1}
	case RIGHT:
		squareToCheck = [2]int{g.x, g.y + 1}
	default:
		panic("no direction")
	}
	if squareToCheck[0] < 0 || squareToCheck[1] < 0 || squareToCheck[0] >= maxX || squareToCheck[1] >= maxY {
		return true
	}
	if blockages[squareToCheck] {
		g.turnRight()

		*turns = append(*turns, [2]int{g.x, g.y})
		return false
	}
	g.x = squareToCheck[0]
	g.y = squareToCheck[1]
	visited[squareToCheck] = true
	return false
}
func (g *guard) moveOnePart2(visited map[guard]bool, maxX, maxY int, blockages map[[2]int]bool, turns *[][2]int) (bool, bool) { //returns true if we are finished
	// fmt.Println(visited)
	var squareToCheck [2]int
	// fmt.Println(g.x, g.y)
	switch g.dir {
	case UP:
		squareToCheck = [2]int{g.x - 1, g.y}
	case DOWN:
		squareToCheck = [2]int{g.x + 1, g.y}
	case LEFT:
		squareToCheck = [2]int{g.x, g.y - 1}
	case RIGHT:
		squareToCheck = [2]int{g.x, g.y + 1}
	default:
		panic("no direction")
	}
	if squareToCheck[0] < 0 || squareToCheck[1] < 0 || squareToCheck[0] >= maxX || squareToCheck[1] >= maxY {
		return true, false
	}
	if blockages[squareToCheck] {
		visited[*g] = true
		g.turnRight()

		*turns = append(*turns, [2]int{g.x, g.y})
		if visited[*g] {
			return true, true
		}
		visited[*g] = true
		return false, false
	}
	visited[*g] = true
	g.x = squareToCheck[0]
	g.y = squareToCheck[1]
	if visited[*g] {
		return true, true
	}
	return false, false

}

func parse() (map[[2]int]bool, int, int, *guard) {
	g := &guard{}
	blockages := make(map[[2]int]bool)
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if line == "" {
			continue
		}

		for j, char := range line {
			switch char {
			case '#':
				blockages[[2]int{i, j}] = true
			case '^':
				g.dir = UP
				g.x = i
				g.y = j
			case '.':
				continue
			}
		}
	}
	return blockages, len(lines), len(lines[0]), g
}
func main() {
	advent()
}
