package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1594D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, v, w int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)

		fa := make([]int, n)
		for i := range fa {
			fa[i] = i
		}
		dis := make([]uint8, n)
		var find func(int) int
		find = func(x int) int {
			if fa[x] != x {
				ffx := find(fa[x])
				dis[x] ^= dis[fa[x]]
				fa[x] = ffx
			}
			return fa[x]
		}
		merge := func(x, y int, d uint8) bool {
			fx, fy := find(x), find(y)
			if fx == fy {
				return dis[x]^dis[y] == d
			}
			dis[fx] = d ^ dis[y] ^ dis[x]
			fa[fx] = fy
			return true
		}

		ok := true
		for range m {
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
		for i := range n {
			cnt[find(i)][dis[i]]++
		}
		ans := 0
		for _, c := range cnt {
			ans += max(c[0], c[1])
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1594D(bufio.NewReader(os.Stdin), os.Stdout) }
