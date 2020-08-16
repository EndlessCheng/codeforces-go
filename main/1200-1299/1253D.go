package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1253D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w, ans int
	Fscan(in, &n, &m)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			fa[x] = f(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) int {
		ff, ft := f(from), f(to)
		if ff > ft {
			ff, ft = ft, ff
		}
		fa[ff] = ft
		return ft
	}

	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		merge(v, w)
	}
	end := -1
	for i := range fa {
		if i > end {
			end = f(i)
		} else if f(i) != end {
			end = merge(i, end)
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { CF1253D(os.Stdin, os.Stdout) }
