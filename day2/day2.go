package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type loc struct {
	depth int
	position int
}

type loc2 struct {
	depth int
	position int
	aim int
}

func main () {
	l := loc{0, 0}
	l2 := loc2{0, 0, 0}

	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Cannot read the file")
		os.Exit(1)
	}
	defer file.Close()

	var x []string
	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		x = append(x, scanner.Text())
	}

	for i := 0; i < len(x); i++{
		split := strings.Split(x[i], " ")
		direction := split[0]

		y, err := strconv.Atoi(split[1])
		if (err != nil){
			log.Fatal("error converting string to number")
		}

		if(direction == "forward"){
			l.position += y
		} else if (direction == "down") {
			l.depth += y
		} else if (direction == "up") {
			l.depth -= y
		}
	}

	fmt.Println(l.depth * l.position)

	for i := 0; i < len(x); i++{
		split := strings.Split(x[i], " ")
		direction := split[0]

		y, err := strconv.Atoi(split[1])
		if (err != nil){
			log.Fatal("error converting string to number")
		}

		if(direction == "forward"){
			l2.position += y
			l2.depth += l2.aim * y
		} else if (direction == "down") {
			l2.aim += y
		} else if (direction == "up") {
			l2.aim -= y
		}
	}

	fmt.Println(l2.depth * l2.position)
}