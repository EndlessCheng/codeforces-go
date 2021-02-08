package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1481C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, w int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		pa := make([]int, n+1)
		pb := make([][]int, n+1)
		for i, v := range a {
			if Fscan(in, &w); w == v {
				pa[w] = i + 1
			} else {
				pb[w] = append(pb[w], i+1)
			}
		}
		c := make([]int, m)
		for i := range c {
			Fscan(in, &c[i])
		}
		ans := make([]interface{}, m)
		for i := m - 1; i >= 0; i-- {
			if v := c[i]; len(pb[v]) > 0 {
				ans[i] = pb[v][0]
				pb[v] = pb[v][1:]
			} else if i < m-1 {
				ans[i] = ans[m-1]
			} else if pa[v] > 0 {
				ans[i] = pa[v]
			} else {
				Fprintln(out, "NO")
				continue o
			}
		}
		for _, p := range pb {
			if len(p) > 0 {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
		Fprintln(out, ans...)
	}
}

//func main() { CF1481C(os.Stdin, os.Stdout) }
