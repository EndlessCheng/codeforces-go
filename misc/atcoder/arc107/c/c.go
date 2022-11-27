package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	var n, k int
	Fscan(in, &n, &k)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	ans := 1
	f := func() {
		fa := make([]int, n)
		sz := make([]int, n)
		for i := range fa {
			fa[i] = i
			sz[i] = 1
		}
		var find func(int) int
		find = func(x int) int {
			if fa[x] != x {
				fa[x] = find(fa[x])
			}
			return fa[x]
		}
		for i, r := range a {
			fi := find(i)
		o:
			for i2, r2 := range a[:i] {
				fi2 := find(i2)
				if fi2 == fi {
					continue
				}
				for j, v := range r {
					if v+r2[j] > k {
						continue o
					}
				}
				fa[fi2] = fi
				sz[fi] += sz[fi2]
				sz[fi2] = 0
			}
		}
		for _, v := range sz {
			for v > 1 {
				ans = ans * v % mod
				v--
			}
		}
	}
	f()
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			a[i][j], a[j][i] = a[j][i], a[i][j]
		}
	}
	f()
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
