package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	for scanner.Scan() {
		row := scanner.Text()

		var stoneInts []int
		var stoneStrings []string

		for _, stone := range strings.Split(row, " ") {
			stoneValue, err := strconv.Atoi(stone)
			check(err)
			stoneInts = append(stoneInts, stoneValue)
			stoneStrings = append(stoneStrings, stone)
		}

		loops := 0
		for {
			if loops >= 25 {
				break
			}

			loops++
			didInsert := false
			index := -1
			for {
				index++
				if didInsert {
					index++
					didInsert = false
				}

				if index >= len(stoneInts) {
					break
				}

				if stoneInts[index] == 0 {
					stoneInts[index] = 1
					stoneStrings[index] = "1"
					continue
				}

				stoneString := stoneStrings[index]
				stoneStringLen := len(stoneStrings[index])
				if stoneStringLen%2 == 0 {
					firstString := stoneString[0:(stoneStringLen / 2)]
					firstString = strings.TrimPrefix(firstString, "0")
					if firstString == "" {
						firstString = "0"
					}
					secondString := stoneString[(stoneStringLen / 2):]
					secondString = strings.TrimPrefix(secondString, "0")
					if secondString == "" {
						secondString = "0"
					}

					firstInt, err := strconv.Atoi(firstString)
					check(err)
					if firstInt == 0 {
						firstString = "0"
					}
					secondInt, err := strconv.Atoi(secondString)
					check(err)
					if secondInt == 0 {
						secondString = "0"
					}

					stoneInts[index] = firstInt
					stoneInts = slices.Insert(stoneInts, index+1, secondInt)

					stoneStrings[index] = firstString
					stoneStrings = slices.Insert(stoneStrings, index+1, secondString)

					didInsert = true

					continue
				}

				new := stoneInts[index] * 2024
				stoneInts[index] = new
				stoneStrings[index] = strconv.Itoa(new)
			}
		}

		fmt.Println("stones", len(stoneInts))
	}
}
