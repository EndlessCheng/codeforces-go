package main

// github.com/EndlessCheng/codeforces-go
func mostVisited(n int, a []int) (ans []int) {
	if x, y := a[0], a[len(a)-1]; x <= y {
		for i := x; i <= y; i++ {
			ans = append(ans, i)
		}
	} else {
		for i := 1; i <= y; i++ {
			ans = append(ans, i)
		}
		for i := x; i <= n; i++ {
			ans = append(ans, i)
		}
	}
	return
}
