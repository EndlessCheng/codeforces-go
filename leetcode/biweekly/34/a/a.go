package main

// github.com/EndlessCheng/codeforces-go
func diagonalSum(a [][]int) (ans int) {
	for i, row := range a {
		ans += row[i]
		if i != len(row)-1-i {
			ans += row[len(row)-1-i]
		}
	}
	return
}
