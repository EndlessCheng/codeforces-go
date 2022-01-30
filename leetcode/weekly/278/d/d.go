package main

// github.com/EndlessCheng/codeforces-go
func groupStrings(words []string) (ans []int) {
	// 并查集模板
	fa := map[int]int{}
	size := map[int]int{}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
			return fa[x]
		}
		return x
	}
	groups, maxSize := len(words), 0
	merge := func(x, y int) {
		if _, ok := fa[y]; !ok {
			return
		}
		x, y = find(x), find(y)
		if x == y {
			return
		}
		fa[x] = y
		size[y] += size[x]
		maxSize = max(maxSize, size[y])
		groups--
	}

	for _, word := range words {
		x := 0
		for _, ch := range word {
			x |= 1 << (ch - 'a') // 计算 word 的二进制表示
		}
		fa[x] = x // 添加至并查集
		size[x]++
		maxSize = max(maxSize, size[x])
		if size[x] > 1 {
			groups--
		}
	}

	for x := range fa { // 枚举所有字符串（二进制表示）
		for i := 0; i < 26; i++ {
			if x>>i&1 == 1 {
				merge(x, x^1<<i) // 删除字符 i
				for j := 0; j < 26; j++ {
					if x>>j&1 == 0 {
						merge(x, x^1<<i|1<<j) // 替换字符 i 为 j
					}
				}
			}
		}
	}
	return []int{groups, maxSize}
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
