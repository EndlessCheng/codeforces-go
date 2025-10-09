package main

// https://space.bilibili.com/206214
var preSum [1001]int

func init() {
	for i := 1; i <= 1000; i++ {
		var dfs func(int, int) bool
		dfs = func(val, sum int) bool {
			if val == 0 { // 递归终点
				return sum == i // i 符合要求
			}
			for x, pow10 := 0, 1; val > 0; val /= 10 {
				x += val % 10 * pow10
				if dfs(val/10, sum+x) {
					return true
				}
				pow10 *= 10
			}
			return false
		}
		preSum[i] = preSum[i-1]
		if dfs(i*i, 0) { // i 符合要求
			preSum[i] += i * i // 计算前缀和
		}
	}
}

func punishmentNumber(n int) int {
	return preSum[n]
}
