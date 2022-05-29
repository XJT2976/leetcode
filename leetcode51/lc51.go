package leetcode51

import "strings"

func SolveNQueens(n int) [][]string {
	chessBoard := make([][]string, n)
	res := make([][]string, 0)
	for i := 0; i < n; i++ {
		chessBoard[i] = make([]string, n)
		for j := 0; j < n; j++ {
			chessBoard[i][j] = "."
		}
	}

	backtracking(chessBoard, 0, &res)
	return res
}

func backtracking(chessBoard [][]string, row int, res *[][]string) {
	if row == len(chessBoard) {
		tmp := make([]string, len(chessBoard))
		for i := 0; i < len(chessBoard); i++ {
			s := strings.Join(chessBoard[i], "")
			tmp[i] = s
		}
		*res = append(*res, tmp)
	}
	for col := 0; col < len(chessBoard); col++ {
		if isValidSquare(chessBoard, row, col) {
			chessBoard[row][col] = "Q"
			backtracking(chessBoard, row+1, res)
			chessBoard[row][col] = "."
		}
	}
}

func isValidSquare(chessBoard [][]string, row, col int) bool {
	for i := 0; i < row; i++ {
		if chessBoard[i][col] == "Q" {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessBoard[i][j] == "Q" {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < len(chessBoard); i, j = i-1, j+1 {
		if chessBoard[i][j] == "Q" {
			return false
		}
	}

	return true
}
