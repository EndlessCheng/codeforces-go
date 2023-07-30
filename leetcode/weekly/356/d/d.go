package main

// https://space.bilibili.com/206214
const mod = 1_000_000_007

func countSteppingNumbers(low, high string) int {
	calc := func(s string) int {
		dp := make([][10]int, len(s))
		for i := range dp {
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		var dfs func(int, int, bool, bool) int
		dfs = func(p, pre int, isLimit, isNum bool) (res int) {
			if p == len(s) {
				if isNum {
					return 1
				}
				return 0
			}
			if !isLimit && isNum {
				dv := &dp[p][pre]
				if *dv >= 0 {
					return *dv
				}
				defer func() { *dv = res }()
			}
			if !isNum {
				res += dfs(p+1, pre, false, false)
			}
			up := 9
			if isLimit {
				up = int(s[p] - '0')
			}
			d := 0
			if !isNum {
				d = 1
			}
			for ; d <= up; d++ {
				if !isNum || abs(d-pre) == 1 {
					res += dfs(p+1, d, isLimit && d == up, true)
					res %= mod
				}
			}
			return
		}
		res := dfs(0, 0, true, false)
		return res
	}
	ansUpper := calc(high)
	ansLower := calc(low)
	ans := ansUpper - ansLower
	for i := 1; i < len(low); i++ {
		pre, cur := int(low[i-1]), int(low[i])
		if abs(pre-cur) != 1 {
			goto o
		}
	}
	ans++
o:
	ans = (ans%mod + mod) % mod
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
