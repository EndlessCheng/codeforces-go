package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1993E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var t, n, m int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		a := make([][]int, n+1)
		for i := range a {
			a[i] = make([]int, m+1)
		}
		for i := range n {
			for j := range m {
				Fscan(in, &a[i][j])
				a[n][j] ^= a[i][j]
			}
		}
		for i := 0; i <= n; i++ {
			for j := range m {
				a[i][m] ^= a[i][j]
			}
		}

		b := make([][][]int, n+1)
		for i := range b {
			b[i] = make([][]int, m+1)
			for j := range b[i] {
				b[i][j] = make([]int, m+1)
			}
		}

		for g := 0; g <= n; g++ {
			for i := 0; i <= m; i++ {
				for j := 0; j <= m; j++ {
					for h := 0; h <= n; h++ {
						if h != g {
							x := a[h][i] - a[h][j]
							if x < 0 {
								x = -x
							}
							b[g][i][j] += x
						}
					}
				}
			}
		}

		u := 1<<(m+1) - 1
		d := make([][]int, 1<<(m+1))
		for i := range d {
			d[i] = make([]int, m+1)
		}

		e := make([][]int, n+1)
		for i := range e {
			e[i] = make([]int, m+1)
		}

		for i := 0; i <= n; i++ {
			for j := 0; j <= m; j++ {
				d[1<<j][j] = 0
			}

			for j := 1; j < u; j++ {
				if bits.OnesCount(uint(j)) > 1 {
					for g := 0; g <= m; g++ {
						if j>>g&1 != 0 {
							x := int(1e9)
							for h := 0; h <= m; h++ {
								if h != g && j>>h&1 != 0 {
									x = min(x, d[j^1<<g][h]+b[i][g][h])
								}
							}
							d[j][g] = x
						}
					}
				}
			}

			for j := 0; j <= m; j++ {
				x := int(1e9)
				for g := 0; g <= m; g++ {
					if g != j {
						x = min(x, d[u^(1<<j)][g])
					}
				}
				e[i][j] = x
			}
		}

		b = make([][][]int, m+1)
		for i := range b {
			b[i] = make([][]int, n+1)
			for j := range b[i] {
				b[i][j] = make([]int, n+1)
			}
		}

		for g := 0; g <= m; g++ {
			for i := 0; i <= n; i++ {
				for j := 0; j <= n; j++ {
					for h := 0; h <= m; h++ {
						if h != g {
							x := a[i][h] - a[j][h]
							if x < 0 {
								x = -x
							}
							b[g][i][j] += x
						}
					}
				}
			}
		}

		u = (1 << (n + 1)) - 1
		d = make([][]int, 1<<(n+1))
		for i := range d {
			d[i] = make([]int, n+1)
		}

		ans := int(1e9)
		for i := 0; i <= m; i++ {
			for j := 0; j <= n; j++ {
				d[1<<j][j] = 0
			}
			for j := 1; j < u; j++ {
				if bits.OnesCount(uint(j)) > 1 {
					for g := 0; g <= n; g++ {
						if j>>g&1 != 0 {
							x := int(1e9)
							for h := 0; h <= n; h++ {
								if h != g && j>>h&1 != 0 {
									x = min(x, d[j^1<<g][h]+b[i][g][h])
								}
							}
							d[j][g] = x
						}
					}
				}
			}

			for j := 0; j <= n; j++ {
				x := int(1e9)
				for g := 0; g <= n; g++ {
					if g != j {
						x = min(x, d[u^1<<j][g])
					}
				}
				ans = min(ans, e[j][i]+x)
			}
		}

		Fprintln(out, ans)
	}
}

//func main() { cf1993E(bufio.NewReader(os.Stdin), os.Stdout) }
