package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w, k, min, extra int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	rg := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v)
	}
	Fscan(in, &k)
	path := make([]int, k)
	for i := range path {
		Fscan(in, &path[i])
		path[i]--
	}

	// 在反图 rg 上跑 BFS，得到每个点到终点的最短距离 dis
	// 这样就能知道有没有跟着导航走了
	dis := make([]int, n)
	vis := make([]bool, n)
	vis[path[k-1]] = true
	q := []int{path[k-1]}
	for len(q) > 0 {
		v, q = q[0], q[1:]
		for _, w := range rg[v] {
			if !vis[w] {
				vis[w] = true
				dis[w] = dis[v] + 1
				q = append(q, w)
			}
		}
	}

	for i, cur := range path[:k-1] {
		next := path[i+1]
		if dis[next] >= dis[cur] { // 偏航了，下一站距终点的距离没有变得更小（没有跟着导航走）
			min++ // 导航一定会更新推荐线路
			continue
		}
		for _, other := range g[cur] {
			if other != next && dis[other] == dis[next] { // 存在另一条最短路径
				extra++
				break
			}
		}
	}
	Fprint(out, min, min+extra)
}

func main() { run(os.Stdin, os.Stdout) }
