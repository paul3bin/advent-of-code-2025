package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input_array() []string {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	return inputs
}

func part1(inputs []string, dialPointer int) int {
	password := 0

	for _, input := range inputs {
		steps, _ := strconv.Atoi(string(input[1:]))
		dir := input[0]

		if dir == 'L' {
			dialPointer -= steps
		} else { // R
			dialPointer += steps
		}

		dialPointer %= 100

		if dialPointer == 0 {
			password++
		}
	}

	return password
}

func part2(inputs []string, dialPointer int) int {
	password := 0

	for _, input := range inputs {
		steps, _ := strconv.Atoi(input[1:])
		prev := dialPointer
		dir := input[0]

		if steps >= 100 {
			password += steps / 100
			steps %= 100
		}

		if dir == 'L' {
			dialPointer -= steps
		} else { // 'R'
			dialPointer += steps
		}

		if prev != 0 && (dialPointer < 0 || dialPointer > 100) {
			password++
		}

		dialPointer = ((dialPointer % 100) + 100) % 100

		if dialPointer == 0 {
			password++
		}
	}

	return password
}

func main() {
	inputs := get_input_array()
	dial_pointer := 50
	fmt.Printf("Part1 password: %d\n", part1(inputs, dial_pointer))
	fmt.Printf("Part2 password: %d\n", part2(inputs, dial_pointer))
}
