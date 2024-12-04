package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	levels []int
}

type Data []Report

func main() {
	content_raw, _ := os.ReadFile("input.txt")
	data := parse(string(content_raw))

	answer_one := partOne(data)
	fmt.Println("Part 1:", answer_one)

	answer_two := partTwo(data)
	fmt.Println("Part 2:", answer_two)
}

func parse(content_raw string) Data {
	scanner := bufio.NewScanner(strings.NewReader(string(content_raw)))

	data := Data{}
	for scanner.Scan() {
		line := scanner.Text()
		levels_string := strings.Split(line, " ")
		levels_int := make([]int, len(levels_string))

		for i, level := range levels_string {
			levelInt, _ := strconv.Atoi(level)
			levels_int[i] = levelInt
		}

		data = append(data, Report{levels_int})
	}

	return data
}

func partOne(input Data) int {
	answer := 0

REPORT_LOOP:
	for _, report := range input {
		first := report.levels[0]
		second := report.levels[1]
		diff := first - second
		var increasing bool
		if diff >= 1 && diff <= 3 {
			increasing = true
		} else if diff <= -1 && diff >= -3 {
			increasing = false
		} else {
			continue
		}

		prev := report.levels[0]
		for _, level := range report.levels[1:] {
			diff := prev - level
			if (increasing && (diff < 1 || diff > 3)) || (!increasing && (diff > -1 || diff < -3)) {
				continue REPORT_LOOP
			}
			prev = level
		}

		answer++
	}

	return answer
}

func partTwo(input Data) int {
	answer := 0

REPORT_LOOP:
	for _, report := range input {
		first := report.levels[0]
		second := report.levels[1]
		diff := first - second
		var increasing bool
		if diff >= 1 && diff <= 3 {
			increasing = true
		} else if diff <= -1 && diff >= -3 {
			increasing = false
		} else {
			continue
		}

		prev := report.levels[0]
		for _, level := range report.levels[1:] {
			diff := prev - level
			if (increasing && (diff < 1 || diff > 3)) || (!increasing && (diff > -1 || diff < -3)) {
				continue REPORT_LOOP
			}
			prev = level
		}

		answer++
	}

	return answer
}
