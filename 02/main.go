package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matsest/advent2023/utils"
)

type Game struct {
	index int
	sets []Set
}

type Set struct {
	Red int
	Green int
	Blue int
}

var (
	RED_LIMIT int = 12
	GREEN_LIMIT int = 13
	BLUE_LIMIT int = 14
)

func possibleGame(game Game) bool {
	for _, set := range game.sets {
		if(set.Red > RED_LIMIT || set.Blue > BLUE_LIMIT || set.Green > GREEN_LIMIT){
			return false
		}
	}
	return true
}

func parseInput(lines []string) (games []Game) {
	for _, line := range lines {
		//fmt.Println(line)
		g := Game{}
		parts := strings.Split(line, ": ")
		index, _ := strconv.Atoi((strings.Split(parts[0], " ")[1]))
		g.index = index
		sets := strings.Split(parts[1], "; ")
		//fmt.Println(index)
		for _, set := range sets {
			newSet := Set{}
			parts := strings.Split(set, ", ")
			//fmt.Println(parts)
			for _, p := range parts {
				//fmt.Println(p)
				cubes := strings.Fields(p)
				cubeCount, _ := strconv.Atoi(cubes[0])
				cubeColor := cubes[1]
				switch cubeColor {
				case "green":
					newSet.Green = cubeCount
				case "red":
					newSet.Red = cubeCount
				case "blue":
					newSet.Blue = cubeCount
				}
			}
			g.sets = append(g.sets, newSet)
		}
		games = append(games, g)
	}
	return games
}

func p1(games []Game) (sum int) {
	for _, g := range games {
		if possibleGame(g){
			sum += g.index
		}
	}
	return sum
}

func p2(input string) int {
	return 2
}

func main(){
	puzzle_input, _ := utils.ReadLines("input.txt")
	games := parseInput(puzzle_input)
	fmt.Println(p1(games))
	//fmt.Println(p2(puzzle_input))
}