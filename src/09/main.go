package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var count1 = 0
var blocks = make([]int, 0)
var empty = make([]int, 0)
var lim = 0

func main() {
	readInput("input.txt")
	reorder()
	calculateChecksum()
	log.Printf("Solution 1: %d", count1)
}

func readInput(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	nextID := 0
	isFile := true
	index := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, c := range strings.Split(scanner.Text(), "") {
			num, _ := strconv.Atoi(c)

			if isFile {
				for range num {
					blocks = append(blocks, nextID)
					index++
				}
				nextID++
			} else {
				for range num {
					blocks = append(blocks, -1)
					empty = append(empty, index)
					index++
				}
			}

			isFile = !isFile
		}
	}
}

func reorder() {
	iFile := len(blocks) - 1

	for _, iEmpty := range empty {
		if iFile <= iEmpty {
			break
		}

		for iFile >= 0 && blocks[iFile] == -1 {
			iFile--
		}

		if iFile >= 0 {
			blocks[iEmpty] = blocks[iFile]
			blocks[iFile] = -1
			iFile--
		}
	}

	lim = iFile
}

func calculateChecksum() {
	for i, num := range blocks {
		if i > lim {
			break
		}

		count1 += i * num
	}
}
