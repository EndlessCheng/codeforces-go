package main

// https://space.bilibili.com/206214
func countComponents(nums []int, threshold int) int {
	n := len(nums)
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	// 非递归并查集
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	// 记录每个数的下标
	idx := make([]int, threshold+1)
	for i, x := range nums {
		if x <= threshold {
			idx[x] = i + 1 // 这里 +1 了，下面减掉
		}
	}

	for g := 1; g <= threshold; g++ {
		minX := -1
		for x := g; x <= threshold; x += g {
			if idx[x] > 0 { // idx[x] == 0 表示不存在
				minX = x
				break
			}
		}
		if minX < 0 {
			continue
		}
		fi := find(idx[minX] - 1)
		for y := minX + g; y <= threshold && y <= g*threshold/minX; y += g {
			if idx[y] > 0 {
				fj := find(idx[y] - 1)
				if fj != fi {
					fa[fj] = fi // 合并 idx[x] 和 idx[y]
					n--
				}
			}
		}
	}
	return n
}
