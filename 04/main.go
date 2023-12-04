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

func p2(cards []Card) int {
	// Find out number of copies
	copies := make(map[int]int)
	for _, c := range cards {
		current := c.index
		currentCopies := copies[current]
		nmatchingNumbers := c.NMatchingNumbers()
		for i := current + 1; i <= current + nmatchingNumbers; i++ {
			copies[i] += (1 + currentCopies)
		}
	}

	// Add initial cards and copies
	total := len(cards)
	for i := 1; i <= len(cards); i++{
		total += copies[i]
	}
	return total
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
			}
		}
	}
	return sum
}

func (g Card) NMatchingNumbers() (sum int){
	for _, c := range g.numbers {
		for _, wc := range g.winningNumbers {
			if c == wc {
					sum += 1
			}
		}
	}
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