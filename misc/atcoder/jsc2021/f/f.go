package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
type fenwick []int

func (f fenwick) add(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick) sum(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, q, ans int
	Fscan(in, &n, &m, &q)
	qs := make([]struct{ op, i, x int }, q)
	d := make([]int, q+1)
	for i := range qs {
		Fscan(in, &qs[i].op, &qs[i].i, &qs[i].x)
		d[i] = qs[i].x
	}
	sort.Ints(d)
	k := 0
	for _, w := range d[1:] {
		if d[k] != w {
			k++
			d[k] = w
		}
	}
	d = d[:k+1]

	ab := [2][]int{make([]int, n+1), make([]int, m+1)}
	sum := [2]fenwick{make(fenwick, k+3), make(fenwick, k+3)}
	cnt := [2]fenwick{make(fenwick, k+3), make(fenwick, k+3)}
	cnt[0].add(1, n)
	cnt[1].add(1, m)
	for _, q := range qs {
		op := q.op - 1
		i := q.i
		a := ab[op]
		old := a[i]
		a[i] = q.x
		preI := sort.SearchInts(d, old) + 2
		curI := sort.SearchInts(d, q.x) + 2
		ans -= cnt[op^1].sum(preI)*old - sum[op^1].sum(preI)
		ans += cnt[op^1].sum(curI)*q.x - sum[op^1].sum(curI) // + - tot 互相抵消
		cnt[op].add(preI, -1)
		cnt[op].add(curI, 1)
		sum[op].add(preI, -old)
		sum[op].add(curI, q.x)
		Fprintln(out, ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }
