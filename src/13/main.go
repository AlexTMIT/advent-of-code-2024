package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type c struct {
	x, y int64
}

type machine struct {
	b1, b2, prize c
}

var machines = make([]machine, 0)
var count1 = int64(0)
var count2 = int64(0)

func main() {
	readInput("input.txt")
	solve()
	log.Printf("Solution 1: %d", count1)
	log.Printf("Solution 2: %d", count2)
}

func readInput(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var m machine
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if parts[0] == "Button" {
			x, _ := strconv.Atoi(parts[2][2:4])
			y, _ := strconv.Atoi(parts[3][2:4])
			if parts[1] == "A:" {
				m.b1.x = int64(x)
				m.b1.y = int64(y)
			} else {
				m.b2.x = int64(x)
				m.b2.y = int64(y)
			}
		} else if parts[0] == "Prize:" {
			x, _ := strconv.Atoi(parts[1][2 : len(parts[1])-1])
			y, _ := strconv.Atoi(parts[2][2:len(parts[2])])
			m.prize.x = int64(x)
			m.prize.y = int64(y)
			machines = append(machines, m)
		}
	}
}

func solve() {
	for _, m := range machines {
		count1 += minimize(m.b1, m.b2, m.prize)
		count2 += minimize(m.b1, m.b2, c{x: m.prize.x + 10000000000000, y: m.prize.y + 10000000000000})
	}
}

// inspired by https://github.com/quintal-william/advent-of-code-2024/blob/main/src/year2024/day13.rs
func minimize(b1, b2, p c) int64 {
	d := b2.y*b1.x - b2.x*b1.y // determinant
	if d == 0 {
		return 0
	}

	bNom := p.y*b1.x - p.x*b1.y
	if bNom%d != 0 {
		return 0
	}
	b := bNom / d

	aNom := p.x - b2.x*b
	if aNom%b1.x != 0 {
		return 0
	}
	a := aNom / b1.x

	return 3*a + b
}
