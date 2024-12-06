package day6

import (
	"strings"

	"github.com/nprimo/aoc_2024/internal/day4"
)

func Solve1(input string) int {
	// matrix
	// seen := make(map[day4.Pos]bool)
	input = strings.TrimRight(input, "\n")
	rows := strings.Split(input, "\n")
	matrix := make([][]string, len(rows))
	for i, row := range rows {
		matrix_r := strings.Split(row, "")
		matrix[i] = matrix_r
	}
	guard := getFirstPosition(matrix)

	walk(guard, matrix)
	return len(guard.Seen)
}

func isValidPos(pos day4.Pos, matrix [][]string) bool {
	if pos.X < 0 || pos.X >= len(matrix[0]) {
		return false
	}
	if pos.Y < 0 || pos.Y >= len(matrix) {
		return false
	}
	return true
}

func currValue(pos day4.Pos, matrix [][]string) string {
	return matrix[pos.Y][pos.X]
}

func walk(guard Guard, matrix [][]string) {
	pos := guard.Pos
	pos.Add(day4.Pos(guard.Dir))
	if !isValidPos(pos, matrix) {
		return
	}
	currVal := currValue(pos, matrix)
	if currVal == "." || strings.Contains(placeHolder, currVal) {
		guard.Pos = pos
		guard.Seen[pos] = true
		walk(guard, matrix)
		return
	}
	if currVal == "#" {
		guard.turnRight()
		walk(guard, matrix)
		return
	}
}

const placeHolder string = "^><v"

type Dir day4.Pos

var (
	Up    = Dir{X: 0, Y: -1}
	Down  = Dir{X: 0, Y: 1}
	Left  = Dir{X: -1, Y: 0}
	Right = Dir{X: 1, Y: 0}
)

type Guard struct {
	Pos   day4.Pos
	Dir   Dir
	Seen  map[day4.Pos]bool
	Steps int
}

func (g *Guard) turnRight() {
	switch g.Dir {
	case Up:
		g.Dir = Right
		return
	case Right:
		g.Dir = Down
		return
	case Down:
		g.Dir = Left
		return
	case Left:
		g.Dir = Up
		return
	}
}

func getFirstPosition(matrix [][]string) Guard {
	for y, row := range matrix {
		for x, val := range row {
			if id := strings.Index(placeHolder, val); id != -1 {
				pos := day4.Pos{X: x, Y: y}
				seen := map[day4.Pos]bool{pos: true}
				switch id {
				case 0:
					return Guard{
						Pos:   pos,
						Dir:   Up,
						Seen:  seen,
						Steps: 0,
					}
				case 1:
					return Guard{
						Pos:   pos,
						Dir:   Right,
						Seen:  seen,
						Steps: 0,
					}
				case 2:
					return Guard{
						Pos:   pos,
						Dir:   Left,
						Seen:  seen,
						Steps: 0,
					}
				case 3:
					return Guard{
						Pos:   pos,
						Dir:   Down,
						Seen:  seen,
						Steps: 0,
					}
				default:
					panic("cannot be here")
				}
			}
		}
	}
	panic("something went wrong in getFirstPosition")
}
