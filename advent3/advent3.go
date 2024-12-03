package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	mulString := "mul"

	possibleMultiplications := strings.Split(parse(), mulString)
	sum := 0

	for _, str := range possibleMultiplications {

		if index := strings.Index(str, ")"); index != -1 {
			str = str[:index+1]
		}
		if str[0] != '(' {
			continue
		}
		numStrings := strings.Split(strings.Trim(str, "()"), ",")
		if len(numStrings) != 2 {
			continue
		}
		n1, err := strconv.Atoi(numStrings[0])
		if err != nil {
			continue
		}
		n2, err := strconv.Atoi(numStrings[1])
		if err != nil {
			continue
		}
		sum += n1 * n2
	}
	fmt.Println(sum)

	input := parse()
	dos := []string{}
	donts := []string{}
	cur := "do"

	for strings.Contains(input, "do()") || strings.Contains(input, "don't()") {
		indexDo := strings.Index(input, "do()")
		indexDont := strings.Index(input, "don't()")
		var addStr string
		if indexDont != -1 && indexDo != -1 {
			if indexDo < indexDont {
				addStr = input[:indexDo]
				input = input[indexDo+4:]
				if cur == "do" {
					dos = append(dos, addStr)
				} else {
					donts = append(donts, addStr)
				}
				cur = "do"
			} else {
				addStr = input[:indexDont]
				input = input[indexDont+7:]
				if cur == "do" {
					dos = append(dos, addStr)
				} else {
					donts = append(donts, addStr)
				}
				cur = "dont"
			}
		} else if indexDo == -1 {
			addStr = input[:indexDont]
			input = input[indexDont+7:]
			if cur == "do" {
				dos = append(dos, addStr)
			} else {
				donts = append(donts, addStr)
			}
			cur = "dont"
		} else if indexDont == -1 {
			addStr = input[:indexDo]
			input = input[indexDo+4:]
			if cur == "do" {
				dos = append(dos, addStr)
			} else {
				donts = append(donts, addStr)
			}
			cur = "do"
		} else {
			if cur == "do" {
				dos = append(dos, input)
			} else {
				donts = append(donts, input)
			}
		}
	}

	sum = 0
	for _, str1 := range dos {

		possibleMultiplications = strings.Split(str1, mulString)

		for _, str := range possibleMultiplications {
			if str == "" {
				continue
			}
			if index := strings.Index(str, ")"); index != -1 {
				str = str[:index+1]
			}
			if str[0] != '(' {
				continue
			}
			numStrings := strings.Split(strings.Trim(str, "()"), ",")
			if len(numStrings) != 2 {
				continue
			}
			n1, err := strconv.Atoi(numStrings[0])
			if err != nil {
				continue
			}
			n2, err := strconv.Atoi(numStrings[1])
			if err != nil {
				continue
			}
			sum += n1 * n2
		}

	}
	fmt.Println(sum)
}

func parse() string {
	content, _ := os.ReadFile("input.txt")
	return string(content)
}
