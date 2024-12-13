package day13

import (
	"regexp"
	"strconv"
	"strings"
)

const extra = 10000000000000

func Solve1(input string) int {
	input = strings.TrimRight(input, "\n")
	instructions := strings.Split(input, "\n\n")
	sum := int64(0)
	for _, instruction := range instructions {
		com := GetButtonComb(instruction)
		sum += com[0]*3 + com[1]*1
	}
	return int(sum)
}

func Solve2(input string) int {
	input = strings.TrimRight(input, "\n")
	instructions := strings.Split(input, "\n\n")
	sum := int64(0)
	for _, instruction := range instructions {
		com := GetButtonComb2(instruction)
		sum += com[0]*3 + com[1]*1
	}
	return int(sum)
}

type Button struct {
	X int64
	Y int64
}

func NewButton(input [2]int64) Button {
	return Button{X: input[0], Y: input[1]}
}

type Target struct {
	X int64
	Y int64
}

func NewTarget(in [2]int64) Target {
	return Target{in[0], in[1]}
}

func NewTargetWithExtra(in [2]int64) Target {
	return Target{in[0] + extra, in[1] + extra}
}

func GetButtonComb(instructions string) [2]int64 {
	// parse instructions
	rows := strings.Split(instructions, "\n")
	a := NewButton(parseRow(rows[0]))
	b := NewButton(parseRow(rows[1]))
	t := NewTarget(parseRow(rows[2]))
	na := (t.X*b.Y - b.X*t.Y) / (a.X*b.Y - a.Y*b.X)
	nb := (t.Y - na*a.Y) / b.Y
	if na >= 100 || nb >= 100 {
		return [2]int64{}
	}
	// Check if it match
	if t.X != na*a.X+nb*b.X || t.Y != na*a.Y+nb*b.Y {
		return [2]int64{}
	}
	return [2]int64{na, nb}
}

func GetButtonComb2(instructions string) [2]int64 {
	// parse instructions
	rows := strings.Split(instructions, "\n")
	a := NewButton(parseRow(rows[0]))
	b := NewButton(parseRow(rows[1]))
	t := NewTargetWithExtra(parseRow(rows[2]))
	na := (t.X*b.Y - b.X*t.Y) / (a.X*b.Y - a.Y*b.X)
	nb := (t.Y - na*a.Y) / b.Y
	// Check if it match
	if t.X != na*a.X+nb*b.X || t.Y != na*a.Y+nb*b.Y {
		return [2]int64{}
	}
	return [2]int64{na, nb}
}
func parseRow(row string) [2]int64 {
	var re = regexp.MustCompile(`(?m)(\d+)`)

	res := [2]int64{}
	for i, match := range re.FindAllString(row, -1) {
		var err error
		res[i], err = strconv.ParseInt(match, 10, 64)
		if err != nil {
			panic(err)
		}
	}
	return res
}
