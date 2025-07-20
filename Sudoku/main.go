package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 9 {
		fmt.Println("Error$")
		return
	}

	board := make([][]int, 9)
	for i := range board {
		board[i] = make([]int, 9)
	}

	for i := 0; i < 9; i++ {
		if len(args[i]) != 9 {
			fmt.Println("Error$")
			return
		}
		for j := 0; j < 9; j++ {
			number := args[i][j]
			if number == '.' {
				board[i][j] = 0
			} else {
				board[i][j] = int(number - '0')
			}

			// print each number in the board
			//fmt.Print(board[i][j])
			// add a space between numbers
			// if j < 8 {
			// 	fmt.Print(" ")
			// }
		}
		// print a new line at the end of each row
		//fmt.Print("$\n")
	}

	if solveSudoku(board) {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Print(board[i][j])
				if j < 8 {
					fmt.Print(" ")
				}
			}
			fmt.Print("$\n")
		}
	} else {
		fmt.Println("No solution found!")
	}
}

func solveSudoku(board [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(board, row, col, num) {
						board[row][col] = num
						if solveSudoku(board) {
							return true
						}
						board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}