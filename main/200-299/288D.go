package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf288D(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := n * (n - 1) / 2
	ans *= ans
	var dfs func(int, int) int
	dfs = func(v, fa int) int {
		size := 1
		turn := 0
		for _, w := range g[v] {
			if w != fa {
				sz := dfs(w, v)
				turn += size * sz
				size += sz
			}
		}
		ans -= turn * (turn + size*(n-size)*2)
		return size
	}
	dfs(1, 0)
	Fprint(out, ans)
}

//func main() { cf288D(bufio.NewReader(os.Stdin), os.Stdout) }
