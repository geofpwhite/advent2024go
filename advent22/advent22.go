package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parse() []int {
	file, _ := os.Open("input.txt")
	// file, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(file)
	ary := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		ary = append(ary, num)
	}
	return ary
}

func prune(num int) int {
	return num % 16777216
}

func mix(num, num2 int) int {
	return num ^ num2
}

func nextSecret(num int) int {
	num = prune(mix(num*64, num))
	num = prune(mix(num/32, num))
	num = prune(mix(num*2048, num))
	// fmt.Println(num)
	return num
}

func changesAndValues(num int) map[[4]int]int {
	changes, values := make([]int, 000), make([]int, 000)
	realChanges := make(map[[4]int]int)
	hold := num
	// changes = append(changes, 0)
	for range 2000 {
		firstVal := hold % 10
		// fmt.Println(hold, hold%10)
		hold = nextSecret(hold)
		// fmt.Println(hold, hold%10, hold%10-firstVal)
		values = append(values, hold%10)
		changes = append(changes, (hold%10)-firstVal)
	}
	fmt.Println(values[:4])
	fmt.Println(changes[:4])
	curChanges := []int{changes[0], changes[1], changes[2], changes[3]}
	keyChanges := [4]int{}
	for i := range changes[:len(changes)-1] {
		keyChanges[0], keyChanges[1], keyChanges[2], keyChanges[3] = curChanges[0], curChanges[1], curChanges[2], curChanges[3]
		if i < 3 {
			continue
		}
		if realChanges[keyChanges] != 0 {
			curChanges = append(curChanges, changes[i+1])
			curChanges = curChanges[1:]
			continue
		}
		// fmt.Println(curChanges, values[i])
		if values[i] == 0 {
			realChanges[keyChanges] = -1
		}
		realChanges[keyChanges] = values[i]
		curChanges = append(curChanges, changes[i+1])
		curChanges = curChanges[1:]
	}
	return realChanges
}

func main() {
	nums := parse()
	sum := 0
	allChanges := make([]map[[4]int]int, 0)
	for _, num := range nums {
		changes := changesAndValues(num)
		allChanges = append(allChanges, changes)
		sum += num
	}

	fmt.Println(allChanges[0])
	maximum := 0
	for i := range allChanges {
		for key := range allChanges[i] {
			if allChanges[0][key] == 0 {
				allChanges[0][key] = -1
			}
		}
	}
	for key, value := range allChanges[0] {
		sum = value
		if value < 0 {
			sum = 0
		}
		for _, changes := range allChanges[1:] {
			if changes[key] < 0 {
				continue
			}
			sum += changes[key]
		}
		// fmt.Println(sum)

		if sum > maximum {
			// fmt.Println(sum)
			maximum = sum
		}
	}
	// for i := range allChanges {
	//
	// 	fmt.Println(allChanges[i][[4]int{-2, 1, -1, 3}])
	// 	fmt.Println(nums[i])
	// }
	fmt.Println(maximum)
	fmt.Println(sum)
}
