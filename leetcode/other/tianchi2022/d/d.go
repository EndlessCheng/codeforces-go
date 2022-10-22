package main

// https://space.bilibili.com/206214
func brilliantSurprise(a [][]int, lim int) (ans int) {
	dp := make([]int, lim+1)
	var f func([][]int, []int)
	f = func(a [][]int, tot []int) {
		if len(a) == 1 {
			s := 0
			for i, v := range a[0] {
				if i >= lim {
					break
				}
				s += v
				ans = max(ans, dp[lim-(i+1)]+s)
			}
			return
		}

		tmp := append([]int{}, dp...)

		m := len(a) / 2
		for i, r := range a[:m] {
			for j := lim; j >= len(r); j-- {
				dp[j] = max(dp[j], dp[j-len(r)]+tot[i])
			}
		}
		f(a[m:], tot[m:])

		dp = tmp
		for i, r := range a[m:] {
			for j := lim; j >= len(r); j-- {
				dp[j] = max(dp[j], dp[j-len(r)]+tot[m+i])
			}
		}
		f(a[:m], tot[:m])
	}

	tot := make([]int, len(a))
	for i, r := range a {
		for _, v := range r {
			tot[i] += v
		}
	}
	f(a, tot)
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
