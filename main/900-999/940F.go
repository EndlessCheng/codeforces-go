package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF940F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, pos, v int
	Fscan(in, &n, &q)
	rk := map[int]int{}
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		if rk[v] == 0 {
			rk[v] = len(rk) + 1
		}
		a[i] = rk[v]
	}
	B := int(math.Round(math.Pow(float64(n), 2.0/3)))
	type query struct{ lb, rb, l, r, t, qid int }
	type modify struct{ pos, val int }
	qs := []query{}
	ms := []modify{}
	for ; q > 0; q-- {
		if Fscan(in, &op); op == 1 {
			var l, r int
			Fscan(in, &l, &r)
			qs = append(qs, query{l / B, (r + 1) / B, l, r + 1, len(ms), len(qs)})
		} else {
			Fscan(in, &pos, &v)
			if rk[v] == 0 {
				rk[v] = len(rk) + 1
			}
			ms = append(ms, modify{pos, rk[v]})
		}
	}
	sort.Slice(qs, func(i, j int) bool {
		a, b := qs[i], qs[j]
		if a.lb != b.lb {
			return a.lb < b.lb
		}
		if a.rb != b.rb {
			return a.rb < b.rb
		}
		if a.rb&1 == 0 {
			return a.t < b.t
		}
		return a.t > b.t
	})

	cnt := make([]int, len(rk)+1)
	cc := make([]int, n+2) // 求 mex 要多开一个空间
	l, r, now := 1, 1, 0
	add := func(val int) {
		cc[cnt[val]]--
		cnt[val]++
		cc[cnt[val]]++
	}
	del := func(val int) {
		cc[cnt[val]]--
		cnt[val]--
		cc[cnt[val]]++
	}
	ans := make([]int, len(qs))
	for _, q := range qs {
		for ; r < q.r; r++ {
			add(a[r])
		}
		for ; l < q.l; l++ {
			del(a[l])
		}
		for l > q.l {
			l--
			add(a[l])
		}
		for r > q.r {
			r--
			del(a[r])
		}
		for ; now < q.t; now++ {
			m := ms[now]
			p, v := m.pos, m.val
			if q.l <= p && p < q.r {
				del(a[p])
				add(v)
			}
			a[p], ms[now].val = v, a[p]
		}
		for now > q.t {
			now--
			m := ms[now]
			p, v := m.pos, m.val
			if q.l <= p && p < q.r {
				del(a[p])
				add(v)
			}
			a[p], ms[now].val = v, a[p]
		}
		for ans[q.qid] = 1; cc[ans[q.qid]] > 0; ans[q.qid]++ {
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { CF940F(os.Stdin, os.Stdout) }
