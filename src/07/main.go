package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var count1 = 0
var count2 = 0

var n = 850 // 850 for input, 9 for test
var equations = make([][]int, n)
var concat = false

func main() {
	readInput("input.txt")
	calibrate()
	log.Printf("Solution 1: %d\n", count1) // 20665830408335
	log.Printf("Solution 2: %d\n", count2) // 354060705047464
}

func calibrate() {
	for _, equation := range equations {
		target := equation[0]
		nums := equation[1:]

		concat = false
		if evaluate(nums[1:], target, nums[0]) {
			if concat {
				count2 += target
			} else {
				count1 += target
				count2 += target
			}
		}
	}
}

func evaluate(nums []int, targ, curr int) bool {
	if len(nums) == 0 {
		return curr == targ
	}
	if evaluate(nums[1:], targ, curr+nums[0]) {
		return true
	}
	if evaluate(nums[1:], targ, curr*nums[0]) {
		return true
	}
	if evaluate(nums[1:], targ, Concat(curr, nums[0])) {
		concat = true
		return true
	}

	return false
}

func Concat(a int, b int) int {
	c := fmt.Sprintf("%d%d", a, b)
	cd, _ := strconv.Atoi(c)
	return cd
}

func readInput(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		parts := strings.Split(strings.TrimSpace(scanner.Text()), ":")
		left, _ := strconv.Atoi(parts[0])
		right := strings.Fields(parts[1])

		equation := append([]int{left}, toIntSlice(right)...)
		equations[i] = equation
	}
}

func toIntSlice(strs []string) []int {
	ints := make([]int, len(strs))
	for i, s := range strs {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}
