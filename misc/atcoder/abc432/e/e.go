package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
type pair struct{ cnt, sum int }
type fenwick []pair

func (t fenwick) update(i, c, v int) {
	for i++; i < len(t); i += i & -i {
		t[i].cnt += c
		t[i].sum += v
	}
}

func (t fenwick) pre(i int) (res pair) {
	for i++; i > 0; i &= i - 1 {
		res.cnt += t[i].cnt
		res.sum += t[i].sum
	}
	return
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n)
	t := make(fenwick, 5e5+2)
	for i := range a {
		Fscan(in, &a[i])
		t.update(a[i], 1, a[i])
	}

	for range q {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			l--
			t.update(a[l], -1, -a[l])
			a[l] = r
			t.update(r, 1, r)
		} else if l >= r {
			Fprintln(out, n*l)
		} else {
			pl := t.pre(l)
			pr := t.pre(r)
			Fprintln(out, pl.cnt*l+pr.sum-pl.sum+(n-pr.cnt)*r)
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
