package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func cf1971G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		pos := map[int][]int{}
		for i := range a {
			Fscan(in, &a[i])
			pos[a[i]] = append(pos[a[i]], i)
		}

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
		merge := func(from, to int) { fa[find(from)] = find(to) }
		for v, pv := range pos {
			i0 := pv[0]
			for _, i := range pv[1:] {
				merge(i, i0)
			}
			for xor := 1; xor < 4; xor++ {
				ps := pos[v^xor]
				if ps != nil {
					merge(ps[0], i0)
				}
			}
		}

		gs := make([][]int, n)
		for i := 0; i < n; i++ {
			rt := find(i)
			gs[rt] = append(gs[rt], i)
		}
		for _, g := range gs {
			if g == nil {
				continue
			}
			b := make([]int, len(g))
			for id, i := range g {
				b[id] = a[i]
			}
			slices.Sort(b)
			for id, i := range g {
				a[i] = b[id]
			}
		}
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1971G(bufio.NewReader(os.Stdin), os.Stdout) }
