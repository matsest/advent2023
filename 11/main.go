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

func p2(input string) int {
	return 2
}

func parseInput(lines []string) (img Image, galaxyMap map[[2]int]int) {
	// Go through rows and expand
	for _, line := range lines {
		row := []string{}
		isEmptyRowCount := 0
		for _, c := range line {
			row = append(row, string(c))
			if string(c) == "." {
				isEmptyRowCount += 1
			}

		}

		img = append(img, row)
		if isEmptyRowCount == len(row) {
			img = append(img, row)
		}
	}

	// Go through columns and expand
	for x := 0; x < len(img[0]); x++ {
		isEmptyColCount := 0
		for y := 0; y < len(img); y++ {
			if img[y][x] == "." {
				isEmptyColCount += 1
			}
		}
		// check columns to expand
		if isEmptyColCount == len(img) {
			newColumn := "."
			columnIndex := x + 1
			expandedImage := make([][]string, len(img))

			// copy in a new column
			for i, row := range img {
				expandedRow := make([]string, len(row)+1)            // Increase the row length by 1 for the new column
				copy(expandedRow[:columnIndex], row[:columnIndex])   // Copy elements before the new column
				expandedRow[columnIndex] = newColumn                 // Insert the new empty column
				copy(expandedRow[columnIndex+1:], row[columnIndex:]) // Copy elements after the new column
				expandedImage[i] = expandedRow                       // Assign the expanded row to the new slice
			}
			img = expandedImage
			x += 1
		}
	}

	// Generate map of all galaxies
	galaxyMap = make(map[[2]int]int)
	galaxyIndex := 1
	for y, row := range img {
		for x, c := range row {
			if c == "#" {
				galaxyMap[[2]int{y,x}] = galaxyIndex
				galaxyIndex ++
			}
		}
	}
	//fmt.Println(galaxyMap)
	return img, galaxyMap
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

func distanceSum(coordinates [][2]int) int {
	distance := 0

	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			dist := manhattanDistance(coordinates[i][1], coordinates[i][0], coordinates[j][1], coordinates[j][0])
			distance += dist
		}
	}

	return distance
}

type Image [][]string

func (i Image) Print(){
	for y := range i {
		for x := range i[y]{
			fmt.Print(i[y][x], " ")
		}
		fmt.Println()
	}
	fmt.Println("Width", len(i[0]))
	fmt.Println("Height", len(i))
}

func main() {
	puzzle_input, _ := utils.ReadLines("input.txt")
	_, galaxies := parseInput(puzzle_input)
	//img.Print()
	fmt.Println(p1(galaxies))
	fmt.Println(p2("2"))
}
