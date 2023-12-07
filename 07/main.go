package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/matsest/advent2023/utils"
)

func p1(hands Hands) (sum int) {
	for i := range hands {
		hands[i].getType()
		//fmt.Println(hands[i])
	}
	//fmt.Println("before sort: ", hands)
	sort.Sort(hands)
	for i := range hands {
		hands[i].rank = i + 1
		sum += hands[i].rank * hands[i].bidAmount
	}
	//fmt.Println("after sort: ", hands)
	return sum
}

func p2(input string) int {
	return 2
}

var cardsToValue map[string]int = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var handTypes map[string]int = map[string]int{
	"fiveOfAKind":  7,
	"fourOfAKind":  6,
	"fullHouse":    5,
	"threeOfAKind": 4,
	"twoPairs":     3,
	"onePair":      2,
	"highCard":     1,
}

type Hand struct {
	cards     [5]string
	bidAmount int
	typeName  string
	rank      int
}

// Implement sort.Interface - Len, Less, and Swap methods
type Hands []Hand

func (hands Hands) Len() int {
	return len(hands)
}

func getStrongestHand(h1, h2 Hand) (h Hand) {
	for i := range h1.cards {
		if h1.cards[i] == h2.cards[i] {
			continue
		} else if cardsToValue[h1.cards[i]] > cardsToValue[h2.cards[i]] {
			h = h1
			break
		} else {
			h = h2
			break
		}
	}
	return h
}

func (hands Hands) Less(i, j int) bool {
	if hands[i].typeName != hands[j].typeName {
		return handTypes[hands[i].typeName] < handTypes[hands[j].typeName]
	} else {
		strongest := getStrongestHand(hands[i], hands[j])
		if strongest == hands[i] {
			return false
		}
	}
	return true
}

func (hands Hands) Swap(i, j int) {
	hands[i], hands[j] = hands[j], hands[i]
}

func (h *Hand) getType() {
	tempMap := map[string]int{}
	var typeName string

	for _, c := range h.cards {
		tempMap[c] += 1
	}

	var pairs int = 0

	if len(tempMap) == 5 { // high card
		typeName = "highCard"
	} else if len(tempMap) == 1 { // five of a kind
		typeName = "fiveOfAKind"
	} else {
		for _, v := range tempMap {
			if v == 4 {
				typeName = "fourOfAKind"
			} else if v == 3 {
				typeName = "threeOfAKind"
			} else if v == 2 {
				pairs += 1
			}
		}
		if pairs == 2 {
			typeName = "twoPairs"
		} else if pairs == 1 && typeName == "threeOfAKind" {
			typeName = "fullHouse"
		} else if pairs == 1 {
			typeName = "onePair"
		}
	}
	h.typeName = typeName
}

func parseInput(lines []string) (hands []Hand) {
	for _, l := range lines {
		hand := Hand{}
		parts := strings.Fields(l)
		for i := range parts[0] {
			hand.cards[i] = string(parts[0][i])
		}
		hand.bidAmount, _ = strconv.Atoi(parts[1])
		hands = append(hands, hand)
	}
	return hands
}

func main() {
	puzzle_input, _ := utils.ReadLines("input.txt")
	hands := parseInput(puzzle_input)
	//fmt.Println(hands)
	fmt.Println(p1(hands))
	fmt.Println(p2("2"))
}
