package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type coords struct {
	x, y int
}

type field struct {
	robotX, robotY int
	walls          map[coords]bool
	boxes          map[coords]bool
}
type moves = string

type field2 struct {
	robotX, robotY int
	walls          map[coords]bool
	boxes          map[[2]coords]bool
}

//up is x--
//down is x++
//right is y++
//left is y--

func parse() (field, moves) {
	m := ""
	f := field{
		walls: make(map[coords]bool),
		boxes: make(map[coords]bool),
	}
	// file, _ := os.Open("input.txt")
	file, _ := os.Open("test.txt")
	// file, _ := os.Open("test2.txt")
	scanner := bufio.NewScanner(file)
	x := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ".") || strings.Contains(line, "O") || strings.Contains(line, "#") || strings.Contains(line, "@") {
			for y, char := range line {
				if char == '@' {
					f.robotX, f.robotY = x, y
				}
				if char == 'O' {
					f.boxes[coords{x, y}] = true
				}
				if char == '#' {
					f.walls[coords{x, y}] = true
				}
			}
		}
		if strings.Contains(line, ">") || strings.Contains(line, "<") || strings.Contains(line, "v") || strings.Contains(line, "^") {
			m += strings.Trim(line, " \n")
		}
		x++
	}
	return f, m
}
func parse2() (field2, moves) {
	m := ""
	f := field2{
		walls: make(map[coords]bool),
		boxes: make(map[[2]coords]bool),
	}
	file, _ := os.Open("input.txt")
	// file, _ := os.Open("test.txt")
	// file, _ := os.Open("test2.txt")
	scanner := bufio.NewScanner(file)
	x := 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(len(line))
		// fmt.Println(line)
		line = strings.ReplaceAll(line, "O", "[]")
		line = strings.ReplaceAll(line, "#", "##")
		line = strings.ReplaceAll(line, ".", "..")
		line = strings.ReplaceAll(line, "@", "@.")
		// fmt.Println(line)
		// fmt.Println(len(line))
		if strings.Contains(line, ".") || strings.Contains(line, "O") || strings.Contains(line, "#") || strings.Contains(line, "@") {
			for y, char := range line {
				if char == '@' {
					f.robotX, f.robotY = x, y
				}
				if char == '[' {
					f.boxes[[2]coords{coords{x, y}, coords{x, y + 1}}] = true
				}
				if char == '#' {
					f.walls[coords{x, y}] = true
				}
			}
		}
		if strings.Contains(line, ">") || strings.Contains(line, "<") || strings.Contains(line, "v") || strings.Contains(line, "^") {
			m += strings.Trim(line, " \n")
		}
		x++
	}
	return f, m
}

func (f *field2) output(step int) {
	fileName := fmt.Sprintf("step%d.txt", step)
	lx, ly := 0, 0
	for coord := range f.walls {
		if coord.x > lx {
			lx = coord.x
		}
		if coord.y > ly {
			ly = coord.y
		}
	}
	str := ""
	for x := range lx {
		for y := range ly {
			if x == f.robotX && y == f.robotY {
				str += "@"
			} else if f.walls[coords{x, y}] {
				str += "#"
			} else if f.boxes[[2]coords{{x, y}, {x, y + 1}}] {
				str += "["
			} else if str[len(str)-1] == '[' {
				str += "]"
			} else {
				str += "."
			}
			// if !f.walls[coords{x,y}]&&f.
		}
		str += "\n"
	}
	file, _ := os.Create(fileName)
	file.WriteString(str)
	file.Close()
}

func main() {
	field, moves := parse()
	// fmt.Println(field, moves)
	for _, char := range moves {
		// fmt.Println(field.boxes, string(char))
		switch char {
		case '>': //y++
			field.right()
		case 'v': //x++
			field.down()
		case '<': //y--
			field.left()
		case '^': //x--
			field.up()
		}
		// fmt.Println(field.boxes)
	}
	sum := 0
	for box, exists := range field.boxes {
		if exists {
			sum += (100 * (box.x)) + box.y
		}
	}
	fmt.Println(sum)
	// fmt.Println(field.boxes)
	field2, moves := parse2()
	ogLength := len(field2.boxes)
	for i, char := range moves {
		// fmt.Println(len(field2.boxes))
		// fmt.Println(field2.robotX, field2.robotY)
		if i < 15 {
			field2.output(i)
		}
		// fmt.Println(field2.boxes[[2]coords{{37, 78}, {37, 79}}], field2.boxes[[2]coords{{37, 79}, {37, 80}}])
		switch char {
		case '>': //y++
			// fmt.Println("right")
			field2.right()
		case 'v': //x++
			// fmt.Println("down")
			field2.down()
		case '<': //y--
			// fmt.Println("left")
			field2.left()
		case '^': //x--
			// fmt.Println("up")
			field2.up()
		}

	}
	sum = 0
	i := 0
	for box, exists := range field2.boxes {
		if exists {
			// fmt.Println(box)
			// fmt.Println(sum)
			i++
			sum += (box[0].x * 100) + box[0].y
		}
	}
	fmt.Println(sum, i, ogLength)
}

func (f *field2) right() {
	rx, ry := f.robotX, f.robotY //right is y++
	boxesBelow := [][2]coords{}
	coordCheck := [2]coords{{rx, ry + 1}, {rx, ry + 2}}
	firstWallBelowCheck := coords{rx, ry + 1}
	if f.walls[firstWallBelowCheck] {
		return
	}
	for f.boxes[coordCheck] {
		boxesBelow = append(boxesBelow, coordCheck)
		coordCheck[0].y += 2
		coordCheck[1].y += 2
	}
	for !f.walls[firstWallBelowCheck] {
		firstWallBelowCheck.y++
		// fmt.Println(firstWallBelowCheck)
	}
	if firstWallBelowCheck.y == ry+(2*len(boxesBelow))+1 {
		return
	}
	if len(boxesBelow) == 0 {
		f.robotY++
		return
	}
	// delete(f.boxes, boxesBelow[0])
	fmt.Println(boxesBelow, "boxesbelow")
	for _, boxCoord := range boxesBelow {
		f.boxes[boxCoord] = false
		boxCoord[0].y++
		boxCoord[1].y++
		f.boxes[boxCoord] = true
	}
	fmt.Println(boxesBelow, "boxesbelow")
	f.robotY++
}
func (f *field2) left() {
	rx, ry := f.robotX, f.robotY //right is y++
	boxesBelow := [][2]coords{}
	coordCheck := [2]coords{{rx, ry - 2}, {rx, ry - 1}}
	firstWallBelowCheck := coords{rx, ry - 1}
	// if f.walls[firstWallBelowCheck] {
	// 	fmt.Println(firstWallBelowCheck)
	// 	return
	// }
	for f.boxes[coordCheck] {
		boxesBelow = append(boxesBelow, coordCheck)
		coordCheck[0].y -= 2
		coordCheck[1].y -= 2
	}
	// fmt.Println(firstWallBelowCheck, f.walls)
	for !f.walls[firstWallBelowCheck] {
		firstWallBelowCheck.y--
	}
	if firstWallBelowCheck.y == ry-(2*len(boxesBelow))-1 {
		return
	}
	if len(boxesBelow) == 0 {
		f.robotY--
		return
	}
	delete(f.boxes, boxesBelow[0])
	for _, boxCoord := range boxesBelow {
		f.boxes[boxCoord] = false
		boxCoord[0].y--
		boxCoord[1].y--
		f.boxes[boxCoord] = true
	}
	f.robotY--
}

func (f *field2) down() {
	connected := f.findConnectedBoxesBelow(f.robotX, f.robotY, [][2]coords{})
	if f.walls[coords{f.robotX + 1, f.robotY}] {
		return
	}
	if len(connected) == 0 && !f.walls[coords{f.robotX + 1, f.robotY}] {
		f.robotX++
		return
	}
	if len(connected) == 0 && f.walls[coords{f.robotX + 1, f.robotY}] {
		return
	}
	lowest := connected[len(connected)-1]
	if f.walls[coords{lowest[0].x + 1, lowest[0].y}] || f.walls[coords{lowest[1].x + 1, lowest[1].y}] {
		return
	}
	// lx := lowest[0].x
	for _, box := range connected {
		if f.walls[coords{box[0].x + 1, box[0].y}] || f.walls[coords{box[1].x + 1, box[1].y}] {
			return
		}
	}
	// if len(connected) > 1 {
	// 	i := 2
	// 	cur := connected[len(connected)-2]
	// 	for len(connected)-i >= 0 && cur[0].x == lx {
	// 		if f.walls[coords{cur[0].x + 1, cur[0].y}] || f.walls[coords{cur[1].x + 1, cur[1].y}] {
	// 			return
	// 		}
	// 		i++
	// 		cur = connected[len(connected)-i]
	// 	}
	// }
	slices.Reverse(connected)
	for _, box := range connected {
		newCoords := [2]coords{{box[0].x + 1, box[0].y}, {box[1].x + 1, box[1].y}}
		f.boxes[box] = false
		// delete(f.boxes, box)
		f.boxes[newCoords] = true
	}
	f.robotX++
}
func (f *field2) up() {
	connected := f.findConnectedBoxesAbove(f.robotX, f.robotY, [][2]coords{})
	fmt.Println(connected)
	if f.walls[coords{f.robotX - 1, f.robotY}] {
		return
	}
	if len(connected) == 0 && !f.walls[coords{f.robotX - 1, f.robotY}] {
		f.robotX--
		return
	}
	if len(connected) == 0 && f.walls[coords{f.robotX - 1, f.robotY}] {
		return
	}
	highest := connected[len(connected)-1]
	// lx := highest[0].x
	if f.walls[coords{highest[0].x - 1, highest[0].y}] || f.walls[coords{highest[1].x - 1, highest[1].y}] {
		return
	}
	for _, box := range connected {
		if f.walls[coords{box[0].x - 1, box[0].y}] || f.walls[coords{box[1].x - 1, box[1].y}] {
			return
		}

	}
	slices.Reverse(connected)
	for _, box := range connected {
		newCoords := [2]coords{{box[0].x - 1, box[0].y}, {box[1].x - 1, box[1].y}}
		f.boxes[box] = false
		// delete(f.boxes, box)
		f.boxes[newCoords] = true
	}
	f.robotX--
}

func (f *field2) findConnectedBoxesBelow(x, y int, curBoxes [][2]coords) [][2]coords { //only for up or down
	ret := make([][2]coords, len(curBoxes))
	copy(ret, curBoxes)
	rightCheck := [2]coords{{x + 1, y}, {x + 1, y + 1}}
	leftCheck := [2]coords{{x + 1, y - 1}, {x + 1, y}}
	if !f.boxes[rightCheck] && !f.boxes[leftCheck] {
		return curBoxes
	}
	if f.boxes[rightCheck] && !f.boxes[leftCheck] {
		ret = append(ret, rightCheck)
		ret = f.findConnectedBoxesBelow(rightCheck[0].x, rightCheck[0].y, ret)
		ret = f.findConnectedBoxesBelow(rightCheck[1].x, rightCheck[1].y, ret)
		ret = slices.Compact(ret)
	}
	if f.boxes[leftCheck] && !f.boxes[rightCheck] {
		ret = append(ret, leftCheck)
		ret = f.findConnectedBoxesBelow(leftCheck[0].x, leftCheck[0].y, ret)
		ret = f.findConnectedBoxesBelow(leftCheck[1].x, leftCheck[1].y, ret)
		ret = slices.Compact(ret)
	}

	if f.boxes[leftCheck] && f.boxes[rightCheck] {
		// fmt.Println(leftCheck, rightCheck)
		panic("overlapping boxes")
	}
	return ret
}
func (f *field2) findConnectedBoxesAbove(x, y int, curBoxes [][2]coords) [][2]coords { //only for up or down
	ret := make([][2]coords, len(curBoxes))
	copy(ret, curBoxes)
	rightCheck := [2]coords{{x - 1, y}, {x - 1, y + 1}}
	leftCheck := [2]coords{{x - 1, y - 1}, {x - 1, y}}
	if !f.boxes[rightCheck] && !f.boxes[leftCheck] {
		return curBoxes
	}
	if f.boxes[rightCheck] && !f.boxes[leftCheck] {
		ret = append(ret, rightCheck)
		ret = f.findConnectedBoxesAbove(rightCheck[0].x, rightCheck[0].y, ret)
		ret = f.findConnectedBoxesAbove(rightCheck[1].x, rightCheck[1].y, ret)
		ret = slices.Compact(ret)
	}
	if f.boxes[leftCheck] && !f.boxes[rightCheck] {
		ret = append(ret, leftCheck)
		ret = f.findConnectedBoxesAbove(leftCheck[0].x, leftCheck[0].y, ret)
		ret = f.findConnectedBoxesAbove(leftCheck[1].x, leftCheck[1].y, ret)
		ret = slices.Compact(ret)
	}

	if f.boxes[leftCheck] && f.boxes[rightCheck] {
		// fmt.Println(leftCheck, rightCheck)
		panic("overlapping boxes")
	}

	return ret
}

func (f *field) down() {
	rx, ry := f.robotX, f.robotY //down is x++
	boxesBelow := []coords{}
	coordCheck := coords{rx + 1, ry}
	firstWallBelowCheck := coords{rx + 1, ry}
	for f.boxes[coordCheck] {
		boxesBelow = append(boxesBelow, coordCheck)
		coordCheck.x++
	}
	for !f.walls[firstWallBelowCheck] {
		firstWallBelowCheck.x++
	}
	if firstWallBelowCheck.x == rx+len(boxesBelow)+1 {
		return
	}
	if len(boxesBelow) == 0 {
		f.robotX++
		return
	}
	delete(f.boxes, boxesBelow[0])
	for _, boxCoord := range boxesBelow {
		// f.boxes[boxCoord] = false
		// fmt.Println("deleted ", boxCoord)
		boxCoord.x++
		f.boxes[boxCoord] = true
		// fmt.Println("added ", boxCoord)
	}
	f.robotX++
}
func (f *field) up() {
	rx, ry := f.robotX, f.robotY //down is x++
	boxesBelow := []coords{}
	coordCheck := coords{rx - 1, ry}
	firstWallBelowCheck := coords{rx - 1, ry}
	for f.boxes[coordCheck] {
		boxesBelow = append(boxesBelow, coordCheck)
		coordCheck.x--
	}
	for !f.walls[firstWallBelowCheck] {
		firstWallBelowCheck.x--
	}
	if firstWallBelowCheck.x == rx-len(boxesBelow)-1 {
		return
	}
	if len(boxesBelow) == 0 {
		f.robotX--
		return
	}
	delete(f.boxes, boxesBelow[0])
	for _, boxCoord := range boxesBelow {
		// f.boxes[boxCoord] = false
		// delete(f.boxes, boxCoord)
		boxCoord.x--
		f.boxes[boxCoord] = true
	}
	f.robotX--
}
func (f *field) right() {
	rx, ry := f.robotX, f.robotY //right is y++
	boxesBelow := []coords{}
	coordCheck := coords{rx, ry + 1}
	firstWallBelowCheck := coords{rx, ry + 1}
	for f.boxes[coordCheck] {
		boxesBelow = append(boxesBelow, coordCheck)
		coordCheck.y++
	}
	for !f.walls[firstWallBelowCheck] {
		firstWallBelowCheck.y++
	}
	if firstWallBelowCheck.y == ry+len(boxesBelow)+1 {
		return
	}
	if len(boxesBelow) == 0 {
		f.robotY++
		return
	}
	delete(f.boxes, boxesBelow[0])
	for _, boxCoord := range boxesBelow {
		// f.boxes[boxCoord] = false
		boxCoord.y++
		f.boxes[boxCoord] = true
	}
	f.robotY++
}
func (f *field) left() {
	rx, ry := f.robotX, f.robotY //left is y--
	boxesBelow := []coords{}
	coordCheck := coords{rx, ry - 1}
	firstWallBelowCheck := coords{rx, ry - 1}
	for f.boxes[coordCheck] {
		boxesBelow = append(boxesBelow, coordCheck)
		coordCheck.y--
	}
	for !f.walls[firstWallBelowCheck] {
		firstWallBelowCheck.y--
	}
	if firstWallBelowCheck.y == ry-len(boxesBelow)-1 {
		return
	}
	if len(boxesBelow) == 0 {
		f.robotY--
		return
	}
	delete(f.boxes, boxesBelow[0])
	for _, boxCoord := range boxesBelow {
		// delete(f.boxes, boxCoord)
		// f.boxes[boxCoord] = false
		boxCoord.y--
		f.boxes[boxCoord] = true
	}
	f.robotY--
}
