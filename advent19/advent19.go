package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parse() ([]string, []string) {
	file, _ := os.Open("input.txt")
	// file, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(file)
	towels := make([]string, 0)
	designs := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.Contains(line, ",") {
			towels = strings.Split(line, ", ")
		} else {
			designs = append(designs, line)
		}
	}

	return towels, designs
}

func main() {
	towels, designs := parse()
	sum := 0
	for _, design := range designs {
		if valid(design, towels) {
			sum++
		}
	}
	fmt.Println(sum)

	sum = 0

	dp := make(map[string]int)
	for _, design := range designs {
		if valid(design, towels) {
			sum += valid2(design, towels, 0, dp)

		}
	}
	fmt.Println(sum)
}

func valid(design string, towels []string) bool {
	for _, towel := range towels {
		if towel == design {
			return true
		}
		if len(towel) < len(design) {
			if towel == design[:len(towel)] {
				if valid(design[len(towel):], towels) {
					return true
				}
			}
		}
	}
	return false
}

func valid2(design string, towels []string, value int, dp map[string]int) int {
	// fmt.Println(design)
	if dp[design] != 0 {
		if dp[design] == -1 {
			return 0
		}
		return dp[design]
	}
	for _, towel := range towels {
		if towel == design {
			if dp[towel] == -1 {
				dp[towel] = 1
			} else {
				dp[towel] += 1
			}
			value += 1
		}
		if len(towel) < len(design) {
			if towel == design[:len(towel)] {
				if matches := valid2(design[len(towel):], towels, 0, dp); matches > 0 {
					value += matches
					dp[design[len(towel):]] = matches
				} else {
					dp[design[len(towel):]] = -1
				}
			}
		}
	}

	dp[design] = value
	return value
}
