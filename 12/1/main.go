package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type pos struct {
	x int
	y int
}

type plot struct {
	value rune
	pos   pos
	r     *region
}

type region struct {
	plots     []*plot
	perimiter int
	area      int
}

func main() {
	file, err := os.Open("../input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	_map := make(map[pos]*plot)

	y := 0
	colCount := 0
	for scanner.Scan() {
		row := scanner.Text()
		if colCount == 0 {
			colCount = len(row)
		}
		for x, col := range row {
			p := pos{x: x, y: y}
			_map[p] = &plot{pos: p, r: nil, value: col}
		}
		y++
	}
	rowCount := y

	for y := 0; y < rowCount; y++ {
		for x := 0; x < colCount; x++ {
			currentPos := pos{x: x, y: y}
			currentPlot := _map[currentPos]

			if currentPlot.r != nil {
				continue
			}

			reg := &region{plots: []*plot{currentPlot}}
			expandRegion(currentPlot, reg, &_map, rowCount, colCount)
		}
	}

	regions := make([]*region, 0)

	for y := 0; y < rowCount; y++ {
		for x := 0; x < rowCount; x++ {
			currentPos := pos{x: x, y: y}
			currentPlot := _map[currentPos]

			currentRegion := currentPlot.r
			if currentRegion == nil {
				panic("aaah")
			}

			for range currentRegion.plots {
				currentRegion.area += 1
				above := pos{x: currentPos.x, y: currentPos.y - 1}
				right := pos{x: currentPos.x + 1, y: currentPos.y}
				below := pos{x: currentPos.x, y: currentPos.y + 1}
				left := pos{x: currentPos.x - 1, y: currentPos.y}

				if pl, exists := _map[above]; (exists && pl.r != currentPlot.r) || !exists {
					currentRegion.perimiter += 1
				}
				if pl, exists := _map[right]; (exists && pl.r != currentPlot.r) || !exists {
					currentRegion.perimiter += 1
				}
				if pl, exists := _map[below]; (exists && pl.r != currentPlot.r) || !exists {
					currentRegion.perimiter += 1
				}
				if pl, exists := _map[left]; (exists && pl.r != currentPlot.r) || !exists {
					currentRegion.perimiter += 1
				}
			}

			if !slices.Contains(regions, currentRegion) {
				regions = append(regions, currentRegion)
			}
		}
	}

	// fmt.Println("count", len(regions))

	total := 0
	for _, region := range regions {
		total += region.area * region.perimiter
	}
	fmt.Println("total", total)
}

func expandRegion(_plot *plot, _region *region, _map *map[pos]*plot, rowCount int, colCount int) {
	_plot.r = _region

	if _plot.pos.y-1 >= 0 {
		// up
		plotAbove := (*_map)[pos{x: _plot.pos.x, y: _plot.pos.y - 1}]
		if plotAbove.r != _plot.r && plotAbove.value == _plot.value {
			expandRegion(plotAbove, _region, _map, rowCount, colCount)
		}
	}

	if _plot.pos.x+1 < colCount {
		// right
		plotRight := (*_map)[pos{x: _plot.pos.x + 1, y: _plot.pos.y}]
		if plotRight.r != _plot.r && plotRight.value == _plot.value {
			expandRegion(plotRight, _region, _map, rowCount, colCount)
		}
	}

	if _plot.pos.y+1 < rowCount {
		// down
		plotBelow := (*_map)[pos{x: _plot.pos.x, y: _plot.pos.y + 1}]
		if plotBelow.r != _plot.r && plotBelow.value == _plot.value {
			expandRegion(plotBelow, _region, _map, rowCount, colCount)
		}
	}

	if _plot.pos.x-1 >= 0 {
		// left
		plotLeft := (*_map)[pos{x: _plot.pos.x - 1, y: _plot.pos.y}]
		if plotLeft.r != _plot.r && plotLeft.value == _plot.value {
			expandRegion(plotLeft, _region, _map, rowCount, colCount)
		}
	}
}
