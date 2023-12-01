package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputString := string(input)
	part01 := Part01(inputString)
	fmt.Println(part01)
	part02 := Part02(inputString)
	fmt.Println(part02)
}
