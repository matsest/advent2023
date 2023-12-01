package main

import (
	"fmt"
	"regexp"
	"sort"
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

var stringsToNumbersMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func getNumbers2(input string) (numbers []int) {

	// find the index the match is found in
	numberMap := make(map[int]int)
	for k, v := range stringsToNumbersMap {
		m := regexp.MustCompile(k)
		// consider multiple matches
		rs := m.FindAllStringIndex(input, -1)
		for _, r := range rs {
			if len(r) != 0 {
				numberMap[r[0]] = v
			}
		}
	}
	// sort the keys of the map
	keys := make([]int, 0, len(numberMap))
	for k := range numberMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	// make a list of numbers from sorted order
	for _, k := range keys {
		numbers = append(numbers, numberMap[k])
	}
	return numbers
}

func p2(input []string) int {
	sum := 0
	for _, line := range input {
		numbers := getNumbers2(line)
		first := numbers[0]
		last := numbers[len(numbers)-1]
		combined := combineNumbers(first, last)
		sum += combined
	}
	return sum
}

func main() {
	puzzle_input, _ := utils.ReadLines("./input.txt")
	fmt.Println(p1(puzzle_input))
	fmt.Println(p2(puzzle_input))
}
