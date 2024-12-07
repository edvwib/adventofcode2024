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
	x := 0
	y := 0
	xVel := 0
	yVel := 0

	mapY := 0
	for scanner.Scan() {
		line := scanner.Text()
		for mapX, value := range line {
			if value == '^' {
				xVel = 0
				yVel = -1
				x = mapX
				y = mapY
				_map[mapY] = append(_map[mapY], '.')
			} else {
				_map[mapY] = append(_map[mapY], value)
			}
		}
		mapY += 1
	}

	xMax := len(_map[0])
	yMax := len(_map)

	total := 1
	for {
		if x+xVel < 0 || x+xVel >= xMax || y+yVel < 0 || y+yVel >= yMax {
			break
		}

		if _map[y+yVel][x+xVel] == '#' {
			turnRight(&xVel, &yVel)
			continue
		}

		_map[y][x] = 'X'

		x = x + xVel
		y = y + yVel

		if _map[y][x] == '.' {
			total += 1
		}
	}

	fmt.Println("total", total)
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
