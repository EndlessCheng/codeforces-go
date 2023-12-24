package main

// https://space.bilibili.com/206214
func minimumCost(source, target string, original, changed []byte, cost []int) (ans int64) {
	dis := [26][26]int{}
	for i := range dis {
		for j := range dis[i] {
			if j != i {
				dis[i][j] = 1e13
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
			for j := range dis {
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}

	for i, b := range source {
		ans += int64(dis[b-'a'][target[i]-'a'])
	}
	if ans >= 1e13 {
		return -1
	}
	return
}
