package main

import (
	"sort"
)

func maxWeight(edges [][]int, values []int) (ans int) {
	quickSelect := func(a []int, k int) []int {
		if k >= len(a)-1 {
			return a
		}
		for l, r := 0, len(a)-1; l < r; {
			v := a[l]
			i, j := l, r+1
			for {
				for i++; i < r && a[i] < v; i++ {
				}
				for j--; j > l && a[j] > v; j-- {
				}
				if i >= j {
					break
				}
				a[i], a[j] = a[j], a[i]
			}
			a[l], a[j] = a[j], v
			if j == k {
				break
			} else if j < k {
				l = j + 1
			} else {
				r = j - 1
			}
		}
		return a[:k+1]
	}

	// 边按照权值从大到小排序
	sort.Slice(edges, func(i, j int) bool { a, b := edges[i], edges[j]; return values[a[0]]+values[a[1]] > values[b[0]]+values[b[1]] })

	g := make([]map[int]int, len(values))
	for i := range g {
		g[i] = map[int]int{}
	}
	for i, e := range edges {
		v, w := e[0], e[1]
		g[v][w] = i + 1
		g[w][v] = i + 1
	}

	for i, m := range g {
		// 找所有能和 i 组成三角形的边
		es := []int{}
		if len(m)*len(m) > len(edges) {
			// 度数大于 √E，枚举所有边
			for j, e := range edges {
				if m[e[0]] > 0 && m[e[1]] > 0 {
					es = append(es, j)
				}
			}
		} else {
			// 否则，枚举相邻的点
			for v := range m {
				for w := range m {
					if eid := g[v][w]; eid > 0 {
						es = append(es, eid-1)
					}
				}
			}
		}

		// 找到权值最大的三条边，对每条边，枚举另一条边
		top3 := quickSelect(es, 2)
		for _, eid := range top3 {
			e := edges[eid]
			v, w := e[0], e[1]
			sum := values[i] + values[v] + values[w]
			for _, eid := range es {
				s := sum
				for _, u := range edges[eid] {
					if u != v && u != w {
						s += values[u]
					}
				}
				if s > ans {
					ans = s
				}
			}
		}
	}
	return
}
