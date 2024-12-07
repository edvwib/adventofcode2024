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

	total := 0
	for scanner.Scan() {
		equation := scanner.Text()
		rawValues := strings.Split(equation, ": ")

		testValue, err := strconv.Atoi(rawValues[0])
		check(err)

		var values []int
		for _, rawValue := range strings.Split(rawValues[1], " ") {
			value, err := strconv.Atoi(rawValue)
			check(err)
			values = append(values, value)
		}

		_, possible := test(testValue, 0, values, len(values))
		if possible {
			total += testValue
		}
	}

	fmt.Println("total", total)
}

var operators = [...]string{"+", "*", "||"}

func test(testValue int, currentValue int, values []int, initialValueCount int) (bool, bool) {
	valueCount := len(values)

	if valueCount == 0 && currentValue == testValue {
		return true, true
	}

	if currentValue > testValue {
		return false, false
	}

	maybeCount := 0
	possibleCount := 0

	if len(values) == 0 {
		return false, false
	}
	nextValues := values[1:]
	nextValue := values[0]

	for _, operator := range operators {
		newValue := 0
		switch operator {
		case "+":
			{
				newValue = currentValue + nextValue
			}
		case "*":
			{
				if currentValue == 0 && initialValueCount == valueCount {
					newValue = nextValue
				} else {
					newValue = currentValue * nextValue
				}
			}
		case "||":
			{
				concatenated, err := strconv.Atoi(strconv.Itoa(currentValue) + strconv.Itoa(nextValue))
				check(err)

				maybe := false
				possible := false
				maybe, possible = test(testValue, concatenated, nextValues[0:], initialValueCount)

				if maybe {
					maybeCount += 1
				}
				if possible {
					possibleCount += 1
				}

				continue
			}
		}

		maybe, possible := test(testValue, newValue, nextValues, initialValueCount)
		if maybe {
			maybeCount += 1
		}
		if possible {
			possibleCount += 1
		}
	}

	return maybeCount != 0, possibleCount != 0
}
