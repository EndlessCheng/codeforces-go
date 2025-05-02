package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1713E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		fa := make([]int, n)
		for i := range fa {
			fa[i] = i
		}
		dis := make([]int, n)
		var find func(int) int
		find = func(x int) int {
			if fa[x] != x {
				ffx := find(fa[x])
				dis[x] ^= dis[fa[x]]
				fa[x] = ffx
			}
			return fa[x]
		}
		merge := func(from, to int, b bool) bool {
			d := 0
			if b {
				d = 1
			}
			ff, ft := find(from), find(to)
			if ff != ft {
				dis[ff] = dis[to] ^ dis[from] ^ d
				fa[ff] = ft
				return true
			}
			return dis[from]^dis[to] == d
		}

		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, n)
			for j := range a[i] {
				Fscan(in, &a[i][j])
			}
		}
		for i, r := range a {
			for j := i + 1; j < n; j++ {
				v, w := r[j], a[j][i]
				if v != w && merge(i, j, v > w) != (v < w) {
					r[j], a[j][i] = w, v
				}
			}
		}
		for _, r := range a {
			for _, v := range r {
				Fprint(out, v, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { cf1713E(bufio.NewReader(os.Stdin), os.Stdout) }
