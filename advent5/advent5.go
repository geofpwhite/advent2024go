package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"slices"
	"strconv"
	"strings"
	"text/scanner"
)

func main() {
	keyBeforeValue, _, updates := parse()
	sum := 0
	needsChanged := [][]int{}
outer:
	for _, update := range updates {
		for i, num := range update {
			for j := i + 1; j < len(update); j++ {
				num2 := update[j]
				if !slices.Contains(keyBeforeValue[num], num2) {
					needsChanged = append(needsChanged, update)
					continue outer
				}
			}
		}
		sum += update[len(update)/2]
	}
	fmt.Println(sum)
	sum = 0
	for _, update := range needsChanged {
		fmt.Println(update)
		newUpdate := []int{}
		for _, num := range update {
			newUpdate = insertInOrder(newUpdate, keyBeforeValue, num)
		}
		fmt.Println(newUpdate)
		sum += newUpdate[len(newUpdate)/2]
	}
	fmt.Println(sum)
}

func insertInOrder(curUpdate []int, keyBeforeValue map[int][]int, value int) []int {
	if len(curUpdate) == 0 {
		curUpdate = append(curUpdate, value)
		return curUpdate
	}
	for i, num := range curUpdate {
		if !slices.Contains(keyBeforeValue[num], value) {
			newPart := make([]int, len(curUpdate)+1)
			copy(newPart[:i], curUpdate[:i])
			newPart[i] = value
			copy(newPart[i+1:], curUpdate[i:])
			clear(curUpdate)
			return newPart
		}
	}
	curUpdate = append(curUpdate, value)
	return curUpdate
}

func parse() (map[int][]int, map[int][]int, [][]int) {
	keyBeforeValue, keyAfterValue := make(map[int][]int), make(map[int][]int)
	updates := make([][]int, 0)

	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, "|") {
			numStrings := strings.Split(line, "|")
			before, err := strconv.Atoi(numStrings[0])
			if err != nil {
				panic("error while reading")
			}
			after, err := strconv.Atoi(numStrings[1])
			if err != nil {
				panic("error while reading")
			}
			if keyBeforeValue[before] == nil {
				keyBeforeValue[before] = []int{after}
			} else {
				keyBeforeValue[before] = append(keyBeforeValue[before], after)
			}
			if keyAfterValue[after] == nil {
				keyAfterValue[after] = []int{before}
			} else {
				keyAfterValue[after] = append(keyAfterValue[after], before)
			}
		}

		if strings.Contains(line, ",") {
			updateNumStrings := strings.Split(line, ",")
			newUpdate := make([]int, len(updateNumStrings))
			for i, str := range updateNumStrings {
				num, err := strconv.Atoi(str)
				if err != nil {
					panic("error reading")
				}
				newUpdate[i] = num
			}
			updates = append(updates, newUpdate)
		}
	}
	return keyBeforeValue, keyAfterValue, updates
}
