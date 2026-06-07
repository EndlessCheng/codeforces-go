package main

import "math/bits"

// https://space.bilibili.com/206214
func sumOfGoodIntegers1(n, k int) (ans int) {
	for x := max(n-k, 1); x <= n+k; x++ {
		if n&x == 0 {
			ans += x
		}
	}
	return
}

func sumOfGoodIntegers(n, k int) int {
	low := max(n-k, 1)
	high := n + k
	m := bits.Len(uint(high))
	type pair struct{ cnt, sum int }
	memo := make([]pair, m)
	for i := range memo {
		memo[i].cnt = -1
	}

	// dfs 返回两个数：子树合法数字个数，子树和
	var dfs func(int, bool, bool) pair
	dfs = func(i int, limitLow, limitHigh bool) (res pair) {
		if i < 0 {
			return pair{1, 0} // 如果没有特殊约束，能递归到终点的都是合法数字
		}
		if !limitLow && !limitHigh {
			p := &memo[i]
			if p.cnt >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow {
			lo = low >> i & 1
		}
		hi := 1
		if limitHigh {
			hi = high >> i & 1
		}

		for d := lo; d <= hi; d++ {
			bit := d << i
			if n&bit > 0 { // 不满足要求
				continue
			}
			sub := dfs(i-1, limitLow && d == lo, limitHigh && d == hi)
			res.cnt += sub.cnt       // 累加子树的合法数字个数
			res.sum += sub.sum       // 累加子树的和
			res.sum += bit * sub.cnt // bit 会出现在 sub.cnt 个数中（贡献法）
		}
		return
	}

	return dfs(m-1, true, true).sum
}
