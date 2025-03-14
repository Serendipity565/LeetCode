package Spring2024Sprint100

func count(x, y, m, n int, board [][]int) int {
	cnt := 0
	directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && nx < m && ny >= 0 && ny < n {
			cnt += board[nx][ny]
		}
	}
	return cnt
}

func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	m, n := len(board), len(board[0])
	ve := make([][]int, m)
	for i := range ve {
		ve[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveNeighbors := count(i, j, m, n, board)
			if board[i][j] == 1 && (liveNeighbors < 2 || liveNeighbors > 3) {
				ve[i][j] = 0
			} else if board[i][j] == 1 || (board[i][j] == 0 && liveNeighbors == 3) {
				ve[i][j] = 1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = ve[i][j]
		}
	}
}
