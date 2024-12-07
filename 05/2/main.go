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

	rules := make(map[int][]int)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, "|")
		v1, err := strconv.Atoi(split[0])
		check(err)
		v2, err := strconv.Atoi(split[1])
		check(err)
		rules[v1] = append(rules[v1], v2)
	}

	updates := make(map[int][]int)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, stringValue := range strings.Split(line, ",") {
			v, err := strconv.Atoi(stringValue)
			check(err)
			updates[i] = append(updates[i], v)
		}
		i += 1
	}

	total := 0

updateLabel:
	for _, update := range updates {
		for pageIndex, page := range update {
			for pageIndex2, page2 := range update {
				if pageIndex2 == pageIndex {
					continue
				}

				if pageIndex > pageIndex2 {
					if slices.Contains(rules[page], page2) {

						slices.SortFunc(update, func(a int, b int) int {
							if slices.Contains(rules[a], b) {
								return -1
							}
							if slices.Contains(rules[b], a) {
								return 1
							}

							return 0
						})

						total += update[len(update)/2]
						continue updateLabel
					}
				}
			}
		}
	}

	fmt.Println("total", total)
}
