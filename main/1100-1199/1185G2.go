package main

import (
	. "fmt"
	"io"
)

func cf1185G2(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	add := func(v *int, w int) {
		*v = (*v + w) % mod
	}

	var n, tot, w, tp int
	Fscan(in, &n, &tot)

	f := make([][]int, n+2)
	for i := range f {
		f[i] = make([]int, tot+1)
	}
	f[0][0] = 1
	g := make([][][]int, n+2)
	for i := range g {
		g[i] = make([][]int, n+2)
		for j := range g[i] {
			g[i][j] = make([]int, tot+1)
		}
	}
	g[0][0][0] = 1

	cnt := [3]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &w, &tp)
		tp--
		if tp == 0 {
			for j := cnt[0]; j >= 0; j-- {
				for t := tot; t >= w; t-- {
					add(&f[j+1][t], f[j][t-w])
				}
			}
		} else {
			is := [3]int{}
			is[tp] = 1
			for j := cnt[1]; j >= 0; j-- {
				for k := cnt[2]; k >= 0; k-- {
					for t := tot; t >= w; t-- {
						add(&g[j+is[1]][k+is[2]][t], g[j][k][t-w])
					}
				}
			}
		}
		cnt[tp]++
	}

	fac := make([]int, n+1)
	fac[0] = 1
	for i := 1; i <= n; i++ {
		fac[i] = fac[i-1] * i % mod
	}

	ans := 0
	c := make([][][][3]int, cnt[0]+2)
	for i := range c {
		c[i] = make([][][3]int, cnt[1]+2)
		for j := range c[i] {
			c[i][j] = make([][3]int, cnt[2]+2)
		}
	}
	c[1][0][0][0] = 1
	c[0][1][0][1] = 1
	c[0][0][1][2] = 1
	for i, mat := range c[:cnt[0]+1] {
		for j, row := range mat[:cnt[1]+1] {
			for k, comb := range row[:cnt[2]+1] {
				sum := 0
				for t, fit := range f[i] {
					sum = (sum + fit*g[j][k][tot-t]) % mod
				}
				add(&ans, fac[i]*fac[j]%mod*fac[k]%mod*(comb[0]+comb[1]+comb[2])%mod*sum)

				for pre := 0; pre < 3; pre++ {
					for cur := 0; cur < 3; cur++ {
						if cur == pre {
							continue
						}
						is := [3]int{}
						is[cur] = 1
						add(&c[i+is[0]][j+is[1]][k+is[2]][cur], comb[pre])
					}
				}
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1185G2(os.Stdin, os.Stdout) }
