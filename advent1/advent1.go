package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	firstNums, secondNums := make([]int, 0), make([]int, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}
		numStrings := strings.Split(line, "   ")
		first, err := strconv.Atoi(numStrings[0])
		if err != nil {
			panic("ah")
		}
		firstNums = append(firstNums, first)
		second, err := strconv.Atoi(numStrings[1])
		if err != nil {
			panic("ah")
		}
		secondNums = append(secondNums, second)
	}

	slices.Sort(firstNums)
	slices.Sort(secondNums)
	sum := 0

	for i := range firstNums {
		sum += int(math.Abs(float64(firstNums[i] - secondNums[i])))
	}
	fmt.Println(sum)
	sum = 0
	part2map := make(map[int]int)

	for _, num := range secondNums {
		part2map[num] += 1
	}

	for _, num := range firstNums {
		sum += (num * part2map[num])
	}
	fmt.Println(sum)
}
