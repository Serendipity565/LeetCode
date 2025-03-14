package Spring2024Sprint100

func rotate48(matrix [][]int) {
	n := len(matrix)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[j][n-i-1] = matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = a[i][j]
		}
	}
	return
}
