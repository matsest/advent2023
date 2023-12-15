package main

import (
	"fmt"
	"strings"

	"github.com/matsest/advent2023/utils"
)

func p1(steps []string) (sum int) {
	for _, step := range steps  {
		sum += hash(step)
	}
	return sum
}

func p2(input string) int {
	return 2
}

func hash(step string) int {
	value := 0
	for _, c := range step {
		asciiCode := int(c)
		value += asciiCode
		value *= 17
		value = value % 256
	}
	return value
}

func parseInput(lines []string)(steps []string){
	for _, line := range lines {
		parts := strings.Split(line, ",")
		steps = append(steps, parts...)
	}
	return steps
}

func main(){
	puzzle_input, _ := utils.ReadLines("input.txt")
	steps := parseInput(puzzle_input)
	fmt.Println(p1(steps))
	fmt.Println(p2("2"))
}