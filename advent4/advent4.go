package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	wordSearch := parse()
	xmasCount := 0
	for i, line := range wordSearch {
		for j := range line {
			xmas := validXMASpathsFromIJ(i, j, wordSearch)
			xmasCount += xmas
		}
	}
	fmt.Println(xmasCount)
	crossCount := 0

	for i, line := range wordSearch {
		for j := range line {
			if validCrossMAS(i, j, wordSearch) {
				crossCount++
			}
		}
	}
	fmt.Println(crossCount)
}

func createNeighborCoords(i, j int) [8][3][2]int {
	return [8][3][2]int{
		{
			{
				i + 1, j,
			},
			{
				i + 2, j,
			},
			{
				i + 3, j,
			},
		},
		{
			{
				i, j + 1,
			},
			{
				i, j + 2,
			},
			{
				i, j + 3,
			},
		},
		{
			{
				i + 1, j + 1,
			},
			{
				i + 2, j + 2,
			},
			{
				i + 3, j + 3,
			},
		},
		{
			{
				i - 1, j + 1,
			},
			{
				i - 2, j + 2,
			},
			{
				i - 3, j + 3,
			},
		},
		{
			{
				i - 1, j - 1,
			},
			{
				i - 2, j - 2,
			},
			{
				i - 3, j - 3,
			},
		},
		{
			{
				i + 1, j - 1,
			},
			{
				i + 2, j - 2,
			},
			{
				i + 3, j - 3,
			},
		},
		{
			{
				i, j - 1,
			},
			{
				i, j - 2,
			},
			{
				i, j - 3,
			},
		},
		{
			{
				i - 1, j,
			},
			{
				i - 2, j,
			},
			{
				i - 3, j,
			},
		},
	}
}

func validCrossMAS(i, j int, wordSearch [][]rune) bool {
	if wordSearch[i][j] != 'A' || i <= 0 || j <= 0 || i >= len(wordSearch)-1 || j >= len(wordSearch[0])-1 {
		return false
	}
	oppositeEnds := [2][2][2]int{
		{
			{
				i - 1, j - 1,
			},
			{
				i + 1, j + 1,
			},
		},
		{
			{
				i + 1, j - 1,
			},
			{
				i - 1, j + 1,
			},
		},
	}

	for _, pair := range oppositeEnds {
		if !((wordSearch[pair[0][0]][pair[0][1]] == 'M' && wordSearch[pair[1][0]][pair[1][1]] == 'S') || (wordSearch[pair[0][0]][pair[0][1]] == 'S' && wordSearch[pair[1][0]][pair[1][1]] == 'M')) {
			return false
		}
	}
	return true
	// return (wordSearch[i-1][j-1] == 'M' && wordSearch[i-1][j+1] == 'M' && wordSearch[i+1][j+1] == 'S' && wordSearch[i+1][j-1] == 'S')||

}

func validXMASpathsFromIJ(i, j int, wordSearch [][]rune) int {
	if wordSearch[i][j] != 'X' {
		return 0
	}
	sum := 0

	neighborPaths := createNeighborCoords(i, j)

outer:
	for _, path := range neighborPaths {
		xmasString := "MAS"
		for z, coords := range path {
			x, y := coords[0], coords[1]
			if x < 0 || y < 0 || x > len(wordSearch)-1 || y > len(wordSearch[0])-1 || wordSearch[x][y] != rune(xmasString[z]) {
				continue outer
			}
		}
		sum++
	}
	return sum
}

func parse() [][]rune {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	wordSearch := make([][]rune, 0)

	for i, line := range lines {
		if line == "" {
			continue
		}
		wordSearch = append(wordSearch, make([]rune, len(line)))
		for j, char := range line {
			wordSearch[i][j] = char
		}
	}
	return wordSearch
}
