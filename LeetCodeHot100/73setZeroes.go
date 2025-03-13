package LeetCodeHot100

func setZeroes(matrix [][]int) {
	var m int = len(matrix)
	var n int = len(matrix[0])

	flagl := false
	flagc := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			flagc = true
			break
		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			flagl = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	if flagc == true {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	if flagl == true {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}

}
