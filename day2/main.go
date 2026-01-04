package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputArray() []string {
	data, err := os.ReadFile("input.txt")
	// data, err := os.ReadFile("test_input.txt")

	if err != nil {
		panic(err)
	}

	inputs := strings.Split(strings.TrimSpace(string(data)), ",")

	return inputs
}

func getInvalidIdsP1(inputs []string) int {
	result := 0
	for _, input := range inputs {
		idRange := strings.Split(input, "-")
		start, _ := strconv.Atoi(string(idRange[0]))
		stop, _ := strconv.Atoi(string(idRange[1]))
		for i := start; i <= stop; i++ {
			numString := strconv.Itoa(i)
			n := len(numString)
			if n%2 == 0 && numString[:n/2] == numString[n/2:] {
				result += i
			}
		}
	}

	return result
}

func main() {
	inputs := getInputArray()
	fmt.Println(getInvalidIdsP1(inputs))
}
