package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/matsest/advent2023/utils"
)

func p1(i []string, nodes map[string]([2]string)) int {
	currentVal := "AAA"
	currentIndex := 0
	count := 0

	for currentVal != "ZZZ" {
		currentIndex = currentIndex % len(i)
		var indexToChoose int
		if i[currentIndex] == "L" {
			indexToChoose = 0
		} else if i[currentIndex] == "R" {
			indexToChoose = 1
		}
		currentVal = nodes[currentVal][indexToChoose]
		currentIndex += 1
		count += 1
	}
	return count
}

func p2(input string) int {
	return 2
}

func parseInput(lines []string) (i []string, nodes map[string]([2]string)) {
	i = strings.Split(lines[0], "")
	nodes = map[string]([2]string){}

	for _, line := range lines[2:] {
		re := regexp.MustCompile(`\b[A-Z]+\b`)
		words := re.FindAllString(line, -1)
		nodes[words[0]] = [2]string{words[1], words[2]}
	}
	return
}

func main() {
	puzzle_input, _ := utils.ReadLines("input.txt")
	instructions, nodes := parseInput(puzzle_input)
	fmt.Println(p1(instructions, nodes))
	fmt.Println(p2("2"))
}
