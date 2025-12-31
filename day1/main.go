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

func part1(inputs []string, dial_pointer int) int {
	password := 0

	for _, input := range inputs {
		num, _ := strconv.Atoi(string(input[1:]))
		if string(input[0]) == "L" {
			dial_pointer -= num
		} else if string(input[0]) == "R" {
			dial_pointer += num
		}

		dial_pointer %= 100

		if dial_pointer == 0 {
			password += 1
		}
	}

	return password
}

func part2(inputs []string, dial_pointer int) int {
	password := 0

	for _, input := range inputs {
		num, _ := strconv.Atoi(string(input[1:]))
		password += int(num / 100)
		if string(input[0]) == "L" {
			dial_pointer -= num
		} else if string(input[0]) == "R" {
			dial_pointer += num
		}
		dial_pointer %= 100

		if dial_pointer == 0 {
			password += 1
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
