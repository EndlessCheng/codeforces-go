package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type xorBasis4151 [60]int

func (b *xorBasis4151) insert(v int) {
	for i := len(b) - 1; i >= 0; i-- {
		if v>>i&1 == 0 {
			continue
		}
		if b[i] == 0 {
			b[i] = v
			return
		}
		v ^= b[i]
	}
}

func (b *xorBasis4151) maxXor(xor int) int {
	for i := len(b) - 1; i >= 0; i-- {
		if xor^b[i] > xor {
			xor ^= b[i]
		}
	}
	return xor
}

func p4151(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n+1)
	for ; m > 0; m-- {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	b := xorBasis4151{}
	dis := make([]int, n+1)
	for i := range dis {
		dis[i] = -1
	}
	var dfs func(int, int)
	dfs = func(v, xor int) {
		dis[v] = xor
		for _, e := range g[v] {
			w := e.to
			if dis[w] < 0 {
				dfs(w, xor^e.wt)
			} else {
				b.insert(xor ^ e.wt ^ dis[w])
			}
		}
	}
	dfs(1, 0)
	Fprint(out, b.maxXor(dis[n]))
}

//func main() { p4151(os.Stdin, os.Stdout) }
