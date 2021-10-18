package main

/* BFS 求次短路，附详细注释
 */

// github.com/EndlessCheng/codeforces-go
func secondMinimum(n int, edges [][]int, time, change int) int {
	g := make([][]int, n)
	for _, e := range edges { // 建图
		v, w := e[0]-1, e[1]-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	// 传入当前时间 d，返回到达下一个节点的时间
	next := func(d int) int {
		if times := d / change; times%2 == 1 { // 如果红绿灯切换次数为奇数，则现在是红灯
			return (times+1)*change + time // 等绿灯再出发
		}
		return d + time // 绿灯，直接出发
	}

	dis := make([][2]int, n) // 距离数组同时存 [最短路, 次短路]
	for i := 1; i < n; i++ {
		dis[i] = [2]int{1e9, 1e9}
	}
	dis[0] = [2]int{}
	type pair struct{ v, d int }
	q := []pair{{}}
	for len(q) > 0 { // BFS 求最短路和次短路
		p := q[0] // 取队首
		q = q[1:]
		for _, w := range g[p.v] {
			d := next(p.d) // 到达节点 w 的时间
			if d < dis[w][0] { // 比最短路小，就更新最短路
				q = append(q, pair{w, d})
				dis[w][0], d = d, dis[w][0] // 将 d 替换为原来的最短路，用于下面更新次短路
			}
			if dis[w][0] < d && d < dis[w][1] { // d 比最短路大又比次短路小，就更新次短路
				q = append(q, pair{w, d})
				dis[w][1] = d
			}
		}
	}
	if dis[n-1][1] < 1e9 {
		return dis[n-1][1]
	}
	return next(next(dis[n-1][0])) // 没有次短路，就在最短路的基础上额外往返一次
}
