package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matsest/advent2023/utils"
)

func p1(cards []Card) (sum int) {
	for _, c := range cards {
		sum += c.Points()
	}
	return sum
}

func p2(games []Card) int {
	return 2
}

type Card struct {
	index int
	winningNumbers []int
	numbers []int
}

func (g Card) Points() (sum int){
	for _, c := range g.numbers {
		for _, wc := range g.winningNumbers {
			if c == wc {
				if sum > 0 {
					sum *= 2
				}
				if sum == 0 {
					sum += 1
				}
				//fmt.Println("winning number ", wc, "sum now ", sum)
			}
		}
	}
	//fmt.Println(g.index, sum)
	return sum
}

func parseInput(lines []string) (cards []Card){
	for _, line := range lines {
		//fmt.Println(line)
		game := Card{}
		parts := strings.Split(line, ": ")

		// set index
		indexS := strings.Fields(parts[0])[1]
		indexN, _ := strconv.Atoi(indexS)
		game.index = indexN

		// set winning cards and cards
		numbers := strings.Split(parts[1], " | ")
		winningCards := strings.Fields(numbers[0])
		gameCards := strings.Fields(numbers[1])
		game.winningNumbers, _ = utils.SliceAtoi(winningCards)
		game.numbers, _ = utils.SliceAtoi(gameCards)

		cards = append(cards, game)
	}
	return cards
}

func main(){
	puzzle_input, _ := utils.ReadLines("input.txt")
	cards := parseInput(puzzle_input)
	fmt.Println(p1(cards))
	fmt.Println(p2(cards))
}