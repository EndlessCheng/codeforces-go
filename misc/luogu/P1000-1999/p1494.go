package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func p1494(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
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

	type query struct{ lb, l, r, qid int }
	qs := make([]query, q)
	blockSize := int(math.Ceil(float64(n) / math.Sqrt(float64(q))))
	for i := range qs {
		var l, r int
		Fscan(in, &l, &r)
		qs[i] = query{l / blockSize, l, r + 1, i}
	}
	sort.Slice(qs, func(i, j int) bool {
		a, b := qs[i], qs[j]
		if a.lb != b.lb {
			return a.lb < b.lb
		}
		if a.lb&1 == 0 {
			return a.r < b.r
		}
		return a.r > b.r
	})

	cnts := make([]int, n)
	pairCnt := 0
	l, r := 1, 1
	move := func(i, delta int) {
		v := a[i-1]
		if delta > 0 {
			pairCnt += 2 * cnts[v]
			cnts[v]++
		} else {
			cnts[v]--
			pairCnt -= 2 * cnts[v]
		}
	}
	ans := make([][2]int, q)
	for _, q := range qs {
		for ; r < q.r; r++ {
			move(r, 1)
		}
		for ; l < q.l; l++ {
			move(l, -1)
		}
		for l > q.l {
			l--
			move(l, 1)
		}
		for r > q.r {
			r--
			move(r, -1)
		}
		sz := q.r - q.l
		if sz == 1 {
			ans[q.qid] = [2]int{0, 1}
		} else {
			a, b := pairCnt, sz*(sz-1)
			g := gcd(a, b)
			ans[q.qid] = [2]int{a / g, b / g}
		}
	}
	for _, v := range ans {
		Fprintf(out, "%d/%d\n", v[0], v[1])
	}
}

//func main() { p1494(os.Stdin, os.Stdout) }
