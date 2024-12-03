package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
}

func main() {
	content_raw, _ := os.ReadFile("input.txt")
	data := parse(content_raw)

	partOne(data)
	// partTwo(data)
}

func parse(content_raw []byte) {
	scanner := bufio.NewScanner(strings.NewReader(string(content_raw)))
}

func partOne(input Data) {
	answer := 0

	fmt.Println("Part 1:", answer)
}

// func partTwo(module_masss []int) {
// 	answer := 0
//
// 	fmt.Println("Part 2:", answer)
// }
