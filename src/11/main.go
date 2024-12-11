package main

import (
	"log"
	"strconv"
)

var stones = map[int]int{
	3028:    1,
	78:      1,
	973951:  1,
	5146801: 1,
	5:       1,
	0:       1,
	23533:   1,
	857:     1,
}

func main() {
	blink(75) // turns negative after 124
	log.Printf("Solution: %d\n", count())
}

func blink(n int) {
	for i := 0; i < n; i++ {
		nextStones := make(map[int]int)
		for num, count := range stones {
			evolveStone(num, count, nextStones)
		}

		stones = nextStones
	}
}

func evolveStone(num, count int, result map[int]int) {
	if num == 0 {
		result[1] += count
		return
	}

	sNum := strconv.Itoa(num)
	if len(sNum)%2 == 0 {
		mid := len(sNum) / 2
		left, _ := strconv.Atoi(sNum[:mid])
		right, _ := strconv.Atoi(sNum[mid:])
		result[left] += count
		result[right] += count
	} else {
		result[num*2024] += count
	}
}

func count() (total int) {
	for _, count := range stones {
		total += count
	}
	return
}
