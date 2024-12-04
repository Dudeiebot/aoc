package main

import (
	"fmt"

	"github.com/dudeiebot/aoc/lib"
)

func main() {
	input := lib.OpenFile("output.txt")

	res := countXMAS(input)
	ans := countXMAS2(input)
	fmt.Println("countXMAS: ", res)
	fmt.Println("countXMAS 2: ", ans)
}

func countXMAS(input []string) int {
	count := 0
	for i := 0; i < len(input); i++ {
		for j, char := range input[i] {
			if char == 'X' {
				count += checkIndex(input, i, j)
			}
		}
	}
	return count
}

func checkIndex(input []string, i, j int) int {
	count := 0
	directions := []struct {
		di, dj int
	}{
		{0, -1},  // horizontal backwards
		{0, 1},   // horizontal forward
		{-1, 0},  // vertical upwards
		{1, 0},   // vertical downwards
		{-1, 1},  // diagonal upwards right
		{-1, -1}, // diagonal upwards left
		{1, 1},   // diagonal downwards right
		{1, -1},  // diagonal downwards left
	}

	target := []byte{'X', 'M', 'A', 'S'}

	// Check in all 8 directions
	for _, dir := range directions {
		valid := true
		for k := 0; k < len(target); k++ {
			ni, nj := i+k*dir.di, j+k*dir.dj
			if ni < 0 || ni >= len(input) || nj < 0 || nj >= len(input[0]) ||
				input[ni][nj] != target[k] {
				valid = false
				break
			}
		}
		if valid {
			count++
		}
	}
	return count
}

func countXMAS2(input []string) int {
	count := 0
	for i := 1; i < len(input)-1; i++ {
		for j, char := range input[i] {
			if char == 'A' {
				count += checkIndex2(input, i, j)
			}
		}
	}
	return count
}

func checkIndex2(input []string, i, j int) int {
	// Early boundary check
	if j < 1 || j > len(input[i])-2 {
		return 0
	}

	patterns := [4][4]byte{
		{byte('M'), byte('M'), byte('S'), byte('S')}, // M.M / .A. / S.S
		{byte('S'), byte('S'), byte('M'), byte('M')}, // S.S / .A. / M.M
		{byte('S'), byte('M'), byte('M'), byte('S')}, // S.M / .A. / S.M
		{byte('M'), byte('S'), byte('S'), byte('M')}, // M.S / .A. / M.S
	}

	count := 0
	for _, pattern := range patterns {
		if input[i-1][j+1] == pattern[0] &&
			input[i-1][j-1] == pattern[1] &&
			input[i+1][j-1] == pattern[2] &&
			input[i+1][j+1] == pattern[3] {
			count++
		}
	}
	return count
}
