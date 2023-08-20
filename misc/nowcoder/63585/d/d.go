package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	isSQ := func(x int) bool {
		res := int(math.Sqrt(float64(x)))
		for i := res - 1; i <= res+1; i++ {
			if i*i == x {
				return true
			}
		}
		return false
	}

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	var dfs func(int, int) (int, int)
	dfs = func(v, fa int) (not, do int) {
		sd := 0
		for _, w := range g[v] {
			if w != fa {
				n, d := dfs(w, v)
				not += max(n, d)
				sd += d
				if isSQ(a[v] * a[w]) {
					do = max(do, n+2-d)
				}
			}
		}
		do += sd
		return
	}
	n, d := dfs(0, -1)
	Fprintln(out, max(n, d))
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
