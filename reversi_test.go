package main

import (
	"Reversi/mypkg"
	"testing"
)

func TestReversi01(t *testing.T) {
	col := 0
	row := 0
	if board[row][col] != board {
		t.Errorf("Error")
	}
}
