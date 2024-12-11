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

type stone struct {
	value int
	count int
}

const maxIterations = 75

func main() {
	file, err := os.Open("../input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()

		stones := make(map[int]int)
		for i, st := range strings.Split(row, " ") {
			stVal, err := strconv.Atoi(st)
			check(err)
			stones[i] = stVal
		}

		cache := make(map[stone]int)

		total := 0
		for _, stoneVal := range stones {
			total += recurseStone(stoneVal, 0, &cache)
		}

		fmt.Println("stones", total)
	}
}

func recurseStone(stoneValue int, iteration int, cache *map[stone]int) int {
	count := 1

	if iteration >= maxIterations {
		return 1
	}

	cachedValue, cached := (*cache)[stone{value: stoneValue, count: iteration}]
	if cached {
		return cachedValue
	}

	if stoneValue == 0 {
		count = recurseStone(1, iteration+1, cache)
	} else if str := strconv.Itoa(stoneValue); len(str)%2 == 0 {
		first, err := strconv.Atoi(str[0 : len(str)/2])
		check(err)
		second, err := strconv.Atoi(str[len(str)/2:])
		check(err)
		count = recurseStone(first, iteration+1, cache) + recurseStone(second, iteration+1, cache)
	} else {
		count = recurseStone(stoneValue*2024, iteration+1, cache)
	}

	(*cache)[stone{value: stoneValue, count: iteration}] = count
	return count
}
