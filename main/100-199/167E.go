package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf167E(in io.Reader, out io.Writer) {
	var n, m, mod int
	Fscan(in, &n, &m, &mod)
	g := make([][]int, n+1)
	inD := make([]int, n+1)
	outD := make([]int, n+1)
	for range m {
		var x, y int
		Fscan(in, &x, &y)
		outD[x]++
		inD[y]++
		g[x] = append(g[x], y)
	}

	p := make([]int, n+1)
	sz := 0
	for i := 1; i <= n; i++ {
		if outD[i] == 0 {
			sz++
			p[i] = sz
		}
	}

	cnt := make([][]int, n+1)
	for i := range cnt {
		cnt[i] = make([]int, sz+1)
	}

	vis := make([]bool, n+1)

	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		if p[v] != 0 {
			cnt[v][p[v]] = 1
			return
		}
		for _, w := range g[v] {
			if !vis[w] {
				dfs(w)
			}
			for i := 1; i <= sz; i++ {
				cnt[v][i] = (cnt[v][i] + cnt[w][i]) % mod
			}
		}
	}

	a := make([][]int, sz+1)
	for i := range a {
		a[i] = make([]int, sz+1)
	}

	k := 0
	for i := 1; i <= n; i++ {
		if inD[i] == 0 {
			k++
			dfs(i)
			for j := 1; j <= sz; j++ {
				a[k][j] = cnt[i][j]
			}
		}
	}

	calc := func(mx int) int {
		ans := 1
		for i := 1; i <= mx; i++ {
			if a[i][i] == 0 {
				for j := i + 1; j <= mx; j++ {
					if a[j][i] != 0 {
						a[i], a[j] = a[j], a[i]
						ans = -ans
						break
					}
				}
			}

			for j := i + 1; j <= mx; j++ {
				for a[j][i] != 0 {
					t := a[i][i] / a[j][i]
					for k := i; k <= mx; k++ {
						a[i][k] = (a[i][k] - a[j][k]*t) % mod
					}
					a[i], a[j] = a[j], a[i]
					ans = -ans
				}
			}
			ans = ans * a[i][i] % mod
		}
		return (ans + mod) % mod
	}

	Fprint(out, calc(sz))
}

//func main() { cf167E(bufio.NewReader(os.Stdin), os.Stdout) }
