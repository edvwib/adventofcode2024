package main

import (
	"bufio"
	// "fmt"
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

func checkForWord(dir string, word string, rows [][]rune, row int, col int) bool {
	if len(word) == 0 {
		return false
	}

	// fmt.Println()
	// fmt.Println("checking for", word, "starting at x", col+1, "y", row+1, dir)

	char := rows[row][col]

	if byte(char) == word[0] && len(word) == 1 {
		return true
	}

	if byte(char) != word[0] {
		// fmt.Printf("skipping %c", char)
		// fmt.Println()
		return false
	}

	wordLength := len(word)

	numCols := len(rows[0])
	numRows := len(rows)

	any := dir == "any"

	// up-left diagonal
	if (any || dir == "up-left") && row+1 >= wordLength && col+1 >= wordLength {
		// fmt.Println("checking up-left")
		return checkForWord("up-left", word[1:], rows, row-1, col-1)
	}

	// up
	if (any || dir == "up") && row+1 >= wordLength {
		// fmt.Println("checking up")
		return checkForWord("up", word[1:], rows, row-1, col)
	}

	// up-right diagonal
	if (any || dir == "up-right") && row+1 >= wordLength && numCols-col+1 > wordLength {
		// fmt.Println("checking up-right")
		return checkForWord("up-right", word[1:], rows, row-1, col+1)
	}

	// right
	if (any || dir == "right") && numCols-col+1 > wordLength {
		// fmt.Println("checking right", word[1:])
		return checkForWord("right", word[1:], rows, row, col+1)
	}

	// down right diagonal
	if (any || dir == "down-right") && numRows-row+1 > wordLength && numCols-col+1 > wordLength {
		// fmt.Println("checking down-right")
		return checkForWord("down-right", word[1:], rows, row+1, col+1)
	}

	// down
	if (any || dir == "down") && numRows-row+1 > wordLength {
		// fmt.Println("checking down")
		return checkForWord("down", word[1:], rows, row+1, col)

	}

	// down-left diagonal
	if (any || dir == "down-left") && numRows-row+1 > wordLength && col+1 >= wordLength {
		// fmt.Println("checking down-left")
		return checkForWord("down-left", word[1:], rows, row+1, col-1)

	}

	// left
	if (any || dir == "left") && col+1 >= wordLength {
		// fmt.Println("checking left")
		return checkForWord("left", word[1:], rows, row, col-1)
	}

	return false
}

func main() {
	rows := buildRows()

	count := 0

	numCols := len(rows[0])
	numRows := (len(rows))

	for row, rowChars := range rows {
		for col := range rowChars {
			if row > 0 && col > 0 && row <= numRows-2 && col <= numCols-2 {
				if checkForWord("any", "A", rows, row, col) {
					if checkForX(rows, row, col) {
						// fmt.Println("found at", row+1, col+1)
						count += 1
					}
				}
			}
		}
	}

	// fmt.Println("count", count)
}

func checkForX(rows [][]rune, row int, col int) bool {
	count := 0

	if checkForWord("down-right", "MAS", rows, row-1, col-1) {
		count += 1
	}

	if checkForWord("down-left", "MAS", rows, row-1, col+1) {
		count += 1
	}

	if checkForWord("up-left", "MAS", rows, row+1, col+1) {
		count += 1
	}

	if checkForWord("up-right", "MAS", rows, row+1, col-1) {
		count += 1
	}

	return count >= 2
}
