package main

// github.com/EndlessCheng/codeforces-go
func numberOfWays(corridor string) int {
	const mod = 1_000_000_007
	ans, cntS, lastS := 1, 0, 0
	for i, ch := range corridor {
		if ch == 'S' {
			cntS++
			// 对于第 3,5,7,... 个座位，可以在其到其左侧最近座位之间的任意空隙放屏风
			if cntS >= 3 && cntS%2 > 0 {
				ans = ans * (i - lastS) % mod
			}
			lastS = i // 记录上一个座位的位置
		}
	}
	if cntS == 0 || cntS%2 > 0 { // 座位个数不能为 0 或奇数
		return 0
	}
	return ans
}
