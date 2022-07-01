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
	gameOver := false
	util.Splash()

	for !gameOver {
		fmt.Print("  PLAYER 1 ○\n" +
			"  ----------\n")
		play("p1", &board)
		util.PrintBoard(board)

		// check for win condition
		// checkForWin("p1")

		fmt.Print("  PLAYER 2 ●\n" +
			"  ----------\n")
		play("p2", &board)
		util.PrintBoard(board)

		// check for win condition
		// checkForWin("p2")

	}

	// fmt.Printf("%v wins! Please play again.\n", winner)

}

func play(player string, board *[15][15]string) {
	for {
		col, row := getPlayerInput()

		if board[row][col] == "" {
			if player == "p1" {
				board[row][col] = "○"
			} else if player == "p2" {
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
