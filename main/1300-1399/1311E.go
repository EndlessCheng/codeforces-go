package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1311E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, d, lastR int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &d)
		p := make([]int, n+1)
		r := make([][]int, n)
		var f func(v, d int)
		f = func(v, d int) {
			if v <= n {
				p[v] = v / 2
				r[d] = append(r[d], v)
				f(2*v, d+1)
				f(2*v+1, d+1)
			}
		}
		f(1, 0)
		for i, vs := range r {
			if len(vs) == 0 {
				break
			}
			lastR = i
			d -= len(vs) * i
		}
		i := n - 1
		for i >= 0 && d > 0 {
			vs := r[i]
			if len(vs) <= 1 {
				i--
				continue
			}
			v := vs[len(vs)-1]
			r[i] = vs[:len(vs)-1]
			if lastR-i+1 < d {
				d -= lastR - i + 1
				p[v] = r[lastR][0]
				lastR++
				r[lastR] = []int{v}
			} else {
				p[v] = r[i+d-1][0]
				d = 0
			}
		}
		if i < 0 || d < 0 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
			for _, v := range p[2:] {
				Fprint(out, v, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { CF1311E(os.Stdin, os.Stdout) }
