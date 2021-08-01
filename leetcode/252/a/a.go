package main

// 统计因子个数

// github.com/EndlessCheng/codeforces-go
func isThree(n int) bool {
	cnt := 0
	for d := 1; d*d <= n; d++ {
		if n%d == 0 {
			cnt++
			if d*d < n {
				cnt++
			}
		}
	}
	return cnt == 3
}
