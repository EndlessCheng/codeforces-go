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

	// borrowed = 1 表示被低位（i+1）借了个 1
	// isNum = 1 表示之前填的数位，两个数都无前导零
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

		// 情况一：两个数位都不为 0
		if isNum > 0 {
			res = dfs(i-1, 0, 1) * twoSumWays(d)     // 不向高位借 1
			res += dfs(i-1, 1, 1) * twoSumWays(d+10) // 向高位借 1
		}

		// 情况二：其中一个数位填前导零
		if i < m-1 { // 不能是最低位
			if d != 0 {
				// 如果 d < 0，必须向高位借 1
				// 如果 isNum = 1，根据对称性，方案数要乘以 2
				res += dfs(i-1, isNeg(d), 0) * (isNum + 1)
			} else if i == 0 { // 两个数位都填 0，只有当 i = 0 的时候才有解
				res++
			}
		}
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
