package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dudeiebot/aoc/lib"
)

func main() {
	input := lib.OpenFile("output.txt")
	fmt.Println(input[0])
	ans := getCurrentStone(input, 75)
	fmt.Println("Total Stone after 75 blinks:", ans)
}

func getCurrentStone(input []string, blinks int) int {
	totalStone := 0
	store := make(map[int]int)

	in := strings.Split(input[0], " ")
	for _, stone := range in {
		num, _ := strconv.Atoi(string(stone))
		store[num]++
	}

	for i := 0; i < blinks; i++ {
		store = blinkStone(store)
	}

	for _, v := range store {
		totalStone += v
	}
	return totalStone
}

func blinkStone(cache map[int]int) map[int]int {
	nextCache := make(map[int]int)

	for k, v := range cache {
		if k == 0 {
			nextCache[1] += v
		} else if len(strconv.Itoa(k))%2 == 0 {
			str := strconv.Itoa(k)
			leftHalf, _ := strconv.Atoi(str[:(len(str) / 2)])
			rightHalf, _ := strconv.Atoi(str[(len(str) / 2):])

			nextCache[leftHalf] += v
			nextCache[rightHalf] += v
		} else {
			nextCache[k*2024] += v
		}
	}
	return nextCache
}
