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

func part1() {
	inputs := get_input_array()
	dial_pointer := 50
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

	fmt.Println(password)
}

func part2() {
	inputs := get_input_array()
	dial_pointer := 50
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

	fmt.Println(password)
}

func main() {
	part1()
	part2()
}
