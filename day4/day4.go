package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bingo struct {
		board [5][5]int64
		rowCount [5]int64
		columnCount [5]int64
}

func main () {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Cannot read the file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	input := strings.Split(scanner.Text(), ",")

	file2, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Cannot read the file")
		os.Exit(1)
	}
	defer file2.Close()

	scanner2 := bufio.NewScanner(file2)

	if err := scanner2.Err(); err != nil {
		log.Fatal(err)
	}

	var bingo []Bingo
	bingo = append(bingo, Bingo{})
	var board int64
	var row int64
	var col int64
	for scanner2.Scan(){
		if row == 5 {
			bingo = append(bingo, Bingo{})
			board++
			row = 0
			continue
		}
		col = 0

		line := strings.Split(scanner2.Text(), " ")
		for i := range line {
			if line[i] == "" { continue }
			item, _ := strconv.ParseInt(strings.TrimSpace(line[i]), 10, 64)
			bingo[board].board[row][col] = item
			col++
		}
		row++
	}

	var total int64
	var already_won []int
	for _, draw := range input {		
		draw_int, _ := strconv.ParseInt(draw, 10,64)
		for b, board := range bingo {
			for r, row := range board.board {
				for c, col := range row {
					if (draw_int == col){
						bingo[b].board[r][c] = -1
						bingo[b].columnCount[c] += 1
						bingo[b].rowCount[r] += 1
					}


					not_won := true
					for _, won := range already_won {
						if (b == won) { not_won = false }
					}
					if ((bingo[b].columnCount[c] == 5 || bingo[b].rowCount[r] == 5) && not_won) {
						for i := 0; i < len(bingo[b].board); i++ {
							for j := 0; j< len(bingo[b].board[i]); j++ {
								if bingo[b].board[i][j] != -1 { total += bingo[b].board[i][j] }
							}
						}


						fmt.Println("winning board")
						fmt.Println(total * draw_int)
						total = 0
						already_won = append(already_won, b)
					}
				}
			}
		}
	}
}