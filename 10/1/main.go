package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type xy struct {
	x int
	y int
}

func main() {
	_map, trailheads := buildMap()

	for trailhead := range trailheads {
		// if trailhead != (xy{x: 2, y: 0}) {
		// 	continue
		// }
		score := 0
		var trailends = make(map[xy]int)
		trailheadScore(trailhead.x, trailhead.y, _map, &score, &trailends)
		trailheads[trailhead] = score
	}

	sum := 0
	for _, trailheadValue := range trailheads {
		sum += trailheadValue
	}

	fmt.Println("sum", sum)
}

func trailheadScore(x int, y int, _map [][]int, score *int, trailends *map[xy]int) {
	xMax := len(_map[0]) - 1
	yMax := len(_map) - 1
	current := _map[y][x]

	if x > 0 {
		// check left
		if _map[y][x-1] == current+1 {
			if current+1 == 9 && (*trailends)[xy{x: x - 1, y: y}] == 0 {
				*score++
				(*trailends)[xy{x: x - 1, y: y}]++
			} else {
				trailheadScore(x-1, y, _map, score, trailends)
			}
		}
	}

	if x < xMax {
		// check right
		if _map[y][x+1] == current+1 {
			if current+1 == 9 && (*trailends)[xy{x: x + 1, y: y}] == 0 {
				*score++
				(*trailends)[xy{x: x + 1, y: y}]++
			} else {
				trailheadScore(x+1, y, _map, score, trailends)
			}
		}
	}

	if y > 0 {
		// check up
		if _map[y-1][x] == current+1 {
			if current+1 == 9 && (*trailends)[xy{x: x, y: y - 1}] == 0 {
				*score++
				(*trailends)[xy{x: x, y: y - 1}]++
			} else {
				trailheadScore(x, y-1, _map, score, trailends)
			}
		}
	}

	if y < yMax {
		// check down
		if _map[y+1][x] == current+1 {
			if current+1 == 9 && (*trailends)[xy{x: x, y: y + 1}] == 0 {
				*score++
				(*trailends)[xy{x: x, y: y + 1}]++
			} else {
				trailheadScore(x, y+1, _map, score, trailends)
			}
		}
	}
}

func buildMap() ([][]int, map[xy]int) {
	file, err := os.Open("../input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var _map [][]int
	var trailheads = make(map[xy]int)

	y := 0
	for scanner.Scan() {
		row := scanner.Text()
		_map = append(_map, make([]int, len(row)))
		for x, col := range row {
			if col == '.' {
				_map[y][x] = -1
				continue
			}
			colValue, err := strconv.Atoi(string(col))
			check(err)

			if colValue == 0 {
				trailheads[xy{x: x, y: y}] = 0
			}

			_map[y][x] = colValue
		}
		y++
	}

	return _map, trailheads
}
