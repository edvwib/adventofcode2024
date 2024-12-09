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

		for i := len(blocks) - 1; i > 0; i-- {
			currentBlock := blocks[i]
			if currentBlock.free {
				continue
			}

			for j, file := range blocks {
				if j >= i {
					break
				}
				if !file.free {
					continue
				}
				if file.size >= currentBlock.size {
					blocks[j] = block{
						value: currentBlock.value,
						size:  currentBlock.size,
						free:  false,
					}

					leftOver := file.size - currentBlock.size
					replaceLastIndex := i
					if leftOver > 0 {
						replaceLastIndex++
						blocks = slices.Insert(blocks, j+1, block{
							value: -1,
							size:  leftOver,
							free:  true,
						})
					}

					blocks[replaceLastIndex] = block{
						value: -1,
						size:  currentBlock.size,
						free:  true,
					}

					break
				}
			}
		}

		checksum := 0
		i := 0
		for _, file := range blocks {
			if file.free {
				i += file.size
				continue
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
		if block.free {
			sb.WriteString(strings.Repeat(".", block.size))
		} else {
			sb.WriteString(strings.Repeat(strconv.Itoa(block.value), block.size))
		}
	}
	fmt.Println("builder", sb.String())
}
