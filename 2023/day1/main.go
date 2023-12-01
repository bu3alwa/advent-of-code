package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var total int = 0
	for scanner.Scan() {
		var firstNumber byte
		var lastNumber byte

		line := scanner.Text()
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
