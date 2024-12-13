package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse() []int {
	ret := make([]int, 0)
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	nums := strings.Split(string(content), " ")
	for _, numStr := range nums {
		num, err := strconv.Atoi(strings.Trim(numStr, "\n"))
		if err != nil {
			continue
		}
		ret = append(ret, num)
	}
	return ret
}
func main() {
	Main()
}

func Main() {
	stones := parse()
	for range 25 {
		// fmt.Println(stones)
		stones = blink(stones)
		// fmt.Println(stones)
	}
	fmt.Println(len(stones))

	newStones := parse2()
	for range 75 {
		newStones = blink2(newStones)
	}
	sum := 0
	for _, freq := range newStones {
		sum += freq
	}
	fmt.Println(sum)
}

func blink(nums []int) []int {
	ret := make([]int, 0)
	for _, num := range nums {
		str := strconv.Itoa(num)
		length := len(str)
		if length%2 == 0 {
			num1, _ := strconv.Atoi(str[:length/2])
			num2, _ := strconv.Atoi(str[length/2:])
			// fmt.Println(num1, num2, num)
			ret = append(ret, num1, num2)
			continue
		}
		if num == 0 {
			ret = append(ret, 1)
			continue
		}
		ret = append(ret, num*2024)
	}
	return ret
}

func blink2(nums map[int]int) map[int]int {
	ret := make(map[int]int)
	for num, frequency := range nums {
		str := strconv.Itoa(num)
		length := len(str)
		if length%2 == 0 {
			num1, _ := strconv.Atoi(str[:length/2])
			num2, _ := strconv.Atoi(str[length/2:])
			// fmt.Println(num1, num2, num)
			ret[num1] += frequency
			ret[num2] += frequency
			continue
		}
		if num == 0 {
			ret[1] += frequency
			continue
		}
		ret[num*2024] += frequency

	}
	return ret
}

func parse2() map[int]int {
	ret := make(map[int]int)
	content, _ := os.ReadFile("input.txt")
	nums := strings.Split(string(content), " ")
	for _, numStr := range nums {
		num, err := strconv.Atoi(strings.Trim(numStr, "\n"))
		if err != nil {
			continue
		}
		ret[num]++
	}
	return ret
}
