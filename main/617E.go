package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol617E(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, q, k, v int
	Fscan(in, &n, &q, &k)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		sum[i+1] = sum[i] ^ v
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

	l, r, _ans := 1, 1, int64(0)
	cnt := make([]int, 1<<20)
	cnt[0] = 1
	update := func(idx, delta int) {
		s := sum[idx]
		tar := k ^ s
		if delta == 1 {
			_ans += int64(cnt[tar])
			cnt[s]++
		} else {
			cnt[s]--
			_ans -= int64(cnt[tar])
		}
	}
	ans := make([]int64, q)
	for _, q := range qs {
		for ; r < q.r; r++ {
			update(r, 1)
		}
		for ; l < q.l; l++ {
			update(l-1, -1)
		}
		for l > q.l {
			l--
			update(l-1, 1)
		}
		for r > q.r {
			r--
			update(r, -1)
		}
		ans[q.idx] = _ans
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() {
//	Sol617E(os.Stdin, os.Stdout)
//}
