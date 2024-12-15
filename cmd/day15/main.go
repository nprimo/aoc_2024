package main

import (
	"fmt"
	"strings"

	"github.com/nprimo/aoc_2024/inputs"
)

func main() {
	input, err := inputs.F.ReadFile(fmt.Sprintf("%s.txt", "day15"))
	if err != nil {
		panic(err)
	}
	fmt.Println(Solve1(string(input)))
}

func Solve1(input string) int {
	input = strings.TrimRight(input, "\n")

	chunks := strings.Split(input, "\n\n")
	rawMatrix := chunks[0]
	instructions := strings.ReplaceAll(chunks[1], "\n", "")

	matrix := parseMatrix(rawMatrix)
	matrix = ExecuteInstructions(strings.Split(instructions, ""), matrix)

	sum := 0
	for r, row := range matrix.Cells {
		for c, val := range row {
			if val == "O" {
				sum += 100*r + c
			}
		}
	}

	return sum
}

type Pos struct {
	X int
	Y int
}

func (p Pos) Add(o Pos) Pos {
	return Pos{
		p.X + o.X,
		p.Y + o.Y,
	}
}

func (p Pos) isValidPos(m Matrix) bool {
	if p.X < 0 || p.X >= len(m.Cells[0]) {
		return false
	}
	if p.Y < 0 || p.Y >= len(m.Cells) {
		return false
	}
	return true
}

type Matrix struct {
	Cells [][]string
	pos   Pos
}

func (m Matrix) ValueAt(p Pos) string {
	return m.Cells[p.Y][p.X]
}

var dir = map[string]Pos{
	"<": {-1, 0},
	">": {1, 0},
	"^": {0, -1},
	"v": {0, 1},
}

func ExecuteInstructions(ins []string, m Matrix) Matrix {
	for _, i := range ins {
		// fmt.Println(i)
		nextPos := m.pos.Add(dir[i])
		if !nextPos.isValidPos(m) || m.ValueAt(nextPos) == "#" {
			continue
		}
		if m.ValueAt(nextPos) == "O" {
			after := nextPos.Add(dir[i])
			if !after.isValidPos(m) || m.ValueAt(nextPos) == "#" {
				continue
			}
			for after.isValidPos(m) &&
				m.ValueAt(after) == "O" {
				after = after.Add(dir[i])
			}
			if m.ValueAt(after) == "#" {
				continue
			}
			m.Cells[after.Y][after.X] = "O"
			m.Cells[nextPos.Y][nextPos.X] = "."
		}
		m.pos = nextPos
		// fmt.Println(nextPos)
		// for _, row := range m.Cells {
		// 	fmt.Println(row)
		// }
	}
	return m
}

func parseMatrix(raw string) Matrix {
	rows := strings.Split(raw, "\n")
	cells := make([][]string, len(rows))
	matrix := Matrix{
		Cells: cells,
	}
	for r, row := range rows {
		cells[r] = make([]string, len(row))
		for c, val := range strings.Split(row, "") {
			if val == "@" {
				matrix.pos.X = c
				matrix.pos.Y = r
				val = "."
			}
			cells[r][c] = val
		}
	}
	return matrix
}
