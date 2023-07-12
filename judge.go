package mypkg

func judge() bool {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == emptyCell {
				return false
			}
		}
	}
	return true
}
