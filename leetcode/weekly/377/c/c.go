package main

import "math"

// https://space.bilibili.com/206214
func minimumCost(source, target string, original, changed []byte, cost []int) (ans int64) {
	const inf = math.MaxInt / 2
	dis := [26][26]int{}
	for i := range dis {
		for j := range dis[i] {
			if j != i {
				dis[i][j] = inf
			}
		}
	}
	for i, c := range cost {
		x := original[i] - 'a'
		y := changed[i] - 'a'
		dis[x][y] = min(dis[x][y], c)
	}
	for k := range dis {
		for i := range dis {
			if dis[i][k] == inf {
				continue // 巨大优化！
			}
			for j := range dis {
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}

	for i, b := range source {
		d := dis[b-'a'][target[i]-'a']
		if d == inf {
			return -1
		}
		ans += int64(d)
	}
	return
}
