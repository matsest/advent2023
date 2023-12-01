package main

import (
	"fmt"
	"strconv"

	"github.com/matsest/advent2023/utils"
)

func getNumbers(input string) (numbers []int) {
	for _, v := range input {
		number, err := strconv.Atoi(string(v))
		if err == nil {
			numbers = append(numbers, number)
		}
	}
	return numbers
}

func combineNumbers(a int, b int) int {
	aString := fmt.Sprint(a)
	bString := fmt.Sprint(b)
	combinedString := aString + bString
	combined, _ := strconv.Atoi(combinedString)
	return combined
}

func p1(input []string) int {
	sum := 0
	for _, line := range input {
		numbers := getNumbers(line)
		first := numbers[0]
		last := numbers[len(numbers)-1]
		combined := combineNumbers(first, last)
		sum += combined
	}
	return sum
}

func p2(input []string) int {
	return 2
}

func main() {
	puzzle_input, _ := utils.ReadLines("./input.txt")
	fmt.Println(p1(puzzle_input))
	//fmt.Println(p2(puzzle_input))
}
