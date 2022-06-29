package main

import (
	"fmt"
)

func main() {

	// 	"15  + + + + + + + + + + + + + + +\n" +
	// 	"14  + + + + + + + + + + + + + + +\n" +
	// 	"13  + + + + + + + + + + + + + + +\n" +
	// 	"12  + + + + + + + + + + + + + + +\n" +
	// 	"11  + + + + + + + + + + + + + + +\n" +
	// 	"10  + + + + + + + + + + + + + + +\n" +
	// 	"9   + + + + + + + + + + + + + + +\n" +
	// 	"8   + + + + + + + + + + + + + + +\n" +
	// 	"7   + + + + + + + + + + + + + + +\n" +
	// 	"6   + + + + + + + + + + + + + + +\n" +
	// 	"5   + + + + + + + + + + + + + + +\n" +
	// 	"4   + + + + + + + + + + + + + + +\n" +
	// 	"3   + + + + + + + + + + + + + + +\n" +
	// 	"2   + + + + + + + + + + + + + + +\n" +
	// 	"1   + + + + + + + + + + + + + + +\n" +
	// 	"\n" +
	// 	"    A B C D E F G H I J K L M N O\n"

	var board [15][15]string

	winner := ""
	gameOver := false

	// playerName := getPlayerName()

	for !gameOver {
		// ask for player input
		var col, row = getPlayerInput()
		fmt.Println(col, row)
		// save player input
		// check for win condition
		//   if win, set winner="Player" and gameOver=true
		//   break loop
		// if no win, print board

		// play computer input
		// save computer input
		// check for win condition
		//   if win, set winner="Computer" and gameOver=true
		//   break loop
		// if no win, print board
		// reset loop
	}

	fmt.Print(board)

	fmt.Printf("%v wins! Please play again.\n", winner)

}

func getPlayerInput() (string, string) {
	var col string
	var row string

	fmt.Print("Please enter a column A-O: ")
	fmt.Scan(&col)

	fmt.Print("Please enter a row 1-15: ")
	fmt.Scan(&row)
	return col, row
}
