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
	var n, p int
	Fscan(in, &n)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		Fscan(in, &p)
		p--
		g[p] = append(g[p], w)
	}

	var dfs func(int) (int, int)
	dfs = func(v int) (a, b int) {
		a = 1
		gs := [2][][2]int{}
		for _, w := range g[v] {
			aa, bb := dfs(w)
			odd := (aa + bb) % 2
			gs[odd] = append(gs[odd], [2]int{aa, bb})
		}
		even, odd := gs[0], gs[1]
		if len(odd)%2 == 0 {
			for _, p := range even {
				a += p[0]
				b += p[1]
			}
		} else {
			for _, p := range even {
				a += min(p[0], p[1])
				b += max(p[0], p[1])
			}
		}
		sort.Slice(odd, func(i, j int) bool {
			a, b := odd[i], odd[j]
			return a[1]-a[0] > b[1]-b[0]
		})
		for i, p := range odd {
			a += p[i&1]
			b += p[i&1^1]
		}
		return
	}
	ans, _ := dfs(0)
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
