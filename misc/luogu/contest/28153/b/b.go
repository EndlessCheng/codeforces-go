package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, buf := 1<<12, make([]byte, 1<<12)
	rc := func() byte {
		if _i == 1<<12 {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}
	type pair struct{ v, i int }
	type match struct{ me, op []pair }

	n, m, delta, cnt := r(), r(), r(), r()
	tp := [1e5 + 1]match{}
	for i := 1; i <= n; i++ {
		c := r()
		tp[c].me = append(tp[c].me, pair{r(), i})
	}
	for i := 0; i < m; i++ {
		c := r()
		tp[c].op = append(tp[c].op, pair{r(), i})
	}
	ans := make([]int, m)
	for _, m := range tp[:] {
		me, op := m.me, m.op
		sort.Slice(me, func(i, j int) bool { return me[i].v < me[j].v })
		sort.Slice(op, func(i, j int) bool { return op[i].v < op[j].v })
		if len(me) > len(op) {
			me = me[len(me)-len(op):]
		}
		for _, p := range op[len(me):] {
			ans[p.i] = -1
		}
		win := len(me) - len(op)
		op = op[:len(me)]
		for _, p := range me {
			cnt += p.v
			if o := op[0]; p.v < o.v {
				win--
				ans[op[len(op)-1].i] = p.i
				op = op[:len(op)-1]
			} else {
				win++
				ans[o.i] = p.i
				op = op[1:]
			}
		}
		cnt += win * delta
	}
	Fprintln(out, cnt)
	for _, v := range ans {
		Fprintln(out, v)
	}
}

func main() { run(os.Stdin, os.Stdout) }
