package main

import (
	"math/bits"
)

// https://space.bilibili.com/206214
func supersequences(words []string) [][]int {
	// 收集有哪些字母，同时建图
	all, mask2 := 0, 0
	g := [26][]int{}
	for _, s := range words {
		x, y := int(s[0]-'a'), int(s[1]-'a')
		all |= 1<<x | 1<<y
		if x == y {
			mask2 |= 1 << x
		}
		g[x] = append(g[x], y)
	}

	// 判断是否有环
	hasCycle := func(sub int) bool {
		color := [26]int8{}
		var dfs func(int) bool
		dfs = func(x int) bool {
			color[x] = 1
			for _, y := range g[x] {
				// 只遍历在 sub 中的字母
				if sub>>y&1 == 0 {
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
			// 只遍历在 sub 中的字母
			if c == 0 && sub>>i&1 > 0 && dfs(i) {
				return true
			}
		}
		return false
	}

	set := map[int]struct{}{}
	maxSize := 0
	mask1 := all ^ mask2
	// 枚举 mask1 的所有子集 sub
	for sub, ok := mask1, true; ok; ok = sub != mask1 {
		size := bits.OnesCount(uint(sub))
		// 剪枝：如果 size < maxSize 就不需要判断了
		if size >= maxSize && !hasCycle(sub) {
			if size > maxSize {
				maxSize = size
				clear(set)
			}
			set[sub] = struct{}{}
		}
		sub = (sub - 1) & mask1
	}

	ans := make([][]int, 0, len(set)) // 预分配空间
	for sub := range set {
		cnt := make([]int, 26)
		for i := range cnt {
			cnt[i] = all>>i&1 + (all^sub)>>i&1
		}
		ans = append(ans, cnt)
	}
	return ans
}

func supersequences2(words []string) (ans [][]int) {
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

	// 快速跳转到下一个要枚举的字母
	nxt := [27]int{26: 26}
	for i := 25; i >= 0; i-- {
		if all>>i&1 > 0 {
			nxt[i] = i
		} else {
			nxt[i] = nxt[i+1]
		}
	}

	type pair struct{ i, sub int }
	q := []pair{{nxt[0], 0}}
	for {
		for _, p := range q {
			if !hasCycle(p.sub) {
				cnt := make([]int, 26)
				for i := range cnt {
					cnt[i] = all>>i&1 + p.sub>>i&1
				}
				ans = append(ans, cnt)
			}
		}
		if ans != nil {
			return
		}
		tmp := q
		q = nil
		for _, p := range tmp {
			for j := p.i; j < 26; j = nxt[j+1] {
				q = append(q, pair{nxt[j+1], p.sub | 1<<j})
			}
		}
	}
}
