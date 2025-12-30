package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input_array() ([]string) {
	data, err := os.ReadFile("input.txt")

	if (err != nil){
		panic(err)
	}

	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	return inputs
}

func main(){
	inputs := get_input_array()
	dial_pointer := 50
	password := 0

	for _, input := range inputs{
		if string(input[0]) == "L"{
			num, _ := strconv.Atoi(string(input[1:]))
			dial_pointer -= num 
			} else if (string(input[0]) == "R"){
				num, _ := strconv.Atoi(string(input[1:]))
				dial_pointer += num
		}

		dial_pointer %= 100
		
		if dial_pointer == 0{
			password += 1
		}
	}

	fmt.Println(password)
}