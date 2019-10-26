package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Invalid sudoku board.")
		return
	}
	// declaring board
	board := make([][]int, 9)
	for i := range board {
		board[i] = make([]int, 9)
	}
	// filling board with initial values
	for i, v := range os.Args {
		if i != 0 {
			for j, r := range v {
				if j == 9 {
					fmt.Println("Too many fields for sudoku board.")
					return
				}
				if r != '.' {
					board[i-1][j] = int(r - '0')
				} 
			}
		}
	}
	if solve(board) {
		printBoard(board)
	} else {
		fmt.Println("No solution found.")
	}
}

func solve(board [][]int) bool {
	row := -1
	col := -1
	isFull := true
	// in our case the length of thr board is 9
	n := len(board)
	// check if we have fields without numbers
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 0 {
				row = i
				col = j
				isFull = false
				break
			}
		}
		if isFull {
			break
		}
	}
	if isFull {
		return true
	}
	// try values in empty fields recursively
	for i := 1; i <= n; i++ {
		if isSafe(board, row, col, i) {
			board[row][col] = i
			if solve(board) {
				return true
			} else {
				board[row][col] = 0
			}
		}
	}
	return false
}

func isSafe(board [][]int, row int, col int, num int) bool {
	// in our case n is 9
	n := len(board)
	// check the uniqueness of num horizontally
	for x := 0; x < n; x++ {
		if board[row][x] == num {
			return false
		}
	}
	// check the uniqueness of num vertically
	for y := 0; y < n; y++ {
		if board[y][col] == num {
			return false
		}
	}
	// apecific value for board of size 9
	sqrtn := 3
	boxRowStart := row - row % sqrtn
	boxColStart := col - col % sqrtn
	// check the uniqueness of num in the corresponding box
	for y := boxRowStart; y < boxRowStart+sqrtn; y++ {
		for x := boxColStart; x < boxColStart + sqrtn; x++ {
			if board[y][x] == num {
				return false
			}
		}
	}
	return true
}

func printBoard(board [][]int) {
	n := len(board)
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			if x != 0 {
				fmt.Print(" ")
			}
			fmt.Print(board[y][x])
		}
		fmt.Println()
	}
}
