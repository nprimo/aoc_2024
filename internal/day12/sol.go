package day12

import (
	"strings"
)

func Solve1(input string) int {
	input = strings.TrimRight(input, "\n")

	// get matrix
	rows := strings.Split(input, "\n")
	cells := make([][]string, len(rows))

	for r, row := range rows {
		values := strings.Split(row, "")
		cells[r] = make([]string, len(values))
		for c, val := range values {
			cells[r][c] = val
		}
	}

	matrix := Matrix{
		cells: cells,
		maxX:  len(cells),
		maxY:  len(cells[0]),
	}

	plots := []Plot{}
	for r := range matrix.maxX {
		for c := range matrix.maxY {
			curr := Pos{X: c, Y: r}
			if matrix.CurrValue(curr) == "." {
				continue
			}
			seen := make(map[Pos]any)
			Walk(curr, matrix, matrix.CurrValue(curr), seen)
			plots = append(plots, NewPlot(seen))
		}
	}

	sum := 0
	for _, p := range plots {
		sum += p.Area * p.Perimeter
	}

	return sum
}

func CalcPerimeter(cells []Pos) int {
	var perimeter int
	for _, cell := range cells {
		perimeter += 4 - CountNeighbours(cell, cells)
	}
	return perimeter
}

func CountNeighbours(cell Pos, cells []Pos) int {
	count := 0
	for _, d := range directions {
		neighbour := cell.Move(d)
		for _, c := range cells {
			if c == neighbour {
				count++
			}
		}
	}
	return count
}

func NewPlot(seen map[Pos]any) Plot {
	cells := make([]Pos, len(seen))
	count := 0
	for i := range seen {
		cells[count] = i
		count++
	}
	return Plot{
		Area:      count,
		Perimeter: CalcPerimeter(cells),
		cells:     cells,
	}
}

type Matrix struct {
	cells [][]string
	maxX  int
	maxY  int
}

func (m Matrix) CurrValue(p Pos) string {
	return m.cells[p.Y][p.X]
}

type Pos struct {
	X int
	Y int
}

func (p Pos) Move(dir Pos) Pos {
	return Pos{
		X: p.X + dir.X,
		Y: p.Y + dir.Y,
	}
}

var directions = []Pos{
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
}

func (p Pos) IsValidPos(m Matrix) bool {
	if p.X < 0 || p.X >= m.maxX {
		return false
	}
	if p.Y < 0 || p.Y >= m.maxY {
		return false
	}
	return true
}

type Plot struct {
	cells     []Pos
	Area      int
	Perimeter int
}

func Walk(curr Pos, m Matrix, val string, seen map[Pos]any) {
	if !curr.IsValidPos(m) || m.CurrValue(curr) != val {
		return
	}
	if _, ok := seen[curr]; ok {
		return
	}

	seen[curr] = true
	m.cells[curr.Y][curr.X] = "."
	for _, dir := range directions {
		Walk(curr.Move(dir), m, val, seen)
	}
}
