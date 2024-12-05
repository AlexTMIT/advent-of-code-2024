/*
For a good solution, see William Ford's. It's 22 lines. Bruh.
https://github.com/quintal-william/advent-of-code-2024/tree/main/src/year2024
*/

package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var count1 = 0
var count2 = 0

var rules = make([][]int, 100)
var seen = make([]bool, 100)
var updates = make([][]int, 0)

func main() {
	content, _ := os.ReadFile("input.txt")
	sections := strings.Split(string(content), "\n\n")

	parseRules(sections[0])
	parseUpdates(sections[1])
	reviewUpdates()

	log.Printf("Solution 1: %d", count1) // 7365
	log.Printf("Solution 2: %d", count2) // 5770
}

func parseRules(data string) {
	for _, line := range strings.Split(data, "\n") {
		if parts := strings.Split(line, "|"); len(parts) == 2 {
			r, _ := strconv.Atoi(parts[1])
			l, _ := strconv.Atoi(parts[0])
			rules[r] = append(rules[r], l)
		}
	}
}

func parseUpdates(data string) {
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}
		var update []int
		for _, p := range strings.Split(line, ",") {
			page, _ := strconv.Atoi(p)
			update = append(update, page)
		}
		updates = append(updates, update)
	}
}

func reviewUpdates() {
	for _, update := range updates {
		isValid := true
		relevantRules := make([][]int, 100)

		for i, page := range update {
			for _, num := range rules[page] {
				if !Contains(update, num) {
					continue
				}

				relevantRules[page] = append(relevantRules[page], num)

				if !Contains(update[:i], num) {
					isValid = false
				}
			}
		}

		if isValid {
			addMiddleIndexPoints(update, &count1)
		} else {
			reorderUpdate(update, relevantRules)
		}
	}
}

func addMiddleIndexPoints(a []int, p *int) {
	mid := (len(a) - 1) / 2
	*p += a[mid]
}

func reorderUpdate(update []int, relevantRules [][]int) {
	remainingPages := make([]int, len(update))
	copy(remainingPages, update)
	newUpdate := make([]int, 0, len(update))

	for len(remainingPages) > 0 {
		nextPage := findCorrectFirst(remainingPages, relevantRules)
		newUpdate = append(newUpdate, nextPage)
		remainingPages = removeUpdate(remainingPages, nextPage)
	}

	addMiddleIndexPoints(newUpdate, &count2)
}

func removeUpdate(a []int, e int) []int {
	result := make([]int, 0, len(a)-1)
	for _, el := range a {
		if el != e {
			result = append(result, el)
		}
	}
	return result
}

func findCorrectFirst(remainingPages []int, relevantRules [][]int) int {
	for _, page := range remainingPages {
		canPlace := true
		for _, rule := range relevantRules[page] {
			if Contains(remainingPages, rule) {
				canPlace = false
				break
			}
		}
		if canPlace {
			return page
		}
	}
	log.Fatal("cyclic dependency found, literally no solution")
	return -1
}

func Contains(a []int, e int) bool {
	for _, el := range a {
		if el == e {
			return true
		}
	}

	return false
}
