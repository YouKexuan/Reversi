package main

import (
	"fmt"
)

const (
	boardSize = 8
	emptyCell = " . "
	playerX   = "X"
	playerO   = "O"
)

var board [boardSize][boardSize]string
var currentPlayer string

func initialBoard() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			board[i][j] = emptyCell
		}
	}
}

func printBoard() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			fmt.Print(board[i][j])
		}
		fmt.Print("\n")
	}
}

func main() {
	initialBoard()
	printBoard()
	
}