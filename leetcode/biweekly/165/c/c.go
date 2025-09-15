package main

import (
	"math/rand"
	"slices"
)

// https://space.bilibili.com/206214
func generateSchedule1(n int) [][]int {
	if n < 5 {
		return nil
	}

	ans := make([][]int, 0, n*(n-1)) // 预分配空间

	// 处理 d=2,3,...,n-2
	for d := 2; d < n-1; d++ {
		for i := range n {
			ans = append(ans, []int{i, (i + d) % n})
		}
	}

	// 交错排列 d=1 与 d=n-1（或者说 d=-1）
	for i := range n {
		ans = append(ans, []int{i, (i + 1) % n}, []int{(i + n - 1) % n, (i + n - 2) % n})
	}

	return ans
}

func gen(perm [][]int) (ans [][]int) {
	ans = append(ans, perm[0])
	perm = perm[1:]
next:
	for len(perm) > 0 {
		// 倒着遍历，这样删除的时候 i 更大，移动的数据少
		for i, p := range slices.Backward(perm) {
			last := ans[len(ans)-1]
			if p[0] != last[0] && p[0] != last[1] && p[1] != last[0] && p[1] != last[1] {
				// p 和上一场比赛无冲突
				ans = append(ans, p)
				perm = append(perm[:i], perm[i+1:]...) // 删除 perm[i]
				continue next // 找下一场比赛
			}
		}
		return nil
	}
	return
}

func generateSchedule(n int) [][]int {
	if n < 5 {
		return nil
	}

	// 赛程排列
	perm := make([][]int, 0, n*(n-1))
	for i := range n {
		for j := range n {
			if j != i {
				perm = append(perm, []int{i, j})
			}
		}
	}

	for {
		rand.Shuffle(len(perm), func(i, j int) { perm[i], perm[j] = perm[j], perm[i] })
		if ans := gen(slices.Clone(perm)); ans != nil {
			return ans
		}
	}
}
