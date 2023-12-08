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

func p2(i []string, nodes map[string]([2]string)) int {
	currentNodes := []string{}

	for k := range nodes {
		parts := strings.Split(k, "")
		if parts[2] == "A" {
			currentNodes = append(currentNodes, k)
		}
	}

	zCounts := []int{}
	currentIndex := 0

	//fmt.Println("startnodes", currentNodes)
Nodes:
	for _, n := range currentNodes {
		currentVal := n
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
			if strings.Split(currentVal, "")[2] == "Z" {
				//fmt.Println(n, currentVal, count)
				zCounts = append(zCounts, count)
				continue Nodes
			}
		}
	}
	//fmt.Println(currentNodes, zCounts)
	return lcm(zCounts...)
}

// greatest common divisor
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// least common multiple
func lcm(integers ...int) int {
	result := integers[0] * integers[1] / gcd(integers[0],integers[1])
	if len(integers) > 2 {
		for i := 0; i < len(integers); i++ {
			result = lcm(result, integers[i])
		}
	}
	return result
}

func parseInput(lines []string) (i []string, nodes map[string]([2]string)) {
	i = strings.Split(lines[0], "")
	nodes = map[string]([2]string){}

	for _, line := range lines[2:] {
		//fmt.Println(line)
		re := regexp.MustCompile(`\b\w{3}\b`)
		words := re.FindAllString(line, -1)
		nodes[words[0]] = [2]string{words[1], words[2]}
	}
	return
}

func main() {
	puzzle_input, _ := utils.ReadLines("input.txt")
	instructions, nodes := parseInput(puzzle_input)
	fmt.Println(p1(instructions, nodes))
	fmt.Println(p2(instructions, nodes))
}
