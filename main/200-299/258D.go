package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf258D(in io.Reader, out io.Writer) {
	var n, m, p, q int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	f := make([][]float64, n)
	for i, x := range a {
		f[i] = make([]float64, n)
		for j, y := range a {
			if x > y {
				f[i][j] = 1
			}
		}
	}

	for range m {
		Fscan(in, &p, &q)
		p--
		q--
		for i := range n {
			if i != p && i != q {
				t := (f[i][p] + f[i][q]) / 2
				f[i][p] = t
				f[i][q] = t
				t = (f[p][i] + f[q][i]) / 2
				f[p][i] = t
				f[q][i] = t
			}
		}
		f[p][q] = 0.5
		f[q][p] = 0.5
	}

	ans := 0.
	for i, row := range f {
		for _, v := range row[i+1:] {
			ans += v
		}
	}
	Fprintf(out, "%.9f", ans)
}

//func main() { cf258D(bufio.NewReader(os.Stdin), os.Stdout) }
