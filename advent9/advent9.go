package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	files, freeSpaceRanges := parse()
	fileIndex := len(files) - 1
	fileSize := len(files[fileIndex]) - 1
outer:
	for len(freeSpaceRanges) > 0 {
		curFreeSpaceRange := freeSpaceRanges[0]
		for curFreeSpaceRange[0] < curFreeSpaceRange[1]+1 {
			if files[fileIndex][fileSize] < curFreeSpaceRange[0] {
				break
			}
			files[fileIndex][fileSize] = curFreeSpaceRange[0]
			fileSize--
			curFreeSpaceRange[0]++
			if fileSize < 0 {
				fileIndex--
				if fileIndex == -1 {
					break outer
				}
				fileSize = len(files[fileIndex]) - 1
			}
		}
		freeSpaceRanges = freeSpaceRanges[1:]
	}
	checkSum := 0

	for fileID, filePositions := range files {
		for _, position := range filePositions {
			checkSum += position * fileID
		}
	}
	fmt.Println(checkSum)

	files, freeSpaceRanges = parse()
	fileRanges := [][2]int{}
	for _, file := range files {
		fileRanges = append(fileRanges, [2]int{file[0], file[len(file)-1]})
	}
	fmt.Println(files)
	fmt.Println(fileRanges)

	fileIndex = len(fileRanges) - 1

	for fileIndex > 0 {
		for i, spaceRange := range freeSpaceRanges {
			if fileRanges[fileIndex][0] < spaceRange[0] {
				break
			}
			if fileRanges[fileIndex][1]-fileRanges[fileIndex][0] == spaceRange[1]-spaceRange[0] {
				// freeSpaceRanges = append(freeSpaceRanges, [2]int{fileRanges[fileIndex][0], fileRanges[fileIndex][1]})
				fileRanges[fileIndex][1], fileRanges[fileIndex][0] = spaceRange[1], spaceRange[0]
				freeSpaceRanges = slices.Delete(freeSpaceRanges, i, i+1)

				// fmt.Println(freeSpaceRanges)
				// fmt.Println(freeSpaceRanges)
				break
			}
			if fileRanges[fileIndex][1]-fileRanges[fileIndex][0] < spaceRange[1]-spaceRange[0] {
				// freeSpaceRanges = append(freeSpaceRanges, [2]int{fileRanges[fileIndex][0], fileRanges[fileIndex][1]})
				diff := fileRanges[fileIndex][1] - fileRanges[fileIndex][0]
				fileRanges[fileIndex][1], fileRanges[fileIndex][0] = spaceRange[0]+diff, spaceRange[0]
				freeSpaceRanges[i][0] += diff + 1
				break
			}
		}
		fileIndex--
	}
	checkSum = 0
	for fileID, fileRange := range fileRanges {
		for i := fileRange[0]; i < fileRange[1]+1; i++ {
			checkSum += i * fileID
		}
	}
	fmt.Println(checkSum)
}

func parse() ([][]int, [][2]int) { //returns file ids by array of positions + free space ranges
	files, freeSpaceRanges := make([][]int, 0), make([][2]int, 0)
	// content, _ := os.ReadFile("input.txt")
	content, _ := os.ReadFile("test.txt")
	line := strings.Trim(string(content), "\n")
	fileID := 0
	position := 0
	for i, char := range line {

		num, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}
		if i%2 == 0 {
			files = append(files, []int{})
			for j := range num {
				files[i/2] = append(files[fileID], position+j)
			}
			fileID++
		} else if num > 0 {
			freeSpaceRanges = append(freeSpaceRanges, [2]int{position, position + num - 1})
		}
		position += num
	}
	return files, freeSpaceRanges
}
