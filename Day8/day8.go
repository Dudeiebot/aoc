package main

import (
	"fmt"

	"github.com/dudeiebot/aoc/lib"
)

type Point struct {
	row int
	col int
}

func symmetricField(a, b Point) Point {
	return Point{2*b.row - a.row, 2*b.col - a.col}
}

func pointOnLineAB(a, b, c Point) bool {
	return float64(c.col-a.col)-float64(b.col-a.col)*float64(c.row-a.row)/float64(b.row-a.row) == 0
}

func isValidCoords(grid [][]rune, p Point) bool {
	return p.row >= 0 && p.row < len(grid) && p.col >= 0 && p.col < len(grid[0])
}

func collectAntennas(grid [][]rune) map[rune][]Point {
	antennas := make(map[rune][]Point)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '.' {
				antennas[grid[i][j]] = append(antennas[grid[i][j]], Point{row: i, col: j})
			}
		}
	}
	return antennas
}

func partOne(grid [][]rune) int {
	antennas := collectAntennas(grid)
	antinode := map[Point]bool{}

	for _, antenna := range antennas {
		for i := 0; i < len(antenna)-1; i++ {
			for j := i + 1; j < len(antenna); j++ {
				an1 := symmetricField(antenna[i], antenna[j])
				an2 := symmetricField(antenna[j], antenna[i])

				if isValidCoords(grid, an1) {
					antinode[an1] = true
				}
				if isValidCoords(grid, an2) {
					antinode[an2] = true
				}

			}
		}
	}
	return len(antinode)
}

func partTwo(grid [][]rune) int {
	antennas := collectAntennas(grid)
	seen := map[Point]bool{}
	for _, antenna := range antennas {
		for i := 0; i < len(antenna)-1; i++ {
			for j := i + 1; j < len(antenna); j++ {
				ax, bx := antenna[i].row, antenna[j].row
				ay, by := antenna[i].col, antenna[j].col

				cx := ax
				cy := ay
				for isValidCoords(grid, Point{cx, cy}) && pointOnLineAB(antenna[i], antenna[j], Point{cx, cy}) {
					if _, ok := seen[Point{cx, cy}]; !ok {
						seen[Point{cx, cy}] = true
					}
					cx -= bx - ax
					cy -= by - ay
				}

				dx := bx
				dy := by
				for isValidCoords(grid, Point{dx, dy}) && pointOnLineAB(antenna[i], antenna[j], Point{dx, dy}) {
					if _, ok := seen[Point{dx, dy}]; !ok {
						seen[Point{dx, dy}] = true
					}
					dx += bx - ax
					dy += by - ay
				}
			}
		}
	}

	return len(seen)
}

func main() {
	input := lib.OpenFile("output.txt")
	grid := [][]rune{}
	for _, line := range input {
		grid = append(grid, []rune(line))
	}

	ans := partOne(grid)
	ans2 := partTwo(grid)

	fmt.Println("Answer for Part 1:", ans)
	fmt.Println("Answer for Part 2:", ans2)
}
