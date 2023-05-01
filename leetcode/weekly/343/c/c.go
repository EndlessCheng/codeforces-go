package main

import "math"

// https://space.bilibili.com/206214
func minimumCost(start, target []int, specialRoads [][]int) int {
	type pair struct{ x, y int }
	t := pair{target[0], target[1]}
	dis := make(map[pair]int, len(specialRoads)+2)
	dis[t] = math.MaxInt
	dis[pair{start[0], start[1]}] = 0
	vis := make(map[pair]bool, len(specialRoads)+1) // 终点不用记
	for {
		v, dv := pair{}, -1
		for p, d := range dis {
			if !vis[p] && (dv < 0 || d < dv) {
				v, dv = p, d
			}
		}
		if v == t { // 到终点的最短路已确定
			return dv
		}
		vis[v] = true
		dis[t] = min(dis[t], dv+t.x-v.x+t.y-v.y) // 更新到终点的最短路
		for _, r := range specialRoads {
			w := pair{r[2], r[3]}
			d := dv + abs(r[0]-v.x) + abs(r[1]-v.y) + r[4]
			if dw, ok := dis[w]; !ok || d < dw {
				dis[w] = d
			}
		}
	}
}

func abs(x int) int { if x < 0 { return -x }; return x }
func min(a, b int) int { if b < a { return b }; return a }
