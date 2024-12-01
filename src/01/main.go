package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var left []int
var right []int
var dist []int

func main() {
	readInput()
	sortInput()

	calculateDist()
	sumDist := Sum(dist)
	log.Printf("Solution 1: %d", sumDist) // 2378066

	similarityScore := calculateSimularityScore()
	log.Printf("Solution 2: %d", similarityScore) // 18934359
}

func readInput() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input: %v", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		leftNum, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatalf("failed to convert %s to int: %v", nums[0], err)
		}
		rightNum, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatalf("failed to convert %s to int: %v", nums[1], err)
		}

		left = append(left, leftNum)
		right = append(right, rightNum)

		i++
	}
}

func sortInput() {
	sort.Ints(left)
	sort.Ints(right)
}

func calculateDist() {
	for i := range left {
		diff := left[i] - right[i]
		dist = append(dist, Abs(diff))
	}
}

func Sum(a []int) int {
	sum := 0

	for i := range a {
		sum += a[i]
	}

	return sum
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateSimularityScore() int {
	score := 0

	for _, e := range left {
		occurrences := findOccurrences(right, e)
		score += e * occurrences
	}

	return score
}

func findOccurrences(right []int, e int) int {
	occurrences := 0

	i, err := findFirstIndex(right, e)
	if err != nil {
		return 0
	}

	for {
		if right[i] != e {
			break
		}

		i++
		occurrences++
	}

	return occurrences
}

func findFirstIndex(slice []int, target int) (int, error) {
	index := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= target
	})

	if slice[index] == target {
		return index, nil
	}

	return -1, errors.New("element not found")
}
