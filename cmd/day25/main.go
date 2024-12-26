package main

import (
	"fmt"
	"strings"

	"github.com/nprimo/aoc_2024/inputs"
)

func main() {
	input, err := inputs.F.ReadFile("day25.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(Solve1(string(input)))
}

func Solve1(input string) int {
	input = strings.TrimRight(input, "\n")

	keysAndLocks := strings.Split(input, "\n\n")
	var locks []KeyLock
	var keys []KeyLock
	for _, kl := range keysAndLocks {
		parsed, isKey := ParseKeyLock(kl)
		if isKey {
			keys = append(keys, parsed)
		} else {
			locks = append(locks, parsed)
		}
	}
	count := 0
	for _, l := range locks {
		for _, k := range keys {
			if keyLockMatch(l, k) {
				count++

			}
		}
	}

	return count
}
func keyLockMatch(key, lock [5]int) bool {
	for i := range 5 {
		if key[i]+lock[i] > 5 {
			return false
		}
	}
	return true
}

func ParseKeyLock(data string) ([5]int, bool) {
	var res [5]int
	var isKey bool
	rows := strings.Split(data, "\n")
	if rows[0][:1] == "." {
		isKey = true
	}
	for _, row := range rows[1 : len(rows)-1] {
		for c, val := range strings.Split(row, "") {
			if val == "#" {
				res[c] += 1
			}
		}
	}
	return res, isKey
}

type KeyLock [5]int
