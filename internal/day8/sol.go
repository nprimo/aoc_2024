package day8

import (
	"strings"

	"github.com/nprimo/aoc_2024/internal/day4"
)

func Solve2(input string) int {
	antennas, maxX, maxY := getAntennasAndRange(input)
	antinodes := make(map[day4.Pos]bool)
	for _, values := range antennas {
		possibleComb := getCombinationsOf2(len(values))
		for _, comb := range possibleComb {
			p1 := values[comb[0]]
			p2 := values[comb[1]]
			currAntL := getAntinodesL(p1, p2, maxX, maxY)
			currAntR := getAntinodesR(p1, p2, maxX, maxY)
			for _, a := range currAntL {
				if _, ok := antinodes[a]; !ok {
					antinodes[a] = true
				}
			}
			for _, a := range currAntR {
				if _, ok := antinodes[a]; !ok {
					antinodes[a] = true
				}
			}
		}
	}
	return len(antinodes)
}

func Solve1(input string) int {
	// parse
	antennas, maxX, maxY := getAntennasAndRange(input)

	// for every combination of points
	antinodes := make(map[day4.Pos]bool)
	for _, values := range antennas {
		possibleComb := getCombinationsOf2(len(values))
		for _, comb := range possibleComb {
			currAnt := getAntinodes(values[comb[0]], values[comb[1]], maxX, maxY)
			for _, a := range currAnt {
				if _, ok := antinodes[a]; !ok {
					antinodes[a] = true
				}
			}
		}
	}
	return len(antinodes)
}

func getCombinationsOf2(a int) [][2]int {
	var res [][2]int
	for i := 0; i < (a - 1); i++ {
		for j := i + 1; j < a; j++ {
			res = append(res, [2]int{i, j})
		}
	}
	return res
}

func getAntinodesL(p1, p2 day4.Pos, maxX, maxY int) []day4.Pos {
	var res []day4.Pos
	deltas := getDeltas(p1, p2)
	for i := 0; true; i++ {
		aNode := day4.Pos{
			X: p1.X + i*deltas.X,
			Y: p1.Y + i*deltas.Y,
		}
		if !isValidPos(aNode, maxX, maxY) {
			break
		}
		res = append(res, aNode)
	}
	return res
}

func getAntinodesR(p1, p2 day4.Pos, maxX, maxY int) []day4.Pos {
	var res []day4.Pos
	deltas := getDeltas(p1, p2)
	for i := 0; true; i++ {
		aNode := day4.Pos{
			X: p2.X - i*deltas.X,
			Y: p2.Y - i*deltas.Y,
		}
		if !isValidPos(aNode, maxX, maxY) {
			break
		}
		res = append(res, aNode)
	}
	return res
}

func getAntinodes(p1, p2 day4.Pos, maxX, maxY int) []day4.Pos {
	deltas := getDeltas(p1, p2)
	antinode1 := day4.Pos{
		X: p1.X + deltas.X,
		Y: p1.Y + deltas.Y,
	}
	antinode2 := day4.Pos{
		X: p2.X - deltas.X,
		Y: p2.Y - deltas.Y,
	}
	var res []day4.Pos
	if isValidPos(antinode1, maxX, maxY) {
		res = append(res, antinode1)
	}
	if isValidPos(antinode2, maxX, maxY) {
		res = append(res, antinode2)
	}
	return res
}

func isValidPos(p day4.Pos, maxX, maxY int) bool {
	if p.X < 0 || p.X > maxX {
		return false
	}
	if p.Y < 0 || p.Y > maxY {
		return false
	}
	return true
}

func getDeltas(p1 day4.Pos, p2 day4.Pos) day4.Pos {
	deltaX := (p1.X - p2.X)
	deltaY := (p1.Y - p2.Y)
	return day4.Pos{X: deltaX, Y: deltaY}
}

func getAntennasAndRange(input string) (map[string][]day4.Pos, int, int) {
	input = strings.TrimRight(input, "\n")
	rows := strings.Split(input, "\n")
	antennas := make(map[string][]day4.Pos)
	var maxX, maxY int
	for r, row := range rows {
		matrix_r := strings.Split(row, "")
		for c, val := range matrix_r {
			if val != "." {
				pos := day4.Pos{X: c, Y: r}
				curr, ok := antennas[val]
				if !ok {
					antennas[val] = []day4.Pos{pos}
				} else {
					antennas[val] = append(curr, pos)
				}
			}
			maxX = c
		}
		maxY = r
	}
	return antennas, maxX, maxY
}
