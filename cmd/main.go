package main

import (
	"fmt"
	"os"

	"github.com/nprimo/aoc_2024/inputs"
	"github.com/nprimo/aoc_2024/internal/day1"
	"github.com/nprimo/aoc_2024/internal/day2"
)

type solutions func(string) int

var AocSolutions = map[string][]solutions{
	"day1": {day1.Solve1, day1.Solve2},
	"day2": {day2.Solve1, day2.Solve2},
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	day := os.Args[1]
	daySolutions, ok := AocSolutions[day]
	if !ok {
		fmt.Println("no soutions for", day)
		os.Exit(1)
	}
	input := fetchDayInput(day)
	fmt.Println("== solution for ", day)
	for i, sol := range daySolutions {
		fmt.Printf("part %d: %d\n", i+1, sol(input))
	}
}

func fetchDayInput(day string) string {
	input, err := inputs.F.ReadFile(fmt.Sprintf("%s.txt", day))
	if err != nil {
		panic(err)
	}
	return string(input)
}
