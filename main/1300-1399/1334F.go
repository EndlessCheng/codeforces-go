package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type fenwick34 []int

func (f fenwick34) update(i int, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] += val
	}
}
func (f fenwick34) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}
func (f fenwick34) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func cf1334F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n)
	a := make([]int, n)
	pos := make([][]int, n+1)
	negSum := make([]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		pos[a[i]] = append(pos[a[i]], i)
	}
	cost := make([]int, n)
	for i := range cost {
		Fscan(in, &cost[i])
		negSum[i+1] = negSum[i]
		if cost[i] < 0 {
			negSum[i+1] += cost[i]
		}
	}
	Fscan(in, &m)
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
	}

	t := make(fenwick34, n+1)
	f := make([]int, n)
	for i := range f {
		f[i] = 1e18
	}
	s := 0
	for i, v := range a {
		if v > b[0] {
			if cost[i] > 0 {
				t.update(i, cost[i])
			}
		} else if v == b[0] {
			f[i] = s
		}
		s += cost[i]
	}

	pid := 1
	for pid <= b[0] {
		pid++
	}

	for bi := 1; bi < m; bi++ {
		p := pos[b[bi-1]]
		j := 0
		mn := int(1e18)
		pre := -1
		for _, i := range pos[b[bi]] {
			if pre >= 0 {
				mn += t.query(pre, i-1) + negSum[i] - negSum[pre]
			}
			for j < len(p) && p[j] < i {
				res := f[p[j]] + t.query(p[j]+1, i-1) + negSum[i] - negSum[p[j]+1]
				mn = min(mn, res)
				j++
			}
			f[i] = mn
			pre = i
		}
		for pid <= b[bi] {
			for _, i := range pos[pid] {
				if cost[i] > 0 {
					t.update(i, -cost[i])
				}
			}
			pid++
		}
	}

	ans := int(1e18)
	s = 0
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		if v == b[m-1] {
			ans = min(ans, f[i]+s)
		}
		if v > b[m-1] || cost[i] < 0 {
			s += cost[i]
		}
	}
	if ans >= 1e17 {
		Fprint(out, "NO")
	} else {
		Fprintln(out, "YES")
		Fprint(out, ans)
	}
}

//func main() { cf1334F(os.Stdin, os.Stdout) }
