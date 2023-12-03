package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/matsest/advent2023/utils"
)

func p2(lines []string) (sum int) {

	// Generate maps to allow lazy traversing
	var numbers schematicNodes = make(schematicNodes)
	var symbols schematicNodes = make(schematicNodes)
	var gears schematicNodes = make(schematicNodes)

	// Populate maps
	for y, line := range lines {
		re := regexp.MustCompile(`\d+`)
		numbersMatch := re.FindAllStringIndex(line, -1)
		for _, xs := range numbersMatch {
			val, _ := strconv.Atoi(line[xs[0]:xs[1]])
			numbers[[2]int{xs[0], y}] = val
		}
		// symbols
		re2 := regexp.MustCompile(`[^0-9.!]`) // find symbols
		symbolsMatch := re2.FindAllStringIndex(line, -1)
		for _, xs := range symbolsMatch {
			symbols[[2]int{xs[0], y}] = 1
		}
		// gears
		re3 := regexp.MustCompile(`\*`) // find gears
		gearsMatch := re3.FindAllStringIndex(line, -1)
		for _, xs := range gearsMatch {
			gears[[2]int{xs[0], y}] = 1
		}

	}

	// Remove part numbers..
	for k, n := range numbers {
		//fmt.Println(k, n)
		isPart := false
	number: // break point for loop
		for x := k[0] - 1; x <= k[0]+len(fmt.Sprint(n)); x++ {
			for y := k[1] - 1; y <= k[1]+1; y++ {
				if symbols[[2]int{x, y}] != 0 {
					//fmt.Println(n, "is part")
					isPart = true
					break number // continue to next number
				}
			}
		}
		if !isPart {
			//fmt.Println(n, "not part, deleting..")
			delete(numbers, k)
		}
	}

	// Go through possible gears
	for k  := range gears {
		//isGear := false
		count := 0
		parts := []int{}
		for x := k[0] - 3; x <= k[0]+1; x++ { // note: start three behind to account for numbers
			for y := k[1] - 1; y <= k[1]+1; y++ {
				if numbers[[2]int{x, y}] != 0 {
					count += 1
					parts = append(parts, numbers[[2]int{x,y}])
					//fmt.Println(numbers[[2]int{x,y}])
				}
			}
		}
		if count == 2 {
			//fmt.Println(k, "is gear")
			//isGear = true
			//fmt.Println("sum", parts[0]*parts[1])
			sum += parts[0]*parts[1]
		}
		//if !isGear {
		//	fmt.Println(k, "not gear,")
		//}
		//fmt.Println(count, "count,")

	}
	return sum
}

type schematicNodes map[[2]int]int

func p1(lines []string) (sum int) {

	// Generate maps to allow lazy traversing
	var numbers schematicNodes = make(schematicNodes)
	var symbols schematicNodes = make(schematicNodes)

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
	puzzle_input, _ := utils.ReadLines("input.txt")
	fmt.Println(p1(puzzle_input))
	fmt.Println(p2(puzzle_input))
}
