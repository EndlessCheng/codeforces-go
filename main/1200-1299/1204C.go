package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1204C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, m int
	var s []byte
	Fscan(in, &n)
	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		Fscan(in, &s)
		for j, b := range s {
			if b == '1' {
				d[i][j] = 1
			} else if j != i {
				d[i][j] = 1e9
			}
		}
	}
	for k := range d {
		for i := range d {
			for j := range d {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	Fscan(in, &m)
	p := make([]int, m)
	for i := range p {
		Fscan(in, &p[i])
	}
	ans := []interface{}{p[0]}
	v := p[0] - 1
	for i := 1; i < m-1; i++ {
		w, u := p[i]-1, p[i+1]-1
		if d[v][w]+d[w][u] > d[v][u] {
			ans = append(ans, w+1)
			v = w
		}
	}
	ans = append(ans, p[m-1])
	Fprintln(out, len(ans))
	Fprint(out, ans...)
}

//func main() { CF1204C(os.Stdin, os.Stdout) }
