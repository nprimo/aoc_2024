package day3

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	doRe   = regexp.MustCompile(`(?m)do\(\)`)
	dontRe = regexp.MustCompile(`(?m)don't\(\)`)
	mulRe  = regexp.MustCompile(`(?m)mul\(\d{1,3},\d{1,3}\)`)
)

func Solve1(input string) int {
	sum := 0
	for _, match := range mulRe.FindAllString(input, -1) {
		sum += solveMul(match)

	}
	return sum
}

func Solve2(input string) int {
	mulIndex := mulRe.FindAllStringIndex(input, -1)
	doIndex := doRe.FindAllStringIndex(input, -1)
	doBegins := make([]int, len(doIndex))
	for i, id := range doIndex {
		doBegins[i] = id[0]
	}
	dontIndex := dontRe.FindAllStringIndex(input, -1)
	dontBegins := make([]int, len(dontIndex))
	for i, id := range dontIndex {
		dontBegins[i] = id[0]
	}

	sum := 0
	for _, mulIds := range mulIndex {
		begin := mulIds[0]
		beforeBegin := func(i int) bool { return i < begin }
		filteredDo := filter(doBegins, beforeBegin)
		filteredDont := filter(dontBegins, beforeBegin)

		if IsDo(filteredDo, filteredDont) {
			sum += solveMul(input[mulIds[0]:mulIds[1]])
		}
	}
	return sum
}

func IsDo(dos []int, donts []int) bool {
	if len(dos) >= 0 && len(donts) == 0 {
		return true
	}
	if len(dos) == 0 && len(donts) > 0 {
		return false
	}
	lastDo := dos[len(dos)-1]
	lastDont := donts[len(donts)-1]
	return lastDo > lastDont
}

func filter(l []int, cond func(i int) bool) []int {
	var res []int
	for _, el := range l {
		if cond(el) {
			res = append(res, el)
		}
	}
	return res
}

func solveMul(op string) int {
	numbers := strings.Split(op, ",")
	// INFO: remove mul(
	a := numbers[0][4:]
	// INFO: remove )
	b := numbers[1][:(len(numbers[1]) - 1)]

	aInt, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	bInt, err := strconv.Atoi(b)
	if err != nil {
		panic(err)
	}
	return aInt * bInt
}
