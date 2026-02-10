// https://adventofcode.com/2025/day/1
// PART 2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countZeroesRight(position, turns int) int {
	return (position + turns) / 100
}

func countZeroesLeft(position, turns int) int {
	if position > 0 && position <= turns {
		return (turns-position)/100 + 1
	} else if position == 0 {
		return turns / 100
	}
	return 0
}

func rotate(position int, direction byte, turns int) (newPos, zeroes int) {
	switch direction {
	case 'R':
		zeroes = countZeroesRight(position, turns)
		newPos = (position + turns) % 100
	case 'L':
		zeroes = countZeroesLeft(position, turns)
		newPos = ((position-turns)%100 + 100) % 100
	}
	return newPos, zeroes
}

func main() {

	var (
		currentPosition = 50
		totalZeroes     = 0
	)

	file, err := os.Open("./2025/day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0]
		turns, _ := strconv.Atoi(line[1:])

		newPos, zeros := rotate(currentPosition, direction, turns)
		totalZeroes += zeros
		currentPosition = newPos
	}
	fmt.Println(totalZeroes)
}
