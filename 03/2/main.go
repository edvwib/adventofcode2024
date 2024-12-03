package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	total := 0
	do := true

	for scanner.Scan() {
		line := scanner.Text()

		matches, err := regexp.Compile("(mul\\((\\d+),(\\d+)\\))|(do\\(\\))|(don't\\(\\))")
		check(err)

		for _, match := range matches.FindAllStringSubmatch(line, -1) {
			if match[0] == "do()" {
				do = true
				continue
			}
			if match[0] == "don't()" {
				do = false
				continue
			}
			if do && strings.HasPrefix(match[0], "mul(") {
				n1, err := strconv.Atoi(match[2])
				check(err)
				n2, err := strconv.Atoi(match[3])
				check(err)

				total += n1 * n2
				continue
			}
		}
	}

	fmt.Println("total", total)
}
