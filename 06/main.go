package main

import (
	"fmt"
	"strings"

	"github.com/matsest/advent2023/utils"
)

func p1(races []Race) (sum int) {
	sum = 1
	for _, r := range races {
		r.CalculateResults()
		rCount := 0
		for _, res := range r.results {
			if res.distance > r.record {
				rCount += 1
			}
		}
		sum*=rCount
	}
	return sum
}

func p2(input string) int {
	return 2
}

type Race struct {
	totalTime    int
	record  int
	results []RaceResult
}

type RaceResult struct {
	holdTime int
	totalTime     int
	distance int
}

func (r *Race) CalculateResults(){
	for i := 1; i <= r.totalTime; i++ {
		holdTime := i //ms
		timeToTravel := r.totalTime - holdTime //ms
		startSpeed := holdTime //mm/ms
		distance := startSpeed*timeToTravel
		//fmt.Println("holdTime:", i, "distance", distance)
		r.results = append(r.results, RaceResult{holdTime: holdTime, totalTime: r.totalTime, distance: distance})
	}
}

func parseInput(lines []string) (races []Race) {
	times, _ := utils.SliceAtoi(strings.Fields(lines[0])[1:])
	records, _ := utils.SliceAtoi(strings.Fields(lines[1])[1:])
	for i := range times {
		races = append(races, Race{times[i], records[i], []RaceResult{}})
	}
	return races
}

func main() {
	puzzle_input, _ := utils.ReadLines("input.txt")
	races := parseInput(puzzle_input)
	fmt.Println(p1(races))
	fmt.Println(p2("2"))
}
