package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type coords struct {
	x, y int
}

type field map[coords]string
type reverseField map[string][]coords

func parse() (field, reverseField, int, int) {
	f := make(field)
	rf := make(reverseField)
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")

	lines := strings.Split(string(content), "\n")
	lines = slices.DeleteFunc(lines, func(line string) bool {
		return line == ""
	})
	for i, line := range lines {
		for j, char := range line {
			if char != '.' {
				if rf[string(char)] == nil {
					rf[string(char)] = make([]coords, 0)
				}
				f[coords{i, j}] = string(char)
				rf[string(char)] = append(rf[string(char)], coords{i, j})
			}
		}
	}
	return f, rf, len(lines), len(lines[0])
}

func main() {
	_, rf, maxX, maxY := parse()
	antinodeSpots := make(map[coords]bool)

	for _, coordList := range rf {
		for i, coord := range coordList {
			for _, coord2 := range coordList[i+1:] {
				newSpots := determineAntiNodePositions(coord, coord2)
				for _, newSpot := range newSpots {
					if newSpot.x < 0 || newSpot.y < 0 || newSpot.x >= maxX || newSpot.y >= maxY {
						continue
					}
					antinodeSpots[newSpot] = true
				}
			}
		}
	}
	fmt.Println(len(antinodeSpots))
	antinodeSpots = make(map[coords]bool)
	for _, coordList := range rf {
		for i, coord := range coordList {
			for _, coord2 := range coordList[i+1:] {
				newSpots := determineAntiNodePositionsPart2(coord, coord2, maxX, maxY)
				for _, newSpot := range newSpots {
					if newSpot.x < 0 || newSpot.y < 0 || newSpot.x >= maxX || newSpot.y >= maxY {
						continue
					}
					antinodeSpots[newSpot] = true
				}
			}
		}
	}
	fmt.Println(len(antinodeSpots))
}

func determineAntiNodePositions(node1, node2 coords) []coords {
	dx1, dy1 := node1.x-node2.x, node1.y-node2.y
	dx2, dy2 := -dx1, -dy1
	return []coords{{node1.x + dx1, node1.y + dy1}, coords{node2.x + dx2, node2.y + dy2}}
}
func determineAntiNodePositionsPart2(node1, node2 coords, maxX, maxY int) []coords {
	antinodes := []coords{node1, node2}
	dx, dy := node1.x-node2.x, node1.y-node2.y
	diffcoord := &coords{node1.x + dx, node1.y + dy}
	for diffcoord.x >= 0 && diffcoord.y >= 0 && diffcoord.x < maxX && diffcoord.y < maxY {
		antinodes = append(antinodes, *diffcoord)
		diffcoord.x += dx
		diffcoord.y += dy
	}
	diffcoord = &coords{node2.x - dx, node2.y - dy}
	for diffcoord.x >= 0 && diffcoord.y >= 0 && diffcoord.x < maxX && diffcoord.y < maxY {
		antinodes = append(antinodes, *diffcoord)
		diffcoord.x -= dx
		diffcoord.y -= dy
	}
	return antinodes
}
