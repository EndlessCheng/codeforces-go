package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2070D(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for w := 1; w < n; w++ {
			var p int
			Fscan(in, &p)
			p--
			g[p] = append(g[p], w)
		}

		row := make([][]int, n)
		var dfs func(int, int)
		dfs = func(v, d int) {
			row[d] = append(row[d], v)
			for _, w := range g[v] {
				dfs(w, d+1)
			}
		}
		dfs(0, 0)

		f := make([]int, n)
		pre := 0
		for d := n - 1; d > 0; d-- {
			cur := 0
			for _, v := range row[d] {
				sum := pre
				for _, w := range g[v] {
					sum -= f[w]
				}
				f[v] = (sum + 1) % mod
				cur += f[v]
			}
			pre = cur
		}
		Fprintln(out, ((pre+1)%mod+mod)%mod)
	}
}

//func main() { cf2070D(bufio.NewReader(os.Stdin), os.Stdout) }
