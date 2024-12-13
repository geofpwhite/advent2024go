package main

import (
	"bufio"
	"fmt"
	"os"
)

type corner struct {
	x, y int
	next *corner
	prev *corner
}

func parse() ([]string, int, int) {
	ret := make([]string, 0)
	f, _ := os.Open("input.txt")
	// f, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret, len(ret), len(ret[0])
}

func main() {
	lines, maxX, maxY := parse()
	visited := map[[2]int]bool{}

	shapes := []map[[2]int]bool{}
	perimeters := []int{}
	for i, line := range lines {
		for j := range line {
			if visited[[2]int{i, j}] {
				continue
			}
			otherShapeCoords, perimeter := getOtherShapeCoords(lines, i, j, maxX, maxY)
			for coord := range otherShapeCoords {
				visited[coord] = true
			}
			shapes = append(shapes, otherShapeCoords)
			perimeters = append(perimeters, perimeter)
		}
	}
	sum := 0
	for i, coordMap := range shapes {
		sum += len(coordMap) * perimeters[i]
	}
	fmt.Println(sum)

	sum = 0

	for _, coordMap := range shapes {
		sides := getSides(coordMap, maxX, maxY)
		sum += sides * len(coordMap)
	}
	fmt.Println(sum)
}

func getOtherShapeCoords(lines []string, x, y, maxX, maxY int) (map[[2]int]bool, int) {
	ret := make(map[[2]int]bool)
	queue := [][2]int{[2]int{x, y}}
	ret[[2]int{x, y}] = true
	perimeter := 0
	for len(queue) > 0 {
		curX, curY := queue[0][0], queue[0][1]
		queue = queue[1:]
		neighbors := [4][2]int{
			{curX + 1, curY},
			{curX - 1, curY},
			{curX, curY + 1},
			{curX, curY - 1},
		}
		for _, coords := range neighbors {
			if coords[0] < 0 || coords[0] >= maxX || coords[1] < 0 || coords[1] >= maxY || lines[coords[0]][coords[1]] != lines[x][y] {
				perimeter++
				continue
			}
			if ret[coords] {
				continue
			}
			ret[coords] = true
			queue = append(queue, coords)
		}
	}
	return ret, perimeter
}

func getSquaresForCorner(coord [2]int) [4][2]int {
	x, y := coord[0], coord[1]
	return [4][2]int{
		{x, y},
		{x + 1, y},
		{x, y + 1},
		{x + 1, y + 1},
	}
}

func getSides(coordMap map[[2]int]bool, maxX, maxY int) int {
	corners := make(map[[2]int]bool)
	for coords := range coordMap {
		curX, curY := coords[0], coords[1]
		corners[[2]int{curX - 1, curY - 1}] = true
		corners[[2]int{curX - 1, curY}] = true
		corners[[2]int{curX, curY - 1}] = true
		corners[[2]int{curX, curY}] = true
	}
	numCorners := 0
	nonEdges := map[[2]int]bool{}
	for coord := range nonEdges {
		delete(corners, coord)
	}

	for corner := range corners {
		squares := getSquaresForCorner(corner)
		goodSquares := make([][2]int, 0)
		count := 0
		for _, sq := range squares {
			if coordMap[sq] {
				goodSquares = append(goodSquares, sq)
				count++
			}
		}
		if count%2 == 1 {
			numCorners++
		}
		if count == 2 {
			if goodSquares[0][0] != goodSquares[1][0] && goodSquares[0][1] != goodSquares[1][1] {
				numCorners += 2
			}
		}
	}
	return numCorners
}
