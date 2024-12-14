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

type clawMachine struct {
	a button
	b button
	p prize

	smallestNumberOfTokens int
}

type button struct {
	xV int
	yV int
}

type prize struct {
	x int
	y int
}

func main() {
	file, err := os.Open("../input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	machines := make([]*clawMachine, 0)
	index := 0
	state := 1
	for scanner.Scan() {
		if state > 3 {
			state = 1
			index++
			continue
		}
		row := scanner.Text()
		if state == 1 {
			split := strings.Split(row, " ")
			x, err := strconv.Atoi(strings.Split(strings.Split(split[2], ",")[0], "X+")[1])
			check(err)
			y, err := strconv.Atoi(strings.Split(split[3], "Y+")[1])
			check(err)
			m := clawMachine{a: button{xV: x, yV: y}}
			machines = append(machines, &m)
		} else if state == 2 {
			split := strings.Split(row, " ")
			x, err := strconv.Atoi(strings.Split(strings.Split(split[2], ",")[0], "X+")[1])
			check(err)
			y, err := strconv.Atoi(strings.Split(split[3], "Y+")[1])
			check(err)
			machines[index].b = button{xV: x, yV: y}
		} else if state == 3 {
			split := strings.Split(row, " ")
			x, err := strconv.Atoi(strings.Split(strings.Split(split[1], ",")[0], "X=")[1])
			check(err)
			y, err := strconv.Atoi(strings.Split(split[2], "Y=")[1])
			check(err)
			machines[index].p = prize{x: x, y: y}
		}

		state++
	}

	fewestTokens := 0

	for _, m := range machines {
		calculateSmallestNumberOfTokens(m)
		if m.smallestNumberOfTokens > 0 {
			fewestTokens += m.smallestNumberOfTokens
		}
	}

	fmt.Println("fewestTokens", fewestTokens)
}

func calculateSmallestNumberOfTokens(m *clawMachine) {
	m.smallestNumberOfTokens = -1
done:
	for a := 1; a <= 100; a++ {
		for b := 1; b <= 100; b++ {
			x := a*m.a.xV + b*m.b.xV
			y := a*m.a.yV + b*m.b.yV

			if x == m.p.x && y == m.p.y {
				m.smallestNumberOfTokens = 3*a + b
				break done
			}
		}
	}
}
