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

	var x []int
	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		y, err := strconv.Atoi(scanner.Text());
		if err != nil {
      log.Fatal(err)
		}

		x = append(x, y)
	}


	var count int = 0
	for i := 0; i < len(x) -1; i++ {
		if x[i] < x[i+1] {
			count++
		}	
	}
	println(count)

	var count2 int = 0
	for i := 0; i < len(x) - 3; i++ {
		var acc int = x[i] + x[i+1] + x[i+2]
		var acc2 int = x[i+1] + x[i+2] + x[i+3]
		if acc < acc2 {
			count2++
		}
	}

	println(count2)
}
