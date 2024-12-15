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

type position struct {
	x int
	y int
}

func main() {
	file, err := os.Open("../input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	_map := make([][]rune, 0)
	y := 0
	robotX := 0
	robotY := 0
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			break
		}

		rowObjects := make([]rune, len(row))
		for x, value := range row {
			rowObjects[x] = value

			if value == '@' {
				robotX = x
				robotY = y
			}
		}
		_map = append(_map, rowObjects)
		y++
	}

	moves := make([]rune, 0)
	for scanner.Scan() {
		movesString := scanner.Text()
		for _, move := range movesString {
			moves = append(moves, move)
		}
	}

	// fmt.Println("Initial state:")
	// printMap(_map)

	for _, m := range moves {
		move(m, &robotX, &robotY, _map)
		// printMap(_map)
	}

	sum := 0

	for y, row := range _map {
		for x, value := range row {
			if value != 'O' {
				continue
			}

			sum += 100*y + x
		}
	}

	fmt.Println("sum", sum)
}

func move(m rune, robotX *int, robotY *int, _map [][]rune) {
	// fmt.Println("move", string(m), *robotX, *robotY)

	xD := 0
	yD := 0
	if m == '^' {
		yD = -1
	} else if m == '>' {
		xD = 1
	} else if m == 'v' {
		yD = 1
	} else if m == '<' {
		xD = -1
	}

	// fmt.Println("xD", xD, "yD", yD)

	count := 1
	for {
		next := _map[*robotY+(yD*count)][*robotX+(xD*count)]
		// fmt.Println("next", string(next))
		if next == 'O' {
			count++
			continue
		}
		if next == '#' {
			count = 0
			break
		}
		// .
		break
	}
	// fmt.Println("count", count)

	if count == 0 {
		// We are hitting a wall and cannot move
		return
	}

	// var next *object
	for i := count; i > 0; i-- {
		// fmt.Println("i", i)
		currX := *robotX + xD*i
		currY := *robotY + yD*i
		nextX := *robotX + xD*(i-1)
		nextY := *robotY + yD*(i-1)

		curr := _map[currY][currX]
		next := _map[nextY][nextX]
		// fmt.Println("curr", string(curr), currX, currY)
		// fmt.Println("next", string(next), nextX, nextY)

		_map[currY][currX] = next
		_map[nextY][nextX] = curr
		if next == '@' {
			*robotX += xD
			*robotY += yD
		}
	}
}

func printMap(_map [][]rune) {
	fmt.Println()
	for _, row := range _map {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}

	fmt.Println()
}
