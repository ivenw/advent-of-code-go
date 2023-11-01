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

func parse(content_raw []byte) {
	content := string(content_raw)
}

func partOne(module_masss []int) {

	fmt.Println("Part 1:", answer)
}

// func partTwo(module_masss []int) {
//
// 	fmt.Println("Part 2:", answer)
// }
