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
	content, _ := os.ReadFile("input.txt")
	data := parse(string(content))

	answer_one := partOne(data)
	fmt.Println("Part 1:", answer_one)

	answer_two := partTwo(data)
	fmt.Println("Part 2:", answer_two)
}

func parse(content string) Data {
	scanner := bufio.NewScanner(strings.NewReader(string(content)))

	return Data{}
}

func partOne(data Data) int {
	answer := 0

	return answer
}

func partTwo(data Data) int {
	answer := 0

	return answer
}
