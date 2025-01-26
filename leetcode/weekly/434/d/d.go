package main

import (
	"math"
	"math/bits"
)

// https://space.bilibili.com/206214
func supersequences(words []string) [][]int {
	// 收集有哪些字母，同时建图
	all := 0
	g := [26][]int{}
	for _, s := range words {
		x, y := int(s[0]-'a'), int(s[1]-'a')
		all |= 1<<x | 1<<y
		g[x] = append(g[x], y)
	}

	// 判断是否有环
	hasCycle := func(sub int) bool {
		color := [26]int8{}
		var dfs func(int) bool
		dfs = func(x int) bool {
			color[x] = 1
			for _, y := range g[x] {
				// 只遍历不在 sub 中的字母
				if sub>>y&1 > 0 {
					continue
				}
				if color[y] == 1 || color[y] == 0 && dfs(y) {
					return true
				}
			}
			color[x] = 2
			return false
		}
		for i, c := range color {
			// 只遍历不在 sub 中的字母
			if c == 0 && sub>>i&1 == 0 && dfs(i) {
				return true
			}
		}
		return false
	}

	set := map[int]struct{}{}
	minSize := math.MaxInt
	for sub, ok := all, true; ok; ok = sub != all {
		size := bits.OnesCount(uint(sub))
		// 剪枝：如果 size > min_size 就不需要判断了
		if size <= minSize && !hasCycle(sub) {
			if size < minSize {
				minSize = size
				clear(set)
			}
			set[sub] = struct{}{}
		}
		sub = (sub - 1) & all
	}

	ans := make([][]int, 0, len(set)) // 预分配空间
	for sub := range set {
		cnt := make([]int, 26)
		for i := range cnt {
			cnt[i] = all>>i&1 + sub>>i&1
		}
		ans = append(ans, cnt)
	}
	return ans
}
