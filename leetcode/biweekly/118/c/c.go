package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func minimumCoins(prices []int) int {
	n := len(prices)
	for i := (n+1)/2 - 1; i > 0; i-- {
		prices[i-1] += slices.Min(prices[i : i*2+1])
	}
	return prices[0]
}

func minimumCoins1(prices []int) int {
	n := len(prices)
	memo := make([]int, (n+1)/2)
	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i*2 >= n {
			return prices[i-1] // i 从 1 开始
		}
		p := &memo[i]
		if *p != 0 { // 之前算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		res = math.MaxInt
		for j := i + 1; j <= i*2+1; j++ {
			res = min(res, dfs(j))
		}
		return res + prices[i-1]
	}
	return dfs(1)
}
