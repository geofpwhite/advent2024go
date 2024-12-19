package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coords struct{ x, y int }
type qCoords struct {
	coords coords
	score  int
}

func parse() []coords {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	blocks := []coords{}
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		if line == "" {
			continue
		}
		xyStr := strings.Split(line, ",")
		xS, _ := strconv.Atoi(xyStr[0])
		yS, _ := strconv.Atoi(xyStr[1])
		blocks = append(blocks, coords{xS, yS})
	}
	return blocks
}

func part1(num int) int {

	blocks := parse()
	blocks = blocks[:num]
	queue := []qCoords{qCoords{coords{0, 0}, 0}}
	co := make(map[coords]bool)
	for _, block := range blocks {
		co[block] = true
	}
	visited := map[coords]int{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		neighbors := [4]coords{
			{cur.coords.x + 1, cur.coords.y},
			{cur.coords.x - 1, cur.coords.y},
			{cur.coords.x, cur.coords.y + 1},
			{cur.coords.x, cur.coords.y - 1},
		}
		for _, coord := range neighbors {
			if co[coord] || coord.x < 0 || coord.x > 70 || coord.y < 0 || coord.y > 70 {
				continue
			}

			if visited[coord] == 0 || visited[coord] > cur.score+1 {
				visited[coord] = cur.score + 1
				queue = append(queue, qCoords{coords: coord, score: cur.score + 1})
			}
		}

	}
	fmt.Println(visited[coords{70, 70}])
	return visited[coords{70, 70}]
}

func main() {
	blocks := parse()
	blocks = blocks[:1024]
	queue := []qCoords{qCoords{coords{0, 0}, 0}}
	co := make(map[coords]bool)
	for _, block := range blocks {
		co[block] = true
	}
	visited := map[coords]int{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		neighbors := [4]coords{
			{cur.coords.x + 1, cur.coords.y},
			{cur.coords.x - 1, cur.coords.y},
			{cur.coords.x, cur.coords.y + 1},
			{cur.coords.x, cur.coords.y - 1},
		}
		for _, coord := range neighbors {
			if co[coord] || coord.x < 0 || coord.x > 70 || coord.y < 0 || coord.y > 70 {
				continue
			}

			if visited[coord] == 0 || visited[coord] > cur.score+1 {
				visited[coord] = cur.score + 1
				queue = append(queue, qCoords{coords: coord, score: cur.score + 1})
			}
		}

	}
	fmt.Println(visited[coords{70, 70}])

	for i := range 3450 {
		score := part1(i)
		if score == 0 {
			fmt.Println(i)
			return
		}
	}
}
