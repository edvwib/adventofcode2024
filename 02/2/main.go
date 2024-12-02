package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("../input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeReports := 0

	for scanner.Scan() {
		report := scanner.Text()
		levelsStrings := strings.Split(report, " ")
		var levels []int
		for _, levelString := range levelsStrings {
			levelInt, err := strconv.Atoi(levelString)
			check(err)
			levels = append(levels, levelInt)
		}

		inc := checkLevels(levels, 0)
		safeReports += inc
		// if inc == 1 {
		// 	fmt.Println("safe")
		// }
		// fmt.Println()
	}

	fmt.Println("safe reports: ", safeReports)
}

func checkLevels(levels []int, offset int) int {
	levelCount := len(levels)

	if offset >= levelCount {
		return 0
	}

	increasing := false

	// fmt.Println("offset", offset, levelCount, "levels", levels)

	currentLevels := remove(levels, offset)
	levelCount = len(currentLevels)
	// fmt.Println("currentLevels", currentLevels)
	// fmt.Println("levels", levels)

	for i, level := range currentLevels {
		if i+1 >= levelCount {
			return 1
		}

		nextLevel := currentLevels[i+1]

		if i == 0 {
			increasing = nextLevel > level
		}

		if level > nextLevel {
			if increasing {
				break
			}
			if level-nextLevel < 1 || level-nextLevel > 3 {
				break
			}
		} else if level < nextLevel {
			if !increasing {
				break
			}
			if nextLevel-level < 1 || nextLevel-level > 3 {
				break
			}
		} else {
			break
		}
	}

	return checkLevels(levels, offset+1)
}

func remove(s []int, index int) []int {
	new := make([]int, 0)
	new = append(new, s[:index]...)
	return append(new, s[index+1:]...)
}
