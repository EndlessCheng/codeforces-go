package main

// github.com/EndlessCheng/codeforces-go
func rotateGrid(a [][]int, k int) [][]int {
	n, m := len(a), len(a[0])
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	lim := min(n, m) / 2
	for d := 0; d < lim; d++ {
		b := make([]int, 0, (n+m-d*4-2)*2)
		for j := d; j < m-d; j++ {
			b = append(b, a[d][j])
		}
		for i := d + 1; i < n-d; i++ {
			b = append(b, a[i][m-1-d])
		}
		for j := m - d - 2; j >= d; j-- {
			b = append(b, a[n-1-d][j])
		}
		for i := n - d - 2; i > d; i-- {
			b = append(b, a[i][d])
		}
		shift := k % len(b)
		b = append(b[shift:], b[:shift]...)
		j := 0
		for i := d; i < m-d; i++ {
			ans[d][i] = b[j]
			j++
		}
		for i := d + 1; i < n-d; i++ {
			ans[i][m-1-d] = b[j]
			j++
		}
		for i := m - d - 2; i >= d; i-- {
			ans[n-1-d][i] = b[j]
			j++
		}
		for i := n - d - 2; i > d; i-- {
			ans[i][d] = b[j]
			j++
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
