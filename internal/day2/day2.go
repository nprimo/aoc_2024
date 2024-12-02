package day2

import (
	"strconv"
	"strings"
)

func Solve1(input string) int {
	levels := strings.Split(input, "\n")
	safeLevelCount := 0
	for _, level := range levels {
		parsed := strings.Split(level, " ")
		if isSafeLevel(parsed) {
			safeLevelCount++
		}
	}
	return safeLevelCount
}

func isSafeLevel(values []string) bool {
	if len(values) < 2 {
		return false
	}
	val1, _ := strconv.Atoi(values[0])
	val2, _ := strconv.Atoi(values[1])
	isIncreasing := val2 > val1
	for i := 0; i < len(values)-1; i++ {
		val1, _ := strconv.Atoi(values[i])
		val2, _ := strconv.Atoi(values[i+1])
		dir := val2 > val1
		delta := val2 - val1
		if dir != isIncreasing || delta == 0 || (delta < -3 || delta > 3) {
			return false
		}
	}

	return true
}

// INFO: does not work
func Solve2(input string) int {
	levels := strings.Split(input, "\n")
	safeLevelCount := 0
	for _, level := range levels {
		parsed := strings.Split(level, " ")
		if len(parsed) < 2 {
			continue
		}
		val1, _ := strconv.Atoi(parsed[0])
		val2, _ := strconv.Atoi(parsed[1])
		isIncreasing := val2 > val1
		if isSafeLevel(parsed[1:]) ||
			isSafeLevelWithException(isIncreasing, parsed, 0) {
			safeLevelCount++
		}
	}
	return safeLevelCount
}

// TODO:: improve "flow"
// - if one check fails, remove an item and try again
// - if it did it twice, return
// Need to ensure I can try to remove all the item in the list (including first and last)
func isSafeLevelWithException(isIncreasing bool, values []string, count int) bool {
	if len(values) < 2 || count > 1 {
		return false
	}
	for i := 1; i < len(values); i++ {
		val1, _ := strconv.Atoi(values[i-1])
		val2, _ := strconv.Atoi(values[i])
		dir := val2 > val1
		delta := val2 - val1
		if dir != isIncreasing || delta == 0 || (delta < -3 || delta > 3) {
			//remove current "val2" and try again
			modifiedLevel := append(values[:i], values[i+1:]...)
			return isSafeLevelWithException(isIncreasing, modifiedLevel, count+1)
		}
	}
	return true
}
