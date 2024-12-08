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

func Solve2(input string) int {
	levels := strings.Split(input, "\n")
	safeLevelCount := 0
	for _, level := range levels {
		parsed := strings.Split(level, " ")
		if len(parsed) < 2 {
			continue
		}
		if isSafeLevel(parsed) {
			safeLevelCount++
			continue
		}
		for i := 0; i < len(parsed); i++ {
			modified := make([]string, len(parsed))
			copy(modified, parsed)
			modified = append(modified[:i], modified[i+1:]...)
			if isSafeLevel(modified) {
				safeLevelCount++
				break
			}
		}
	}
	return safeLevelCount
}

func isSafeLevelWithException(values []string, count int) bool {
	if len(values) < 2 || count > 1 {
		return false
	}
	val1, _ := strconv.Atoi(values[0])
	val2, _ := strconv.Atoi(values[1])
	isIncreasing := val2 > val1
	for i := 1; i < len(values); i++ {
		val1, _ := strconv.Atoi(values[i-1])
		val2, _ := strconv.Atoi(values[i])
		dir := val2 > val1
		delta := val2 - val1
		if dir != isIncreasing || delta == 0 || (delta < -3 || delta > 3) {
			// remove current "val2" and try again
			modifiedLevel := append(values[:i], values[i+1:]...)
			return isSafeLevelWithException(modifiedLevel, count+1)
		}
	}
	return true
}
