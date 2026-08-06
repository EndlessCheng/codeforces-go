package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2113F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T int
	for Fscan(in, &T); T > 0; T-- {
		var n int
		Fscan(in, &n)
		a := make([]int, n+1)
		b := make([]int, n+1)
		buc := make([][]int, n*2+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			buc[a[i]] = append(buc[a[i]], i)
		}
		for i := 1; i <= n; i++ {
			Fscan(in, &b[i])
			buc[b[i]] = append(buc[b[i]], i+n)
		}

		ans := 0
		for i := 1; i <= n*2; i++ {
			ans += min(len(buc[i]), 2)
		}
		Fprintln(out, ans)

		g := make([][]int, n*2+1)
		for i := 1; i <= n; i++ {
			g[i] = append(g[i], i+n)
			g[i+n] = append(g[i+n], i)
		}
		for i := 1; i <= n*2; i++ {
			if len(buc[i]) > 1 {
				g[buc[i][0]] = append(g[buc[i][0]], buc[i][1])
				g[buc[i][1]] = append(g[buc[i][1]], buc[i][0])
			}
		}

		col := make([]int8, n*2+1)
		vis := make([]int8, n*2+1)
		var dfs func(int)
		dfs = func(x int) {
			vis[x] = 1
			for _, v := range g[x] {
				if vis[v] == 0 {
					col[v] = col[x] ^ 1
					dfs(v)
				}
			}
		}
		for i := 1; i <= n*2; i++ {
			if vis[i] == 0 {
				dfs(i)
			}
		}

		for i := 1; i <= n; i++ {
			if col[i] != 0 {
				a[i], b[i] = b[i], a[i]
			}
		}

		for i := 1; i <= n; i++ {
			Fprint(out, a[i], " ")
		}
		Fprintln(out)
		for i := 1; i <= n; i++ {
			Fprint(out, b[i], " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2113F(bufio.NewReader(os.Stdin), os.Stdout) }
