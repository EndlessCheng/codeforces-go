package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF859E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7

	var m, v, w int
	Fscan(in, &m)
	n := 2 * m
	fa := make([]int, n)
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
	sl := make([]bool, n)
	c := make([]bool, n)
	sz := make([]int, n)
	for i := range sz {
		sz[i] = 1
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		if v == w {
			sl[v] = true
			continue
		}
		v, w = find(v), find(w)
		if v == w {
			c[v] = true
			continue
		}
		fa[v] = w
		sz[w] += sz[v]
		if sl[v] {
			sl[w] = true
		}
	}
	ans := int64(1)
	for i, b := range sl {
		if !b && i == find(i) {
			if c[i] {
				ans = ans * 2 % mod
			} else {
				ans = ans * int64(sz[i]) % mod
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF859E(os.Stdin, os.Stdout) }
