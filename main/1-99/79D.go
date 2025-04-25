package main

import (
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://github.com/EndlessCheng
func cf79D(in io.Reader, out io.Writer) {
	var n, k, l int
	Fscan(in, &n, &k, &l)
	x := make([]int, k)
	for i := range x {
		Fscan(in, &x[i])
	}
	t := []int{}
	for i := 0; i < k; {
		t = append(t, x[i])
		for i++; i < k && x[i] == x[i-1]+1; i++ {
		}
		t = append(t, x[i-1]+1)
	}
	x = t
	k = len(x)

	a := make([]int, l)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	dis := make([]int, n+2)
	bfs := func(st int) []int {
		for i := range dis {
			dis[i] = 1e9
		}
		q := []int{st}
		dis[st] = 0
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, w := range a {
				if v > w && dis[v]+1 < dis[v-w] {
					dis[v-w] = dis[v] + 1
					q = append(q, v-w)
				}
				if v+w < len(dis) && dis[v]+1 < dis[v+w] {
					dis[v+w] = dis[v] + 1
					q = append(q, v+w)
				}
			}
		}
		dx := make([]int, k)
		for i, v := range x {
			dx[i] = dis[v]
		}
		return dx
	}
	dx := make([][]int, k)
	for i, v := range x {
		dx[i] = bfs(v)
	}

	f := make([]int, 1<<k)
	for i := 1; i < len(f); i++ {
		if bits.OnesCount(uint(i))&1 > 0 {
			continue
		}
		f[i] = 1e9
		tz := bits.TrailingZeros(uint(i))
		for m := uint(i ^ 1<<tz); m > 0; m &= m - 1 {
			j := bits.TrailingZeros(m)
			f[i] = min(f[i], f[i^1<<tz^1<<j]+dx[tz][j])
		}
	}
	ans := f[1<<k-1]
	if ans == 1e9 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { cf79D(os.Stdin, os.Stdout) }
