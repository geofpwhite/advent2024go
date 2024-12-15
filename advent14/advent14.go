package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"
)

type robot struct {
	px, py, vx, vy int
}

func parse() []robot {
	robots := make([]robot, 0)
	f, _ := os.Open("input.txt")
	// f, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(f)
	r := &robot{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		positions := line[2:strings.Index(line, "v")]
		velocities := line[strings.Index(line, "v=")+2:]
		// fmt.Println(positions)
		// fmt.Println(velocities)
		ps := strings.Split(positions, ",")
		px, err := strconv.Atoi(ps[0])
		if err != nil {
			panic("ah")
		}
		py, err := strconv.Atoi(strings.Trim(ps[1], " "))
		if err != nil {
			fmt.Println(ps[1], err)
			panic("ah")
		}
		r.px, r.py = px, py
		vs := strings.Split(velocities, ",")
		vx, err := strconv.Atoi(vs[0])
		if err != nil {
			panic("ah")
		}
		vy, err := strconv.Atoi(vs[1])
		if err != nil {
			panic("ah")
		}
		r.vx, r.vy = vx, vy

		robots = append(robots, *r)
	}
	return robots
}
func main() {
	robots := parse()
	fmt.Println(robots)
	maxX, maxY := 101, 103

	// maxX, maxY := 11, 7
	for range 100 {
		for i := range robots {
			robots[i].px = (robots[i].px + robots[i].vx) % maxX
			robots[i].py = (robots[i].py + robots[i].vy) % maxY
			if robots[i].px < 0 {
				robots[i].px = maxX + robots[i].px
			}
			if robots[i].py < 0 {
				robots[i].py = maxY + robots[i].py
			}
			// robots[i].px = (robots[i].px + robots[i].vx) % 11
			// robots[i].py = (robots[i].py + robots[i].vy) % 7
		}
	}
	quarters := [4]int{}
	for _, r := range robots {
		if r.px < maxX/2 && r.py < maxY/2 {
			quarters[0]++
		}
		if r.px > maxX/2 && r.py < maxY/2 {
			quarters[1]++
		}
		if r.px < maxX/2 && r.py > maxY/2 {

			quarters[2]++
		}
		if r.px > maxX/2 && r.py > maxY/2 {
			quarters[3]++
		}

	}
	// fmt.Println(robots)
	// fmt.Println(quarters[0] * quarters[1] * quarters[2] * quarters[3])
	// fmt.Println(quarters)
	// fmt.Println(len(robots))
	robots = parse()
	coordMap := make(map[[2]int][]*robot)
	for _, r := range robots {
		coordMap[[2]int{r.px, r.py}] = []*robot{&r}
	}
	ogMap := make(map[[2]int][]*robot)
	maps.Copy(ogMap, coordMap)
	for i := range 10000 {
		fileName := fmt.Sprintf("step_%d", i)
		content := ""
		for j := range maxX {
			for k := range maxY {
				if coordMap[[2]int{k, j}] != nil {
					content += strconv.Itoa(len(coordMap[[2]int{j, k}]))
					continue
				}
				content += "."
			}
			content += "\n"
		}
		// fmt.Println(coordMap)
		newMap := make(map[[2]int][]*robot)
		for _, robots := range coordMap {
			for _, r := range robots {
				r.px = (r.px + r.vx) % maxX
				r.py = (r.py + r.vy) % maxY
				if r.px < 0 {
					r.px = maxX + r.px
				}
				if r.py < 0 {
					r.py = maxY + r.py
				}
				// fmt.Println(coord, robot.py, robot.py)

				if newMap[[2]int{r.px, r.py}] == nil {
					newMap[[2]int{r.px, r.py}] = []*robot{r}
				} else {
					newMap[[2]int{r.px, r.py}] = append(newMap[[2]int{r.px, r.py}], r)
				}
				// robots[i].px = (robots[i].px + robots[i].vx) % 11
				// robots[i].py = (robots[i].py + robots[i].vy) % 7
			}
		}

		if noOverlaps(coordMap) {
			file, _ := os.Create(fileName)
			file.WriteString(content)
			file.Close()
		} else {
			os.Remove(fileName)
		}
		coordMap = newMap
		newMap = make(map[[2]int][]*robot)
	}
}
func noOverlaps(coordMap map[[2]int][]*robot) bool {
	for _, ary := range coordMap {
		if len(ary) > 1 {
			// fmt.Println(ary[0], ary[1])
			return false
		}
	}
	return true
}
