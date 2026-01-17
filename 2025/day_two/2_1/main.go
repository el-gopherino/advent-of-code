// https://adventofcode.com/2025/day/2
// part 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./2025/day_two/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	parseContent(extractString(file))
}

func extractString(file *os.File) (content string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = scanner.Text()
	}
	return content
}

func parseContent(content string) (sum int) {
	for _, part := range strings.Split(content, ",") {
		nums := strings.Split(part, "-")

		firstID, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		secondID, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		sum += sumInvalidInRange(firstID, secondID)
	}
	fmt.Println(sum)
	return sum
}

func sumInvalidInRange(start, end int) (sum int) {
	for n := start; n <= end; n++ {
		if isInvalid(n) {
			sum += n
		}
	}
	return sum
}

func isInvalid(n int) bool {
	s := strconv.Itoa(n)

	if len(s)%2 != 0 {
		return false
	}

	mid := len(s) / 2
	return s[:mid] == s[mid:]
}
