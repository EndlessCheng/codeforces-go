package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF901C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, l, r, q int
	Fscan(in, &n, &m)
	g := make([][]int, n+1)
	for ; m > 0; m-- {
		Fscan(in, &l, &r)
		g[l] = append(g[l], r)
		g[r] = append(g[r], l)
	}

	left := make([]int, n+1)
	id := 0
	dfn := make([]int, n+1)
	fa := make([]int, n+1)
	var f func(int)
	f = func(v int) {
		id++
		dfn[v] = id
		for _, w := range g[v] {
			if dfn[w] == 0 {
				fa[w] = v
				f(w)
			} else if w != fa[v] && dfn[w] < dfn[v] {
				mi, mx := v, v
				for x := v; x != w; {
					x = fa[x]
					if x < mi {
						mi = x
					} else if x > mx {
						mx = x
					}
				}
				left[mx] = mi
			}
		}
	}
	for i, d := range dfn {
		if d == 0 {
			f(i)
		}
	}
	for i := 2; i <= n; i++ {
		if left[i-1] > left[i] {
			left[i] = left[i-1]
		}
	}
	sum := make([]int64, n+2)
	for i, v := range left {
		sum[i+1] = sum[i] + int64(v)
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r)
		p := sort.SearchInts(left[l:r+1], l) + l
		// 梯形 - 矩形 - 前缀和
		Fprintln(out, int64(l+r)*int64(r-l+1)/2-int64(p-l)*int64(l-1)-(sum[r+1]-sum[p]))
	}
}

//func main() { CF901C(os.Stdin, os.Stdout) }
