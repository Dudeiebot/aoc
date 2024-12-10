package main

import (
	"fmt"
	"strconv"

	"github.com/dudeiebot/aoc/lib"
)

type Point struct {
	x, y int
}

func main() {
	var grid [][]int
	input := lib.OpenFile("example.txt")
	for _, line := range input {
		row := make([]int, len(line))
		for j, char := range line {
			num, _ := strconv.Atoi(string(char))
			row[j] = num
		}
		grid = append(grid, row)
	}

	ans, res := findTrails(grid)
	fmt.Println(ans, res)
}

func findTrails(grids [][]int) (int, int) {
	countScore := 0
	countRating := 0
	for j, grid := range grids {
		for i, each := range grid {
			if each == 0 {
				score, rating := findScore(grids, Point{i, j}, make(map[Point]struct{}), 0)
				countScore += len(score)
				countRating += rating
			}
		}
	}
	return countScore, countRating
}

func findScore(
	grid [][]int,
	start Point,
	trailHeads map[Point]struct{},
	count int,
) (map[Point]struct{}, int) {
	if grid[start.y][start.x] == 9 {
		if _, ok := trailHeads[start]; !ok {
			trailHeads[start] = struct{}{}
		}
		return trailHeads, count + 1
	}
	nextSteps := findNext(grid, start)
	if len(nextSteps) == 0 {
		return trailHeads, count
	}

	for _, step := range nextSteps {
		trailHeads, count = findScore(grid, step, trailHeads, count)
	}
	return trailHeads, count
}

func findNext(input [][]int, current Point) []Point {
	validNextSteps := []Point{}
	if current.x > 0 && input[current.y][current.x-1] == input[current.y][current.x]+1 {
		validNextSteps = append(validNextSteps, Point{current.x - 1, current.y})
	}

	// can check right
	if current.x < len(input[0])-1 &&
		input[current.y][current.x+1] == input[current.y][current.x]+1 {
		validNextSteps = append(validNextSteps, Point{current.x + 1, current.y})
	}

	// can check up
	if current.y > 0 && input[current.y-1][current.x] == input[current.y][current.x]+1 {
		validNextSteps = append(validNextSteps, Point{current.x, current.y - 1})
	}

	// can check down
	if current.y < len(input)-1 && input[current.y+1][current.x] == input[current.y][current.x]+1 {
		validNextSteps = append(validNextSteps, Point{current.x, current.y + 1})
	}
	return validNextSteps
}
