package main

import (
	"math/bits"
)

// https://space.bilibili.com/206214
const mx = 31
var c [mx][mx]int

func init() {
	for i := 0; i < mx; i++ {
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
}

func waysToReachStair(k int) (ans int) {
	for j := bits.Len(uint(max(k-1, 0))); 1<<j-k <= j+1; j++ {
		ans += c[j+1][1<<j-k]
	}
	return
}

func main() {
	for i := 0; i <= 1e9; i++ {
		waysToReachStair(i)
	}
}

func waysToReachStair2(k int) int {
	type args struct {
		i, j    int
		preDown bool
	}
	memo := map[args]int{}
	var dfs func(int, int, bool) int
	dfs = func(i, j int, preDown bool) int {
		if i > k+1 {
			return 0
		}
		p := args{i, j, preDown}
		if v, ok := memo[p]; ok {
			return v
		}
		res := dfs(i+1<<j, j+1, false)
		if !preDown {
			res += dfs(i-1, j, true)
		}
		if i == k {
			res++
		}
		memo[p] = res
		return res
	}
	return dfs(1, 0, false)
}
