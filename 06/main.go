package main

import (
	"fmt"
	"strconv"
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
		sum *= rCount
	}
	return sum
}

func p2(r Race) (sum int) {
	r.CalculateResults()
	for _, res := range r.results {
		if res.distance > r.record {
			sum += 1
		}
	}
	return sum
}

type Race struct {
	totalTime int
	record    int
	results   []RaceResult
}

type RaceResult struct {
	holdTime  int
	totalTime int
	distance  int
}

func (r *Race) CalculateResults() {
	for i := 1; i <= r.totalTime; i++ {
		holdTime := i                          //ms
		timeToTravel := r.totalTime - holdTime //ms
		startSpeed := holdTime                 //mm/ms
		distance := startSpeed * timeToTravel  //mm
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

func parseInput2(lines []string) (race Race) {
	times, _ := strconv.Atoi(strings.Join(strings.Fields(lines[0])[1:], ""))
	records, _ := strconv.Atoi(strings.Join(strings.Fields(lines[1])[1:], ""))
	return Race{times, records, []RaceResult{}}
}

func main() {
	puzzle_input, _ := utils.ReadLines("input.txt")
	races := parseInput(puzzle_input)
	fmt.Println(p1(races))
	races2 := parseInput2(puzzle_input)
	fmt.Println(p2(races2))
}
