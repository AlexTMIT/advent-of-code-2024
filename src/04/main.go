package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var n = 140 // 10 for test, 140 for input
var a = make([][]string, n)
var count = 0

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "")
		a[i] = parts
		i++
	}

	for y, s := range a {
		for x, c := range s {
			if c == "X" {
				log.Printf("Checking %s in %d, %d\n", c, x, y)

				if x+3 < n && a[y][x+1] == "M" && a[y][x+2] == "A" && a[y][x+3] == "S" { // right
					count++
				}
				if x-3 >= 0 && a[y][x-1] == "M" && a[y][x-2] == "A" && a[y][x-3] == "S" { // left
					count++
				}
				if y+3 < n && a[y+1][x] == "M" && a[y+2][x] == "A" && a[y+3][x] == "S" { // down
					count++
				}

				if y-3 >= 0 && a[y-1][x] == "M" && a[y-2][x] == "A" && a[y-3][x] == "S" { // up
					count++
				}

				if y+3 < n && x-3 >= 0 && a[y+1][x-1] == "M" && a[y+2][x-2] == "A" && a[y+3][x-3] == "S" { // bottom left
					count++
				}

				if y-3 >= 0 && x-3 >= 0 && a[y-1][x-1] == "M" && a[y-2][x-2] == "A" && a[y-3][x-3] == "S" { // top left
					count++
				}

				if y+3 < n && x+3 < n && a[y+1][x+1] == "M" && a[y+2][x+2] == "A" && a[y+3][x+3] == "S" { // bottom right
					count++
				}

				if y-3 >= 0 && x+3 < n && a[y-1][x+1] == "M" && a[y-2][x+2] == "A" && a[y-3][x+3] == "S" { // top right
					count++
				}
			}
		}
	}

	log.Println(count)
}
