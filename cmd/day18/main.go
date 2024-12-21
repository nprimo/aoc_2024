package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nprimo/aoc_2024/inputs"
	"github.com/nprimo/aoc_2024/pkg"
)

func main() {
	data, err := inputs.F.ReadFile("day18.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(
		Solve1(string(data)),
	)
}

const (
	dim   = 70
	bytes = 1024
)

func Solve1(input string) int {
	input = strings.TrimRight(input, "\n")

	rows := strings.Split(input, "\n")
	corr := make([]pkg.Pos, len(rows))
	for i, r := range rows {
		coord := strings.Split(r, ",")
		x, err := strconv.Atoi(coord[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(coord[1])
		if err != nil {
			panic(err)
		}
		corr[i] = pkg.Pos{X: x, Y: y}

	}
	m := pkg.Matrix{
		Start: pkg.Pos{X: 0, Y: 0},
		End:   pkg.Pos{X: dim, Y: dim},
	}
	cells := make([][]string, dim+1)
	for i := range cells {
		cells[i] = make([]string, dim+1)
	}
	for i := 0; i < bytes; i++ {
		pos := corr[i]
		cells[pos.Y][pos.X] = "#"
	}

	m.Cells = cells

	return Walk(m.Start, m)
}

type State struct {
	Pos  pkg.Pos
	Seen []pkg.Pos
}

func Walk(curr pkg.Pos, m pkg.Matrix) int {
	states := []State{{curr, []pkg.Pos{}}}
	visited := make(map[pkg.Pos]bool)
	for len(states) != 0 {
		c := states[0]
		states = states[1:]
		if c.Pos == m.End {
			return len(c.Seen)
		}
		if !c.Pos.IsValid(m) {
			continue
		}
		if visited[c.Pos] {
			continue
		}
		if m.At(c.Pos) == "#" {
			continue
		}

		visited[c.Pos] = true
		c.Seen = append(c.Seen, c.Pos)

		for _, d := range pkg.Dir {
			newPos := c.Pos.Add(d)
			states = append(states, State{newPos, c.Seen})
		}
	}
	return -1
}
