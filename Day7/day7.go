package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dudeiebot/aoc/lib"
)

func main() {
	input := lib.OpenFile("output.txt")
	ans, ans2 := getAns(input)

	fmt.Println(ans, ans2)
}

func getAns(input []string) (int, int) {
	var ans int
	var ans2 int

	for _, line := range input {
		parts := strings.Split(line, ":")
		expectedSum, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		numStrings := strings.Fields(parts[1])
		var nums []int
		for _, numStr := range numStrings {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		if isCorrectMatch(expectedSum, 0, nums, false) {
			ans += expectedSum
		}

		if isCorrectMatch(expectedSum, 0, nums, true) {
			ans2 += expectedSum
		}

	}

	return ans, ans2
}

func calculate(a, b int, operation byte) int {
	calculation := 0

	switch operation {
	case '+':
		calculation = a + b
	case '*':
		calculation = a * b
	case '|':
		mul := 10
		for b/mul > 0 {
			mul *= 10
		}
		return (a * mul) + b
	default:
		return 0
	}
	return calculation
}

func isCorrectMatch(expectedSum, sum int, input []int, allowConcat bool) bool {
	if len(input) == 0 {
		return sum == expectedSum
	}
	if sum > expectedSum {
		return false
	}

	// Recursive check
	return (allowConcat &&
		isCorrectMatch(
			expectedSum,
			calculate(sum, input[0], '|'),
			input[1:],
			allowConcat,
		)) || isCorrectMatch(expectedSum, calculate(sum, input[0], '+'), input[1:], allowConcat) ||
		isCorrectMatch(expectedSum, calculate(sum, input[0], '*'), input[1:], allowConcat)
}
