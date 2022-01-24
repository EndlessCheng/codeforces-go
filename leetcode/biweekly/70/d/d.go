package main

// 按照题意模拟 + O(1) 额外空间

// github.com/EndlessCheng/codeforces-go
func numberOfWays(corridor string) int {
	ans, cntS, pre := 1, 0, 0
	for i, ch := range corridor {
		if ch == 'S' {
			// 对第 3,5,7,... 个座位，可以在其到其左侧最近座位之间的任意一个放置屏风
			cntS++
			if cntS >= 3 && cntS%2 == 1 {
				ans = ans * (i - pre) % (1e9 + 7)
			}
			pre = i // 记录上一个座位的位置
		}
	}
	if cntS == 0 || cntS%2 == 1 { // 座位个数不能为 0 或奇数
		return 0
	}
	return ans
}
