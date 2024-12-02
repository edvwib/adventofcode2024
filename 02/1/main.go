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

		increasing := false
		levelCount := len(levels)

		for i, level := range levels {
			if i+1 >= levelCount {
				safeReports += 1
				break
			}

			nextLevel := levels[i+1]

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

	}

	fmt.Println("safe reports: ", safeReports)
}
