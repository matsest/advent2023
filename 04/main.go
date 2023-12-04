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

func (c Card) Points() (sum int){
	for _, n := range c.numbers {
		for _, wn := range c.winningNumbers {
			if n == wn {
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

func (c Card) NMatchingNumbers() (sum int){
	for _, n := range c.numbers {
		for _, wn := range c.winningNumbers {
			if n == wn {
					sum += 1
			}
		}
	}
	return sum
}

func parseInput(lines []string) (cards []Card){
	for _, line := range lines {
		//fmt.Println(line)
		card := Card{}
		parts := strings.Split(line, ": ")

		// set index
		indexS := strings.Fields(parts[0])[1]
		indexN, _ := strconv.Atoi(indexS)
		card.index = indexN

		// set winning numbers and numbers
		numbers := strings.Split(parts[1], " | ")
		winningNumbers := strings.Fields(numbers[0])
		cardNumbers := strings.Fields(numbers[1])
		card.winningNumbers, _ = utils.SliceAtoi(winningNumbers)
		card.numbers, _ = utils.SliceAtoi(cardNumbers)

		cards = append(cards, card)
	}
	return cards
}

func main(){
	puzzle_input, _ := utils.ReadLines("input.txt")
	cards := parseInput(puzzle_input)
	fmt.Println(p1(cards))
	fmt.Println(p2(cards))
}