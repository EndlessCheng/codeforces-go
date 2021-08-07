package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

/*
设从 1 出发到各点的距离数组为 ds，从 n 出发到各点的距离数组为 dt

对于两个特殊点 v 和 w，连边后形成的最短路如果经过 v-w 的话，有两种情况：
1---v-w---n
1---w-v---n

如果枚举所有的 v 和 w，复杂度是 O(k^2) 的，无法接受

如果能保证最短路是第一种情况，在此基础上，还能保证枚举 w 的时候能知道 1 到哪个 v 的距离是最大的，就能优化枚举

若要保证最短路是第一种情况，则需满足 ds[v]+1+dt[w] < ds[w]+1+dt[v]
即 ds[v]-dt[v] < dt[v]-dt[w]
可以据此来对特殊点排序

然后枚举 w，同时维护 ds[v] 的最大值，就可以计算最短路的最大值了

最后注意，最短路的最大值不能超过 1 到 n 的距离
*/

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, k, v, w int
	Fscan(in, &n, &m, &k)
	sp := make([]int, k)
	for i := range sp {
		Fscan(in, &sp[i])
		sp[i]--
	}
	g := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	bfs := func(st int) []int {
		d := make([]int, n)
		for i := range d {
			d[i] = -1
		}
		d[st] = 0
		q := []int{st}
		for len(q) > 0 {
			v, q = q[0], q[1:]
			for _, w := range g[v] {
				if d[w] == -1 {
					d[w] = d[v] + 1
					q = append(q, w)
				}
			}
		}
		return d
	}
	ds, dt := bfs(0), bfs(n-1)

	sort.Slice(sp, func(i, j int) bool { v, w := sp[i], sp[j]; return ds[v]-dt[v] < dt[v]-dt[w] })

	ans := 0
	maxds := ds[sp[0]]
	for _, v := range sp[1:] {
		ans = max(ans, maxds+1+dt[v])
		maxds = max(maxds, ds[v])
	}
	if ans > dt[0] {
		ans = dt[0]
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
