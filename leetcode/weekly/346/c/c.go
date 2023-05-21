package main

import "strconv"

// https://space.bilibili.com/206214
var preSum [1001]int

func init() {
	for i := 1; i <= 1000; i++ {
		s := strconv.Itoa(i * i)
		n := len(s)
		var f func(int, int) bool
		f = func(p, sum int) bool {
			if p == n { // 递归终点
				return sum == i // i 是惩罚数
			}
			x := 0
			for j := p; j < n; j++ { // 从 s[p] 到 s[j] 组成的子串
				x = x*10 + int(s[j]-'0') // 对应的整数值
				if f(j+1, sum+x) {
					return true
				}
			}
			return false
		}
		preSum[i] = preSum[i-1]
		if f(0, 0) { // i 是惩罚数
			preSum[i] += i * i // 计算前缀和
		}
	}
}

func punishmentNumber(n int) (ans int) {
	return preSum[n]
}
