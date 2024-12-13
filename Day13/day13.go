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
	ans1, ans2 := getTokenCount(getPrizes(input))
	fmt.Println("part 1: ", ans1)
	fmt.Println("part 2: ", ans2)
}

type Coordinates struct {
	X, Y int
}

func getPrizes(input []string) [][]Coordinates {
	values := []Coordinates{}
	prizes := [][]Coordinates{}

	// find number in string
	re := regexp.MustCompile(`([+-]?\d+)`)

	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		nums := re.FindAllString(line, -1)

		var x, y int
		if len(nums) >= 2 {
			x, _ = strconv.Atoi(nums[0])
			y, _ = strconv.Atoi(nums[1])
		}
		values = append(values, Coordinates{x, y})

		if strings.Contains(line, "Prize") {
			prizes = append(prizes, values)
			values = []Coordinates{}
		}
	}
	return prizes
}

func SolveLinearEquation(a, b, c Coordinates) (bool, Coordinates) {
	x := ((b.X * (-c.Y)) - (b.Y * (-c.X))) / ((a.X * b.Y) - (a.Y * b.X))
	y := (((-c.X) * a.Y) - ((-c.Y) * a.X)) / ((a.X * b.Y) - (a.Y * b.X))
	if (((b.X*(-c.Y))-(b.Y*(-c.X)))%((a.X*b.Y)-(a.Y*b.X)) == 0) &&
		((((-c.X)*a.Y)-((-c.Y)*a.X))%((a.X*b.Y)-(a.Y*b.X)) == 0) {
		return true, Coordinates{x, y}
	}
	return false, Coordinates{x, y}
}

func getButtonPressesIfValid(a, b, final Coordinates, add int) (bool, Coordinates) {
	final.X += add
	final.Y += add
	return SolveLinearEquation(a, b, final)
}

func getTokenCount(input [][]Coordinates) (int, int) {
	count1, count2 := 0, 0
	for _, prize := range input {
		isValid1, tokens1 := getButtonPressesIfValid(prize[0], prize[1], prize[2], 0)
		if isValid1 {
			count1 += tokens1.X*3 + tokens1.Y
		}

		isValid2, tokens2 := getButtonPressesIfValid(prize[0], prize[1], prize[2], 10000000000000)
		if isValid2 {
			count2 += tokens2.X*3 + tokens2.Y
		}
	}
	return count1, count2
}
