package leetcode37

func SolveSudoku(board [][]byte) {
	backtracking(board)
}

func backtracking(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == '.' {
				for k := byte('1'); k <= '9'; k++ {
					if isValidNumber(board, i, j, k) {
						board[i][j] = k
						if backtracking(board) {
							return true
						}
						board[i][j] = '.'
					}
				}

				return false
			}
		}
	}

	return true
}

func isValidNumber(board [][]byte, row, col int, num byte) bool {
	for i := 0; i < len(board); i++ {
		if board[i][col] == num {
			return false
		}
	}

	for i := 0; i < len(board); i++ {
		if board[row][i] == num {
			return false
		}
	}

	for i := row / 3 * 3; i < row/3*3+3; i++ {
		for j := col / 3 * 3; j < col/3*3+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	return true
}
