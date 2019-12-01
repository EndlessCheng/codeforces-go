package main

import "fmt"

func palindromePartition(str string, k int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	const mx int = 105
	dp := [mx][mx][mx]int{}
	vis := [mx][mx][mx]bool{}
	var f func(st, end int, left int) int
	f = func(st, end int, seg int) int {
		s := str[st:end]
		if len(s) < seg {
			return -1
		}
		if len(s) == 1 {
			return 0
		}
		if vis[st][end][seg] {
			return dp[st][end][seg]
		}
		vis[st][end][seg] = true

		if seg == 1 {
			ans := 0
			for i, j := 0, len(s)-1; i < j; {
				if s[i] != s[j] {
					ans++
				}
				i++
				j--
			}
			dp[st][end][seg] = ans
			return ans
		}

		ans := int(1e9)
		for mid := st + 1; mid < end; mid++ {
			for segL := 1; segL < seg; segL++ {
				ansL := f(st, mid, segL)
				if ansL == -1 {
					continue
				}
				ansR := f(mid, end, seg-segL)
				if ansR == -1 {
					continue
				}
				ans = min(ans, ansL+ansR)
			}
		}
		dp[st][end][seg] = ans
		return ans
	}

	return f(0, len(str), k)
}

// GoLand debug 失败！
func main() {
	fmt.Println(palindromePartition("abc", 2))
}
