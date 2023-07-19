package mypkg

import (
	"fmt"
)

const (
	boardSize   = 8
	emptyCell   = " . "
	playerBlack = " ● "
	playerWhite = " ○ "
)

func ChoicePos(row int, col int, NowPlayer string, board [8][8]string) [8][8]string {
	//SET CHESS
	if isEmpty(row, col, board) {
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

func ReverChess(raw int, col int, NowPlayer string, board [8][8]string) [8][8]string {
	//left
	anotherPlayer := playerChange(NowPlayer)

	for dr, dc := raw, col; dr >= 0 && dc >= 0 && dr < 8 && dc < 8; {
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
	for dr, dc := raw, col; dr >= 0 && dc >= 0 && dr < 8 && dc < 8; {
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
	for dr, dc := raw, col; dr >= 0 && dc >= 0 && dr < 8 && dc < 8; {
		dr = dr - 1
		if board[dr][dc] == emptyCell {
			break
		}
		if board[dr][dc] == NowPlayer {
			for i := raw; i < dr; i = i - 1 {
				if board[raw][i] == anotherPlayer {
					board[raw][i] = NowPlayer
				}
			}
		}
	}

	//down
	for dr, dc := raw, col; dr >= 0 && dc >= 0 && dr < 8 && dc < 8; {
		dr = dr + 1
		if board[dr][dc] == emptyCell {
			break
		}
		if board[dr][dc] == NowPlayer {
			for i := raw; i > dr; i = i + 1 {
				if board[raw][i] == anotherPlayer {
					board[raw][i] = NowPlayer
				}
			}
		}
	}

	//up-left
	for dr, dc := raw, col; dr >= 0 && dc >= 0 && dr < 8 && dc < 8; {
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
	for dr, dc := raw, col; dr >= 0 && dc >= 0 && dr < 8 && dc < 8; {
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
	for dr, dc := raw, col; dr >= 0 && dc >= 0 && dr < 8 && dc < 8; {
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
	for dr, dc := raw, col; dr >= 0 && dc >= 0 && dr < 8 && dc < 8; {
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
