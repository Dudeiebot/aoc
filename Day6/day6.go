package main

import (
	"fmt"

	"github.com/dudeiebot/aoc/lib"
)

type Position struct {
	x, y int
}

const UP = 0

var directions = [4][2]int{
	{0, -1}, // UP
	{1, 0},  // RIGHT
	{0, 1},  // DOWN
	{-1, 0}, // LEFT
}

func main() {
	input := lib.OpenFile("output.txt")
	grid := getByteGrid(input)
	startPos := findStart(grid)
	pos := getPath(grid, startPos)

	ans := countLoopPositions(grid, startPos, pos)
	fmt.Println("Total Position:", len(pos))
	fmt.Println("Count Loop", ans)
}

func getByteGrid(input []string) [][]byte {
	grid := make([][]byte, len(input))
	for i, line := range input {
		grid[i] = []byte(line)
	}
	return grid
}

func findStart(grid [][]byte) Position {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '^' {
				return Position{x, y}
			}
		}
	}
	return Position{}
}

func getPath(grid [][]byte, pos Position) []Position {
	visited := make(map[Position]bool)
	path := []Position{}
	dir := UP

	for {
		if !visited[pos] {
			visited[pos] = true
			path = append(path, pos)
		}

		nextPos := Position{
			x: pos.x + directions[dir][0],
			y: pos.y + directions[dir][1],
		}

		// Boundary check
		if nextPos.x < 0 || nextPos.x >= len(grid[0]) || nextPos.y < 0 || nextPos.y >= len(grid) {
			return path
		}

		// Obstacle check
		if grid[nextPos.y][nextPos.x] == '#' {
			dir = (dir + 1) % 4 // Rotate clockwise
		} else {
			pos = nextPos
		}
	}
}

func hasLoop(grid [][]byte, start Position) bool {
	visited := make(map[Position]map[int]bool)
	pos := start
	dir := UP

	for {
		// Create a map for this position if it doesn't exist
		if _, exists := visited[pos]; !exists {
			visited[pos] = make(map[int]bool)
		}

		// Check if this direction has been visited before
		if visited[pos][dir] {
			return true
		}

		// Mark this direction as visited
		visited[pos][dir] = true

		// Calculate next position
		nextPos := Position{
			x: pos.x + directions[dir][0],
			y: pos.y + directions[dir][1],
		}

		// Check boundaries
		if nextPos.x < 0 || nextPos.x >= len(grid[0]) || nextPos.y < 0 || nextPos.y >= len(grid) {
			return false
		}

		// Handle obstacle or direction change
		if grid[nextPos.y][nextPos.x] == '#' {
			dir = (dir + 1) % 4 // Rotate clockwise
		} else {
			pos = nextPos
		}
	}
}

func countLoopPositions(grid [][]byte, start Position, initialPath []Position) int {
	count := 0
	for _, pos := range initialPath {
		if grid[pos.y][pos.x] == '.' {
			grid[pos.y][pos.x] = '#'
			if hasLoop(grid, start) {
				count++
			}
			grid[pos.y][pos.x] = '.'
		}
	}
	return count
}
