package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		r, err := regexp.Compile("mul\\((\\d+),(\\d+)\\)")
		check(err)

		fmt.Println()

		for _, match := range r.FindAllStringSubmatch(line, -1) {
			n1, err := strconv.Atoi(match[1])
			check(err)
			n2, err := strconv.Atoi(match[2])
			check(err)

			total += n1 * n2
		}
	}

	fmt.Println("total", total)
}
