package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/matsest/advent2023/utils"
)

func p2(input string) int {
	return 2
}
type conversionRange struct {
	dstStart int
	srcStart int
	length   int
}
type Conversions struct {
	seedToSoil            []conversionRange
	soilToFertilizer      []conversionRange
	fertilizerToWater     []conversionRange
	waterToLight          []conversionRange
	lightToTemperate      []conversionRange
	temperatureToHumidity []conversionRange
	humidityToLocation    []conversionRange
}

func parseInput(inputFile string) (seeds []int, c Conversions) {

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var section string

	for scanner.Scan() {
		line := scanner.Text()

		switch line {
		case "seed-to-soil map:":
			section = "seedToSoil"
		case "soil-to-fertilizer map:":
			section = "soilToFertilizer"
		case "fertilizer-to-water map:":
			section = "fertilizerToWater"
		case "water-to-light map:":
			section = "waterToLight"
		case "light-to-temperature map:":
			section = "lightToTemperature"
		case "temperature-to-humidity map:":
			section = "temperatureToHumidity"
		case "humidity-to-location map:":
			section = "humidityToLocation"
		default:
			if line != "" {

				if line[0:5] == "seeds" {
					numbers, err := utils.SliceAtoi(strings.Fields(line[6:]))
					//fmt.Println(line[6:], numbers, err)
					if err == nil {
						seeds = numbers
					}
				} else {
					numbers, err := utils.SliceAtoi(strings.Fields(line))
					//fmt.Println("as section", line, numbers, err)
					if err == nil {
						switch section {
						case "seedToSoil":
							c.seedToSoil = append(c.seedToSoil, conversionRange{numbers[0], numbers[1], numbers[2]})
						case "soilToFertilizer":
							c.soilToFertilizer = append(c.soilToFertilizer, conversionRange{numbers[0], numbers[1], numbers[2]})
						case "fertilizerToWater":
							c.fertilizerToWater = append(c.fertilizerToWater, conversionRange{numbers[0], numbers[1], numbers[2]})
						case "waterToLight":
							c.waterToLight = append(c.waterToLight, conversionRange{numbers[0], numbers[1], numbers[2]})
						case "lightToTemperature":
							c.lightToTemperate = append(c.lightToTemperate, conversionRange{numbers[0], numbers[1], numbers[2]})
						case "temperatureToHumidity":
							c.temperatureToHumidity = append(c.temperatureToHumidity, conversionRange{numbers[0], numbers[1], numbers[2]})
						case "humidityToLocation":
							c.humidityToLocation = append(c.humidityToLocation, conversionRange{numbers[0], numbers[1], numbers[2]})
						}
					}
				}
			}
		}
	}
	return
}

func resolve(input int, c []conversionRange) (new int) {
	for _, convRange := range c {
		if input >= convRange.srcStart && input < convRange.srcStart+convRange.length {
			diff := input - convRange.srcStart
			new = convRange.dstStart + diff
			break
		} else {
			new = input
		}
	}
	return new
}

func p1(seeds []int, c Conversions) int {
	lowest := math.MaxInt
	var soil, fertilizer, water, light, temperature, humidity, location int
	for _, s := range seeds {
		//fmt.Println("\nseed ", s)

		soil = resolve(s, c.seedToSoil)
		//fmt.Println("soil", soil)

		fertilizer = resolve(soil, c.soilToFertilizer)
		//fmt.Println("fertilizer", fertilizer)

		water = resolve(fertilizer, c.fertilizerToWater)
		//fmt.Println("water", water)

		light = resolve(water, c.waterToLight)
		//fmt.Println("light", light)

		temperature = resolve(light, c.lightToTemperate)
		//fmt.Println("temperature", temperature)

		humidity = resolve(temperature, c.temperatureToHumidity)
		//fmt.Println("humidity", humidity)

		location = resolve(humidity, c.humidityToLocation)
		//fmt.Println("location", location)

		if location < lowest {
			lowest = location
		}
	}
	return lowest
}

func main() {
	seeds, conversions := parseInput("input.txt")
	fmt.Println(p1(seeds, conversions))
	fmt.Println(p2("2"))
}
