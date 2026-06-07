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

func sumOfGoodIntegers2(n, k int) int {
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

// 计算小于 high 的正整数中，AND n 等于 0 的数之和
func calc(high, n int) (res int) {
	m := bits.Len(uint(high))
	freeMask := (1<<m - 1) &^ n
	freeCnt := bits.OnesCount(uint(freeMask))
	prefix := 0

	for i := m - 1; i >= 0; i-- {
		if n>>i&1 == 0 {
			freeCnt--
			freeMask ^= 1 << i
		}
		if high>>i&1 > 0 {
			// 这一位填 0
			res += prefix << freeCnt            // 前缀的贡献：后面 freeCnt 个位置，0 和 1 随便填
			res += 1 << freeCnt >> 1 * freeMask // 后缀的贡献：每个 free 位置固定为 1 时，其余 freeCnt-1 个位置 0 和 1 随便填

			// 这一位填 1，继续计算
			if n>>i&1 > 0 { // 这一位不能填 1
				break
			}
			prefix |= 1 << i
		}
	}

	return
}

func sumOfGoodIntegers(n, k int) int {
	low := max(n-k, 1)
	high := n + k
	return calc(high+1, n) - calc(low, n)
}
