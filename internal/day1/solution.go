package day1

import (
	"sort"
	"strconv"
	"strings"
)

func Solve1(input string) int {
	list1, list2 := parseInput(input)
	sort.Ints(list1)
	sort.Ints(list2)
	sum := 0
	for i := 0; i < len(list1); i++ {
		delta := list1[i] - list2[i]
		if delta < 0 {
			delta *= -1
		}
		sum += delta
	}
	return sum
}

func Solve2(input string) int {
	LList, RList := parseInput(input)
	sum := 0
	for _, n := range LList {
		rep := repetition(n, RList)
		sum += n * rep
	}
	return sum
}

func repetition(n int, list []int) int {
	res := 0
	for i := range list {
		if list[i] == n {
			res++
		}
	}
	return res
}

func parseInput(in string) ([]int, []int) {
	rows := strings.Split(in, "\n")
	list1 := make([]int, len(rows))
	list2 := make([]int, len(rows))
	for i, r := range rows {
		ids := strings.Split(r, " ")
		id1, _ := strconv.Atoi(ids[0])
		id2, _ := strconv.Atoi(ids[len(ids)-1])
		list1[i] = id1
		list2[i] = id2
	}
	return list1, list2
}
