// https://adventofcode.com/2025/day/1
// PART 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calibrate(originalPosition, turns int) (result int) {
	result = (originalPosition + turns) % 100
	if result < 0 {
		result += 100
	}
	return result
}

func calculate(currentLine string, currentPosition int) int {
	rotationNum, _ := strconv.Atoi(currentLine[1:])

	switch currentLine[0] {
	case 'R':
		return calibrate(currentPosition, rotationNum)
	case 'L':
		return calibrate(currentPosition, -rotationNum)
	}

	return currentPosition
}

func main() {

	file, err := os.Open("./2025/day_one/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// init scanner object to read file
	scanner := bufio.NewScanner(file)

	// split each scanned line into words
	scanner.Split(bufio.ScanWords)

	const startPosition = 50
	var currentPosition = startPosition
	var zeroes int

	for scanner.Scan() {
		currentLine := scanner.Text()
		nextPos := calculate(currentLine, currentPosition)

		if nextPos == 0 {
			zeroes++
			fmt.Printf("Found zero.\t\t zeroes encountered:  %d\n", zeroes)
		}
		currentPosition = nextPos
	}

	fmt.Println(zeroes)

}
