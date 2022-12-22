package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const inf int = 1e18
	var n, m, k, v int
	Fscan(in, &n, &m, &k)
	es := make([]struct{ x, y, wt int }, m)
	for i := range es {
		Fscan(in, &es[i].x, &es[i].y, &es[i].wt)
	}
	dis := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dis[i] = inf
	}
	for ; k > 0; k-- {
		Fscan(in, &v)
		e := es[v-1]
		dis[e.y] = min(dis[e.y], dis[e.x]+e.wt)
	}
	if dis[n] == inf {
		Fprint(out, -1)
	} else {
		Fprint(out, dis[n])
	}
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
