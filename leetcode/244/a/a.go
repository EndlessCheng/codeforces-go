package main

// github.com/EndlessCheng/codeforces-go
func findRotation(a, b [][]int) bool {
	for k := 0; k < 4; k++ {
		a = rotate(a)
		if equalMatrix(a, b) {
			return true
		}
	}
	return false
}

func rotate(a [][]int) [][]int {
	n := len(a)
	b := make([][]int, n)
	for i := range b {
		b[i] = make([]int, n)
	}
	for i, r := range a {
		for j, v := range r {
			b[j][n-1-i] = v
		}
	}
	return b
}

func equalMatrix(a, b [][]int) bool {
	for i, r := range a {
		for j, v := range r {
			if v != b[i][j] {
				return false
			}
		}
	}
	return true
}
