package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf581F(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	var dfs func(int, int) (int, []int)
	dfs = func(v, fa int) (size int, f []int) {
		if len(g[v]) == 1 {
			size = 1
		}
		f = make([]int, n+1)
		for i := 1; i <= n; i++ {
			f[i] = 1e9
		}

		for _, w := range g[v] {
			if w == fa {
				continue
			}
			sz, fw := dfs(w, v)
			for i := size; i >= 0; i-- {
				for j := range sz + 1 {
					f[i+j] = min(f[i+j], f[i]+fw[j])
				}
			}
			size += sz
		}

		for i := range size + 1 {
			f[size-i] = min(f[size-i], f[i]+1)
		}
		return
	}

	sz, f := dfs(1, 0)
	Fprint(out, f[sz/2])
}

//func main() { cf581F(bufio.NewReader(os.Stdin), os.Stdout) }
