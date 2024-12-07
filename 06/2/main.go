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

func main() {
	file, err := os.Open("../input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	_map := make(map[int][]rune)
	xStart := 0
	yStart := 0
	xVel := 0
	yVel := 0

	mapY := 0
	for scanner.Scan() {
		line := scanner.Text()
		for mapX, value := range line {
			if value == '^' {
				xVel = 0
				yVel = -1
				xStart = mapX
				yStart = mapY
				_map[mapY] = append(_map[mapY], '.')
			} else {
				_map[mapY] = append(_map[mapY], value)
			}
		}
		mapY += 1
	}

	// fmt.Println("starting at", x+1, "x", y+1)

	total := 0
	for y := range len(_map) {
		for x := range len(_map[y]) {
			if _map[y][x] == '#' {
				continue
			}

			testMap := copyMap(_map)
			testMap[y][x] = '#'

			_, loop := walk(testMap, xStart, yStart, xVel, yVel)
			if loop {
				total += 1
			}
		}
	}

	fmt.Println("total", total)
}

func copyMap(_map map[int][]rune) map[int][]rune {
	copy := make(map[int][]rune)

	for y := range _map {
		for _, value := range _map[y] {
			copy[y] = append(copy[y], value)
		}
	}

	return copy
}

func walk(_map map[int][]rune, xStart int, yStart int, xVel int, yVel int) (int, bool) {
	x := xStart
	y := yStart
	xMax := len(_map[0])
	yMax := len(_map)
	total := 1
	totalMax := xMax * yMax
	loop := false

	for {
		if total > totalMax {
			loop = true
			break
		}
		if x+xVel < 0 || x+xVel >= xMax || y+yVel < 0 || y+yVel >= yMax {
			// fmt.Println("stopping")
			break
		}

		// fmt.Println("next", string(_map[y+yVel][x+xVel]))

		if _map[y+yVel][x+xVel] == '#' {
			turnRight(&xVel, &yVel)
			// fmt.Println("turned right at", x+1, "x", y+1)
			// fmt.Println("xVel", xVel, "yVel", yVel)

			continue
		}

		_map[y][x] = 'X'

		x = x + xVel
		y = y + yVel

		total += 1

		// fmt.Println("moved to", x+1, "x", y+1)
	}

	return total, loop
}

func turnRight(x *int, y *int) {
	up := *y == -1 && *x == 0
	right := *x == 1 && *y == 0
	down := *y == 1 && *x == 0
	// left := *x == -1 && *y == 0

	if up {
		*y = 0
		*x = 1
	} else if down {
		*y = 0
		*x = -1
	} else if right {
		*y = 1
		*x = 0
	} else {
		// left
		*y = -1
		*x = 0
	}
}
