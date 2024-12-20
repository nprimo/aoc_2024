package main

import (
	"fmt"
	"strings"

	"github.com/nprimo/aoc_2024/inputs"
)

func main() {
	input, err := inputs.F.ReadFile(fmt.Sprintf("%s.txt", "day20"))
	if err != nil {
		panic(err)
	}
	fmt.Println(Solve1(string(input)))
}

func Solve1(input string) int {
	input = strings.TrimRight(input, "\n")
	m := ParseMatrix(input)
	// for _, row := range m.cells {
	// 	fmt.Println(row)
	// }
	// fmt.Println("start at", m.start)
	seenOg := Walk(m.start, m, []Pos{}, 0)
	fmt.Println()
	fmt.Println("enOg", len(seenOg[0]))
	seen := Walk(m.start, m, []Pos{}, 1)
	fmt.Println()
	count := 0
	for _, s := range seen {
		diff := len(seenOg[0]) - len(s)
		if diff >= 100 {
			count++
		}
	}
	return count
}

type Pos struct {
	X int
	Y int
}

func (p Pos) Add(o Pos) Pos {
	return Pos{
		X: p.X + o.X,
		Y: p.Y + o.Y,
	}
}
func (p Pos) IsValid(m Matrix) bool {
	if p.X < 0 || p.X >= len(m.cells[0]) {
		return false
	}
	if p.Y < 0 || p.Y >= len(m.cells) {
		return false
	}
	return true
}

func (p Pos) isSeen(seen []Pos) bool {
	for _, s := range seen {
		if s == p {
			return true
		}
	}
	return false
}

type Matrix struct {
	cells [][]string
	start Pos
	end   Pos
}

func (m Matrix) At(p Pos) string {
	// assume p is inside the range
	return m.cells[p.Y][p.X]
}

func ParseMatrix(input string) Matrix {
	m := Matrix{}
	rows := strings.Split(input, "\n")
	cells := make([][]string, len(rows))
	for r, row := range rows {
		cells[r] = make([]string, len(row))
		for c, val := range strings.Split(row, "") {
			if val == "S" {
				m.start = Pos{X: c, Y: r}
				val = "."
			}
			if val == "E" {
				m.end = Pos{X: c, Y: r}
				val = "."
			}
			cells[r][c] = val
		}
	}
	m.cells = cells
	return m
}

var dir = []Pos{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func Walk(curr Pos, m Matrix, seen []Pos, skipCount int) [][]Pos {
	if curr == m.end {
		fmt.Print(".")
		return [][]Pos{seen}
	}
	if !curr.IsValid(m) || curr.isSeen(seen) {
		return [][]Pos{}
	}
	if m.At(curr) == "#" {
		skipCount = skipCount - 1
		if skipCount == -1 {
			return [][]Pos{}
		}
	}
	seen = append(seen, curr)
	res := [][]Pos{}
	for _, d := range dir {
		newPos := curr.Add(d)
		res = append(res, Walk(newPos, m, seen, skipCount)...)
	}
	return res
}
