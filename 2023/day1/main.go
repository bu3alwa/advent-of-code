package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	var total int = 0
	for scanner.Scan() {
		var firstNumber byte
		var lastNumber byte

		line := scanner.Text()
		// part 2
		line = strings.ReplaceAll(line, "nine", "n9e")
		line = strings.ReplaceAll(line, "eight", "e8t")
		line = strings.ReplaceAll(line, "seven", "s7n")
		line = strings.ReplaceAll(line, "six", "s6x")
		line = strings.ReplaceAll(line, "five", "f5e")
		line = strings.ReplaceAll(line, "four", "f4r")
		line = strings.ReplaceAll(line, "three", "t3e")
		line = strings.ReplaceAll(line, "two", "t2o")
		line = strings.ReplaceAll(line, "one", "o1e")

		println(line)

		for c := 0; c < len(line); c++ {
			if line[c] >= '\x30' && line[c] <= '\x39' {

				if firstNumber == '\x00' {
					firstNumber = line[c]
				}
				lastNumber = line[c]

			}
		}

		buff := []byte("")
		buff = append(buff, firstNumber)
		buff = append(buff, lastNumber)

		lineNumber, err := strconv.Atoi(string(buff))
		if err != nil {
			log.Fatal(err)
		}

		total = total + lineNumber
	}

	println(total)
}
