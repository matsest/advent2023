package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/matsest/advent2023/utils"
)

type schematicNodes map[[2]int]int

func resolvePartNumbers(lines []string) schematicNodes {
	// Generate maps to allow lazy traversing without index nightmare
	var numbers schematicNodes = make(schematicNodes)
	var symbols schematicNodes = make(schematicNodes)

	// Populate maps
	for y, line := range lines {
		// numbers
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
	}

	// Remove non-part numbers..
Numbers:
	for k, n := range numbers {
		isPart := false
		for x := k[0] - 1; x <= k[0]+len(fmt.Sprint(n)); x++ {
			for y := k[1] - 1; y <= k[1]+1; y++ {
				if symbols[[2]int{x, y}] != 0 {
					isPart = true
					continue Numbers // continue to next number
				}
			}
		}
		if !isPart {
			delete(numbers, k)
		}
	}
	return numbers
}

func p1(lines []string) (sum int) {
	parts := resolvePartNumbers(lines)

	// Find sum of all part numbers
	for _, n := range parts {
		sum += n
	}
	return sum
}

func p2(lines []string) (sum int) {
	parts := resolvePartNumbers(lines)
	var gears schematicNodes = make(schematicNodes)

	// Populate map of possible gears
	for y, line := range lines {
		// gears
		re3 := regexp.MustCompile(`\*`) // find gears
		gearsMatch := re3.FindAllStringIndex(line, -1)
		for _, xs := range gearsMatch {
			gears[[2]int{xs[0], y}] = 1
		}
	}
	// Go through possible gears
	for k := range gears {
		count := 0
		partsToAdd := []int{}
		for x := k[0] - 3; x <= k[0]+1; x++ { // note: start three behind to account for three digit numbers
			for y := k[1] - 1; y <= k[1]+1; y++ {
				if parts[[2]int{x, y}] != 0 {
					count += 1
					partsToAdd = append(partsToAdd, parts[[2]int{x, y}])
				}
			}
		}
		if count == 2 { // is gear!
			sum += partsToAdd[0] * partsToAdd[1]
		}
	}
	return sum
}

func main() {
	puzzle_input, _ := utils.ReadLines("input.txt")
	fmt.Println(p1(puzzle_input))
	fmt.Println(p2(puzzle_input))
}
