package main

import (
	"fmt"
	"strings"

	"github.com/matsest/advent2023/utils"
)

func p1(dataset [][]int) (sum int) {
	for _, history := range dataset {
		extrapolated := extraPolate(history)
		sum += extrapolated[len(extrapolated)-1]
	}
	return sum
}

func p2(input string) int {
	return 2
}

func extraPolate(history []int) (extrapolated []int){
	diffs := [][]int{} // need to find out how large
	diffs = append(diffs, history)
	isAllZeros := false
	//fmt.Println("diffs:", diffs)

	for !isAllZeros {
		for i := 1; i < len(history); i++ { // need to find how many
			tmpdiffs := make([]int, len(history)-i)
			//fmt.Println(i)
			for j := 1; j <= len(tmpdiffs); j++ {
				tmpdiffs[j-1] = diffs[i-1][j] - diffs[i-1][j-1]
			}
			diffs = append(diffs, tmpdiffs)
			if allZeros(tmpdiffs) { // need to have a check
				isAllZeros = true
				diffs[i] = append(diffs[i], 0)
				break
			}
		}
	}
	//fmt.Println("diffs:", diffs, len(diffs))
	//fmt.Println("start to add")
	for i:= len(diffs)-1; i > 0; i-- {
		//fmt.Println(i)
		last := diffs[i][len(diffs[i])-1]
		prev := diffs[i-1][len(diffs[i-1])-1]
		diffs[i-1] = append(diffs[i-1], prev + last)
	}
	//fmt.Println("diffs:", diffs)
	return diffs[0]
}

func allZeros(slice []int) bool {
    for _, v := range slice {
        if v != 0 {
            return false
        }
    }
    return true
}

func parseInput(lines []string) (dataset [][]int) {
	for _, line := range lines {
		//fmt.Println(line)
		values, _ := utils.SliceAtoi(strings.Fields(line))
		dataset = append(dataset, values)
	}
	return dataset
}

func main() {
	puzzle_input, _ := utils.ReadLines("input.txt")
	dataset := parseInput(puzzle_input)
	fmt.Println(p1(dataset))
	fmt.Println(p2("2"))
}
