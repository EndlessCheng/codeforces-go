package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1758E(in io.Reader, _w io.Writer) {
	const mod = 1_000_000_007
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, h int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &h)
		sz := n + m + 1
		type pair struct{ to, val int }
		g := make([][]pair, sz)
		d := make([]int, sz)
		for i := range d {
			d[i] = -1
		}
		k := 0

		var dfs func(int)
		dfs = func(x int) {
			for _, e := range g[x] {
				to, val := e.to, e.val
				if d[to] != -1 {
					if (d[x]+d[to])%h != val {
						k = -sz
					}
				} else {
					d[to] = (val + h - d[x]) % h
					dfs(to)
				}
			}
		}

		for i := 1; i <= n; i++ {
			for j := 1; j <= m; j++ {
				var x int
				Fscan(in, &x)
				if x != -1 {
					u := i
					v := j + n
					g[u] = append(g[u], pair{v, x})
					g[v] = append(g[v], pair{u, x})
				}
			}
		}

		k = 0
		valid := true

		for i := 1; i <= n+m; i++ {
			if d[i] == -1 {
				d[i] = 0
				dfs(i)
				k++
			}
			if k < 0 {
				valid = false
			}
		}

		if !valid {
			Fprintln(out, 0)
			continue
		}

		ans := 1
		for i := 1; i < k; i++ {
			ans = ans * h % mod
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1758E(bufio.NewReader(os.Stdin), os.Stdout) }
