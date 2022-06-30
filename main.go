package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mattmuroya/go.moku/common"
)

func main() {

	// var board [15][15]string

	board := [][]string{
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 15
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 14
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 13
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 12
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 11
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 10
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 09
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 08
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 07
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 06
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 05
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 04
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 03
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 02
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // 01
		// A   B   C   D   E   F   G   H   I   J   K   L   M   N   O
	}

	// winner := ""
	gameOver := false

	fmt.Print("\nWelcome to go.moku!\n\n")

	for !gameOver {
		// ask for player input and convert to col and row indices
		col, row := getPlayerInput()
		// save player input
		board[row][col] = "○"
		// print board
		printBoard(board)
		// check for win condition
		//   if win, set winner="Player" and gameOver=true
		//   break loop

		// if no win, print board

		// play computer input
		// save computer input
		// print board
		// check for win condition
		//   if win, set winner="Computer" and gameOver=true
		//   break loop
		// if no win, print board
		// reset loop
	}

	// fmt.Printf("%v wins! Please play again.\n", winner)

}

func getPlayerInput() (int, int) {

	c := []string{
		"a", "b", "c", "d", "e",
		"f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o",
	}

	r := []string{
		"15", "14", "13", "12", "11",
		"10", "9", "8", "7", "6",
		"5", "4", "3", "2", "1",
	}

	var col string
	var row string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter column A-O: ")
		scanner.Scan()
		col = strings.ToLower(scanner.Text())
		if common.IndexOf(col, c) > -1 {
			break
		} else {
			fmt.Print("Invalid input. Please try again.\n")
		}
	}

	for {
		fmt.Print("Enter row 1-15:   ")
		scanner.Scan()
		row = scanner.Text()
		if common.IndexOf(row, r) > -1 {
			break
		} else {
			fmt.Print("Invalid input. Please try again.\n")
		}
	}

	fmt.Print("\nPlayer input: " + strings.ToUpper(col) + row + "\n\n")

	return common.IndexOf(col, c), common.IndexOf(row, r)
}

func printBoard(board [][]string) {
	for i, row := range board {
		fmt.Print(strconv.FormatInt(15-int64(i), 10) + "  ")
		if i > 5 {
			fmt.Print(" ")
		}
		for _, col := range row {
			if col != "" {
				fmt.Print(col + " ")
			} else {
				fmt.Print("· ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n    A B C D E F G H I J K L M N O\n\n")
}

// "15 · · · · · · · · · · · · · · ·\n" +
// "14 · · · · · · · · · · · · · · ·\n" +
// "13 · · · · · · · · · · · · · · ·\n" +
// "12 · · · · · · · · · · · · · · ·\n" +
// "11 · · · · · · · · · · · · · · ·\n" +
// "10 · · · · · · · ● · · · · · · ·\n" +
// "9  · · · · · · ● · ● ○ · · · · ·\n" +
// "8  · · · · · ○ ● ○ · ○ · · · · ·\n" +
// "7  · · · · · · ○ ● · · · · · · ·\n" +
// "6  · · · · · · · · · · · · · · ·\n" +
// "5  · · · · · · · · · · · · · · ·\n" +
// "4  · · · · · · · · · · · · · · ·\n" +
// "3  · · · · · · · · · · · · · · ·\n" +
// "2  · · · · · · · · · · · · · · ·\n" +
// "1  · · · · · · · · · · · · · · ·\n\n" +
// "    A B C D E F G H I J K L M N O"
