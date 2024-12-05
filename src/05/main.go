package main

import (
	"bufio"
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
	readInput("input.txt")
	reviewUpdates()

	log.Printf("Solution 1: %d", count1)
	log.Printf("Solution 2: %d", count2)
}

func readInput(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		} else if line[2] == '|' {
			addRule(line)
		} else {
			addUpdate(line)
		}
	}
}

func addUpdate(line string) {
	parts := strings.Split(line, ",")
	update := make([]int, len(parts))

	for i, part := range parts {
		page, _ := strconv.Atoi(part)
		update[i] = page
	}

	updates = append(updates, update)
}

func addRule(line string) {
	parts := strings.Split(line, "|")
	l, _ := strconv.Atoi(parts[0])
	r, _ := strconv.Atoi(parts[1])
	rules[r] = append(rules[r], l)
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
