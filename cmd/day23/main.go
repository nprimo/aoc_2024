package main

import (
	"fmt"
	"strings"

	"github.com/nprimo/aoc_2024/inputs"
)

func main() {
	input, err := inputs.F.ReadFile("day23.txt")
	if err != nil {
		panic(err)
	}
	// 2298 is too high
	fmt.Println(
		Solve1(string(input)),
	)

}

func Solve1(input string) int {
	input = strings.TrimRight(input, "\n")
	uniquePc := make(map[string]bool)
	links := make(map[string][]string)
	for _, row := range strings.Split(input, "\n") {
		pcs := strings.Split(row, "-")
		links[pcs[0]] = append(links[pcs[0]], pcs[1])
		links[pcs[1]] = append(links[pcs[1]], pcs[0])
		uniquePc[pcs[0]] = true
		uniquePc[pcs[1]] = true
	}
	uPCs := make([]string, len(uniquePc))
	count := 0
	for k := range uniquePc {
		uPCs[count] = k
		count++
	}
	count = 0
	for i := range uPCs {
		for j := i + 1; j < len(uPCs); j++ {
			for k := j + 1; k < len(uPCs); k++ {
				curr := [3]string{
					uPCs[i], uPCs[j], uPCs[k],
				}
				if isValidSet(curr, links) && hasT(curr) {
					count++
				}
			}
		}
	}
	return count
}

func In(a string, l []string) bool {
	for _, v := range l {
		if a == v {
			return true
		}
	}
	return false
}

func isValidSet(set [3]string, links map[string][]string) bool {
	return In(set[0], links[set[1]]) &&
		In(set[0], links[set[2]]) &&
		In(set[1], links[set[2]])

}

func hasT(set [3]string) bool {
	for _, s := range set {
		if s[:1] == "t" {
			return true
		}
	}
	return false
}
