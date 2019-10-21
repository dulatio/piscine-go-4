package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	board := [...]int{-1, -1, -1, -1, -1, -1, -1, -1}
	recursiveThing(board, 0)
}

func recursiveThing(board [8]int, col int) bool {
	if col == 8 {
		printSolution(board)
		return false
	}
	flag := false
	for row := 0; row < 8; row++ {
		board[col] = row
		if canBe(board, col) {
			board[col] = row
			flag = recursiveThing(board, col + 1) || flag
			board[col] = -1
		}
	}
	return flag
}

func canBe(board [8]int, col int) bool {
	for i := 0; i < col; i++ {
		if board[i] == board[col] {
			return false
		}
		if (col-i == board[col]-board[i]) || (col-i == board[i]-board[col]) {
			return false
		}
	}
	return true
}

func printSolution(board [8]int) {
	for i := 0; i < 8; i++ {
		if board[i] == 0 {
			z01.PrintRune('1')
		} else if board[i] == 1 {
			z01.PrintRune('2')
		} else if board[i] == 2 {
			z01.PrintRune('3')
		} else if board[i] == 3 {
			z01.PrintRune('4')
		} else if board[i] == 4 {
			z01.PrintRune('5')
		} else if board[i] == 5 {
			z01.PrintRune('6')
		} else if board[i] == 6 {
			z01.PrintRune('7')
		} else if board[i] == 7 {
			z01.PrintRune('8')
		}
	}
	z01.PrintRune(10)
}
