package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func buildRows() [][]rune {
	file, err := os.Open("../input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rows [][]rune
	for scanner.Scan() {
		var row []rune
		for _, char := range scanner.Text() {
			row = append(row, char)
		}
		rows = append(rows, row)
	}

	return rows
}

func checkForWord(dir string, word string, rows [][]rune, row int, col int) int {
	if len(word) == 0 {
		return 0
	}

	// fmt.Println()
	// fmt.Println("checking for", word, "starting at x", col+1, "y", row+1, dir)

	char := rows[row][col]

	if byte(char) == word[0] && len(word) == 1 {
		return 1
	}

	if byte(char) != word[0] {
		// fmt.Printf("skipping %c", char)
		// fmt.Println()
		return 0
	}

	count := 0
	wordLength := len(word)

	numCols := len(rows[0])
	numRows := len(rows)

	any := dir == "any"

	// up-left diagonal
	if (any || dir == "up-left") && row+1 >= wordLength && col+1 >= wordLength {
		// fmt.Println("checking up-left")
		count += checkForWord("up-left", word[1:], rows, row-1, col-1)
	}

	// up
	if (any || dir == "up") && row+1 >= wordLength {
		// fmt.Println("checking up")
		count += checkForWord("up", word[1:], rows, row-1, col)

	}

	// up-right diagonal
	if (any || dir == "up-right") && row+1 >= wordLength && numCols-col+1 > wordLength {
		// fmt.Println("checking up-right")
		count += checkForWord("up-right", word[1:], rows, row-1, col+1)

	}

	// right
	if (any || dir == "right") && numCols-col+1 > wordLength {
		// fmt.Println("checking right", word[1:])
		count += checkForWord("right", word[1:], rows, row, col+1)

	}

	// down right diagonal
	if (any || dir == "down-right") && numRows-row+1 > wordLength && numCols-col+1 > wordLength {
		// fmt.Println("checking down-right")
		count += checkForWord("down-right", word[1:], rows, row+1, col+1)

	}

	// down
	if (any || dir == "down") && numRows-row+1 > wordLength {
		// fmt.Println("checking down")
		count += checkForWord("down", word[1:], rows, row+1, col)

	}

	// down-left diagonal
	if (any || dir == "down-left") && numRows-row+1 > wordLength && col+1 >= wordLength {
		// fmt.Println("checking down-left")
		count += checkForWord("down-left", word[1:], rows, row+1, col-1)

	}

	// left
	if (any || dir == "left") && col+1 >= wordLength {
		// fmt.Println("checking left")
		count += checkForWord("left", word[1:], rows, row, col-1)
	}

	if any {
		return count
	} else if count != 0 {
		return 1
	}

	return 0
}

func main() {
	rows := buildRows()

	count := 0

	for row, rowChars := range rows {
		for col := range rowChars {
			found := checkForWord("any", "XMAS", rows, row, col)
			if found != 0 {
				// fmt.Println("found", found, "starting at", col+1, row+1)
			}
			count += found
		}
	}

	fmt.Println("count", count)
}
