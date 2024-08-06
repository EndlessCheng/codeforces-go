package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, q int
	Fscan(in, &n, &m, &q)
	qs := make([]struct{ op, l, r, x int }, q)
	for i := range qs {
		Fscan(in, &qs[i].op)
		if qs[i].op == 1 {
			Fscan(in, &qs[i].l, &qs[i].r, &qs[i].x)
		} else if qs[i].op == 2 {
			Fscan(in, &qs[i].l, &qs[i].x)
		} else {
			Fscan(in, &qs[i].l, &qs[i].r)
		}
	}

	t := make(fenwick, m+1)
	todo := make([][]int, n+1)
	for i := q - 1; i >= 0; i-- { // 倒着处理询问
		q := qs[i]
		if q.op == 1 {
			t.update(q.l, q.x)
			t.update(q.r+1, -q.x) // 差分更新
		} else if q.op == 2 {
			for _, j := range todo[q.l] { // 回答这一行的 op=3
				qs[j].x += t.pre(qs[j].r) + q.x
			}
			todo[q.l] = todo[q.l][:0]
		} else {
			qs[i].x = -t.pre(q.r) // 先减掉，因为后面（大于 i）的区间加不会算进来
			todo[q.l] = append(todo[q.l], i)
		}
	}
	// 更新剩下的
	for _, td := range todo {
		for _, j := range td {
			qs[j].x += t.pre(qs[j].r)
		}
	}
	for _, q := range qs {
		if q.op == 3 {
			Fprintln(out, q.x)
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
