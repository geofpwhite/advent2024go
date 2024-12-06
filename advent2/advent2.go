package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	INCR = iota
	DECR
)

func main() {
	reports := parse()
	sum := 0
outer:
	for _, report := range reports {
		rate := INCR
		if report[0] > report[1] {
			rate = DECR
		}
		cur := report[0]
		for _, num := range report[1:] {
			if diff := int(math.Abs(float64(num - cur))); diff > 3 || diff < 1 {
				continue outer
			}
			realDiff := cur - num
			if (realDiff < 0 && rate == DECR) || (realDiff > 0 && rate == INCR) {
				continue outer
			}
			cur = num
		}
		sum++
	}
	fmt.Println(sum)
	sum = 0
outer2:
	for _, report := range reports {
		rate := INCR
		if report[0] > report[1] {
			rate = DECR
		}
		cur := report[0]
		for _, num := range report[1:] {
			if diff := int(math.Abs(float64(num - cur))); diff > 3 || diff < 1 {
				if checkValidByRemoving(report) {
					sum++
				}
				continue outer2
			}
			realDiff := cur - num
			if (realDiff < 0 && rate == DECR) || (realDiff > 0 && rate == INCR) {
				if checkValidByRemoving(report) {
					sum++
				}
				continue outer2
			}
			cur = num
		}
		sum++
	}
	fmt.Println(sum)

}

func checkValidByRemoving(report []int) bool {
	reportChecks := make([][]int, len(report))
	for i := range report {
		newReport := make([]int, i)
		copy(newReport, report[:i])
		newReport = append(newReport, report[i+1:]...)
		reportChecks[i] = newReport
	}

outer3:
	for _, newReport := range reportChecks {
		rate := INCR
		if newReport[0] > newReport[1] {
			rate = DECR
		}
		cur := newReport[0]
		for _, num := range newReport[1:] {
			if diff := int(math.Abs(float64(num - cur))); diff > 3 || diff < 1 {
				continue outer3
			}
			realDiff := cur - num
			if (realDiff < 0 && rate == DECR) || (realDiff > 0 && rate == INCR) {
				continue outer3
			}
			cur = num
		}
		return true
	}
	return false
}

func parse() [][]int {
	reports := make([][]int, 0)
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	ind := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		reports = append(reports, make([]int, 0))
		for _, numStr := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(numStr)
			reports[ind] = append(reports[ind], num)
		}
		ind++
	}
	return reports
}
