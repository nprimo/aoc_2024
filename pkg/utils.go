package pkg

import "strings"

var Dir = []Pos{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

type Pos struct {
	X int
	Y int
}

func (p Pos) In(seen []Pos) bool {
	for _, s := range seen {
		if s == p {
			return true
		}
	}
	return false
}

func (p Pos) Add(o Pos) Pos {
	return Pos{
		X: p.X + o.X,
		Y: p.Y + o.Y,
	}
}

func (p Pos) IsValid(m Matrix) bool {
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
	Start Pos
	End   Pos
}

func (m Matrix) At(p Pos) string {
	// assume p is inside the range
	return m.Cells[p.Y][p.X]
}

func ParseMatrix(input string) Matrix {
	m := Matrix{}
	rows := strings.Split(input, "\n")
	cells := make([][]string, len(rows))
	for r, row := range rows {
		cells[r] = make([]string, len(row))
		for c, val := range strings.Split(row, "") {
			if val == "S" {
				m.Start = Pos{X: c, Y: r}
				val = "."
			}
			if val == "E" {
				m.End = Pos{X: c, Y: r}
				val = "."
			}
			cells[r][c] = val
		}
	}
	m.Cells = cells
	return m
}
