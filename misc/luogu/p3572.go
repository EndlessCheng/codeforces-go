package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p3572(in io.Reader, out io.Writer) {
	var n, t, k int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &k)
		type pair struct{ f, i int }
		q := []pair{{}}
		for i := 1; i < n; i++ {
			if q[0].i < i-k {
				q = q[1:]
			}
			f := q[0].f
			if a[i] >= a[q[0].i] {
				f++
			}
			for len(q) > 0 && f <= q[len(q)-1].f {
				q = q[:len(q)-1]
			}
			q = append(q, pair{f, i})
		}
		Fprintln(out, q[len(q)-1].f)
	}
}

//func main() { p3572(bufio.NewReader(os.Stdin), os.Stdout) }
