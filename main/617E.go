package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol617E(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, q, k int
	Fscan(in, &n, &q, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] ^ v
	}

	ans := make([]int, q)
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
		return qi.blockIdx < qj.blockIdx || qi.blockIdx == qj.blockIdx && qi.r < qj.r
	})

	l, r, _ans := 1, 1, 0
	cnt := make([]int, 1<<20)
	update := func(idx, delta int) {
		cnt[sum[idx-1]] += delta
		tar := k ^ sum[idx]
		_ans += cnt[tar] * delta
	}
	for _, q := range qs {
		for ; l < q.l; l++ {
			update(l, -1)
		}
		for ; r < q.r; r++ {
			update(r, 1)
		}
		for l > q.l {
			l--
			update(l, 1)
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

func main() {
	Sol617E(os.Stdin, os.Stdout)
}
