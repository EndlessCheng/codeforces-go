package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF86D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type query struct{ bid, l, r, qid int }
	qs := make([]query, q)
	blockSize := int(math.Round(math.Sqrt(float64(n))))
	for i := range qs {
		Fscan(in, &l, &r)
		qs[i] = query{l / blockSize, l, r + 1, i}
	}
	sort.Slice(qs, func(i, j int) bool {
		a, b := qs[i], qs[j]
		if a.bid != b.bid {
			return a.bid < b.bid
		}
		if a.bid&1 == 0 {
			return a.r < b.r
		}
		return a.r > b.r
	})

	sum, cnt, l, r := int64(0), [1e6 + 1]int{}, 1, 1
	update := func(i, d int) {
		v := a[i-1]
		sum += int64(2*d*cnt[v]+1) * int64(v)
		cnt[v] += d
	}
	ans := make([]int64, q)
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
		ans[q.qid] = sum
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { CF86D(os.Stdin, os.Stdout) }
