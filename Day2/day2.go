package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func openFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into string numbers
		strNumbers := strings.Fields(line)

		// Convert string numbers to integers
		var lineNumbers []int
		for _, strNum := range strNumbers {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Println("Error converting number:", err)
				continue
			}
			lineNumbers = append(lineNumbers, num)
		}

		// Add the line of numbers to our lines
		lines = append(lines, lineNumbers)
	}
	return lines
}

func main() {
	lines := openFile("output.txt")

	ans := getSafe(lines)
	res := getSafeWithHelp(lines)
	fmt.Println("Get Safe:", ans)
	fmt.Println("Get Safe:", res)
}

func getSafe(lines [][]int) int {
	safe := 0
	for _, line := range lines {
		if isLineSafe(line) {
			safe++
		}
	}
	return safe
}

func isLineSafe(line []int) bool {
	increasing := 0
	decreasing := 0

	for i := 1; i < len(line); i++ {
		diff := line[i] - line[i-1]

		// Check if diff is within the allowed range
		if abs(diff) > 3 {
			return false
		}

		// Track increasing and decreasing trends
		if diff > 0 {
			if decreasing > 0 {
				return false
			}
			increasing++
		} else if diff < 0 {
			if increasing > 0 {
				return false
			}
			decreasing++
		} else {
			// If the numbers are the same, it's not safe
			return false
		}
	}

	return true
}

// Helper function to get absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getSafeWithHelp(lines [][]int) int {
	safe := 0
	for _, line := range lines {
		if isSafeWithRemoval(line) {
			safe++
		}
	}
	return safe
}

func isSafeWithRemoval(line []int) bool {
	if isLineSafe(line) {
		return true
	}

	for i := 0; i < len(line); i++ {
		if isLineSafe(removeIdx(line, i)) {
			return true
		}
	}

	return false
}

func removeIdx(s []int, i int) []int {
	r := make([]int, 0)
	r = append(r, s[:i]...)
	return append(r, s[i+1:]...)
}
