// https://adventofcode.com/2025/day/3
// pt. 1
package main

import (
	"bufio"
	"fmt"
	"os"
)

// findLargestPair uses suffix max algorithm
func findLargestPair(currentLine string) int {
	best := 0
	suffixMax := int(currentLine[len(currentLine)-1] - '0')

	for i := len(currentLine) - 2; i >= 0; i-- {
		d := int(currentLine[i] - '0')
		val := d*10 + suffixMax
		if val > best {
			best = val
		}
		if d > suffixMax {
			suffixMax = d
		}
	}
	return best
}

func main() {

	file, err := os.Open("./2025/day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	n := 0
	for scanner.Scan() {
		n += findLargestPair(scanner.Text())
	}

	fmt.Println(n)
}
