package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/dudeiebot/aoc/lib"
)

type robot struct {
	pos xy
	vel xy
}

type xy struct {
	x int
	y int
}

const (
	height = 103
	width  = 101
)

func main() {
	input := lib.OpenFile("output.txt")
	robots := getRobots(input)
	robots = getPositionAfterSec(robots, 100, width, height)
	ans := getScore(robots, height, width)
	robots2 := getRobots(input)
	res := findTree(robots2, width, height)
	fmt.Println("Ans: ", ans)
	fmt.Println("Ans: ", res)
}

func getRobots(input []string) []robot {
	re := regexp.MustCompile(`([+-]?\d+)`)

	robots := []robot{}
	for _, line := range input {
		nums := re.FindAllString(line, -1)
		px, _ := strconv.Atoi(nums[0])
		py, _ := strconv.Atoi(nums[1])
		vx, _ := strconv.Atoi(nums[2])
		vy, _ := strconv.Atoi(nums[3])
		each := robot{pos: xy{px, py}, vel: xy{vx, vy}}
		robots = append(robots, each)
	}
	return robots
}

func (p xy) add(p2 xy) xy {
	return xy{p.x + p2.x, p.y + p2.y}
}

func getPositionAfterSec(robots []robot, sec, width, height int) []robot {
	for i := 0; i < sec; i++ {
		for index, each := range robots {
			nVel := each.pos.add(each.vel)
			if nVel.x >= width {
				nVel.x = nVel.x % width
			} else if nVel.x < 0 {
				nVel.x = width + nVel.x
			}
			if nVel.y >= height {
				nVel.y = nVel.y % height
			} else if nVel.y < 0 {
				nVel.y = height + nVel.y
			}
			robots[index].pos = nVel
		}
	}
	return robots
}

func getScore(robots []robot, height, width int) int {
	seen := make(map[int]int)
	total := 1
	for _, robot := range robots {
		k := robot.pos
		if k.x < width/2 && k.y < height/2 {
			seen[0]++
		} else if k.x > width/2 && k.y < height/2 {
			seen[1]++
		} else if k.x < width/2 && k.y > height/2 {
			seen[2]++
		} else if k.x > width/2 && k.y > height/2 {
			seen[3]++
		}
	}
	for _, v := range seen {
		total *= v
	}
	return total
}

func rootMeanSquare(robots []robot) float64 {
	var ans int
	for i := 0; i < len(robots); i++ {
		for j := 0; j < len(robots); j++ {
			if i == j {
				continue
			}
			x1, y1 := robots[i].pos.x, robots[i].pos.y
			x2, y2 := robots[j].pos.x, robots[j].pos.y

			ans += (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
		}
	}
	fnum := float64(len(robots))
	return math.Sqrt(float64(ans) / (fnum * fnum))
}

func print(robots []robot, width, height int) {
	m := make(map[xy]robot)
	for _, b := range robots {
		m[b.pos] = b
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pt := xy{x, y}
			if _, ok := m[pt]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func findTree(robots []robot, width, height int) int {
	n := 1
	for {
		robots = getPositionAfterSec(robots, 1, width, height)
		if rootMeanSquare(robots) < 42 {
			print(robots, width, height)
			break
		}
		n++
	}
	return n
}
