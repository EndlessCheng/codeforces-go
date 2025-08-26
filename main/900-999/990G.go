package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf990G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	const mx = 200001
	var cnt, time, fa, size, f [mx]int
	var n, now int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		cnt[a[i]]++
	}
	es := [mx][][2]int{}
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g := gcd(a[v], a[w])
		es[g] = append(es[g], [2]int{v, w})
	}

	find := func(x int) int {
		rt := x
		for {
			if time[rt] != now {
				time[rt] = now
				fa[rt] = rt
				size[rt] = 1
			}
			if fa[rt] == rt {
				break
			}
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	for i := mx - 1; i > 0; i-- {
		now = i
		for j := i; j < mx; j += i {
			for _, e := range es[j] {
				x := find(e[0])
				y := find(e[1])
				if x != y {
					f[i] += size[x] * size[y]
					size[x] += size[y]
					fa[y] = x
				}
			}
		}
		for j := i * 2; j < mx; j += i {
			f[i] -= f[j]
		}
	}

	for i, v := range f[:] {
		v += cnt[i]
		if v > 0 {
			Fprintln(out, i, v)
		}
	}
}

//func main() { cf990G(bufio.NewReader(os.Stdin), os.Stdout) }
