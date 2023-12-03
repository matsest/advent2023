package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/matsest/advent2023/utils"
)

func p2(input string) int {
	return 2
}

func p1(lines []string) (sum int) {

	// Generate maps to allow lazy traversing
	var numbers map[[2]int]int = make(map[[2]int]int)
	var symbols map[[2]int]int = make(map[[2]int]int)

	// Populate maps
	for y, line := range lines {
		re := regexp.MustCompile(`\d+`)
		numbersMatch := re.FindAllStringIndex(line, -1)
		re2 := regexp.MustCompile(`[^0-9.!]`) // find symbols
		for _, xs := range numbersMatch {
			val, _ := strconv.Atoi(line[xs[0]:xs[1]])
			numbers[[2]int{xs[0], y}] = val
		}
		symbolsMatch := re2.FindAllStringIndex(line, -1)
		for _, xs := range symbolsMatch {
			symbols[[2]int{xs[0], y}] = 1
		}
	}

	// Find all part numbers
	for k, n := range numbers {
		//fmt.Println(k, n)
		number: // break point for loop
		for x := k[0] - 1; x <= k[0]+len(fmt.Sprint(n)); x++ {
			for y := k[1] - 1; y <= k[1]+1; y++ {
				if symbols[[2]int{x, y}] != 0 {
					//fmt.Println(n, "is part")
					sum += n
					break number // continue to next number
				}
			}
		}
	}

	return sum
}

func main() {
	puzzle_input, _ := utils.ReadLines("test.txt")
	fmt.Println(p1(puzzle_input))
	fmt.Println(p2("hei"))
}
