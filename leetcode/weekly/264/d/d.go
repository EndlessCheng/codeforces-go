package main

// https://space.bilibili.com/206214
func minimumTime(n int, relations [][]int, time []int) (ans int) {
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range relations {
		x, y := e[0]-1, e[1]-1
		g[x] = append(g[x], y) // 建图
		deg[y]++ // 可以理解为 y 的先修课的个数
	}

	q := []int{}
	for i, d := range deg {
		if d == 0 { // 没有先修课
			q = append(q, i)
		}
	}
	f := make([]int, n)
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		// x 出队，意味着 x 的所有先修课都上完了
		f[x] += time[x] // 加上当前课程的时间，就得到了最终的 f[x]
		ans = max(ans, f[x])
		for _, y := range g[x] { // 遍历 x 的邻居
			f[y] = max(f[y], f[x]) // 更新 f[y] 的所有先修课程耗时的最大值
			if deg[y]--; deg[y] == 0 { // y 的先修课已上完
				q = append(q, y)
			}
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
