package main

import (
	"fmt"
	"os"

	"github.com/nprimo/aoc_2024/inputs"
	"github.com/nprimo/aoc_2024/internal/day1"
	"github.com/nprimo/aoc_2024/internal/day10"
	"github.com/nprimo/aoc_2024/internal/day11"
	"github.com/nprimo/aoc_2024/internal/day12"
	"github.com/nprimo/aoc_2024/internal/day13"
	"github.com/nprimo/aoc_2024/internal/day14"
	"github.com/nprimo/aoc_2024/internal/day2"
	"github.com/nprimo/aoc_2024/internal/day3"
	"github.com/nprimo/aoc_2024/internal/day4"
	"github.com/nprimo/aoc_2024/internal/day5"
	"github.com/nprimo/aoc_2024/internal/day6"
	"github.com/nprimo/aoc_2024/internal/day7"
	"github.com/nprimo/aoc_2024/internal/day8"
	"github.com/nprimo/aoc_2024/internal/day9"
)

type solutions func(string) int

var AocSolutions = map[string][]solutions{
	"day1":  {day1.Solve1, day1.Solve2},
	"day2":  {day2.Solve1, day2.Solve2},
	"day3":  {day3.Solve1, day3.Solve2},
	"day4":  {day4.Solve1},
	"day5":  {day5.Solve1},
	"day6":  {day6.Solve1},
	"day7":  {day7.Solve1, day7.Solve2},
	"day8":  {day8.Solve1, day8.Solve2},
	"day9":  {day9.Solve1},
	"day10": {day10.Solve1},
	"day11": {day11.Solve1, day11.Solve2},
	"day12": {day12.Solve1},
	"day13": {day13.Solve1, day13.Solve2},
	"day14": {day14.Solve1},
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
