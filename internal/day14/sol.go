package day14

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	// maxX = 11
	// maxY = 7
	maxX = 101
	maxY = 103
)

func Solve1(input string) int {
	input = strings.TrimRight(input, "\n")
	rows := strings.Split(input, "\n")

	guards := make([]Guard, len(rows))
	for i, row := range rows {
		guards[i] = ParseRow(row)
		guards[i] = guards[i].Move(100, maxX, maxY)
	}
	// debug
	matrix := make([][]int, maxY)
	for y := range maxY {
		matrix[y] = make([]int, maxX)
		for x := range maxX {
			for _, g := range guards {
				if g.pos.X == x && g.pos.Y == y {
					matrix[y][x] += 1
				}
			}
		}
	}
	counts := [4]int{}
	for _, g := range guards {
		if g.pos.X < maxX/2 {
			if g.pos.Y < maxY/2 {
				counts[0]++
			}
			if g.pos.Y > maxY/2 {
				counts[1]++
			}
		} else if g.pos.X > maxX/2 {
			if g.pos.Y < maxY/2 {
				counts[2]++
			}
			if g.pos.Y > maxY/2 {
				counts[3]++
			}
		}
	}
	return counts[0] * counts[1] * counts[2] * counts[3]
}

type Pos struct {
	X int
	Y int
}

type Guard struct {
	pos Pos
	vel Pos
}

func (g Guard) Move(time, maxX, maxY int) Guard {
	newPos := g.pos
	for range time {
		newPos = Pos{
			X: (newPos.X + g.vel.X) % maxX,
			Y: (newPos.Y + g.vel.Y) % maxY,
		}
		if newPos.X < 0 {
			newPos.X += maxX
		}
		if newPos.Y < 0 {
			newPos.Y += maxY
		}
	}

	return Guard{
		vel: g.vel,
		pos: newPos,
	}
}

func ParseRow(row string) Guard {
	re := regexp.MustCompile(`(?m)-*\d+`)
	matches := re.FindAllString(row, -1)

	num := make([]int, len(matches))
	for i, m := range matches {
		var err error
		num[i], err = strconv.Atoi(m)
		if err != nil {
			panic(err)
		}
	}
	return Guard{
		pos: Pos{X: num[0], Y: num[1]},
		vel: Pos{X: num[2], Y: num[3]},
	}
}
