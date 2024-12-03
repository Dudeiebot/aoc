package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dudeiebot/aoc/lib"
)

func main() {
	input := lib.OpenFile("output.txt")
	sum := 0
	otherSum := 0
	for _, each := range input {
		// each := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
		sum += findValidMuls(each)
	}
	enabled := true
	for _, each := range input {
		// each := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
		otherSum += findValidMulsDoDonts(each, &enabled)
	}
	fmt.Println("Answer Part 1:", sum)
	fmt.Println("Answer Part 2:", otherSum)
}

func findValidMuls(input string) int {
	sum := 0
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		if len(match) == 3 {
			num1, err1 := strconv.Atoi(match[1])
			num2, err2 := strconv.Atoi(match[2])

			if err1 == nil && err2 == nil {
				sum += num1 * num2
			}
		}
	}
	return sum
}

func findValidMulsDoDonts(input string, enabled *bool) int {
	sum := 0

	r := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	matches := r.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		switch {
		case strings.Contains(match[0], "don't"):
			*enabled = false
		case strings.Contains(match[0], "do("):
			*enabled = true
		case *enabled && strings.HasPrefix(match[0], "mul"):
			if match[1] != "" && match[2] != "" {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				sum += num1 * num2
			}
		}
	}

	return sum
}
