package main

import (
	"fmt"

	"github.com/matsest/advent2023/utils"
)

func p1(galaxyMap map[[2]int]int) int {
	coords := [][2]int{}
	for k := range galaxyMap {
		coords = append(coords, k)
	}
	distance := distanceSum(coords)
	return distance
}

func p2(galaxyMap map[[2]int]int) int {
	return p1(galaxyMap)
}

func parseInput(lines []string, expansionCoeff int) (img Image, galaxyMap map[[2]int]int) {
	// Go through rows and index which to expand
	emptyRowsIndexes := []int{}
	for y, line := range lines {
		row := []string{}
		isEmptyRowCount := 0
		for _, c := range line {
			row = append(row, string(c))
			if string(c) == "." {
				isEmptyRowCount += 1
			}
		}
		// check row to expand
		img = append(img, row)
		if isEmptyRowCount == len(row) {
			emptyRowsIndexes = append(emptyRowsIndexes, y)
		}
	}

	// Go through columns and index which to expand
	emptyColsIndexes := []int{}
	for x := 0; x < len(img[0]); x++ {
		isEmptyColCount := 0
		for y := 0; y < len(img); y++ {
			if img[y][x] == "." {
				isEmptyColCount += 1
			}
		}
		// check column to expand
		if isEmptyColCount == len(img) {
			emptyColsIndexes = append(emptyColsIndexes, x)
		}
	}
	//fmt.Println("empty row", emptyRowsIndexes)
	//fmt.Println("empty cols", emptyColsIndexes)

	// Generate map of all galaxies with expansions
	galaxyMap = make(map[[2]int]int)
	galaxyIndex := 1
	multiplier := expansionCoeff
	for y, row := range img {
		ydiff := countBelow(emptyRowsIndexes, y)
		for x, c := range row {
			xdiff := countBelow(emptyColsIndexes, x)
			if c == "#" {
				galaxyMap[[2]int{y + ydiff*(multiplier-1), x + xdiff*(multiplier-1)}] = galaxyIndex
				galaxyIndex++
			}
		}
	}
	//fmt.Println(galaxyMap)
	return img, galaxyMap
}

func countBelow(indexes []int, index int) (count int) {
	for _, i := range indexes {
		if i < index {
			count += 1
		}
	}
	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func distanceSum(coordinates [][2]int) (distance int) {
	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			distance += manhattanDistance(coordinates[i][1], coordinates[i][0], coordinates[j][1], coordinates[j][0])
		}
	}
	return distance
}

type Image [][]string

func (i Image) Print() {
	for y := range i {
		for x := range i[y] {
			fmt.Print(i[y][x], " ")
		}
		fmt.Println()
	}
	fmt.Println("Width", len(i[0]))
	fmt.Println("Height", len(i))
}

func main() {
	puzzle_input, _ := utils.ReadLines("input.txt")
	_, galaxies := parseInput(puzzle_input, 2)
	//img.Print()
	fmt.Println(p1(galaxies))
	_, galaxies2 := parseInput(puzzle_input, 1000000)
	//img.Print()
	//fmt.Println(galaxies2)
	fmt.Println(p2(galaxies2))
}
