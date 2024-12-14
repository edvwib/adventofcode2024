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

type position struct {
	x int
	y int
}

type velocity struct {
	x int
	y int
}

type robot struct {
	pos position
	vel velocity
}

func (bot *robot) move(maxX int, maxY int) {
	nextPosition := position{
		x: bot.pos.x + bot.vel.x,
		y: bot.pos.y + bot.vel.y,
	}

	if nextPosition.x >= maxX {
		nextPosition.x = 0 + (nextPosition.x - maxX)
	} else if nextPosition.x < 0 {
		nextPosition.x = maxX - (nextPosition.x * -1)
	}

	if nextPosition.y >= maxY {
		nextPosition.y = 0 + nextPosition.y - maxY
	} else if nextPosition.y < 0 {
		nextPosition.y = maxY - (nextPosition.y * -1)
	}

	bot.pos = nextPosition
}

func main() {
	file, err := os.Open("../input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	robots := make([]*robot, 0)

	for scanner.Scan() {
		row := scanner.Text()

		split := strings.Split(row, " ")
		p := strings.Split(split[0], "=")[1]
		v := strings.Split(split[1], "=")[1]
		x, err := strconv.Atoi(strings.Split(p, ",")[0])
		check(err)
		y, err := strconv.Atoi(strings.Split(p, ",")[1])
		check(err)
		xV, err := strconv.Atoi(strings.Split(v, ",")[0])
		check(err)
		yV, err := strconv.Atoi(strings.Split(v, ",")[1])
		check(err)

		pos := position{
			x: x,
			y: y,
		}

		robots = append(robots, &robot{
			pos: pos,
			vel: velocity{
				x: xV,
				y: yV,
			},
		})
	}

	maxX := 101
	maxY := 103
	// maxX := 11
	// maxY := 7
	maxSeconds := 100

	for second := 1; second <= maxSeconds; second++ {
		// fmt.Println("second", second)
		for _, bot := range robots {
			bot.move(maxX, maxY)
			// fmt.Println("bot moved to ", bot.pos.x+1, bot.pos.y+1)
		}

		// for y := 0; y < maxY; y++ {
		// 	row := strings.Repeat(".", maxX)
		// 	for x := 0; x < maxX; x++ {
		// 		pos := position{x, y}
		// 		for _, bot := range robots {
		// 			if bot.pos != pos {
		// 				continue
		// 			}
		// 			cur := row[x : x+1]
		// 			v, err := strconv.Atoi(cur)
		// 			if err != nil {
		// 				row = replaceAtIndex(row, '1', x)
		// 			} else {
		// 				row = replaceAtIndex(row, []rune(strconv.Itoa(v + 1))[0], x)
		// 			}
		// 		}
		// 	}
		// 	fmt.Println(row)
		// }
		// fmt.Println()
	}

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	midX := maxX / 2
	midY := maxY / 2

	for _, bot := range robots {
		if bot.pos.x == midX || bot.pos.y == midY {
			continue
		}

		if bot.pos.x < midX && bot.pos.y < midY {
			q1++
			continue
		}
		if bot.pos.x > midX && bot.pos.y < midY {
			q2++
			continue
		}
		if bot.pos.x < midX && bot.pos.y > midY {
			q3++
			continue
		}
		if bot.pos.x > midX && bot.pos.y > midY {
			q4++
			continue
		}
	}

	safetyFactor := q1 * q2 * q3 * q4

	fmt.Println("safety factor", safetyFactor)
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
