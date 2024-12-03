package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func OpenFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error Opening File: ", err)
	}
	defer file.Close()
	var output []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}
