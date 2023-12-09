package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Pos struct {
	x, y int
}

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

	var arr [][]rune
	var lineNumber int

	var partList []string
	for scanner.Scan() {
		arr = append(arr, []rune{})
		for _, elem := range scanner.Text() {
			arr[lineNumber] = append(arr[lineNumber], elem)
		}
		lineNumber = lineNumber + 1
	}

	// part 1
	for x, row := range arr {
		number := ""
		valid := false
		for y, char := range row {
			// reset if period
			if char == '.' || checkSymbol(char) {
				if valid {
					partList = append(partList, number)
					valid = false
				}
				number = ""
			} else {
				if checkAdjacent(arr, x, y) {
					valid = true
				}
				number = number + string(char)
			}
		}

		// check if valid on end of line
		if valid {
			partList = append(partList, number)
			valid = false
		}
		number = ""
	}

	hashmap := make(map[Pos][]string)
	// part2
	for x, row := range arr {
		valid := false
		pos := Pos{0, 0}
		number := ""
		for y, char := range row {
			if char == '.' || checkSymbol(char) {
				if valid {
					hashmap[pos] = append(hashmap[pos], number)
					valid = false
				}
				number = ""
			} else {

				p, checkGear := checkGearAdjacent(arr, x, y)
				if checkGear {
					valid = true
					pos = p
				}

				number = number + string(char)
			}
		}

		if valid {
			hashmap[pos] = append(hashmap[pos], number)
			valid = false
		}
		number = ""
	}

	sumRatio := 0
	for _, partList := range hashmap {
		if len(partList) > 1 {
			part1, _ := strconv.Atoi(partList[0])
			part2, _ := strconv.Atoi(partList[1])
			println(part1, part2)

			sumRatio = sumRatio + (part1 * part2)
		}
	}

	println("Part1", sum(partList))
	println("Part2", sumRatio)
}

func sum(list []string) int {
	count := 0
	for _, numAsString := range list {
		n1, _ := strconv.Atoi(numAsString)
		count = count + n1
	}

	return count
}

func checkAdjacent(list [][]rune, x int, y int) bool {
	up := Pos{x - 1, y}
	down := Pos{x + 1, y}
	right := Pos{x, y + 1}
	left := Pos{x, y - 1}
	uRight := Pos{x - 1, y + 1}
	dRight := Pos{x + 1, y + 1}
	uLeft := Pos{x - 1, y - 1}
	dLeft := Pos{x + 1, y - 1}

	// check up left
	if uLeft.x >= 0 && uLeft.y >= 0 && checkSymbol(list[uLeft.x][uLeft.y]) {
		return true
	}

	// check down left
	if dLeft.x < len(list) && dLeft.y >= 0 && checkSymbol(list[dLeft.x][dLeft.y]) {
		return true
	}

	// check left
	if left.y >= 0 && checkSymbol(list[left.x][left.y]) {
		return true
	}

	// check up right
	if uRight.x >= 0 && uRight.y < len(list[uRight.x]) && checkSymbol(list[uRight.x][uRight.y]) {
		return true
	}

	// check down right
	if dRight.x < len(list) && dRight.y < len(list[dRight.x]) && checkSymbol(list[dRight.x][dRight.y]) {
		return true
	}

	// check right
	if right.y < len(list[right.x]) && checkSymbol(list[right.x][right.y]) {
		return true
	}

	// check up
	if up.x >= 0 && checkSymbol(list[up.x][up.y]) {
		return true
	}

	// check down
	if down.x < len(list) && checkSymbol(list[down.x][down.y]) {
		return true
	}

	return false
}

func checkSymbol(symbol rune) bool {
	if symbol == '\x2E' {
		return false
	}

	if symbol >= '\x30' && symbol <= '\x39' {
		return false
	}

	return true
}

func checkGear(symbol rune) bool {
	if symbol == '*' {
		return true
	}
	return false
}

func checkGearAdjacent(list [][]rune, x int, y int) (Pos, bool) {
	up := Pos{x - 1, y}
	down := Pos{x + 1, y}
	right := Pos{x, y + 1}
	left := Pos{x, y - 1}
	uRight := Pos{x - 1, y + 1}
	dRight := Pos{x + 1, y + 1}
	uLeft := Pos{x - 1, y - 1}
	dLeft := Pos{x + 1, y - 1}

	// check up left
	if uLeft.x >= 0 && uLeft.y >= 0 && checkSymbol(list[uLeft.x][uLeft.y]) {
		return uLeft, false
	}

	// check down left
	if dLeft.x < len(list) && dLeft.y >= 0 && checkSymbol(list[dLeft.x][dLeft.y]) {
		return dLeft, true
	}

	// check left
	if left.y >= 0 && checkSymbol(list[left.x][left.y]) {
		return left, true
	}

	// check up right
	if uRight.x >= 0 && uRight.y < len(list[uRight.x]) && checkSymbol(list[uRight.x][uRight.y]) {
		return uRight, true
	}

	// check down right
	if dRight.x < len(list) && dRight.y < len(list[dRight.x]) && checkSymbol(list[dRight.x][dRight.y]) {
		return dRight, true
	}

	// check right
	if right.y < len(list[right.x]) && checkSymbol(list[right.x][right.y]) {
		return right, true
	}

	// check up
	if up.x >= 0 && checkSymbol(list[up.x][up.y]) {
		return up, true
	}

	// check down
	if down.x < len(list) && checkSymbol(list[down.x][down.y]) {
		return down, true
	}

	return Pos{0, 0}, false
}
