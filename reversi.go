package main

import (
	"Reversi/mypkg"
	"fmt"
	"math/rand"
	"time"
)

const (
	boardSize   = 8
	emptyCell   = " . "
	playerBlack = " x "
	playerWhite = " o "
)

var board [boardSize][boardSize]string
var currentPlayer string

func initialBoard() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			board[i][j] = emptyCell
		}
	}
	board[3][3] = playerWhite
	board[4][4] = playerWhite
	board[3][4] = playerBlack
	board[4][3] = playerBlack
}

func printBoard() {
	fmt.Println("---------------------------")
	fmt.Println("    0  1  2  3  4  5  6  7 ")
	for i := 0; i < boardSize; i++ {
		fmt.Print(" ", i, " ")
		for j := 0; j < boardSize; j++ {
			fmt.Print(board[i][j])
		}
		fmt.Print("\n")
	}
}

func counts() (int, int) {
	blackCount, whiteCount := 0, 0
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			if board[row][col] == playerBlack {
				blackCount++
			} else if board[row][col] == playerWhite {
				whiteCount++
			}
		}
	}
	return blackCount, whiteCount
}

func winner(blackCount int, whiteCount int) string {
	if blackCount > whiteCount {
		return "Black Win"
	} else if blackCount < whiteCount {
		return "White Win"
	} else {
		return "Draw"
	}
}

func getRandomMove(currentPlayer string, board [boardSize][boardSize]string) (int, int) {
	rand.Seed(time.Now().UnixNano()) // Initialize random seed
	for {
		row := rand.Intn(boardSize)
		col := rand.Intn(boardSize)
		if mypkg.IsValidMove(row, col, board, currentPlayer) {
			return row, col
		}
	}
}

// TODO:Computer Player
// TODO:Time Limit

func main() {
	initialBoard()
	printBoard()
	// Round Start
	var col, raw int
	var flag int
	currentPlayer = playerBlack
	fmt.Println("Chose a mode: 1.normal  2.computer")
	fmt.Scan(&flag)
	if flag == 1 {
		for !mypkg.IsGameOver(playerBlack, playerWhite, board, currentPlayer) {
			blackCount, whiteCount := counts()
			fmt.Println("Black Count: ", blackCount)
			fmt.Println("White Count: ", whiteCount)
			fmt.Println("Current Player: ", currentPlayer)
			fmt.Println("Input Chess(X,Y): ")
			fmt.Scan(&raw, &col)
			board = mypkg.ChoicePos(raw, col, currentPlayer, board)
			printBoard()
			if currentPlayer == playerBlack {
				currentPlayer = playerWhite
			} else {
				currentPlayer = playerBlack
			}
			// if blackCount+whiteCount == boardSize*boardSize {
			// 	winner(blackCount, whiteCount)
			// }
		}
		// Round End
		fmt.Print("---------------------------\n")
		printBoard()
		blackCount, whiteCount := counts()
		fmt.Printf("Game Over! Winner: %s\n", winner(blackCount, whiteCount))
		fmt.Printf("Black Pieces: %d\n", blackCount)
		fmt.Printf("White Pieces: %d\n", whiteCount)
	} else {
		for !mypkg.IsGameOver(playerBlack, playerWhite, board, currentPlayer) {
			blackCount, whiteCount := counts()
			fmt.Println("Black Count: ", blackCount)
			fmt.Println("White Count: ", whiteCount)
			fmt.Println("Current Player: ", currentPlayer)

			if currentPlayer == playerBlack {
				// Human player's turn
				fmt.Println("Input Chess(X,Y): ")
				fmt.Scan(&raw, &col)
			} else {
				// Computer player's turn
				raw, col = getRandomMove(currentPlayer, board)
				fmt.Printf("Computer plays at position (%d, %d)\n", raw, col)
			}

			board = mypkg.ChoicePos(raw, col, currentPlayer, board)
			printBoard()

			if currentPlayer == playerBlack {
				currentPlayer = playerWhite
			} else {
				currentPlayer = playerBlack
			}

			// if blackCount+whiteCount == boardSize*boardSize {
			// 	winner(blackCount, whiteCount)
			// }
		}
		// Round End
		fmt.Print("---------------------------\n")
		printBoard()
		blackCount, whiteCount := counts()
		fmt.Printf("Game Over! Winner: %s\n", winner(blackCount, whiteCount))
		fmt.Printf("Black Pieces: %d\n", blackCount)
		fmt.Printf("White Pieces: %d\n", whiteCount)
	}

}
