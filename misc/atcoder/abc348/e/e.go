package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	sum := make([]int, n)
	var dfs func(int, int, int)
	dfs = func(x, fa, depth int) {
		ans += depth * a[x]
		sum[x] = a[x]
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x, depth+1)
				sum[x] += sum[y]
			}
		}
	}
	dfs(0, -1, 0)

	var reroot func(int, int, int)
	reroot = func(x, fa, res int) {
		ans = min(ans, res)
		for _, y := range g[x] {
			if y != fa {
				reroot(y, x, res+sum[0]-sum[y]*2)
			}
		}
	}
	reroot(0, -1, ans)
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
