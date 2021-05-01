package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF746G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, t, k, m int
	Fscan(in, &n, &t, &k)
	fa, cur := 1, 2
	a := make([][]int, t)
	p := make([]int, n+1)
	l := 1
	for i := range a {
		tmp := cur
		Fscan(in, &m)
		l += m - 1
		for ; m > 0; m-- {
			a[i] = append(a[i], cur)
			p[cur] = fa
			cur++
		}
		fa = tmp
	}
	if l < k {
		Fprint(out, -1)
		return
	}
	output := func() {
		Fprintln(out, n)
		for i := 2; i <= n; i++ {
			Fprintln(out, i, p[i])
		}
	}
	if l == k {
		output()
		return
	}
	for i := 1; i < t; i++ {
		for j := 1; j < len(a[i-1]) && j < len(a[i]); j++ {
			p[a[i][j]] = a[i-1][j]
			if l--; l == k {
				output()
				return
			}
		}
	}
	Fprint(out, -1)
}

//func main() { CF746G(os.Stdin, os.Stdout) }
