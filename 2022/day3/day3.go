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

	var x []string
	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		x = append(x, scanner.Text())
	}

	var count [12]int
	for i := 0; i < len(x); i++{
		for j:= 0; j < 12; j++ {
			if (x[i][j] == '0') {
				count[j] = count[j] + 1
			} 
		}
	}

	var gamma int64 = 0
	var epsilon int64 = 0
	// bit 1 most common gamma, 0 is epsilon
	// 1110000
	// 0123456
	eps := ""
	gam := ""
	for j := range x[0] {		
		if (count[j] > len(x)/2) {
			eps = eps + string('0')
			gam = gam + string('1')
		} else {
			eps = eps + string('1')
			gam = gam + string('0')
		}
	}

	if eps != "" { 
		if e, err := strconv.ParseInt(eps, 2, 64); err != nil {
			log.Fatal(err)
		} else {
			epsilon = epsilon + e
		}
	}

	if gam != "" { 
		if g, err := strconv.ParseInt(gam, 2, 64); err != nil {
			log.Fatal(err)
		} else {
			gamma =  gamma + g
		}
	}

	fmt.Println(epsilon * gamma)


	oxy := oxy_rating(x, 0)	
	co2 := co2_rating(x, 0)
	o, _ := strconv.ParseInt(oxy, 2, 64)
	c, _ := strconv.ParseInt(co2, 2, 64)

	fmt.Println(o * c)
}

func oxy_rating(array []string, start_position int) string {
	length := len(array)
	bit_length := 12
	var newArr []string
	var count int

	if (start_position > bit_length) { log.Fatal("error")}

	if (length == 1) { return array[0] }
	for i := 0; i < length; i++{
		temp := array[i]
		if (temp[start_position] == '0') {
			count = count + 1
		}
	}

	if ((count > length/2)){
		// keep with 0 in start position
		for i := 0; i < length; i++{
			temp := array[i]
			if (temp[start_position] == '0') {
				newArr = append(newArr, array[i])
			}
		}
	}

	if (count <= length/2) {
		// keep 1 in start position
		for i := 0; i < length; i++{
			temp := array[i]
			if (temp[start_position] == '1') {
				newArr = append(newArr, array[i])
			}
		}
	}

	return oxy_rating(newArr, start_position + 1)
}

func co2_rating(array []string, start_position int) string {
	length := len(array)
	bit_length := 12
	var newArr []string
	var count int

	if (start_position > bit_length) { log.Fatal("error")}

	if (length == 1) { return array[0] }
	for i := 0; i < length; i++{
		temp := array[i]
		if (temp[start_position] == '0') {
			count = count + 1
		}
	}

	if ((count <= length/2)){
		// keep with 0 in start position
		for i := 0; i < length; i++{
			temp := array[i]
			if (temp[start_position] == '0') {
				newArr = append(newArr, array[i])
			}
		}
	}

	if (count > length/2) {
		// keep 1 in start position
		for i := 0; i < length; i++{
			temp := array[i]
			if (temp[start_position] == '1') {
				newArr = append(newArr, array[i])
			}
		}
	}

	return co2_rating(newArr, start_position + 1)
}