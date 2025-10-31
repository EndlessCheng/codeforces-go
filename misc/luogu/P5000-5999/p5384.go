package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func p5384(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12)
	rc := func() byte {
		if _i == _n {
			_n, _ = in.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int32) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int32(b-'0')
		}
		return
	}

	n, q := r(), r()
	g := make([][]int32, n)
	for w := int32(1); w < n; w++ {
		v := r() - 1
		g[v] = append(g[v], w)
	}

	// 需要把 dfs 写到 main 外面 https://www.luogu.com.cn/record/176578013
	const mx = 20
	pa := make([][mx]int32, n)
	type info struct{ in, out, d int32 }
	is := make([]info, n)
	depT := make([][]int32, n)
	t := int32(0)
	var f func(v, p, d int32)
	f = func(v, p, d int32) {
		pa[v][0] = p
		t++
		is[v].in = t
		is[v].d = d
		depT[d] = append(depT[d], t)
		for _, w := range g[v] {
			if w != p {
				f(w, v, d+1)
			}
		}
		is[v].out = t
	}
	f(0, -1, 0)
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}

	uptoKthPa := func(v, k int32) int32 {
		for i := 0; i < mx && v != -1; i++ {
			if k>>i&1 == 1 {
				v = pa[v][i]
			}
		}
		return v
	}
	query := func(v, d int32) int {
		i := is[v]
		d += i.d
		a := depT[d]
		l := sort.Search(len(a), func(j int) bool { return a[j] >= i.in })
		r := sort.Search(len(a), func(j int) bool { return a[j] >= i.out+1 })
		return r - l
	}
	for ; q > 0; q-- {
		v, k := r()-1, r()
		if v = uptoKthPa(v, k); v == -1 {
			Fprint(out, "0 ")
		} else {
			Fprint(out, query(v, k)-1, " ")
		}
	}
}

//func main() { p5384(os.Stdin, os.Stdout) }
