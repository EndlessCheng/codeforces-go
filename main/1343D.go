package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1343D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, k int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ds := make([]int, 2*k+2)
		for i, v := range a[:n/2] {
			w := a[n-1-i]
			if v > w {
				v, w = w, v
			}
			ds[v+1]--
			ds[v+w]--
			ds[v+w+1]++
			ds[w+k+1]++
		}
		ans := n
		c := n
		for _, d := range ds {
			c += d
			if c < ans {
				ans = c
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1343D(os.Stdin, os.Stdout) }
