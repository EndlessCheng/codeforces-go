package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1882D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		ans := make([]int64, n)
		size := make([]int, n)
		var f func(int, int)
		f = func(x, fa int) {
			size[x] = 1
			for _, y := range g[x] {
				if y != fa {
					f(y, x)
					ans[0] += int64(a[x]^a[y]) * int64(size[y])
					size[x] += size[y]
				}
			}
		}
		f(0, -1)

		var r func(int, int)
		r = func(x, fa int) {
			for _, y := range g[x] {
				if y != fa {
					ans[y] = ans[x] + int64(a[x]^a[y])*int64(n-size[y]*2)
					r(y, x)
				}
			}
		}
		r(0, -1)
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1882D(os.Stdin, os.Stdout) }
