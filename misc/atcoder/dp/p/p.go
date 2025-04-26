package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n int
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
	var dfs func(int, int) [2]int
	dfs = func(v, fa int) [2]int {
		f := [2]int{1, 1}
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			fw := dfs(w, v)
			f[0] = f[0] * (fw[0] + fw[1]) % mod
			f[1] = f[1] * fw[0] % mod
		}
		return f
	}
	ans := dfs(0, -1)
	Fprint(out, (ans[0]+ans[1])%mod)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
