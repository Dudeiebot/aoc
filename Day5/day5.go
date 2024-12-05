package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/dudeiebot/aoc/lib"
)

func getRulesAndUpdate() ([][]string, [][]string) {
	// input := []string{
	// 	"47|53",
	// 	"97|13",
	// 	"97|61",
	// 	"97|47",
	// 	"75|29",
	// 	"61|13",
	// 	"75|53",
	// 	"29|13",
	// 	"97|29",
	// 	"53|29",
	// 	"61|53",
	// 	"97|53",
	// 	"61|29",
	// 	"47|13",
	// 	"75|47",
	// 	"97|75",
	// 	"47|61",
	// 	"75|61",
	// 	"47|29",
	// 	"75|13",
	// 	"53|13",
	// 	"",
	// 	"75,47,61,53,29",
	// 	"97,61,53,29,13",
	// 	"75,29,13",
	// 	"75,97,47,61,53",
	// 	"61,13,29",
	// 	"97,13,75,29,47",
	// }

	input := lib.OpenFile("output.txt")
	isRule := true
	var rules, updates [][]string

	for _, each := range input {

		// Check for an empty line as the delimiter
		if len(each) == 0 {
			isRule = false
			continue
		}

		// Parse the integers in the current line
		intStrings := strings.Fields(each) // Split line into words
		var lineStr []string
		for _, str := range intStrings {
			lineStr = append(lineStr, str)
		}

		if isRule {
			rules = append(rules, lineStr)
		} else {
			updates = append(updates, lineStr)
		}
	}
	return rules, updates
}

func rulesMap(rules [][]string) map[int][]int {
	ruleMap := make(map[int][]int)

	for _, rule := range rules {

		ruleInt := splitString(rule)

		if _, ok := ruleMap[ruleInt[0]]; ok {
			ruleMap[ruleInt[0]] = append(ruleMap[ruleInt[0]], ruleInt[1])
		} else {
			ruleMap[ruleInt[0]] = []int{ruleInt[1]}
		}
	}
	return ruleMap
}

func getMiddleNumber(ruleMap map[int][]int, updates [][]string) (int, int) {
	sum := 0
	correctSum := 0

	for _, update := range updates {
		var individualUpdates []int
		individualUpdates = splitString(update)

		sortedUpdate := make([]int, len(individualUpdates))
		copy(sortedUpdate, individualUpdates)

		sort.Slice(sortedUpdate, func(i, j int) bool {
			return validUpdate(ruleMap, sortedUpdate, i, j)
		})

		if reflect.DeepEqual(individualUpdates, sortedUpdate) {
			sum += individualUpdates[len(individualUpdates)/2]
		} else {
			correctSum += sortedUpdate[len(sortedUpdate)/2]
		}
	}

	return sum, correctSum
}

func validUpdate(ruleMap map[int][]int, update []int, i, j int) bool {
	// custom sorting
	if _, ok := ruleMap[update[i]]; ok {
		for _, each := range ruleMap[update[i]] {
			if each == update[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	rules, updates := getRulesAndUpdate()

	ans, res := getMiddleNumber(rulesMap(rules), updates)
	fmt.Println(ans, res)
}

func splitString(input []string) []int {
	numInts := make([]int, 0)
	for _, str := range input {
		parts := strings.Split(str, ",")
		if len(parts) == 1 {
			parts = strings.Split(str, "|")
		}

		for _, p := range parts {
			num, err := strconv.Atoi(p)
			if err == nil {
				numInts = append(numInts, num)
			}
		}
	}
	return numInts
}
