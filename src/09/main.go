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

var Blocks = make([]int, 0)
var Empty = make([]int, 0)
var indexToFirstEmpty = make([]int, 0)

func main() {
	readInput("input.txt")

	blocksOne := partOne()
	calculateChecksum(blocksOne, &count1)
	blocksTwo := partTwo()
	calculateChecksum(blocksTwo, &count2)

	log.Printf("Solution 1: %d", count1)
	log.Printf("Solution 2: %d", count2)
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
					Blocks = append(Blocks, nextID)
					index++
				}
				nextID++
			} else {
				indexToFirstEmpty = append(indexToFirstEmpty, index)
				for range num {
					Blocks = append(Blocks, num*-1)
					Empty = append(Empty, index)
					index++
				}
			}

			isFile = !isFile
		}
	}
}

func partOne() []int {
	var blocks = make([]int, len(Blocks))
	copy(blocks, Blocks)

	iFile := len(blocks) - 1

	for _, iEmpty := range Empty {
		if iFile <= iEmpty {
			break
		}

		for iFile >= 0 && blocks[iFile] < 0 {
			iFile--
		}

		if iFile >= 0 {
			blocks[iEmpty] = blocks[iFile]
			blocks[iFile] = -11
			iFile--
		}
	}

	return blocks
}

func partTwo() []int {
	var blocks = make([]int, len(Blocks))
	copy(blocks, Blocks)

	emptySpaces := make([][2]int, 0) // [0] = index, [1] = space
	for _, iEmpty := range indexToFirstEmpty {
		emptySpaces = append(emptySpaces, [2]int{iEmpty, blocks[iEmpty] * -1})
	}

	for iFile := len(blocks) - 1; iFile >= 0; {
		if blocks[iFile] < 0 {
			iFile--
			continue
		}

		fileID := blocks[iFile]
		fileStart := iFile
		for fileStart >= 0 && blocks[fileStart] == fileID {
			fileStart--
		}
		fileStart++

		fileLength := iFile - fileStart + 1

		for i, e := range emptySpaces { // find first empty space that fits file
			iEmpty, emptyLength := e[0], e[1]

			if fileLength <= emptyLength && iEmpty < fileStart { // file fits?
				for j := 0; j < fileLength; j++ { // move file
					blocks[iEmpty+j] = fileID
					blocks[fileStart+j] = -11
				}

				if fileLength < emptyLength { // update empty space tracking
					newEmptyStart := iEmpty + fileLength
					newEmptyLength := emptyLength - fileLength
					blocks[newEmptyStart] = -newEmptyLength
					emptySpaces[i] = [2]int{newEmptyStart, newEmptyLength}
				} else {
					emptySpaces = append(emptySpaces[:i], emptySpaces[i+1:]...)
				}
				break
			}
		}

		iFile = fileStart - 1
	}

	return blocks
}

func calculateChecksum(blocks []int, count *int) {
	*count = 0
	for i, num := range blocks {
		if num >= 0 {
			*count += i * num
		}
	}
}
