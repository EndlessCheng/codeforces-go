package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1490D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		var f func(l, r, d int)
		f = func(l, r, d int) {
			if l == r {
				return
			}
			p := l
			for i := l; i < r; i++ {
				if a[i] > a[p] {
					p = i
				}
			}
			a[p] = d
			f(l, p, d+1)
			f(p+1, r, d+1)
		}
		f(0, n, 0)
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1490D(os.Stdin, os.Stdout) }
