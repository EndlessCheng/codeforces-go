package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2053F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		a := make([][]int, n+2)
		b := make([]int, n+2)
		for i := 1; i <= n; i++ {
			for range m {
				var x int
				Fscan(in, &x)
				if x != -1 {
					a[i] = append(a[i], x)
				} else {
					b[i]++
				}
			}
		}

		d := make([]int, k+1)
		u, v, s := 0, 0, 0
		for i := 1; i <= n; i++ {
			x := b[i-1] * b[i]
			u += x
			v = max(v+x, s)
			s += x
			for _, x := range a[i-1] {
				t := max(d[x]+u, v) + b[i]
				s = max(s, t)
				d[x] = t - u
			}
			for _, x := range a[i+1] {
				t := max(d[x]+u, v) + b[i]
				s = max(s, t)
				d[x] = t - u
			}
		}

		c := make([]int, k+1)
		for i := 1; i < n; i++ {
			for _, x := range a[i] {
				c[x]++
			}
			for _, x := range a[i+1] {
				s += c[x]
			}
			for _, x := range a[i] {
				c[x]--
			}
		}
		Fprintln(out, s)
	}
}

//func main() { cf2053F(bufio.NewReader(os.Stdin), os.Stdout) }
