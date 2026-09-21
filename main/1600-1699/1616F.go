package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1616F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		g := make([][]int, n+1)
		for i := range g {
			g[i] = make([]int, n+1)
		}
		f := make([][]int, n*m+1)
		for i := range f {
			f[i] = make([]int, m+2)
		}
		p := make([]int, m+1)
		cnt := 0
		for i := 1; i <= m; i++ {
			var x, y, z int
			Fscan(in, &x, &y, &z)
			g[x][y], g[y][x] = i, i
			if z != -1 {
				cnt++
				f[cnt][i] = 1
				f[cnt][m+1] = z % 3
			}
		}

		for i := 1; i <= n; i++ {
			for j := i + 1; j <= n; j++ {
				for k := j + 1; k <= n; k++ {
					if g[i][j] != 0 && g[j][k] != 0 && g[k][i] != 0 {
						cnt++
						f[cnt][g[i][j]] = 1
						f[cnt][g[j][k]] = 1
						f[cnt][g[k][i]] = 1
					}
				}
			}
		}

		ans := make([]int, m+1)
		d := 0
		for i := 1; i <= m; i++ {
			u := 0
			for j := d + 1; j <= cnt; j++ {
				if f[j][i] != 0 {
					u = j
				}
			}
			if u == 0 {
				ans[i] = 0
				continue
			}

			d++
			p[d] = i
			f[u], f[d] = f[d], f[u]

			if f[d][i] != 1 {
				for j := i; j <= m+1; j++ {
					f[d][j] = 3 - f[d][j]
				}
			}

			for j := 1; j <= cnt; j++ {
				if j != d && f[j][i] != 0 {
					x := f[j][i]
					for k := i; k <= m+1; k++ {
						f[j][k] = (f[j][k] - x*f[d][k] + 9) % 3
					}
				}
			}
		}

		for i := d + 1; i <= cnt; i++ {
			if f[i][m+1] != 0 {
				Fprintln(out, -1)
				continue o
			}
		}

		for i := 1; i <= d; i++ {
			ans[p[i]] = f[i][m+1]
		}
		for i := 1; i <= m; i++ {
			if ans[i] != 0 {
				Fprint(out, ans[i], " ")
			} else {
				Fprint(out, "3 ")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf1616F(bufio.NewReader(os.Stdin), os.Stdout) }
