package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1594D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n, m, v, w int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		fa := make([]int, n)
		for i := range fa {
			fa[i] = i
		}
		dis := make([]uint8, n)
		var f func(int) int
		f = func(x int) int {
			if fa[x] != x {
				ffx := f(fa[x])
				dis[x] ^= dis[fa[x]]
				fa[x] = ffx
			}
			return fa[x]
		}
		merge := func(x, y int, d uint8) bool {
			if fx, fy := f(x), f(y); fx != fy {
				dis[fx] = d ^ dis[y] ^ dis[x]
				fa[fx] = fy
				return true
			}
			f(x)
			f(y)
			return dis[x]^dis[y] == d
		}
		ok := true
		for ; m > 0; m-- {
			Fscan(in, &v, &w, &s)
			if ok && !merge(v-1, w-1, s[1]&1) {
				ok = false
			}
		}
		if !ok {
			Fprintln(out, -1)
			continue
		}
		cnt := make([][2]int, n)
		for i := range dis {
			cnt[f(i)][dis[i]]++
		}
		ans := 0
		for _, p := range cnt {
			if p[0] > 0 || p[1] > 0 {
				ans += max(p[0], p[1])
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1594D(os.Stdin, os.Stdout) }
