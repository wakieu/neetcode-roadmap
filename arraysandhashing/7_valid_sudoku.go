package arraysandhashing

// https://leetcode.com/problems/valid-sudoku/

func IsValidSudoku(board [][]byte) bool {
	for x, line := range board {
		for y, num := range line {
			if num == '.' {
				continue
			}

			xQuadInit := (x / 3) * 3
			yQuadInit := (y / 3) * 3

			board[x][y] = '.'

			if num == board[xQuadInit+0][yQuadInit+0] ||
				num == board[xQuadInit+1][yQuadInit+0] ||
				num == board[xQuadInit+2][yQuadInit+0] ||
				num == board[xQuadInit+0][yQuadInit+1] ||
				num == board[xQuadInit+1][yQuadInit+1] ||
				num == board[xQuadInit+2][yQuadInit+1] ||
				num == board[xQuadInit+0][yQuadInit+2] ||
				num == board[xQuadInit+1][yQuadInit+2] ||
				num == board[xQuadInit+2][yQuadInit+2] ||
				Contains(line, num) ||
				num == board[0][y] ||
				num == board[1][y] ||
				num == board[2][y] ||
				num == board[3][y] ||
				num == board[4][y] ||
				num == board[5][y] ||
				num == board[6][y] ||
				num == board[7][y] ||
				num == board[8][y] {
				return false
			}

			board[x][y] = num
		}
	}
	return true
}

func Contains(a []byte, x byte) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
