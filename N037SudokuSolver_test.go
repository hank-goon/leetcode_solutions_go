package main

import (
	"fmt"
	"testing"
)

func TestN037SudokuSolver(t *testing.T) {
	var n037 N037SudokuSolver
	board := [][]byte{
		[]byte{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		[]byte{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		[]byte{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		[]byte{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		[]byte{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		[]byte{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	}
	n037.solveSudoku(board)
	fmt.Println("N037SudokuSolver: ")
	n037.Print(board)
}
