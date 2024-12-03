package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
	n1 := 0
	n2 := 0
	do := true

	for scanner.Scan() {
		line := scanner.Text()
		re, _ := regexp.Compile(pattern)

		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if strings.Contains(match[0], "do()") {
				do = true
			} else if strings.Contains(match[0], "don't()") {
				do = false
			} else {
				l, _ := strconv.Atoi(match[1])
				r, _ := strconv.Atoi(match[2])
				if do {
					n2 += l * r
					n1 += l * r
				} else {
					n1 += l * r
				}
			}
		}
	}

	log.Printf("Solution 1: %d\n", n1) // 192767529
	log.Printf("Solution 1: %d\n", n2) // 104083373
}
