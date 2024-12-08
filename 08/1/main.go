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

	_map := buildMap(file)

	antinodes := make(map[int]map[int]bool)
	for i := range _map {
		antinodes[i] = make(map[int]bool)
	}

	for y1, row := range _map {
		for x1, antenna1 := range row {
			if antenna1 == '.' {
				continue
			}

			for y2, row := range _map {
				for x2, antenna2 := range row {
					if antenna2 != antenna1 {
						continue
					}
					if y2 == y1 && x2 == x1 {
						continue
					}
					xDistance := (x2 - x1)
					yDistance := (y2 - y1)

					antX := x2 + xDistance
					antY := y2 + yDistance

					if antX >= 0 && antX < len(row) && antY >= 0 && antY < len(_map) {
						antinodes[antY][antX] = true
						continue
					}
				}
			}

		}
	}

	antinodesCount := 0
	for _, row := range antinodes {
		for _, antinode := range row {
			if antinode {
				antinodesCount += 1
			}
		}
	}

	fmt.Println("antinodes", antinodesCount)
}

func buildMap(file *os.File) map[int][]rune {
	scanner := bufio.NewScanner(file)

	_map := make(map[int][]rune)

	y := 0
	for scanner.Scan() {
		for _, col := range scanner.Text() {
			_map[y] = append(_map[y], col)
		}
		y += 1
	}

	return _map
}
