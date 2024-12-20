package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/nprimo/aoc_2024/inputs"
)

func main() {
	input, err := inputs.F.ReadFile(fmt.Sprintf("%s.txt", "day16"))
	if err != nil {
		panic(err)
	}
	fmt.Println(Solve1(string(input)))
}

func Solve1(input string) int {
	input = strings.TrimRight(input, "\n")
	m := ParseMatrix(input)

	fmt.Println("walking...")
	res := Walk(m.start, m, []Pos{}, 0, Pos{0, 0})
	fmt.Println()

	best := math.MaxInt
	for _, r := range res {
		if r < best {
			best = r
		}
	}
	return best
}

var dir = []Pos{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func Walk(curr Pos, m Matrix, seen []Pos, turn int, prevD Pos) []int {
	fmt.Println(len(seen))
	if !curr.IsValid(m) ||
		m.At(curr) == "#" ||
		curr.IsSeen(seen) {
		return []int{}
	}
	if curr == m.end {
		fmt.Print(".")
		sum := len(seen) + turn*1000
		return []int{sum}
	}
	seen = append(seen, curr)
	res := []int{}
	for _, d := range dir {
		nexPos := curr.Add(d)
		deltaTurn := 0
		if d != prevD {
			deltaTurn = 1
		}
		res = append(res, Walk(nexPos, m, seen, turn+deltaTurn, d)...)
	}
	return res

}

type Pos struct {
	X int
	Y int
}

func (p Pos) IsSeen(seen []Pos) bool {
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
	if p.X < 0 || p.X >= len(m.cells[0]) {
		return false
	}
	if p.Y < 0 || p.Y >= len(m.cells) {
		return false
	}
	return true
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
