package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type c struct {
	x, y int
}

type robot struct {
	p, v c
}

var X = 101
var Y = 103

var robots = make([]robot, 0)
var qs = make([]int, 4)

func main() {
	readInput("input.txt")
	log.Printf("Solution 1: %d", partOne())
	log.Printf("Solution 2: %d", partTwo())
}

func readInput(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		pc := getC(parts[0], "p=")
		vc := getC(parts[1], "v=")
		robots = append(robots, robot{pc, vc})
	}
}

func getC(part, split string) c {
	splits := strings.Split(strings.Split(part, split)[1], ",")
	x, _ := strconv.Atoi(splits[0])
	y, _ := strconv.Atoi(splits[1])
	return c{x, y}
}

func partOne() int {
	for i := range robots {
		r := moveRobot(&robots[i], 100)
		incrementQuadrant(r)
	}

	return qs[0] * qs[1] * qs[2] * qs[3]
}

func partTwo() int {
	robots = make([]robot, 0)
	readInput("input.txt")

	var positions map[c]bool
	for i := 1; true; i++ {
		positions = make(map[c]bool)
		for i := range robots {
			r := moveRobot(&robots[i], 1)
			positions[r.p] = true
		}
		if len(robots) == len(positions) {
			return i
		}

	}
	return 0
}

func moveRobot(r *robot, s int) *robot {
	r.p.x = (r.p.x + r.v.x*s) % X
	if r.p.x < 0 {
		r.p.x += X
	}

	r.p.y = (r.p.y + r.v.y*s) % Y
	if r.p.y < 0 {
		r.p.y += Y
	}
	return r
}

func incrementQuadrant(r *robot) {
	midX := X / 2
	midY := Y / 2
	if r.p.x < midX && r.p.y < midY {
		qs[0]++ // top-left
	} else if r.p.x > midX && r.p.y < midY {
		qs[1]++ // top-right
	} else if r.p.x < midX && r.p.y > midY {
		qs[2]++ // bottom-left
	} else if r.p.x > midX && r.p.y > midY {
		qs[3]++ // bottom-right
	}
}
