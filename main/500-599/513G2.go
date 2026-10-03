package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf513G2(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range n {
		Fscan(in, &a[i])
	}

	f := make([][]float64, n)
	nf := make([][]float64, n)
	for i := range n {
		f[i] = make([]float64, n)
		nf[i] = make([]float64, n)
	}

	for i := range n {
		for j := range n {
			if a[i] < a[j] {
				f[i][j] = 1
			} else {
				f[j][i] = 1
			}
		}
	}

	for range k {
		for x := range n {
			for y := range n {
				if x == y {
					continue
				}
				for l := range n {
					for r := l; r < n; r++ {
						fx, fy := x, y
						if x >= l && x <= r {
							fx = l + r - x
						}
						if y >= l && y <= r {
							fy = l + r - y
						}
						nf[fx][fy] += f[x][y] * (2 / float64(n) / float64(n+1))
					}
				}
			}
		}
		f, nf = nf, f
		for i := range nf {
			clear(nf[i])
		}
	}

	ans := 0.
	for i := range n {
		for j := 0; j < i; j++ {
			ans += f[i][j]
		}
	}
	Fprintf(out, "%.9f", ans)
}

//func main() { cf513G2(bufio.NewReader(os.Stdin), os.Stdout) }
