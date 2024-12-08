package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkCurPermutations(test int, values []int, cur int, index int) bool {

	if index >= len(values) && test == cur {
		return true
	}
	if index >= len(values) && test != cur {
		return false
	}
	return checkCurPermutations(test, values, cur+values[index], index+1) || checkCurPermutations(test, values, cur*values[index], index+1)
}
func checkCurPermutationsPart2(test int, values []int, cur int, index int) bool {

	if index >= len(values) && test == cur {
		return true
	}
	if index >= len(values) && test != cur {
		return false
	}
	str1, str2 := strconv.Itoa(cur), strconv.Itoa(values[index])
	combinedNum, err := strconv.Atoi(str1 + str2)
	// fmt.Println(combinedNum, test)
	if err != nil {
		panic("ah")
	}
	return checkCurPermutationsPart2(test, values, cur+values[index], index+1) ||
		checkCurPermutationsPart2(test, values, cur*values[index], index+1) ||
		checkCurPermutationsPart2(test, values, combinedNum, index+1)

}

func main() {
	sum := 0
	testByValues := parse()
	// fmt.Println(len(testByValues))
	// fmt.Println(testByValues[4322902])
	for test, values := range testByValues {
		if checkCurPermutations(*test, values[1:], values[0], 0) {
			// fmt.Println(test, values)
			sum += *test
		}
	}
	fmt.Println(sum)
	sum = 0
	for test, values := range testByValues {
		if checkCurPermutationsPart2(*test, values[1:], values[0], 0) {
			// fmt.Println(test, values)
			sum += *test
		}
	}
	fmt.Println(sum)
}

func parse() map[*int][]int {
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	testByValues := make(map[*int][]int)
	lines := strings.Split(string(content), "\n")
	fmt.Println(len(lines))
	// if lines[len(lines)-1] == "" {
	// 	lines = lines[:len(lines)-1]
	// }
	for _, line := range lines {
		fmt.Println(line)
		if line == "" {
			fmt.Println("c")
			continue
		}
		colonIndex := strings.Index(line, ":")
		testStr := line[:colonIndex]
		// fmt.Println(testStr)
		test, err := strconv.Atoi(testStr)
		if err != nil {
			fmt.Println(err)
			panic("ah")
		}
		numStrs := strings.Split(line[colonIndex+2:], " ")
		nums := make([]int, 0, len(numStrs))
		for _, str := range numStrs {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println(err)
				panic("ah")
			}
			nums = append(nums, num)
		}
		testByValues[&test] = nums
	}
	return testByValues
}
