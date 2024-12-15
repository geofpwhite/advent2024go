package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type claw struct {
	AX, AY, BX, BY, X, Y int
}

func parse() []claw {
	claws := make([]claw, 0)
	curClaw := &claw{}
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if index := strings.Index(line, "Button A:"); index != -1 {
			line := line[index+10:]
			strs := strings.Split(line, ", ")
			ax, e := strconv.Atoi(strs[0][2:])
			if e != nil {
				fmt.Println(e)
			}
			ay, _ := strconv.Atoi(strs[1][2:])
			curClaw.AX, curClaw.AY = ax, ay
		}
		if index := strings.Index(line, "Button B:"); index != -1 {
			line := line[index+10:]
			strs := strings.Split(line, ", ")
			bx, _ := strconv.Atoi(strs[0][2:])
			by, _ := strconv.Atoi(strs[1][2:])
			curClaw.BX, curClaw.BY = bx, by
		}
		if index := strings.Index(line, "Prize:"); index != -1 {
			line := line[index+7:]
			strs := strings.Split(line, ", ")
			x, _ := strconv.Atoi(strs[0][2:])
			y, _ := strconv.Atoi(strs[1][2:])
			curClaw.X, curClaw.Y = x, y
			claws = append(claws, *curClaw)
		}
	}
	return claws
}
func main() {
	claws := parse()
	fmt.Println(claws)
	sum := 0
	for _, claw := range claws {
		x0, y0 := claw.AX, claw.AY
		x1, y1 := claw.BX, claw.BY
		x2, y2 := claw.X, claw.Y

		b := (y2 - (y0 * x2 / x0)) / (y1 - (x1 * y0 / x0))
		_b := ((y2 * x0) - (y0 * x2)) / ((y1 * x0) - (x1 * y0))
		a := (x2 - (b * x1)) / x0
		_a := (x2 / x0) - (b * x1 / x0)
		aSource := ((y2 * x1) - (y1 * x2)) / ((y0 * x1) - (x0 * y1))
		bSource := (y2 - (aSource * y0)) / y1
		// a2 := ((y2) - (y1 * x2 / x0)) / (y0 - (y1 * x0 / x0))
		// b2 := (x2 - (a2 * x0)) / x0
		// fmt.Println(a, b, _b, (a*3)+_b, (a*x0)+(_b*x1), (a*y0)+(_b*y1), x2, y2)
		//
		// // fmt.Println(a2, b2, (a2*3)+b2, (a2*x0)+(b2*x1))
		// fmt.Println(a >= 0, _b >= 0, a < 100, _b < 100, ((a*x0)+(_b*x1)) == x2)
		// fmt.Println(a, x0, _b, x1, x2)
		// fmt.Println(a*x0 + (_b * x1))
		// fmt.Println(a >= 0 && _b >= 0 && a < 100 && _b < 100 && ((a*x0)+(_b+x1)) == x2)
		ax0bx1 := (a * x0) + (_b * x1)
		ax0Bx1 := (a * x0) + (b * x1)
		Ax0bx1 := (_a * x0) + (b * x1)
		Ax0Bx1 := (_a * x0) + (_b * x1)
		asx0bsx0 := (aSource * x0) + (bSource * x1)
		// fmt.Println(ax0bx1, x2)
		fmt.Println(a, _a, aSource, "a")
		fmt.Println(b, _b, bSource, "b")
		fmt.Println(sum)
		if a >= 0 && _b >= 0 && a < 100 && _b < 100 && ax0bx1 == x2 {
			// fmt.Println(ax0bx1, x2, ax0bx1 == x2)
			// fmt.Println("true for", a, _b)
			sum += (a * 3) + _b
			// fmt.Println(sum)
		} else if a >= 0 && b >= 0 && a < 100 && b < 100 && ax0Bx1 == x2 {
			// fmt.Println(b, _b)
			// fmt.Println(ax0Bx1, x2, ax0Bx1 == x2)
			// fmt.Println("true for", a, b)
			sum += (a * 3) + b
			// fmt.Println(sum)
		} else if _a >= 0 && b >= 0 && _a < 100 && b < 100 && Ax0bx1 == x2 {
			// fmt.Println(b, _b)
			// fmt.Println(ax0Bx1, x2, ax0Bx1 == x2)
			// fmt.Println("true for", a, b)
			sum += (_a * 3) + b
			// fmt.Println(sum)
		} else if _a >= 0 && _b >= 0 && _a < 100 && _b < 100 && Ax0Bx1 == x2 {
			// fmt.Println(b, _b)
			// fmt.Println(ax0Bx1, x2, ax0Bx1 == x2)
			// fmt.Println("true for", a, b)
			sum += (_a * 3) + _b
			// fmt.Println(sum)
		} else if aSource >= 0 && bSource >= 0 && aSource < 100 && bSource < 100 && asx0bsx0 == x2 {
			// fmt.Println(b, _b)
			// fmt.Println(ax0Bx1, x2, ax0Bx1 == x2)
			// fmt.Println("true for", a, b)
			fmt.Println(x2, asx0bsx0)
			sum += (aSource * 3) + bSource
			// fmt.Println(sum)
		} else {
			fmt.Println("no solution for", claw)
			fmt.Println(ax0bx1, ax0Bx1, Ax0bx1, Ax0Bx1, asx0bsx0)
			fmt.Println(a, _a, b, _b)
			fmt.Println(97*23 + (126 * 64))
			fmt.Println(96*23 + (126 * 64))
			fmt.Println(95*23 + (129 * 64))
			// ax0Bx1 := (a * x0) + (b * x1)

		}
		fmt.Println(sum)

	}
	fmt.Println(sum)
	sum = 0
	for i := range claws {
		claws[i].X += 10000000000000
		claws[i].Y += 10000000000000
	}
	for i, claw := range claws {
		x0, y0 := claw.AX, claw.AY
		x1, y1 := claw.BX, claw.BY
		x2, y2 := claw.X, claw.Y

		b := (y2 - (y0 * x2 / x0)) / (y1 - (x1 * y0 / x0))
		_b := ((y2 * x0) - (y0 * x2)) / ((y1 * x0) - (x1 * y0))
		a := (x2 - (b * x1)) / x0
		_a := (x2 / x0) - (b * x1 / x0)
		aSource := ((y2 * x1) - (y1 * x2)) / ((y0 * x1) - (x0 * y1))
		bSource := (y2 - (aSource * y0)) / y1
		// a2 := ((y2) - (y1 * x2 / x0)) / (y0 - (y1 * x0 / x0))
		// b2 := (x2 - (a2 * x0)) / x0
		// fmt.Println(a, b, _b, (a*3)+_b, (a*x0)+(_b*x1), (a*y0)+(_b*y1), x2, y2)
		//
		// // fmt.Println(a2, b2, (a2*3)+b2, (a2*x0)+(b2*x1))
		// fmt.Println(a >= 0, _b >= 0, a < 100, _b < 100, ((a*x0)+(_b*x1)) == x2)
		// fmt.Println(a, x0, _b, x1, x2)
		// fmt.Println(a*x0 + (_b * x1))
		// fmt.Println(a >= 0 && _b >= 0 && a < 100 && _b < 100 && ((a*x0)+(_b+x1)) == x2)
		ax0bx1 := (a * x0) + (_b * x1)
		ax0Bx1 := (a * x0) + (b * x1)
		Ax0bx1 := (_a * x0) + (b * x1)
		Ax0Bx1 := (_a * x0) + (_b * x1)
		asx0bsx0 := (aSource * x0) + (bSource * x1)
		ay0by1 := (a * y0) + (_b * y1)
		ay0By1 := (a * y0) + (b * y1)
		Ay0by1 := (_a * y0) + (b * y1)
		Ay0By1 := (_a * y0) + (_b * y1)
		asy0bsy0 := (aSource * y0) + (bSource * y1)
		// fmt.Println(ax0bx1, x2)
		fmt.Println(a, _a, aSource, "a")
		fmt.Println(b, _b, bSource, "b")
		fmt.Println(sum)
		if a >= 0 && _b >= 0 && ax0bx1 == x2 && ay0by1 == y2 {
			// fmt.Println(ax0bx1, x2, ax0bx1 == x2)
			// fmt.Println("true for", a, _b)
			fmt.Println(sum, "+=", (a*3)+_b)
			sum += (a * 3) + _b
		} else if a >= 0 && b >= 0 && ax0Bx1 == x2 && ay0By1 == y2 {
			// fmt.Println(b, _b)
			// fmt.Println(ax0Bx1, x2, ax0Bx1 == x2)
			// fmt.Println("true for", a, b)
			fmt.Println(sum, "+=", (a*3)+b)
			sum += (a * 3) + b
			// fmt.Println(sum)
		} else if _a >= 0 && b >= 0 && Ax0bx1 == x2 && Ay0by1 == y2 {
			// fmt.Println(b, _b)
			// fmt.Println(ax0Bx1, x2, ax0Bx1 == x2)
			// fmt.Println("true for", a, b)
			fmt.Println(sum, "+=", (_a*3)+b)
			sum += (_a * 3) + b
			// fmt.Println(sum)
		} else if _a >= 0 && _b >= 0 && _a < 100 && _b < 100 && Ax0Bx1 == x2 && Ay0By1 == y2 {
			// fmt.Println(b, _b)
			// fmt.Println(ax0Bx1, x2, ax0Bx1 == x2)
			// fmt.Println("true for", a, b)
			fmt.Println(sum, "+=", (_a*3)+_b)
			sum += (_a * 3) + _b
			// fmt.Println(sum)
		} else if aSource >= 0 && bSource >= 0 && asx0bsx0 == x2 && asy0bsy0 == y2 {
			// fmt.Println(b, _b)
			// fmt.Println(ax0Bx1, x2, ax0Bx1 == x2)
			// fmt.Println("true for", a, b)
			fmt.Println(x2, asx0bsx0)
			fmt.Println(sum, "+=", (aSource*3)+bSource)
			sum += (aSource * 3) + bSource
			// fmt.Println(sum)
		} else {
			fmt.Println("no solution for", i, claw)
			fmt.Println(ax0Bx1, ax0bx1, Ax0Bx1, asx0bsx0, Ax0bx1, x2)
			// ax0Bx1 := (a * x0) + (b * x1)
			continue
		}
		fmt.Println("solution for ", i, ax0Bx1, ax0bx1, Ax0Bx1, asx0bsx0, Ax0bx1, x2)
		fmt.Println(sum)

	}
}
