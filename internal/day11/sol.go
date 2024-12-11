package day11

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func Solve1(input string) int {
	input = strings.TrimRight(input, "\n")
	stones := strings.Split(input, " ")

	stones = blink(stones, 25)
	return len(stones)
}

func Solve2(input string) int {
	input = strings.TrimRight(input, "\n")
	stones := strings.Split(input, " ")

	for i := range 3 {
		fmt.Println(i)
		var wg sync.WaitGroup
		numWorkers := (i + 1) * 10
		wg.Add(numWorkers)
		ch := make(chan []string, numWorkers)
		for i := range numWorkers {
			delta := len(stones) / numWorkers
			start := i * delta
			end := i*delta + delta
			chunk := stones[start:end]
			if i == numWorkers-1 {
				chunk = stones[start:]
			}
			go func() {
				defer wg.Done()
				ch <- blink(chunk, 25)
			}()
		}
		fmt.Println("waiting..")
		wg.Wait()

		var update []string
		for range numWorkers {
			update = append(update, <-ch...)
		}
		close(ch)

		stones = update
	}

	return len(stones)
}

func processStone(stone string) []string {
	if stone == "0" {
		return []string{"1"}
	}
	if len(stone)%2 == 0 {
		half := len(stone) / 2
		leftStone := strings.TrimLeft(stone[:half], "0")
		rightStone := strings.TrimLeft(stone[half:], "0")
		if len(leftStone) == 0 {
			leftStone = "0"
		}
		if len(rightStone) == 0 {
			rightStone = "0"
		}
		return []string{leftStone, rightStone}
	}
	num, err := strconv.ParseInt(stone, 10, 64)
	if err != nil {
		panic(err)
	}
	return []string{strconv.FormatInt(num*2024, 10)}
}

func blink(stones []string, count int) []string {
	// fmt.Println(count)
	if count == 0 {
		return stones
	}
	var updated []string
	for _, stone := range stones {
		updated = append(updated, processStone(stone)...)
	}
	return blink(updated, count-1)
}
