package main

import (
	"bufio"
	"fmt"
	"log"
	_ "log"
	"os"
)

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}
func main() {
	// open file for reading
	// read line by line
	lines, err := readLines("data/100.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	// print file contents
	var a1 [100]int
	for j, line := range lines {
		fmt.Println(line)

		a1 = append(a1, line)
	}
}
