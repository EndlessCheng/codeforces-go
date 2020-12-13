package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

func p1494(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	calcGCD := func(a, b int) int {
		for a > 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, q int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i]--
	}

	type query struct {
		blockIdx  int
		l, r, idx int
	}
	qs := make([]query, q)
	blockSize := int(math.Round(math.Sqrt(float64(n))))
	for i := range qs {
		var l, r int
		Fscan(in, &l, &r)
		qs[i] = query{l / blockSize, l, r + 1, i}
	}
	sort.Slice(qs, func(i, j int) bool {
		qi, qj := qs[i], qs[j]
		if qi.blockIdx != qj.blockIdx {
			return qi.blockIdx < qj.blockIdx
		}
		if qi.blockIdx&1 == 0 {
			return qi.r < qj.r
		}
		return qi.r > qj.r
	})

	cnts := make([]int, n)
	pairCnt := 0
	l, r := 1, 1
	update := func(idx, delta int) {
		v := a[idx-1]
		if delta == 1 {
			pairCnt += 2 * cnts[v]
			cnts[v]++
		} else {
			cnts[v]--
			pairCnt -= 2 * cnts[v]
		}
	}
	ans := make([][2]int, q)
	getAns := func(q query) [2]int {
		l := q.r - q.l
		if l == 1 {
			return [2]int{0, 1}
		}
		a, b := pairCnt, l*(l-1)
		g := calcGCD(a, b)
		return [2]int{a / g, b / g}
	}
	for _, q := range qs {
		for ; r < q.r; r++ {
			update(r, 1)
		}
		for ; l < q.l; l++ {
			update(l, -1)
		}
		for l > q.l {
			l--
			update(l, 1)
		}
		for r > q.r {
			r--
			update(r, -1)
		}
		ans[q.idx] = getAns(q)
	}
	for _, v := range ans {
		Fprintf(out, "%d/%d\n", v[0], v[1])
	}
}

//func main() { p1494(os.Stdin, os.Stdout) }
