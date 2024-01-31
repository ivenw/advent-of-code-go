package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content_raw, _ := os.ReadFile("input.txt")
	content_parsed := parse(content_raw)

	// partOne(content_parsed)
	// partTwo(content_parsed)
}

func parse(content_raw []byte) []int {
	content := string(content_raw)
	content = strings.TrimSpace(content)
	numbers := strings.Split(content, ",")
	var numbers_parsed []int
	for _, number := range numbers {
		number_parsed, _ := strconv.Atoi(number)
		numbers_parsed = append(numbers_parsed, number_parsed)
	}
	return numbers_parsed
}

func partOne(module_masss []int) {

	fmt.Println("Part 1:", answer)
}

// func partTwo(module_masss []int) {
//
// 	fmt.Println("Part 2:", answer)
// }
