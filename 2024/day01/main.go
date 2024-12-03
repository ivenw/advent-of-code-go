package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content_raw, _ := os.ReadFile("input.txt")

	data := parse(content_raw)

	answer_one := partOne(data)
	fmt.Println("Part 1:", answer_one)

	answer_two := partTwo(data)
	fmt.Println("Part 2:", answer_two)
}

type Data struct {
	left  []int
	right []int
}

func parse(content_raw []byte) Data {
	scanner := bufio.NewScanner(strings.NewReader(string(content_raw)))

	var left []int
	var right []int

	for scanner.Scan() {
		line := scanner.Text()
		split_line := strings.Split(line, "   ")
		left_value, _ := strconv.Atoi(split_line[0])
		right_value, _ := strconv.Atoi(split_line[1])
		left = append(left, left_value)
		right = append(right, right_value)
	}

	return Data{left, right}
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func partOne(data Data) int {
	answer := 0

	slices.Sort(data.left)
	slices.Sort(data.right)

	for i := 0; i < len(data.left); i++ {
		distance := absInt(data.left[i] - data.right[i])
		answer += distance
	}

	return answer
}

func partTwo(data Data) int {
	answer := 0

	tally := make(map[int]int)

	for _, value := range data.left {
		tally[value] = 0
	}

	for _, number := range data.right {
		if _, ok := tally[number]; ok {
			tally[number]++
		}
	}

	for _, number := range data.left {
		answer += number * tally[number]
	}

	return answer
}
