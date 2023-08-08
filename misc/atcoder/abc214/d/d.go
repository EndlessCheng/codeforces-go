package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	es := make([]struct{ x, y, wt int }, n-1)
	for i := range es {
		Fscan(in, &es[i].x, &es[i].y, &es[i].wt)
	}

	fa := make([]int, n+1)
	size := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	sort.Slice(es, func(i, j int) bool { return es[i].wt < es[j].wt })
	for _, e := range es {
		x, y := find(e.x), find(e.y)
		ans += size[x] * size[y] * e.wt
		fa[x] = y
		size[y] += size[x]
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
