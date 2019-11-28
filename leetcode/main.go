package main

import (
	. "fmt"
)

func collections() {
	const mod int = 1e9 + 7
	dirOffset4 := [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ifElseI := func(cond bool, r1, r2 int) int {
		if cond {
			return r1
		}
		return r2
	}
	ifElseS := func(cond bool, r1, r2 string) string {
		if cond {
			return r1
		}
		return r2
	}

	_ = []interface{}{Print, ifElseI, ifElseS, dirOffset4, min, max}
}

// slice 先提取出 n, m 等信息
// 注意，若 slice 对自己切片，n, m 需要更新

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	type pair struct {
		v      int
		isBlue bool
		length int
	}
	g := make([][]pair, n)
	for _, p := range redEdges {
		g[p[0]] = append(g[p[0]], pair{p[1], false, 0})
	}
	for _, p := range blueEdges {
		g[p[0]] = append(g[p[0]], pair{p[1], true, 0})
	}

	ans := make([]int, n)
	for i := 1; i < n; i++ {
		ans[i] = -1
	}
	visB := make([]bool, n)
	visR := make([]bool, n)
	q := []pair{{}}
	for len(q) > 0 {
		var p pair
		p, q = q[0], q[1:]
		for _, e := range g[p.v] {
			w, isBlue := e.v, e.isBlue
			if p.v > 0 && isBlue == p.isBlue {
				continue
			}
			if isBlue && visB[w] || !isBlue && visR[w] {
				continue
			}
			if isBlue {
				visB[w] = true
			} else {
				visR[w] = true
			}
			if ans[w] == -1 {
				ans[w] = p.length + 1
			} else {
				ans[w] = min(ans[w], p.length+1)
			}
			q = append(q, pair{w, isBlue, p.length + 1})
		}
	}
	return ans
}

func main() {
	toBytes := func(g [][]string) [][]byte {
		n, m := len(g), len(g[0])
		bytes := make([][]byte, n)
		for i := range bytes {
			bytes[i] = make([]byte, m)
			for j := range bytes[i] {
				bytes[i][j] = g[i][j][0]
			}
		}
		return bytes
	}
	_ = toBytes

	Println(shortestAlternatingPaths(3, [][]int{{0, 1}, {1, 2}}, [][]int{}))
	Println(shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{2, 1}}))
	Println(shortestAlternatingPaths(3, [][]int{{1, 0}}, [][]int{{2, 1}}))
	Println(shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{1, 2}}))
	Println(shortestAlternatingPaths(3, [][]int{{0, 1}, {0, 2}}, [][]int{{1, 0}}))
	Println(shortestAlternatingPaths(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, [][]int{{1, 2}, {2, 3}, {3, 1}}))
}
