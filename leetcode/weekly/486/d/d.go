package main

// https://space.bilibili.com/206214
const mx = 50

var comb [mx + 1][mx + 1]int64

func init() {
	// 预处理组合数
	for i := range comb {
		comb[i][0] = 1
		for j := 1; j <= i; j++ {
			comb[i][j] = comb[i-1][j-1] + comb[i-1][j]
		}
	}
}

func nthSmallest(n int64, k int) (ans int64) {
	for i := mx - 1; k > 0; i-- {
		c := comb[i][k] // 第 i 位填 0 的方案数
		if n > c { // n 比较大，第 i 位必须填 1
			n -= c
			ans |= 1 << i
			k-- // 维护剩余的 1 的个数
		}
	}
	return
}
