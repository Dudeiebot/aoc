package main

import (
	"fmt"
	"unicode"

	"github.com/dudeiebot/aoc/lib"
)

func main() {
	input := lib.OpenFile("output.txt")
	str := calculateChecksum(moveFileBlocks(generateFileBlock(input[0])))
	ans := calculateChecksum(moveFileBlocks2(generateFileBlock(input[0])))
	fmt.Println(str)
	fmt.Println(ans)
}

func generateFileBlock(input string) []int {
	fileBlock := []int{}
	for index, char := range input {
		// Check if the character is a digit
		if unicode.IsDigit(char) {
			// Extract the numeric value directly
			digit := int(char - '0')

			// Add to fileBlock based on the index
			for i := 0; i < digit; i++ { // Repeat based on the numeric value
				if index%2 == 0 {
					fileBlock = append(fileBlock, index/2) // Even index: add index/2
				} else {
					fileBlock = append(fileBlock, -1) // Odd index: add -1
				}
			}
		}
	}
	return fileBlock
}

func moveFileBlocks(block []int) []int {
	start, end := 0, len(block)-1
	for start < end {
		if block[start] == -1 && block[end] != -1 {
			block[start], block[end] = block[end], block[start]
			start++
			end--
			continue
		}
		if block[start] != -1 {
			start++
		}
		if block[end] == -1 {
			end--
		}
	}
	return block
}

func writeFile(block []int, fileNum int, fileLength int, index int) {
	for i := range fileLength {
		block[index+i] = fileNum
	}
}

func clearFile(block []int, length int, index int) {
	for i := range length {
		block[index+i] = -1
	}
}

func moveFile(fileBlock []int, length int, originalStartIndex int) {
	freeSpaceCount := 0
	freeSpaceStartIndex := -1
	for i := 1; i <= originalStartIndex; i++ {
		if fileBlock[i-1] != -1 && fileBlock[i] == -1 {
			if freeSpaceCount >= length {
				writeFile(fileBlock, fileBlock[originalStartIndex], length, freeSpaceStartIndex)
				clearFile(fileBlock, length, originalStartIndex)
			}
			freeSpaceStartIndex = i
			freeSpaceCount = 1
			continue
		}

		if fileBlock[i-1] == -1 && fileBlock[i] != -1 {
			if freeSpaceCount >= length {
				writeFile(fileBlock, fileBlock[originalStartIndex], length, freeSpaceStartIndex)
				clearFile(fileBlock, length, originalStartIndex)
			}
			freeSpaceStartIndex = -1
			freeSpaceCount = 0
		}

		if fileBlock[i] == -1 {
			freeSpaceCount++
		}
	}
}

func moveFileBlocks2(fileBlock []int) []int {
	currentFile := -1
	currentFileLength := 0
	for i := len(fileBlock) - 1; i > 0; i-- {
		if fileBlock[i] != -1 {
			currentFileLength += 1
			currentFile = i
			if fileBlock[i] != fileBlock[i-1] {
				if currentFileLength != 0 {
					moveFile(fileBlock, currentFileLength, currentFile)
					currentFileLength = 0
				}
			}
		} else {
			if currentFileLength != 0 {
				moveFile(fileBlock, currentFileLength, currentFile)
				currentFileLength = 0
			}
		}
	}
	return fileBlock
}

func calculateChecksum(input []int) int {
	sum := 0
	for index, fileNumber := range input {
		if fileNumber == -1 {
			continue
		}
		sum += index * fileNumber
	}
	return sum
}
