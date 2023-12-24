package main

import "math"

// https://space.bilibili.com/206214
func minimumCost(source, target string, original, changed []string, cost []int) int64 {
	const inf = math.MaxInt / 2
	type node struct {
		son [26]*node
		sid int // 字符串的编号
	}
	root := &node{}
	sid := 0
	put := func(s string) int {
		o := root
		for _, b := range s {
			b -= 'a'
			if o.son[b] == nil {
				o.son[b] = &node{sid: -1}
			}
			o = o.son[b]
		}
		if o.sid < 0 {
			o.sid = sid
			sid++
		}
		return o.sid
	}

	// 初始化距离矩阵
	m := len(cost)
	dis := make([][]int, m*2)
	for i := range dis {
		dis[i] = make([]int, m*2)
		for j := range dis[i] {
			if j != i {
				dis[i][j] = inf
			}
		}
	}
	for i, c := range cost {
		x := put(original[i])
		y := put(changed[i])
		dis[x][y] = min(dis[x][y], c)
	}

	// Floyd 求任意两点最短路
	for k := 0; k < sid; k++ {
		for i := 0; i < sid; i++ {
			if dis[i][k] == inf {
				continue
			}
			for j := 0; j < sid; j++ {
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}

	n := len(source)
	f := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		if source[i] == target[i] {
			f[i] = f[i+1] // 不修改 source[i]
		} else {
			f[i] = inf
		}
		p, q := root, root
		for j := i; j < n; j++ {
			p = p.son[source[j]-'a']
			q = q.son[target[j]-'a']
			if p == nil || q == nil {
				break
			}
			if p.sid >= 0 && q.sid >= 0 {
				// 修改从 i 到 j 的这一段
				f[i] = min(f[i], dis[p.sid][q.sid]+f[j+1])
			}
		}
	}
	if f[0] == inf {
		return -1
	}
	return int64(f[0])
}
