package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func p5906(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := slices.Clone(a)
	slices.Sort(b)
	b = slices.Compact(b)
	for i, v := range a {
		a[i] = sort.SearchInts(b, v)
	}
	nb := len(b)

	posL := make([]int, nb)
	for i := range posL {
		posL[i] = 1e9
	}
	posR := make([]int, nb)
	var mx int
	add := func(i int) {
		v := a[i]
		posL[v] = min(posL[v], i)
		posR[v] = max(posR[v], i)
		mx = max(mx, posR[v]-posL[v])
	}

	Fscan(in, &m)
	ans := make([]int, m)
	blockSize := int(math.Ceil(float64(n) / math.Sqrt(float64(m))))
	type query struct{ bid, l, r, qid int }
	qs := []query{}
	for i := 0; i < m; i++ {
		var l, r int
		Fscan(in, &l, &r)
		l--
		if r-l > blockSize {
			qs = append(qs, query{l / blockSize, l, r, i})
			continue
		}
		for j := l; j < r; j++ {
			add(j)
		}
		ans[i] = mx
		for _, v := range a[l:r] {
			posL[v] = 1e9
			posR[v] = 0
		}
		mx = 0
	}
	slices.SortFunc(qs, func(a, b query) int {
		if a.bid != b.bid {
			return a.bid - b.bid
		}
		return a.r - b.r
	})

	r := 0
	tmpPosL := make([]int, nb)
	tmpPosR := make([]int, nb)
	for i, q := range qs {
		l0 := (q.bid + 1) * blockSize
		if i == 0 || q.bid > qs[i-1].bid {
			r = l0
			for j := range posL {
				posL[j] = 1e9
			}
			clear(posR)
			mx = 0
		}

		for ; r < q.r; r++ {
			add(r)
		}

		tmp := mx
		for j := q.l; j < l0; j++ {
			v := a[j]
			tmpPosL[v] = posL[v]
			tmpPosR[v] = posR[v]
		}

		for j := q.l; j < l0; j++ {
			add(j)
		}
		ans[q.qid] = mx

		mx = tmp
		for _, v := range a[q.l:l0] {
			posL[v] = tmpPosL[v]
			posR[v] = tmpPosR[v]
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { p5906(bufio.NewReader(os.Stdin), os.Stdout) }
