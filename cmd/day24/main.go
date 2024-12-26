package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/nprimo/aoc_2024/inputs"
)

func main() {
	input, err := inputs.F.ReadFile("day24.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(
		Solve1(string(input)),
	)
}

func Solve1(input string) int {
	input = strings.TrimRight(input, "\n")

	parts := strings.Split(input, "\n\n")

	wireMapping := make(map[string]bool)
	for _, row := range strings.Split(parts[0], "\n") {
		data := strings.Split(row, ": ")
		wireMapping[data[0]] = data[1] == "1"
	}

	rows := strings.Split(parts[1], "\n")
	var res []Z
	for len(rows) > 0 {
		row := rows[0]
		rows = rows[1:]
		data := strings.Split(row, " ")
		arg1, ok := wireMapping[data[0]]
		if !ok {
			rows = append(rows, row)
			continue
		}
		arg2, ok := wireMapping[data[2]]
		if !ok {
			rows = append(rows, row)
			continue
		}
		arg3 := data[4]
		var opRes bool
		switch data[1] {
		case "AND":
			opRes = arg1 && arg2
			break
		case "OR":
			opRes = arg1 || arg2
			break
		case "XOR":
			opRes = arg1 != arg2
			break
		}
		wireMapping[arg3] = opRes
		if arg3[:1] == "z" {
			pos, err := strconv.Atoi(arg3[1:])
			if err != nil {
				panic(err)
			}
			res = append(res, Z{opRes, pos})
		}
	}

	sort.Sort(ZList(res))
	var b strings.Builder
	for i := len(res) - 1; i >= 0; i-- {
		if res[i].val {
			b.WriteString("1")

		} else {
			b.WriteString("0")
		}
	}
	result, err := strconv.ParseInt(b.String(), 2, 64)
	if err != nil {
		panic(err)
	}
	return int(result)
}

type Z struct {
	val bool
	pos int
}

type ZList []Z

func (s ZList) Len() int           { return len(s) }
func (s ZList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ZList) Less(i, j int) bool { return s[i].pos < s[j].pos }
