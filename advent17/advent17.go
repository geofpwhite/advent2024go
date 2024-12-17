package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type computer struct {
	a, b, c          int
	commands         []int
	instrPointer     int
	output           []int
	outputInTermsOfA []string
}

type computer2 struct {
	a, b, c      string
	commands     []int
	instrPointer int
	output       []string
}

func find(input int, comp computer, index int, values []int) []int {
	hold := comp
	for i := range 2 {
		for j := range 2 {
			for k := range 2 {
				comp = hold
				newInp := (input << 3) + (i << 2) + (j << (1)) + (k)
				// fmt.Println(leng)
				comp.a = newInp
				leng := len(comp.output)
				fmt.Println(comp.output)
				fmt.Println(newInp)
				for leng == len(comp.output) && comp.instrPointer < len(comp.commands) {
					comp.advance()
				}
				if index == 0 {
					if comp.output[0] == comp.commands[0] {
						values = append(values, newInp)
					}
				} else if comp.commands[index] == comp.output[len(comp.output)-1] {
					values = find(newInp, comp, index-1, values)
				}
			}
		}
	}
	return values
}

// if (ind == 0) {
//                   if (Integer.parseInt(run(t, prog)) == prog[0]) {
//                       val.add(t);
//                   }
//               } else if (Integer.parseInt(run(t, prog)) == prog[ind]) {
//                   val = find(s, prog, ind - 1, val);
//               }

func main() {
	comp := parse()
	fmt.Println(comp)
	// comp.a = 70000000
	for comp.instrPointer < len(comp.commands) {
		comp.advance()
		fmt.Println(comp)
	}
	// fmt.Println(comp.output)
	str := ""
	for _, num := range comp.output {
		str += strconv.Itoa(num) + ","
	}
	fmt.Println(strings.Trim(str, ","))
	// correct := comp.commands
	//
	// possibles := [][]int{{1 << (3 * (len(comp.commands) - 1))}}
	// // possibles := [][]int{{0}}
	// possibleIndex := 0
	// index := len(comp.commands) - 1
	// for index != 0 {
	// 	K := 1 << (3 * index)
	// 	fmt.Println(K, "k")
	// 	possibles = append(possibles, make([]int, 0))
	// 	for _, num := range possibles[possibleIndex] {
	// 		for mul := range 8 {
	//
	// 			for i := range 2 {
	// 				for j := range 2 {
	// 					for k := range 2 {
	// 						// comp = og
	// 						comp = og
	// 						comp.a = (mul * ((i << (3 * (index))) + (j << ((3 * (index)) - 1)) + (k << ((3 * (index)) - 2))))
	// 						// fmt.Println((i << (3 * index)) + (j << (3*(index) - 1)) + (k << ((3 * index) - 2)))
	// 						// fmt.Println(comp.a)
	// 						// fmt.Println("---")
	// 						for len(comp.output) == 0 && comp.instrPointer < len(comp.commands) {
	// 							comp.advance()
	// 							// if len(comp.output) > 0 {
	// 							// 	break
	// 							// }
	// 						}
	// 						// if len(comp.output) < index || len(comp.output) > len(correct) {
	// 						// 	continue
	// 						// }
	// 						// if comp.output[index] = correct[index]{}
	// 						fmt.Println(comp.output, correct[index])
	// 						if comp.output[0] == correct[index] {
	// 							possibles[len(possibles)-1] = append(possibles[len(possibles)-1], num+(mul*((i<<(3*index))+(j<<((3*index)-1))+(k<<((3*index)-2)))))
	// 						}
	// 						if slices.Equal(comp.output, correct) {
	// 							// fmt.Println()
	// 							// fmt.Println()
	// 							fmt.Println(num + (mul * ((i << (3 * index)) + (j << (3*(index) - 1)) + (k << ((3 * index) - 2)))))
	// 							panic("")
	// 						}
	//
	// 					}
	// 				}
	// 			}
	// 		}
	// 		// for i := range 8 {
	// 		// 	comp = parse()
	// 		// 	comp.a = num + (K * (i))
	// 		// 	for comp.instrPointer < len(comp.commands) {
	// 		// 		comp.advance()
	// 		// 	}
	// 		// 	fmt.Println(K, num+(K*i), comp.output)
	// 		// 	fmt.Println(num)
	// 		// 	// fmt.Println(comp.output[index], correct[index])
	// 		// 	if comp.output[index] == correct[index] {
	// 		// 		possibles[len(possibles)-1] = append(possibles[len(possibles)-1], num+(K*(i)))
	// 		// 	}
	// 		// }
	// 	}
	// 	index--
	// 	possibleIndex++
	// }

	comp = parse()
	fmt.Println(find(0, comp, len(comp.commands)-1, []int{}))
}

func (comp *computer) comboOperand(num int) int {
	switch num {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return comp.a
	case 5:
		return comp.b
	case 6:
		return comp.c
	}
	return -1
}

func (comp *computer) determineModulusForOutputToBeSame() {
	// commands := comp.commands
	//what value of a makes it so the first output is equal to the first instr
	// a = x
	//given

}

func (comp *computer) advance() {
	instr := comp.commands[comp.instrPointer]
	operand := comp.commands[comp.instrPointer+1]
	switch instr {
	case 0:
		num := comp.a
		denom := math.Pow(2, float64(comp.comboOperand(operand)))
		comp.a = num / int(denom)
	case 1:
		comp.b = comp.b ^ operand
	case 2:
		comp.b = comp.comboOperand(operand) % 8
	case 3:
		if comp.a != 0 {
			comp.instrPointer = operand
			return
		}
	case 4:
		comp.b = comp.b ^ comp.c
	case 5: //outputs itself if combo operand is correct mod 8
		comp.output = append(comp.output, comp.comboOperand(operand)%8)
	case 6:
		num := comp.a
		denom := math.Pow(2, float64(comp.comboOperand(operand)))
		comp.b = num / int(denom)
	case 7:
		num := comp.a
		denom := math.Pow(2, float64(comp.comboOperand(operand)))
		comp.c = num / int(denom)
	}
	comp.instrPointer += 2
}

func parse() computer {
	file, _ := os.Open("input.txt")
	// file, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(file)
	var (
		a, b, c      int
		commands     []int
		instrPointer int = 0
	)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		switch i {
		case 0:

			aStr := line[strings.Index(line, ":")+2:]
			a, _ = strconv.Atoi(aStr)
			i++
			continue
		case 1:
			bStr := line[strings.Index(line, ":")+2:]
			b, _ = strconv.Atoi(bStr)
			i++
			continue
		case 2:
			cStr := line[strings.Index(line, ":")+1:]
			c, _ = strconv.Atoi(cStr)
			i++
			continue
		}
		if line == "" {
			continue
		}
		numStrs := strings.Split(line[strings.Index(line, ":")+2:], ",")
		for _, str := range numStrs {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println(err)
			}
			commands = append(commands, num)
		}
		i++
	}
	return computer{
		a, b, c, commands, instrPointer, make([]int, 0), make([]string, 0),
	}
}
