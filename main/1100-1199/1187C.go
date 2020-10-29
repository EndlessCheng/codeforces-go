package main

import (
	"bufio"
	. "fmt"
	"io"
)

// O(n+m) https://www.luogu.com.cn/blog/emptyset/solution-cf1187c

// github.com/EndlessCheng/codeforces-go
func CF1187C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type fact struct{ l, r int }

	var n, m, t, l, r int
	Fscan(in, &n, &m)
	fa := make([]int, 2*n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) { fa[find(from)] = find(to) }
	fs := []fact{}
	for ; m > 0; m-- {
		Fscan(in, &t, &l, &r)
		l--
		if t > 0 {
			tar := n + l
			for i := l; i < r; i++ {
				if find(i) != i {
					tar = find(i)
					break
				}
			}
			for i := l; i < r; i++ {
				merge(i, tar)
			}
		} else {
			fs = append(fs, fact{l, r})
		}
	}
o:
	for _, f := range fs {
		for i := f.l; i < f.r; i++ {
			if find(i) != find(f.l) {
				continue o
			}
		}
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	for i := 0; i < n; {
		st := i
		for ; i < n && find(i) == find(st); i++ {
			Fprint(out, n-st, " ")
		}
	}
}

//func main() { CF1187C(os.Stdin, os.Stdout) }
