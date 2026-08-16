package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1383D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	r := make([]int, n)
	c := make([]int, m)
	b := make([]int, n*m+1)
	a := make([][]int, n)
	ans := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		ans[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			Fscan(in, &a[i][j])
			r[i] = max(r[i], a[i][j])
			c[j] = max(c[j], a[i][j])
		}
	}

	for i := 0; i < n; i++ {
		b[r[i]] |= 1
	}
	for j := 0; j < m; j++ {
		b[c[j]] |= 2
	}

	p := 1
	gen := func() int {
		for b[p] != 0 {
			p++
		}
		p++
		return p - 1
	}

	for i, x, y := 1, 0, 0; i <= n*m; i++ {
		if b[i] != 0 {
			ans[x][y] = i
			if b[i]&1 != 0 {
				for t := m - 1; t > y; t-- {
					ans[x][t] = gen()
				}
			}
			if b[i]&2 != 0 {
				for t := n - 1; t > x; t-- {
					ans[t][y] = gen()
				}
			}
			if b[i]&1 != 0 {
				x++
			}
			if b[i]&2 != 0 {
				y++
			}
		}
	}

	for _, r := range ans {
		for _, v := range r {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1383D(bufio.NewReader(os.Stdin), os.Stdout) }
