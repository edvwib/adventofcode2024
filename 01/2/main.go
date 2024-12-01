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

	var list1 []int
	var list2 []int

	for scanner.Scan() {
		row := scanner.Text()
		values := strings.Split(row, "   ")

		v1, err := strconv.Atoi(values[0])
		check(err)
		v2, err := strconv.Atoi(values[1])
		check(err)

		list1 = append(list1, v1)
		list2 = append(list2, v2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	var total int

	for _, value := range list1 {
		count := 0

		for _, value2 := range list2 {
			if value2 == value {
				count += 1
				continue
			}
			if value2 > value {
				break
			}
		}

		total += value * count
	}

	fmt.Println("score: ", total)
}
