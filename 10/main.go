package main

import (
	"fmt"
	"strings"

	"github.com/matsest/advent2023/utils"
)

func p1(diagram Diagram, start Position) (largest int) {
	directionstoCheck := diagram.checkNeighbours(start)
	largest = 0
	for _, dir := range directionstoCheck[0:1] { // we only check one direction - go around the loop...
		largest = diagram.findLongestDistance(start, dir)
	}
	return largest
}

func (diagram Diagram) findLongestDistance(position Position, direction Direction) (distance int) {
	distance = 0
	currentPos := position
	directionsToCheck := []Direction{direction}

	for {
		//fmt.Println("currentpos", currentPos, diagram.Coordinates[currentPos].instruction, "looking", directionsToCheck[0])
		nextInstruction := diagram.getInstructionFromDirection(currentPos, directionsToCheck[0])
		nextPosition, _ := diagram.Move(currentPos, nextInstruction, directionsToCheck[0])
		//fmt.Println("moving to pos", nextPosition, diagram.Coordinates[nextPosition].instruction, diagram.Coordinates[nextPosition].distanceFromS)
		if diagram.Coordinates[nextPosition].instruction == "S" {
			//fmt.Println("next is S")
			break
		}
		currentPos = Position{nextPosition[0], nextPosition[1]}
		directionsToCheck = diagram.checkNeighbours(currentPos)
		if len(directionsToCheck) == 0 {
			//fmt.Println("no directions to hceck")
			break
		}
		//fmt.Println("can now go to ", directionsToCheck)
	}
	return diagram.Coordinates[position].distanceFromS / 2 // return halfway point
}

func (d Diagram) getInstructionFromDirection(position Position, direction Direction) (instruction string) {
	switch direction {
	case Right:
		instruction = d.Coordinates[[2]int{position[0], position[1] + 1}].instruction
	case Left:
		instruction = d.Coordinates[[2]int{position[0], position[1] - 1}].instruction
	case Up:
		instruction = d.Coordinates[[2]int{position[0] - 1, position[1]}].instruction
	case Down:
		instruction = d.Coordinates[[2]int{position[0] + 1, position[1]}].instruction
	}
	return instruction
}

func (d Diagram) Move(position Position, instruction string, direction Direction) (newPosition Position, stepsMoved int) {
	var xdiff, ydiff int
	stepsMoved = 2 // always move from x => x + past instruction

	switch instruction {
	case "|":
		xdiff = 0
		if direction == Down {
			ydiff = 2
		} else if direction == Up {
			ydiff = -2
		}
	case "-":
		ydiff = 0
		if direction == Right {
			xdiff = 2
		} else if direction == Left {
			xdiff = -2
		}
	case "L":
		if direction == Down {
			ydiff = 1
			xdiff = 1
		} else if direction == Left {
			ydiff = -1
			xdiff = -1
		}
	case "J":
		if direction == Down {
			ydiff = 1
			xdiff = -1
		} else if direction == Right {
			ydiff = -1
			xdiff = 1
		}
	case "7":
		if direction == Up {
			ydiff = -1
			xdiff = -1
		} else if direction == Right {
			ydiff = 1
			xdiff = 1
		}
	case "F":
		if direction == Up {
			xdiff = 1
			ydiff = -1
		} else if direction == Left {
			xdiff = -1
			ydiff = 1
		}
	case ".":
		xdiff = 0
		ydiff = 0
	}
	newPosition = Position{position[0] + ydiff, position[1] + xdiff}

	// set position of current instruction as visited
	var positionOfInstruction Position
	switch direction {
	case Left:
		positionOfInstruction = Position{position[0], position[1] - 1}
	case Right:
		positionOfInstruction = Position{position[0], position[1] + 1}
	case Up:
		positionOfInstruction = Position{position[0] - 1, position[1]}
	case Down:
		positionOfInstruction = Position{position[0] + 1, position[1]}
	}
	if c, ok := d.Coordinates[positionOfInstruction]; ok {
		c.visited = true
		c.distanceFromS = d.Coordinates[position].distanceFromS + 1 // add one
		d.Coordinates[positionOfInstruction] = c
	}

	// set newPosition as visited
	if c, ok := d.Coordinates[newPosition]; ok {
		c.visited = true
		c.distanceFromS = d.Coordinates[position].distanceFromS + stepsMoved // add steps
		d.Coordinates[newPosition] = c
	}

	return newPosition, stepsMoved
}

func (d Diagram) checkNeighbours(position Position) (directionsToCheck []Direction) {
	up := d.Coordinates[[2]int{position[0] - 1, position[1]}].instruction
	upVisited := d.Coordinates[[2]int{position[0] - 1, position[1]}].visited
	down := d.Coordinates[[2]int{position[0] + 1, position[1]}].instruction
	downVisited := d.Coordinates[[2]int{position[0] + 1, position[1]}].visited
	left := d.Coordinates[[2]int{position[0], position[1] - 1}].instruction
	leftVisited := d.Coordinates[[2]int{position[0], position[1] - 1}].visited
	right := d.Coordinates[[2]int{position[0], position[1] + 1}].instruction
	rightVisited := d.Coordinates[[2]int{position[0], position[1] + 1}].visited

	current := d.Coordinates[position].instruction

	if (current == "|" || current == "S" || current == "J" || current == "L") && !upVisited && (up == "|" || up == "F" || up == "7") {
		directionsToCheck = append(directionsToCheck, Up)
	}
	if (current == "|" || current == "S" || current == "7" || current == "F") && !downVisited && (down == "|" || down == "L" || down == "J") {
		directionsToCheck = append(directionsToCheck, Down)
	}
	if (current == "-" || current == "S" || current == "F" || current == "L") && !rightVisited && (right == "-" || right == "J" || right == "7") {
		directionsToCheck = append(directionsToCheck, Right)
	}
	if (current == "-" || current == "S" || current == "J" || current == "7") && !leftVisited && (left == "-" || left == "L" || left == "F") {
		directionsToCheck = append(directionsToCheck, Left)
	}

	// will return empty slice if no neighbours are found
	return directionsToCheck
}

func parseInput(lines []string) (diagram Diagram, start Position) {
	diagram = Diagram{}
	diagram.Coordinates = make(map[[2]int]Coordinate)
	for y := range lines {
		parts := strings.Split(lines[y], "")
		for x := range lines[y] {
			diagram.Coordinates[[2]int{y, x}] = Coordinate{Position{y, x}, parts[x], false, -1}
			if parts[x] == "S" {
				diagram.Coordinates[[2]int{y, x}] = Coordinate{Position{y, x}, parts[x], true, 0}
				start = Position{y, x}
			}
		}
	}
	diagram.Height = len(lines)
	diagram.Width = len(lines[0])
	return diagram, start
}

// TYPES
type Diagram struct {
	Coordinates map[[2]int]Coordinate
	Height      int
	Width       int
}

type Position [2]int

type Coordinate struct {
	position      Position // y,x
	instruction   string
	visited       bool
	distanceFromS int
}

// type for directions
type Direction string

const (
	Left  Direction = "left"
	Right Direction = "right"
	Up    Direction = "up"
	Down  Direction = "down"
)

func (d Diagram) Print() {
	for i := 0; i < d.Height; i++ {
		for j := 0; j < d.Width; j++ {
			current := d.Coordinates[Position{i, j}]
			if current.visited {
				fmt.Print(current.instruction, "")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	for i := 0; i < d.Height; i++ {
		for j := 0; j < d.Width; j++ {
			current := d.Coordinates[Position{i, j}]
			if current.visited {
				fmt.Print(current.distanceFromS, " ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}

}

func main() {
	puzzle_input, _ := utils.ReadLines("input.txt")
	//diagram, start := parseInput(puzzle_input)
	//fmt.Println(diagram, start)
	//fmt.Println(p1(diagram, start))
	diagram, start := parseInput(puzzle_input)
	fmt.Println(p1(diagram, start))
	//diagram.Print()
}
