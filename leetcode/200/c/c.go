package main

// github.com/EndlessCheng/codeforces-go
func minSwaps(mat [][]int) (ans int) {
	n := len(mat)
	a := make([]int, n)
	for i, row := range mat {
		for j := len(row) - 1; j >= 0; j-- {
			if row[j] > 0 {
				a[i] = j + 1
				break
			}
		}
	}
	for i := range a {
		j := i
		for ; j < n && a[j] > i+1; j++ {
		}
		if j == n {
			return -1
		}
		ans += j - i
		for ; j > i; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
	return
}
