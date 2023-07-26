package mypkg

import (
	"fmt"
)

const (
	BoardSize   = 8
	emptyCell   = " . "
	PlayerBlack = " ● "
	PlayerWhite = " ○ "
)

func ChoicePos(row int, col int, NowPlayer string, board [8][8]string) [8][8]string {
	//SET CHESS
	if !isEmpty(row, col, board) {
		fmt.Print("!!!Please set chess in emptyCell!!!")
		fmt.Print("\n")
		return board
	}
	board_new, Score := ReverChess(row, col, NowPlayer, board)
	if Score > 1 {
		board_new[row][col] = NowPlayer
		return board_new
	} else {
		fmt.Println("!!! Please choose a selectable slot !!!")
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
	if currentPlayer == PlayerWhite {
		currentPlayer = PlayerBlack
	} else {
		currentPlayer = PlayerWhite
	}
	return currentPlayer
}

func isValidMove(row int, col int, board [8][8]string, currentPlayer string) bool {
	if row < 0 || row >= BoardSize || col < 0 || col >= BoardSize || board[row][col] != emptyCell {
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
	if r >= 0 && r < BoardSize && c >= 0 && c < BoardSize && board[r][c] == anotherPlayer {
		// Keep moving in the direction until an empty cell or the current player's piece is found
		for r >= 0 && r < BoardSize && c >= 0 && c < BoardSize {
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
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if isValidMove(row, col, board, currentPlayer) && board[row][col] == emptyCell {
				return true
			}
		}
	}
	return false
}

//game over
func isGameOver(PlayerBlack string, PlayerWhite string, board [8][8]string, currentPlayer string) bool {
	return !hasValidMove(PlayerBlack, board, currentPlayer) && !hasValidMove(PlayerWhite, board, currentPlayer)
}

func ReverChess(row int, col int, NowPlayer string, board [8][8]string) ([8][8]string, int) {
	flag := false
	Score := 0
	anotherPlayer := playerChange(NowPlayer)
	for flag != true {
		//left
		for dr, dc := row, col; dr >= 1 && dc >= 1 && dr < 8 && dc < 8; {
			dc = dc - 1
			if board[dr][dc] == emptyCell {
				break
			}
			if board[dr][dc] == NowPlayer {
				//fmt.Println("Score-1:", Score)
				for i := col; i > dc; i = i - 1 {
					if board[row][i] == anotherPlayer {
						Score = Score + 1
						board[row][i] = NowPlayer
						flag = true
					}
				}
			}
		}

		//right
		for dr, dc := row, col; dr >= 0 && dc >= 0 && dr < 7 && dc < 7; {
			dc = dc + 1
			if board[dr][dc] == emptyCell {
				break
			}
			if board[dr][dc] == NowPlayer {
				for i := col; i < dc; i = i + 1 {
					if board[row][i] == anotherPlayer {
						Score = Score + 1
						board[row][i] = NowPlayer
						flag = true
					}
				}
			}
		}

		//up
		for dr, dc := row, col; dr >= 1 && dc >= 1 && dr < 8 && dc < 8; {
			dr = dr - 1
			if board[dr][dc] == emptyCell {
				break
			}
			if board[dr][dc] == NowPlayer {
				for i := row; i > dr; i = i - 1 {
					if board[i][col] == anotherPlayer {
						Score = Score + 1
						board[i][col] = NowPlayer
						flag = true
					}
				}
			}
		}

		//down
		for dr, dc := row, col; dr >= 0 && dc >= 0 && dr < 7 && dc < 7; {
			dr = dr + 1
			if board[dr][dc] == emptyCell {
				break
			}
			if board[dr][dc] == NowPlayer {
				for i := row; i < dr; i = i + 1 {
					if board[i][col] == anotherPlayer {
						Score = Score + 1
						board[i][col] = NowPlayer
						flag = true
					}
				}
			}
		}

		//up-left
		for dr, dc := row, col; dr >= 1 && dc >= 1 && dr < 8 && dc < 8; {
			dr = dr - 1
			dc = dc - 1
			if board[dr][dc] == emptyCell {
				break
			}
			if board[dr][dc] == NowPlayer {
				//fmt.Println("Score3:", Score)
				for i, j := row, col; i > dr && j > dc; {
					if board[i][j] == anotherPlayer {
						Score = Score + 1
						board[i][j] = NowPlayer
						flag = true
					}
					i = i - 1
					j = j - 1
				}
			}
		}
		//down-right
		for dr, dc := row, col; dr >= 0 && dc >= 0 && dr < 7 && dc < 7; {
			dr = dr + 1
			dc = dc + 1
			if board[dr][dc] == emptyCell {
				break
			}
			if board[dr][dc] == NowPlayer {
				for i, j := row, col; i < dr && j < dc; {
					if board[i][j] == anotherPlayer {
						Score = Score + 1
						board[i][j] = NowPlayer
						flag = true
					}
					i = i + 1
					j = j + 1
				}
			}
		}

		//up-right
		for dr, dc := row, col; dr >= 1 && dc >= 0 && dr < 8 && dc < 7; {
			dr = dr - 1
			dc = dc + 1
			if board[dr][dc] == emptyCell {
				break
			}
			if board[dr][dc] == NowPlayer {
				for i, j := row, col; i > dr && j < dc; {
					if board[i][j] == anotherPlayer {
						Score = Score + 1
						board[i][j] = NowPlayer
						flag = true
					}
					i = i - 1
					j = j + 1
				}
			}
		}

		//down-left
		for dr, dc := row, col; dr >= 0 && dc >= 1 && dr < 7 && dc < 8; {
			dr = dr + 1
			dc = dc - 1
			if board[dr][dc] == emptyCell {
				break
			}
			if board[dr][dc] == NowPlayer {
				//fmt.Println("Score6:", Score)
				for i, j := row, col; i < dr && j > dc; {
					if board[i][j] == anotherPlayer {
						Score = Score + 1
						board[i][j] = NowPlayer
						flag = true
					}
					i = i + 1
					j = j - 1
				}
			}
		}
		return board, Score + 1
	}
	return board, Score + 1
}
