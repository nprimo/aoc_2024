package day9

import (
	"strconv"
	"strings"
)

func Solve1(input string) int {
	input = strings.TrimRight(input, "\n")
	values := strings.Split(input, "")
	var res []int
	for i, v := range values {
		rep, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			for range rep {
				res = append(res, i/2)
			}
		} else {
			for range rep {
				//INFO: use -1 as "empty"
				res = append(res, -1)
			}
		}

	}
	newOrder := rearrange(res)
	sum := 0
	for i, v := range newOrder {
		if v == -1 {
			break
		}
		sum += i * v
	}
	return sum
}

func rearrange(curr []int) []int {
	freePos := 0
	for ; freePos < len(curr); freePos++ {
		if curr[freePos] == -1 {
			break
		}
	}
	for i := len(curr) - 1; i > len(curr)/2; i-- {
		if freePos > i {
			break
		}
		if curr[i] != -1 {
			//Swap with free pos
			curr[freePos] = curr[i]
			curr[i] = -1
			for ; freePos < len(curr); freePos++ {
				if curr[freePos] == -1 {
					break
				}
			}
		}
	}
	return curr
}
