package main

// https://space.bilibili.com/206214
func removeStones(stones [][]int) int {
	// 哈希表并查集
	fa := map[int]int{}
	var find func(int) int
	find = func(x int) int {
		fx, ok := fa[x]
		if !ok {
			return x
		}
		if fx != x {
			fa[x] = find(fx)
			return fa[x]
		}
		return x
	}

	const offset int = 3e4
	for _, p := range stones {
		fa[find(p[0])] = find(p[1] + offset)
	}

	// 记录不同的代表元
	set := map[int]struct{}{}
	for _, p := range stones {
		set[find(p[0])] = struct{}{}
	}

	return len(stones) - len(set)
}
