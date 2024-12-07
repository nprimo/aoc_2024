package day7

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve1(input string) int {
	input = strings.TrimRight(input, "\n")
	rows := strings.Split(input, "\n")
	sum := 0
	for _, row := range rows {
		res, nums := parseRow(row)
		possibilities := possibleResults2Ops([]int{nums[0]}, nums[1:])
		if isValidRes(res, possibilities) {
			sum += res
		}
	}
	return sum
}

func possibleResults2Ops(res []int, nums []int) []int {
	if len(nums) == 0 {
		return res
	}
	updated := make([]int, len(res)*2)
	for i, v := range res {
		// 0, 1 -> 2, 3 -> ...
		updated[i*2] = v + nums[0]
		updated[i*2+1] = v * nums[0]
	}
	return possibleResults2Ops(updated, nums[1:])
}

func Solve2(input string) int {
	input = strings.TrimRight(input, "\n")
	rows := strings.Split(input, "\n")
	sum := 0
	for _, row := range rows {
		res, nums := parseRow(row)
		possibilities := possibleResults3Ops([]int{nums[0]}, nums[1:])
		if isValidRes(res, possibilities) {
			sum += res
		}
	}
	return sum
}

func possibleResults3Ops(res []int, nums []int) []int {
	if len(nums) == 0 {
		return res
	}
	updated := make([]int, len(res)*3)
	for i, v := range res {
		// 0, 1, 2 -> 3, 4, 5 -> ...
		updated[i*3] = v + nums[0]
		updated[i*3+1] = v * nums[0]
		merged, err := strconv.Atoi(fmt.Sprintf("%v%v", v, nums[0]))
		if err != nil {
			panic(err)
		}
		updated[i*3+2] = merged
	}
	return possibleResults3Ops(updated, nums[1:])
}

func parseRow(row string) (int, []int) {
	values := strings.Split(row, ": ")
	res, err := strconv.Atoi(values[0])
	if err != nil {
		panic(err)
	}
	toCombine := strings.Split(values[1], " ")
	nums := make([]int, len(toCombine))
	for i, val := range toCombine {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	return res, nums
}

func isValidRes(res int, possibilities []int) bool {
	for _, p := range possibilities {
		if p == res {
			return true
		}
	}
	return false
}
