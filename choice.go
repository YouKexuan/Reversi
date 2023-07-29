package mypkg

import (
	"fmt"
)

const (
	boardSize   = 8
	emptyCell   = " . "
	playerBlack = " x "
	playerWhite = " o "
)

func ChoicePos(row int, col int, NowPlayer string, board [8][8]string) [8][8]string {
	//SET CHESS
	if isValidMove(row, col, board, NowPlayer) {
		board[row][col] = NowPlayer
		board = ReverChess(row, col, NowPlayer, board)
		return board
	} else {
		fmt.Print("!!!Please set chess in emptyCell!!!")
		fmt.Print("\n")
		return board
	}
}

func isEmpty(row int, col int, board [8][8]string) bool {
	if row < 8 && col < 8 && board[row][col] == emptyCell {
		return true
	} else {
		return false

	}
}

func playerChange(currentPlayer string) string {
	if currentPlayer == playerWhite {
		currentPlayer = playerBlack
	} else {
		currentPlayer = playerWhite
	}
	return currentPlayer
}

func isValidMove(row int, col int, board [8][8]string, currentPlayer string) bool {
	if row < 0 || row >= boardSize || col < 0 || col >= boardSize || board[row][col] != emptyCell {
		return false
	}
	// Check if at least one piece can be flipped in any direction
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			if checkDirection(row, col, dr, dc, currentPlayer, board) {
				return true
			}
		}
	}
	return false
}

// checkDirection checks if there are pieces to flip in a specific direction.
func checkDirection(row int, col int, dr int, dc int, currentPlayer string, board [8][8]string) bool {
	anotherPlayer := playerChange(currentPlayer)

	r, c := row+dr, col+dc

	// Check if the first piece in the direction is the other player's piece
	if r >= 0 && r < boardSize && c >= 0 && c < boardSize && board[r][c] == anotherPlayer {
		// Keep moving in the direction until an empty cell or the current player's piece is found
		for r >= 0 && r < boardSize && c >= 0 && c < boardSize {
			if board[r][c] == emptyCell {
				return false
			} else if board[r][c] == currentPlayer {
				return true
			}
			r += dr
			c += dc
		}
	}
	return false
}

// hasValidMove checks if the current player has a valid move.
func hasValidMove(player string, board [8][8]string, currentPlayer string) bool {
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			if isValidMove(row, col, board, currentPlayer) && board[row][col] == emptyCell {
				return true
			}
		}
	}
	return false
}

func IsGameOver(playerBlack string, playerWhite string, board [8][8]string, currentPlayer string) bool {
	return !hasValidMove(playerBlack, board, currentPlayer) && !hasValidMove(playerWhite, board, currentPlayer)
}

func ReverChess(raw int, col int, NowPlayer string, board [8][8]string) [8][8]string {
	//left
	anotherPlayer := playerChange(NowPlayer)

	for dr, dc := raw, col; dr >= 0 && dc > 0 && dr < 8 && dc < 8; {
		dc = dc - 1
		if board[dr][dc] == emptyCell {
			break
		}
		if board[dr][dc] == NowPlayer {
			for i := col; i > dc; i = i - 1 {
				if board[raw][i] == anotherPlayer {
					board[raw][i] = NowPlayer
				}
			}
		}
	}

	//right
	for dr, dc := raw, col; dr >= 0 && dc >= 0 && dr < 8 && dc < 7; {
		dc = dc + 1
		if board[dr][dc] == emptyCell {
			break
		}
		if board[dr][dc] == NowPlayer {
			for i := col; i < dc; i = i + 1 {
				if board[raw][i] == anotherPlayer {
					board[raw][i] = NowPlayer
				}
			}
		}
	}

	//up
	for dr, dc := raw, col; dr > 0 && dc >= 0 && dr < 8 && dc < 8; {
		dr = dr - 1
		if board[dr][dc] == emptyCell {
			break
		}
		if board[dr][dc] == NowPlayer {
			for i := raw; i > dr; i = i - 1 {
				if board[i][col] == anotherPlayer {
					board[i][col] = NowPlayer
				}
			}
		}
	}

	//down
	for dr, dc := raw, col; dr >= 0 && dc >= 0 && dr < 7 && dc < 8; {
		dr = dr + 1
		if board[dr][dc] == emptyCell {
			break
		}
		if board[dr][dc] == NowPlayer {
			for i := raw; i < dr; i = i + 1 {
				if board[i][col] == anotherPlayer {
					board[i][col] = NowPlayer
				}
			}
		}
	}

	//up-left
	for dr, dc := raw, col; dr > 0 && dc > 0 && dr < 8 && dc < 8; {
		dr = dr - 1
		dc = dc - 1
		if board[dr][dc] == emptyCell {
			break
		}
		if board[dr][dc] == NowPlayer {
			for i, j := raw, col; i > dr && j > dc; {
				if board[i][j] == anotherPlayer {
					board[i][j] = NowPlayer
				}
				i = i - 1
				j = j - 1
			}
		}
	}
	//down-right
	for dr, dc := raw, col; dr >= 0 && dc >= 0 && dr < 7 && dc < 7; {
		dr = dr + 1
		dc = dc + 1
		if board[dr][dc] == emptyCell {
			break
		}
		if board[dr][dc] == NowPlayer {
			for i, j := raw, col; i < dr && j < dc; {
				if board[i][j] == anotherPlayer {
					board[i][j] = NowPlayer
				}
				i = i + 1
				j = j + 1
			}
		}
	}

	//up-right
	for dr, dc := raw, col; dr > 0 && dc >= 0 && dr < 8 && dc < 7; {
		dr = dr - 1
		dc = dc + 1
		if board[dr][dc] == emptyCell {
			break
		}
		if board[dr][dc] == NowPlayer {
			for i, j := raw, col; i > dr && j < dc; {
				if board[i][j] == anotherPlayer {
					board[i][j] = NowPlayer
				}
				i = i - 1
				j = j + 1
			}
		}
	}

	//down-left
	for dr, dc := raw, col; dr >= 0 && dc > 0 && dr < 7 && dc < 8; {
		dr = dr + 1
		dc = dc - 1
		if board[dr][dc] == emptyCell {
			break
		}
		if board[dr][dc] == NowPlayer {
			for i, j := raw, col; i < dr && j > dc; {
				if board[i][j] == anotherPlayer {
					board[i][j] = NowPlayer
				}
				i = i + 1
				j = j - 1
			}
		}
	}
	return board
}
