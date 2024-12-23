package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type numberRobot struct {
	x, y int
}

type directionRobot struct {
	x, y int
}

const A = -127

var numericKeyPad = [][]int{
	{7, 8, 9},
	{4, 5, 6},
	{1, 2, 3},
	{-1, 0, A},
}

var directionalKeyPad = [][]rune{
	{' ', '^', 'A'}, {'<', 'v', '>'},
}

var numKeyCoords = map[rune][2]int{
	'A': {3, 2},
	'0': {3, 1},
	'1': {2, 0},
	'2': {2, 1},
	'3': {2, 2},
	'4': {1, 0},
	'5': {1, 1},
	'6': {1, 2},
	'7': {0, 0},
	'8': {0, 1},
	'9': {0, 2},
}

var dirKeyCoords = map[rune][2]int{
	'A': {0, 2},
	'^': {0, 1},
	'<': {1, 0},
	'v': {1, 1},
	'>': {1, 2},
}

var directionCache = map[rune]string{
	'A': "A",
	'^': "<A",
	'<': "v<<A",
	'v': "v<A",
	'>': "vA",
}

func main() {
	// part1()
	part2()
}

func part2() {
	seq2seq := map[string]map[string]int{}
	num, dir, combos := parse()
	human := directionRobot{dirKeyCoords['A'][0], dirKeyCoords['A'][1]}
	cache := make(map[cacheObj]string)
	dirs := [4]directionRobot{}
	for i := range dirs {
		dirs[i].y = 2
	}
	// humanX, humanY := 0, 2
	moves := []int{}
	for _, str := range combos {
		numRobotMoves := ""
		dirRobotMoves := ""
		humanMoves := ""
		num = numberRobot{3, 2}
		dir = directionRobot{0, 2}
		human = directionRobot{0, 2}
		for _, char := range str {

			numRobotMoves += num.determineMovement(char, directionRobot{0, 2})
			num.x, num.y = numKeyCoords[char][0], numKeyCoords[char][1]
		}
		fmt.Println(numRobotMoves)
		seqHold := make(map[string]int)
		for _, char := range numRobotMoves {
			// fmt.Println(len(cache))
			// cache[cacheObj{directionRobot{0, 2}, str[:i]}] = dirRobotMoves
			move := dir.determineMovement(char, directionRobot{3, 2}, cache)
			dirRobotMoves += move
			seqHold[move] += 1
			fmt.Println(seqHold)
			// dir.x, dir.y = dirKeyCoords[char][0], dirKeyCoords[char][1]
			// dirRobotMoves += dir.determineMovement(char, directionRobot{3, 2}, cache)
			dir.x, dir.y = dirKeyCoords[char][0], dirKeyCoords[char][1]
		}
		keyMap := map[rune]int{
			'^': strings.Count(dirRobotMoves, "^"),
			'A': strings.Count(dirRobotMoves, "A"),
			'v': strings.Count(dirRobotMoves, "v"),
			'<': strings.Count(dirRobotMoves, "<"),
			'>': strings.Count(dirRobotMoves, ">"),
		}
		newMap := map[rune]int{}
		for key, value := range keyMap {
			for _, char := range directionCache[key] {
				newMap[char] += value
			}
		}
		fmt.Println(newMap, "keymap")
		// for i := range dirs {
		// 	h := dirRobotMoves
		// 	dirRobotMoves = ""
		// 	fmt.Println(h, len(h))
		// 	for _, char := range h {
		// 		dirRobotMoves += dirs[i].determineMovement(char, directionRobot{3, 2}, cache)
		// 		dirs[i].x, dirs[i].y = dirKeyCoords[char][0], dirKeyCoords[char][1]
		// 	}
		// }
		lastA := 0
		for i, char := range dirRobotMoves {
			move := human.determineMovement(char, directionRobot{3, 2}, cache)
			human.x, human.y = dirKeyCoords[char][0], dirKeyCoords[char][1]
			humanMoves += move
			seqHold[move] += 1
			if char == 'A' {
				seq2seq[dirRobotMoves[lastA:i+1]] = seqHold
				seqHold = make(map[string]int)
				lastA = i
			}
		}

		// fmt.Println(numRobotMoves, dirRobotMoves, humanMoves, str)
		// fmt.Println(len(humanMoves))
		moves = append(moves, len(humanMoves))
	}

	// sum := 0
	// for i, str := range combos {
	// 	num, _ := strconv.Atoi(str[:len(str)-1])
	// 	sum += (num * moves[i])
	// }
	// fmt.Println(num, dir, combos)
	// fmt.Println(sum)
	// for key, val := range cache {
	// 	fmt.Println(string(key.key), val)
	// }
	// fmt.Println(seq2seq)
}
func part1() {
	num, dir, combos := parse()
	human := directionRobot{dirKeyCoords['A'][0], dirKeyCoords['A'][1]}
	// humanX, humanY := 0, 2
	moves := []int{}
	for _, str := range combos {
		numRobotMoves := ""
		dirRobotMoves := ""
		humanMoves := ""
		num = numberRobot{3, 2}
		dir = directionRobot{0, 2}
		human = directionRobot{0, 2}
		for _, char := range str {

			numRobotMoves += num.determineMovement(char, directionRobot{0, 2})
			num.x, num.y = numKeyCoords[char][0], numKeyCoords[char][1]
		}

		for _, char := range numRobotMoves {
			dirRobotMoves += dir.determineMovement(char, directionRobot{3, 2}, make(map[cacheObj]string))
			dir.x, dir.y = dirKeyCoords[char][0], dirKeyCoords[char][1]
		}
		for _, char := range dirRobotMoves {
			humanMoves += human.determineMovement(char, directionRobot{3, 2}, make(map[cacheObj]string))
			human.x, human.y = dirKeyCoords[char][0], dirKeyCoords[char][1]
		}

		// fmt.Println(numRobotMoves, dirRobotMoves, humanMoves, str)
		// fmt.Println(len(humanMoves))
		moves = append(moves, len(humanMoves))
	}

	sum := 0
	for i, str := range combos {
		num, _ := strconv.Atoi(str[:len(str)-1])
		sum += (num * moves[i])
	}
	fmt.Println(num, dir, combos)
	fmt.Println(sum)
}

// /v<<A >>^A <A >A vA <^A A>A<vAAA>^A
// /v<<A ^>>A <A >A <A A v>A^Av<AAA^>A
func parse() (numberRobot, directionRobot, []string) {
	number := numberRobot{3, 2}
	dir := directionRobot{0, 2}
	// file, _ := os.Open("input.txt")
	file, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(file)
	combos := []string{}
	for scanner.Scan() {
		combos = append(combos, strings.Trim(scanner.Text(), "\n"))
		fmt.Println(len(scanner.Text()))
	}
	file.Close()
	return number, dir, combos
}

//      ^       A           <    <        ^ ^      A     >   >     A          v  v v         A
// <    A    >  A  v   <  < A    A >  ^   A  A  >  A  v  A   A  ^  A  <    v  A  A A >  ^    A
// <v<A >>^A vA ^A <vA <A A >>^A A vA <^A >A A vA ^A <vA >^A A <A >A <v<A >A >^A A A vA <^A >A

//	^       A       ^    ^           < <           A       > >     A           v v v      A
//
// <    A    >  A  <    A    A    <  v   A A >  >   ^  A   v   A A  ^  A    <  v   A A A >  ^ A
// <<vA >>^A vA ^A <<vA >>^A A <<vA >A >^A A vA A <^A >A <vA >^A A <A >A <<vA >A >^A A A vA <^A >A
// return something like "<<<vv>^A"
// if choosing between <^ vs. ^<, choose the one where the first key is closest to the directionBot

type permQueueNode struct {
	value      string
	numx, numy int
}

// ^^<<A>A>AvvA

func (n *numberRobot) determineMovement(key rune, dir directionRobot) string {

	movementString := ""
	coords := numKeyCoords[key]

	if coords[0] == n.x && coords[1] == n.y {
		return "A"
	}
	dirStringx, dirStringy := "", ""
	x, y := n.x, n.y
	dx := coords[0] - n.x
	dy := coords[1] - n.y
	if dx > 0 {
		dx = 1
		dirStringx = "v"
	} else if dx < 0 {
		dx = -1
		dirStringx = "^"
	}
	if dy > 0 {
		dy = 1
		dirStringy = ">"
	} else if dy < 0 {
		dy = -1
		dirStringy = "<"
	}

	distanceFromDirX1 := 0
	distanceFromDirY1 := 0
	distanceFromDirX2 := 0
	distanceFromDirY2 := 0
	if dirStringx != "" {
		distanceFromDirX1 = dir.x - dirKeyCoords[rune(dirStringx[0])][0]
		distanceFromDirY1 = dir.y - dirKeyCoords[rune(dirStringx[0])][1]

	}
	if dirStringy != "" {
		distanceFromDirX2 = dir.x - dirKeyCoords[rune(dirStringy[0])][0]
		distanceFromDirY2 = dir.y - dirKeyCoords[rune(dirStringy[0])][1]
	}

	dist1 := max(distanceFromDirX1, -distanceFromDirX1) + max(distanceFromDirY1, -distanceFromDirY1)
	dist2 := max(distanceFromDirX2, -distanceFromDirX2) + max(distanceFromDirY2, -distanceFromDirY2)
holdUp:
	hx, hy := x, y
	// dist1 := directionSums[dirStringx+dirStringy]
	// dist2 := directionSums[dirStringy+dirStringx]

	// fmt.Println(dist1, dist2)
	// fmt.Println(distanceFromDirX1, distanceFromDirY1)
	// fmt.Println(distanceFromDirX2, distanceFromDirY2)
	if dist2 > dist1 {

		for y != coords[1] {
			// fmt.Println(x, y)
			// fmt.Println(movementString)
			if x == 3 && y+dy == 0 {
				dist2 = -1
				x, y = hx, hy
				movementString = ""
				goto holdUp
			}
			movementString += dirStringy
			y += dy
		}
		for x != coords[0] {
			// fmt.Println(movementString)
			movementString += dirStringx
			x += dx
		}
	} else {
		for x != coords[0] {
			if x == 3-dx && y == 0 {
				dist1 = -1
				movementString = ""
				x, y = hx, hy
				goto holdUp
			}
			movementString += dirStringx
			x += dx
		}
		for y != coords[1] {
			// fmt.Println(movementString)
			movementString += dirStringy
			y += dy
		}
	}

	// fmt.Println(movementString+"A", "to move to", numKeyCoords[key], "from", n.x, n.y, dx, dy, dirStringx, dirStringy)
	return movementString + "A"
}

type cacheObj struct {
	dir directionRobot
	key string
}

func (d *directionRobot) determineMovement(key rune, dir directionRobot, cache map[cacheObj]string) string {
	movementString := ""
	coords := dirKeyCoords[key]
	dirStringx, dirStringy := "", ""
	if cache[cacheObj{*d, string(key)}] != "" {
		return cache[cacheObj{*d, string(key)}]
	}
	// defer func() { cache[cacheObj{*d, string(key)}] = movementString + "A" }()

	x, y := d.x, d.y
	if x == coords[0] && y == coords[1] {
		// fmt.Println(movementString+"A", "to move to", dirKeyCoords[key], "from", d.x, d.y, dirStringx, dirStringy)
		return "A"
	}
	dx := coords[0] - d.x
	dy := coords[1] - d.y
	if dx > 0 {
		dx = 1
		dirStringx = "v"
	} else if dx < 0 {
		dx = -1
		dirStringx = "^"
	}
	if dy > 0 {
		dy = 1
		dirStringy = ">"
	} else if dy < 0 {
		dy = -1
		dirStringy = "<"
	}
	for y != coords[1] {
		movementString += dirStringy
		y += dy
	}
	for x != coords[0] {
		movementString += dirStringx
		x += dx
	}
	distanceFromDirX1 := 0
	distanceFromDirY1 := 0
	distanceFromDirX2 := 0
	distanceFromDirY2 := 0
	if dirStringx != "" {
		distanceFromDirX1 = dir.x - dirKeyCoords[rune(dirStringx[0])][0]
		distanceFromDirY1 = dir.y - dirKeyCoords[rune(dirStringx[0])][1]
	}
	if dirStringy != "" {
		distanceFromDirX2 = dir.x - dirKeyCoords[rune(dirStringy[0])][0]
		distanceFromDirY2 = dir.y - dirKeyCoords[rune(dirStringy[0])][1]
	}

	dist1 := max(distanceFromDirX1, -distanceFromDirX1) + max(distanceFromDirY1, -distanceFromDirY1)
	dist2 := max(distanceFromDirX2, -distanceFromDirX2) + max(distanceFromDirY2, -distanceFromDirY2)
	// dist1 := directionSums[dirStringx+dirStringy]
	// dist2 := directionSums[dirStringy+dirStringx]

	if dist2 > dist1 {

		for y != coords[1] {
			// fmt.Println(x, y)
			// if x == 0 && y+dy == 2 {
			// 	movementString += dirStringx
			// 	x += dx
			// 	continue
			// }
			movementString += dirStringy
			y += dy
			// fmt.Println(movementString)
		}
		for x != coords[0] {
			movementString += dirStringx
			x += dx
		}
	} else {
		for x != coords[0] {
			// if x == 0-dx && y == 2 {
			// 	movementString += dirStringy
			// 	y += dy
			// 	continue
			// }
			movementString += dirStringx
			x += dx
		}
		for y != coords[1] {
			movementString += dirStringy
			y += dy
		}
	}
	// fmt.Println(movementString+"A", "to move to", dirKeyCoords[key], "from", d.x, d.y, dx, dy, dirStringx, dirStringy)
	return movementString + "A"

}
