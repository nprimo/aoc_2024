package day10

import (
	"strconv"
	"strings"

	"github.com/nprimo/aoc_2024/internal/day4"
)

func Solve1(input string) int {
	matrix := getMatrix(input)
	xMax := len(matrix[0])
	yMax := len(matrix)
	startingPoints := getStartingPoints(matrix)

	sum := 0
	for _, p := range startingPoints {
		w := Walker{
			Top: make(map[day4.Pos]bool),
		}
		w.walk(p, matrix, xMax, yMax)
		sum += len(w.Top)
	}
	return sum
}

type Walker struct {
	Top map[day4.Pos]bool
}

func (w *Walker) walk(pos day4.Pos, matrix [][]int, xMax, yMax int) {
	currVal := matrix[pos.Y][pos.X]
	if currVal == 9 {
		w.Top[pos] = true
	}
	for _, dir := range directions {
		newPos := day4.Pos{
			X: pos.X + dir.X,
			Y: pos.Y + dir.Y,
		}
		if !IsValidPos(newPos, xMax, yMax) {
			continue
		}
		newVal := matrix[newPos.Y][newPos.X]
		if (newVal - currVal) == 1 {
			w.walk(newPos, matrix, xMax, yMax)
		}
	}
}

var directions = []day4.Pos{
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: -1, Y: 0},
}

func IsValidPos(pos day4.Pos, xMax, yMax int) bool {
	if pos.X < 0 || pos.X >= xMax {
		return false
	}
	if pos.Y < 0 || pos.Y >= yMax {
		return false
	}
	return true
}

func getMatrix(input string) [][]int {
	input = strings.TrimRight(input, "\n")
	rows := strings.Split(input, "\n")
	matrix := make([][]int, len(rows))
	for r, row := range rows {
		values := strings.Split(row, "")
		matrix[r] = make([]int, len(values))
		for c, val := range values {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			matrix[r][c] = num
		}
	}
	return matrix
}

func getStartingPoints(matrix [][]int) []day4.Pos {
	var res []day4.Pos
	for r, rows := range matrix {
		for c, val := range rows {
			if val == 0 {
				res = append(res, day4.Pos{X: c, Y: r})
			}
		}
	}
	return res
}
