package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/matsest/advent2023/utils"
)

func p1(steps []string) (sum int) {
	for _, step := range steps {
		sum += hash(step)
	}
	return sum
}

func p2(steps []string) (sum int) {

	boxes := Boxes{}
	// initialize
	for _, step := range steps {
		boxes.process(step)
		//fmt.Println("After ", step)
		//for b, lenses := range boxes {
		//	if len(lenses) > 0 {
		//	fmt.Printf("Box %v: %v\n", b, lenses)
		//	}
		//}
		//fmt.Println()
	}

	// calculate
	for box, lenses := range boxes {
		for i, lens := range lenses {
			val := (box + 1) * (i + 1) * lens.focalLength
			//fmt.Println(lens.label, val)
			sum += val
		}
	}
	return sum
}

func (boxes *Boxes) process(step string) {
	re := regexp.MustCompile(`\b[a-z]+\b`)
	label := re.FindString(step)
	box := hash(label)
	//fmt.Println(step, label, box)
	// if step contains a number, then it's an equal
	//  get focal Length (the number)
	re = regexp.MustCompile(`\d+`)
	digits := re.FindAllString(step, -1)
	var focalLength int
	if len(digits) > 0 {
		focalLength, _ = strconv.Atoi(digits[0])
		//fmt.Println("get focal length", focalLength)
		boxToProcess := box
		lenses := (*boxes)[boxToProcess]
		lenses.AddOrReplace(label, focalLength)
		(*boxes)[boxToProcess] = lenses
		// if lens with same label is in the box, replace it
		// else append it to the box
	} else {
		// remove the lens with this label from the box it is in
		//fmt.Println("is removal")
		boxToRemoveFrom := box
		lenses := (*boxes)[boxToRemoveFrom]
		lenses.Remove(label)
		(*boxes)[boxToRemoveFrom] = lenses
	}
}

func (lenses *Lenses) Remove(label string) {
	var result Lenses
	for _, v := range *lenses {
		if v.label != label {
			result = append(result, v)
		}
	}
	*lenses = result
}

func (lenses *Lenses) AddOrReplace(label string, focalLength int) {
	var result Lenses
	var isAdded bool = false
	for _, v := range *lenses {
		if v.label == label {
			// add the lens at this index
			result = append(result, Lens{label, focalLength})
			isAdded = true
		} else {
			result = append(result, v)
		}
	}
	if !isAdded {
		result = append(result, Lens{label, focalLength})
	}
	*lenses = result
}

type Boxes map[int]Lenses

type Lenses []Lens
type Lens struct {
	label       string
	focalLength int
}

func hash(step string) int {
	value := 0
	for _, c := range step {
		asciiCode := int(c)
		value += asciiCode
		value *= 17
		value = value % 256
	}
	return value
}

func parseInput(lines []string) (steps []string) {
	for _, line := range lines {
		parts := strings.Split(line, ",")
		steps = append(steps, parts...)
	}
	return steps
}

func main() {
	puzzle_input, _ := utils.ReadLines("input.txt")
	steps := parseInput(puzzle_input)
	fmt.Println(p1(steps))
	fmt.Println(p2(steps))
}
