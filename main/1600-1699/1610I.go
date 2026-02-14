package main

import (
	"bytes"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1610I(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	pa := make([]int, n)
	sg := make([]int, n)
	var dfs func(int, int)
	dfs = func(v, fa int) {
		for _, w := range g[v] {
			if w != fa {
				pa[w] = v
				dfs(w, v)
				sg[v] ^= sg[w] + 1
			}
		}
	}
	dfs(0, -1)

	ans := bytes.Repeat([]byte{'2'}, n)
	res := sg[0]
	if res > 0 {
		ans[0] = '1'
	}

	vis := make([]bool, n)
	vis[0] = true
	for i := 1; i < n; i++ {
		for j := i; !vis[j]; j = pa[j] {
			vis[j] = true
			res ^= sg[j] ^ (sg[j] + 1) ^ 1
		}
		if res > 0 {
			ans[i] = '1'
		}
	}
	Fprintf(out, "%s", ans)
}

//func main() { cf1610I(bufio.NewReader(os.Stdin), os.Stdout) }
