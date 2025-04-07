package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1791F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, op, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		fa := make([]int, n+1)
		for i := range fa {
			fa[i] = i
		}
		find := func(x int) int {
			rt := x
			for fa[rt] != rt {
				rt = fa[rt]
			}
			for fa[x] != rt {
				fa[x], x = rt, fa[x]
			}
			return rt
		}

		for range q {
			Fscan(in, &op, &l)
			l--
			if op == 2 {
				Fprintln(out, a[l])
				continue
			}
			Fscan(in, &r)
			for i := find(l); i < r; i = find(i + 1) {
				s := 0
				for v := a[i]; v > 0; v /= 10 {
					s += v % 10
				}
				a[i] = s
				if s < 10 {
					fa[i] = i + 1
				}
			}
		}
	}
}

//func main() { cf1791F(bufio.NewReader(os.Stdin), os.Stdout) }
