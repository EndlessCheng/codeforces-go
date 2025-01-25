package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, k, v, w int
	Fscan(in, &n, &k)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := k
	var dfs func(int, int, int)
	dfs = func(v, fa, dep int) {
		c := 0
		for _, w := range g[v] {
			if w != fa {
				ans = ans * (k - min(dep+1, 2) - c) % mod // 如果题目改成距离 <= 3，这里改成和 3 取 min
				dfs(w, v, dep+1)
				c++ // 如果题目改成距离 <= 3，这里改成 c += len(g[w])
			}
		}
	}
	dfs(1, 0, 0)
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
