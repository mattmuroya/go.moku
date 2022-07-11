package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mattmuroya/go.moku/util"
)

func main() {
	var board [15][15]string

	util.Splash()

	util.PrintBoard(board)

	// need to set current player, initialize with P1
	player := "PLAYER 1"

	// loop through single play, check for win condition and if it returns true,
	// set break the loop and let main() continue
	// otherwise set the new player and restart loop

	for {
		fmt.Printf("\n  %v\n  -------------------\n", player)
		play(player, &board)
		util.PrintBoard(board)

		// check for win condition
		// if winner, break

		if player == "PLAYER 1" {
			player = "PLAYER 2"
		} else {
			player = "PLAYER 1"
		}
	}

	// fmt.Printf("  %v wins! Please play again.\n\n", winner)

}

func play(player string, board *[15][15]string) {
	for {
		col, row := getPlayerInput()

		if board[row][col] == "" {
			if player == "PLAYER 1" {
				board[row][col] = "○"
			} else {
				board[row][col] = "●"
			}
			break
		} else {
			fmt.Print("  Space unavailable. Please try again.\n\n")
		}
	}
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

	var x, y string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("  Enter column A-O: ")
		scanner.Scan()
		x = strings.ToLower(scanner.Text())

		if util.IndexOf(x, c) > -1 {
			break
		} else {
			fmt.Print("  Invalid input. Please try again.\n")
		}
	}

	for {
		fmt.Print("  Enter row 1-15:   ")
		scanner.Scan()
		y = scanner.Text()

		if util.IndexOf(y, r) > -1 {
			break
		} else {
			fmt.Print("  Invalid input. Please try again.\n")
		}
	}

	fmt.Print("\n  Player input: " + strings.ToUpper(x) + y + "\n\n")
	return util.IndexOf(x, c), util.IndexOf(y, r)
}

// "15  · · · · · · · · · · · · · · ·\n" +
// "14  · · · · · · · · · · · · · · ·\n" +
// "13  · · · · · · · · · · · · · · ·\n" +
// "12  · · · · · · · · · · · · · · ·\n" +
// "11  · · · · · · · · · · · · · · ·\n" +
// "10  · · · · · · · ● · · · · · · ·\n" +
// "9   · · · · · · ● · ● ○ · · · · ·\n" +
// "8   · · · · · ○ ● ○ · ○ · · · · ·\n" +
// "7   · · · · · · ○ ● · · · · · · ·\n" +
// "6   · · · · · · · · · · · · · · ·\n" +
// "5   · · · · · · · · · · · · · · ·\n" +
// "4   · · · · · · · · · · · · · · ·\n" +
// "3   · · · · · · · · · · · · · · ·\n" +
// "2   · · · · · · · · · · · · · · ·\n" +
// "1   · · · · · · · · · · · · · · ·\n" +
// "\n"
// "    A B C D E F G H I J K L M N O"
