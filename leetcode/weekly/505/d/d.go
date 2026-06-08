package main

import (
	"math"
	"sort"
)

// https://space.bilibili.com/206214
type pair struct{ f, cnt int }

// 相等的时候，子数组个数 cnt 更大的劣
func less(a, b pair) bool {
	return a.f < b.f || a.f == b.f && a.cnt > b.cnt
}

func maximumSum(nums []int, m, l, r int) int64 {
	n := len(nums)
	s := make([]int, n+1) // nums 的前缀和
	posSum := 0           // nums 中的正数之和
	for i, x := range nums {
		s[i+1] = s[i] + x
		if x > 0 {
			posSum += x
		}
	}

	// 没有 m 约束，但每选一个子数组就要把元素和减少 k
	dpWithoutLimit := func(k int) pair {
		f := make([]pair, n+1)
		q := []int{}
		res := pair{math.MinInt, 0}

		for i := l; i <= n; i++ {
			// 1. 入
			j := i - l
			v := pair{f[j].f - s[j], f[j].cnt}
			for len(q) > 0 && less(pair{f[q[len(q)-1]].f - s[q[len(q)-1]], f[q[len(q)-1]].cnt}, v) {
				q = q[:len(q)-1]
			}
			q = append(q, j)

			// 2. 更新答案
			choose := pair{f[q[0]].f - s[q[0]] + s[i] - k, f[q[0]].cnt + 1}
			if less(res, choose) {
				// choose 保证我们至少选了一个子数组
				res = choose
			}

			// 更新 DP
			if less(f[i-1], choose) {
				f[i] = choose
			} else { // 不选
				f[i] = f[i-1]
			}

			// 3. 出，下一轮循环队首离开窗口
			if q[0] <= i-r {
				q = q[1:]
			}
		}

		return res
	}

	res0 := dpWithoutLimit(0)
	if res0.cnt <= m { // 直接满足题目要求
		return int64(res0.f)
	}

	// 现在专注于解决「选恰好 m 个子数组」的问题
	ans := 0
	sort.Search(posSum, func(k int) bool {
		k++
		res := dpWithoutLimit(k)
		if res.cnt <= m {
			ans = res.f + m*k // 直接算，二分会缩小到凸函数中的 x=m 所在的那条线段
			return true
		}
		return false
	})
	return int64(ans)
}
