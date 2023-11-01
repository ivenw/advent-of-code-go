package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	module_masses := parse(content)

	partOne(module_masses)
	partTwo(module_masses)
}

func parse(input []byte) []int {
	text := string(input)
	lines := strings.Fields(text)
	module_masses := make([]int, len(lines))
	for i, line := range lines {
		module_masses[i], _ = strconv.Atoi(line)
	}
	return module_masses
}

func partOne(module_masss []int) {

	total_fuel := 0
	for _, mass := range module_masss {
		total_fuel += calcFuel(mass)
	}

	answer := total_fuel
	fmt.Println("Part 1:", answer)
}

func partTwo(module_masss []int) {

	total_fuel := 0
	for _, mass := range module_masss {
		calc_fuel := calcFuel(mass)
		for calc_fuel > 0 {
			total_fuel += calc_fuel
			calc_fuel = calcFuel(calc_fuel)
		}
	}

	answer := total_fuel
	fmt.Println("Part 2:", answer)
}

func calcFuel(mass int) int {
	return mass/3 - 2
}
