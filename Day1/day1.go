package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func openFile(fileName string) ([]int, []int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var leftDigits []int
	var rightDigits []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())

		if len(parts) == 2 {
			leftDigit, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Error converting left digit:", err)
				continue
			}

			rightDigit, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error converting right digit:", err)
				continue
			}

			leftDigits = append(leftDigits, leftDigit)
			rightDigits = append(rightDigits, rightDigit)
		}
	}
	return leftDigits, rightDigits, nil

	// bytes, _ := os.ReadFile(fileName)
	// lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	// fmt.Print(lines)
	// n := len(lines)
	// fmt.Println(n)
	//
	// l := make([]int, n)
	// r := make([]int, n)
	//
	// for i := 0; i < n; i++ {
	// 	fmt.Sscanf(lines[i], "%d %d", &l[i], &r[i])
	// }
	// return l, r, nil
}

func findDistance(ll []int, rr []int) int {
	if len(ll) != len(rr) {
		return 0
	}

	slices.Sort(ll)
	slices.Sort(rr)

	ans := 0
	for i := 0; i < len(ll); i++ {
		ans += int(math.Abs(float64(ll[i] - rr[i])))
	}

	return ans
}

func main() {
	leftDigits, rightDigits, err := openFile("output.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ans := findDistance(leftDigits, rightDigits)
	fmt.Println("Distance ans", ans)

	res := findSimilarities(rightDigits, leftDigits)
	fmt.Println("Similarities ans", res)
}

func findSimilarities(rr, ll []int) int {
	rmap := make(map[int]int)
	ans := 0

	for i := 0; i < len(rr); i++ {
		rmap[rr[i]]++
	}
	for _, l := range ll {
		for k, v := range rmap {
			if k == l {
				ans += l * v
			}
		}
	}
	return ans
}
