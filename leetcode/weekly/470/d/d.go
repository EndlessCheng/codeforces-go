package main

import "strconv"

// https://space.bilibili.com/206214

// 返回两个 1~9 的整数和为 target 的方案数
func twoSumWays(target int) int {
	return max(min(target-1, 19-target), 0) // 保证结果非负
}

func countNoZeroPairs(n int64) int64 {
	s := strconv.FormatInt(n, 10)
	m := len(s)
	memo := make([][2][2]int, m)
	for i := range memo {
		memo[i] = [2][2]int{{-1, -1}, {-1, -1}} // -1 表示没有计算过
	}

	// borrow = 1 表示被低位（i+1）借位
	// isNum = 1 表示之前填的数位，两个数都不为 0（无前导零）
	var dfs func(int, int, int) int
	dfs = func(i, borrowed, isNum int) (res int) {
		if i < 0 {
			// borrowed 必须为 0
			return borrowed ^ 1
		}

		p := &memo[i][borrowed][isNum]
		if *p >= 0 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化

		d := int(s[i]-'0') - borrowed
		// 其中一个数必须填前导零
		if isNum == 0 {
			// 在 i > 0 的情况下，另一个数必须不为 0（否则可以为 0，即两个数的最高位都是 0）
			if i > 0 && d == 0 {
				return 0
			}
			// 如果 d < 0，必须向高位借位
			return dfs(i-1, isNeg(d), 0)
		}

		// 令其中一个数从当前位置开始往左都是 0（前导零）
		if i < m-1 {
			if d != 0 { // 另一个数不为 0
				res = dfs(i-1, isNeg(d), 0) * 2 // 根据对称性乘以 2
			} else if i == 0 { // 最高位被借走
				res = 1 // 两个数都是 0
			} // else res = 0
		}

		// 两个数位都不为 0
		res += dfs(i-1, 0, 1) * twoSumWays(d)    // 不向 i-1 借位
		res += dfs(i-1, 1, 1) * twoSumWays(d+10) // 向 i-1 借位
		return
	}

	return int64(dfs(m-1, 0, 1))
}

// 返回 d 是否为负数
func isNeg(d int) int {
	if d < 0 {
		return 1
	}
	return 0
}
