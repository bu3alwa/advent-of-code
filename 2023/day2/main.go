package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Cannot read the file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	redMax := 12
	greenMax := 13
	blueMax := 14
	var total int
	var power int

	for scanner.Scan() {
		line := scanner.Text()
		id := getGameId(line)
		red := getRedMax(line)
		blue := getBlueMax(line)
		green := getGreenMax(line)

		if red <= redMax && blue <= blueMax && green <= greenMax {
			total = total + id
		}

		greenF := getGreenFewest(line, greenMax)
		redF := getRedFewest(line, redMax)
		blueF := getBlueFewest(line, blueMax)
		power = power + (greenF * redF * blueF)
	}
	println("part1", total)
	println("part2", power)
}

func getNumberFromString(line string) string {
	re2 := regexp.MustCompile("([0-9]+)")
	num := re2.FindStringSubmatch(line)
	return num[0]
}

func getGameId(line string) int {
	re := regexp.MustCompile("Game ([0-9]+):")
	substr := re.FindStringSubmatch(line)

	id, _ := strconv.Atoi(getNumberFromString(substr[0]))

	return id
}
func getBlueMatch(line string) [][]string {
	re := regexp.MustCompile("[0-9]+ blue")
	substr := re.FindAllStringSubmatch(line, -1)
	return substr
}

func getBlueFewest(line string, maxCubes int) int {
	substr := getBlueMatch(line)
	fewest := getFewestFromMatch(substr, maxCubes)
	return fewest
}

func getBlueMax(line string) int {
	substr := getBlueMatch(line)

	max := getMaxNumberFromMatch(substr)
	return max
}

func getRedMatch(line string) [][]string {
	re := regexp.MustCompile("[0-9]+ red")
	substr := re.FindAllStringSubmatch(line, -1)
	return substr
}

func getRedMax(line string) int {
	substr := getRedMatch(line)

	max := getMaxNumberFromMatch(substr)
	return max
}

func getRedFewest(line string, maxCubes int) int {
	substr := getRedMatch(line)
	fewest := getFewestFromMatch(substr, maxCubes)
	return fewest
}

func getGreenMatch(line string) [][]string {
	re := regexp.MustCompile("[0-9]+ green")
	substr := re.FindAllStringSubmatch(line, -1)
	return substr
}

func getGreenMax(line string) int {
	substr := getGreenMatch(line)
	max := getMaxNumberFromMatch(substr)

	return max
}

func getGreenFewest(line string, maxCubes int) int {
	substr := getGreenMatch(line)
	fewest := getFewestFromMatch(substr, maxCubes)
	return fewest
}

func getFewestFromMatch(substr [][]string, maxCubes int) int {
	var fewest int
	for _, value := range substr {
		val, _ := strconv.Atoi(getNumberFromString(value[0]))

		if val > fewest {
			fewest = val
		}

	}
	return fewest

}

func getMaxNumberFromMatch(substr [][]string) int {
	var max int

	for x := 0; x < len(substr); x++ {
		numAsString := getNumberFromString((substr[x][0]))
		number, _ := strconv.Atoi(numAsString)

		if number > max {
			max = number
		}
	}

	return max
}
