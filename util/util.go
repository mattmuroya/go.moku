package util

import (
	"fmt"
	"strconv"
)

func IndexOf(input string, data []string) int {
	for i, e := range data {
		if e == input {
			return i
		}
	}
	return -1
}

func PrintBoard(board [15][15]string) {
	for i, row := range board {
		fmt.Print("  " + strconv.FormatInt(15-int64(i), 10) + "   ")
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
	fmt.Print("\n       A B C D E F G H I J K L M N O\n\n")
}

func Splash() {
	fmt.Print(
		"\n  let's play ...                 \n" +
			"                         _       \n" +
			"   ___ ___     _____ ___| |_ _ _ \n" +
			"  | ○ | ● |   |     | ● | '_| | |\n" +
			"  |_  |___| ○ |_|_|_|___|_,_|___|\n" +
			"  |___|                          \n" +
			"                                 \n" +
			"  ...but it's a command line app written in go\n\n")
}
