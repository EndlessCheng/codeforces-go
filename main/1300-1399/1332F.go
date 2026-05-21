package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1332F(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	var dfs func(int, int) [3]int
	dfs = func(v, fa int) [3]int {
		f := [3]int{1, 1, 1}
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			fw := dfs(w, v)
			f[0] = f[0] * (fw[0] + fw[1] + fw[2]) % mod
			f[1] = f[1] * (fw[0] + fw[2]) % mod
			f[2] = f[2] * fw[2] % mod
		}
		f[2] = (f[0] + f[1] - f[2]) % mod
		return f
	}

	ans := dfs(1, 0)[2] - 1
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { cf1332F(bufio.NewReader(os.Stdin), os.Stdout) }
