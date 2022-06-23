package main

import (
	"bufio"
	"fmt"
	"io"
	_ "log"
	"os"
	"time"
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
	file, err := os.Open("data/100.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var perline int
	var nums []int
	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		nums = append(nums, perline)
	}
	inicio := time.Now()
	BubbleSort(nums)
	duracion := time.Since(inicio)
	fmt.Println("Tiempo en Microsegundos: ", duracion.Nanoseconds())

}
