package main

// github.com/EndlessCheng/codeforces-go
func findRotation(a, tar [][]int) bool {
o:
	for k := 0; k < 4; k++ {
		a = rotate(a)
		for i, r := range a {
			for j, v := range r {
				if v != tar[i][j] {
					continue o
				}
			}
		}
		return true
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
