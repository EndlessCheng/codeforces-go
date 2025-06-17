package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
var g25 [][]uint32
var mn25 []uint32

func dfs25(v, fa, x uint32) {
	x = min(x, v)
	mn25[v] = x
	for _, w := range g25[v] {
		if w != fa {
			dfs25(w, v, x)
		}
	}
}

func cf825G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, v, w, last uint32
	Fscan(in, &n, &q)
	g25 = make([][]uint32, n+1)
	for range n - 1 {
		Fscan(in, &v, &w)
		g25[v] = append(g25[v], w)
		g25[w] = append(g25[w], v)
	}
	Fscan(in, &v, &v)
	v = v%n + 1
	mn25 = make([]uint32, n+1)
	dfs25(v, 0, v)

	ans := v
	for range q - 1 {
		Fscan(in, &op, &v)
		v = (v+last)%n + 1
		if op == 1 {
			ans = min(ans, mn25[v])
		} else {
			last = min(ans, mn25[v])
			Fprintln(out, last)
		}
	}
}

//func main() { cf825G(bufio.NewReader(os.Stdin), os.Stdout) }
