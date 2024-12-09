package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type block struct {
	value int
	size  int
	free  bool
}

func main() {
	file, err := os.Open("../input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		diskMap := scanner.Text()

		fileId := 0
		var blocks []block

		for i, char := range diskMap {
			size, err := strconv.Atoi(string(char))
			check(err)
			if i%2 == 0 {
				// file
				blocks = append(blocks, block{
					value: fileId,
					size:  size,
					free:  false,
				})
				fileId += 1
			} else {
				// free space
				blocks = append(blocks, block{
					value: -1,
					size:  size,
					free:  true,
				})
			}
		}

		lastFile := blocks[len(blocks)-1]
		lastFileIndex := len(blocks) - 1

		for blockIndex, file := range blocks {
			if blockIndex >= lastFileIndex {
				break
			}
			if !file.free {
				continue
			}
			if file.size == lastFile.size {
				// lastFile fits exactly
				blocks[blockIndex] = lastFile
				blocks = blocks[:lastFileIndex]
			} else if file.size < lastFile.size {
				// lastFile will not fit entirely
				fileSizeLeft := lastFile.size - file.size

				blocks[blockIndex] = block{
					value: lastFile.value,
					size:  file.size,
					free:  lastFile.free,
				}
				blocks[lastFileIndex] = block{
					value: lastFile.value,
					size:  fileSizeLeft,
					free:  lastFile.free,
				}
				lastFile = blocks[lastFileIndex]
				continue
			} else {
				// lastFile will fit entirely, with space left over
				fileSizeLeft := file.size - lastFile.size

				blocks[blockIndex] = lastFile
				blocks = blocks[:lastFileIndex]
				blocks = slices.Insert(blocks, blockIndex+1, block{
					value: file.value,
					size:  fileSizeLeft,
					free:  file.free,
				})
			}

			for {
				lastFileIndex -= 1
				lastFile = blocks[lastFileIndex]
				if !lastFile.free {
					break
				}
			}
		}

		checksum := 0
		i := 0
		for _, file := range blocks {
			if file.free {
				break
			}

			for j := 0; j < file.size; j++ {
				checksum = checksum + file.value*i
				i++
			}
		}

		fmt.Println("checksum", checksum)
	}
}

func printBlocks(blocks []block) {
	var sb strings.Builder
	for _, block := range blocks {
		if block.value == -1 {
			sb.WriteString(strings.Repeat(".", block.size))
		} else {
			sb.WriteString(strings.Repeat(strconv.Itoa(block.value), block.size))
		}
	}
	fmt.Println("builder", sb.String())
}
