package day5

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve1(input string) int {
	input = strings.TrimRight(input, "\n")
	parts := strings.Split(input, "\n\n")
	part1 := parts[0]
	part2 := parts[1]
	rules := parseRules(part1)
	updates := parseUpdates(part2)
	sum := 0
	for _, update := range updates {
		if isValidUpdate(update, rules) {
			sum += update[len(update)/2]
		}
	}
	return sum
}

func isValidUpdate(update []int, rules map[int][]int) bool {
	//backward - if any of the previous number is inside rules[num] it means
	// the update is not valid
	// no need to check first number (??)
	for i := len(update) - 1; i > 0; i-- {
		curr := update[i]
		musBeAfter, ok := rules[curr]
		// TODO: double check
		// skip if number not in rules
		if !ok {
			continue
		}
		for j := range i {
			for _, num := range musBeAfter {
				if update[j] == num {
					return false
				}
			}
		}
	}
	return true
}

func parseUpdates(data string) [][]int {
	var res [][]int
	for i, row := range strings.Split(data, "\n") {
		res = append(res, []int{})
		for _, rawNum := range strings.Split(row, ",") {
			num, err := strconv.Atoi(rawNum)
			if err != nil {
				panic(err)
			}
			res[i] = append(res[i], num)
		}
	}
	return res
}

func parseRules(data string) map[int][]int {
	rows := strings.Split(data, "\n")
	res := make(map[int][]int)
	for _, row := range rows {
		values := strings.Split(row, "|")
		if len(values) != 2 {
			panic(fmt.Sprintf("values: %v - with row %s\n", values, row))
		}
		v1, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		v2, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		if _, ok := res[v1]; !ok {
			res[v1] = []int{v2}
		} else {
			res[v1] = append(res[v1], v2)
		}

	}
	return res
}
