package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binarySearch(array []int, target int) int {
	left, rigth := 0, len(array)-1
	result := -1
	for left <= rigth {
		mid := left + (rigth-left)/2

		if array[mid] == target {
			result = mid
			left = mid + 1
		} else if array[mid] < target {
			left = mid + 1
		} else if array[mid] > target {
			rigth = mid - 1
		}
	}

	return result
}

func main() {
	// Initialization of the scanner buffer
	scanner := bufio.NewScanner(os.Stdin)

	// Read array data
	println("leng and queries:")
	scanner.Scan()
	data := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(data[0])
	Q, _ := strconv.Atoi(data[1])

	// Read array content
	println("content:")
	scanner.Scan()
	content := strings.Fields(scanner.Text())
	array := make([]int, N)
	for i, c := range content {
		array[i], _ = strconv.Atoi(c)
	}

	for i := 0; i < Q; i++ {
		fmt.Println("get query:")
		scanner.Scan()
		query, _ := strconv.Atoi(scanner.Text())
		fmt.Println("position:")
		fmt.Println(binarySearch(array, query))
	}
	fmt.Scanln()
}
