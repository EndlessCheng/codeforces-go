package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
var nodes40 [22 * 3e5]struct{ lo, ro, sum int32 }
var pid40 int32 = -1

func build40(l, r int32) int32 {
	pid40++
	o := pid40
	if l == r {
		return pid40
	}
	m := (l + r) >> 1
	nodes40[o].lo = build40(l, m)
	nodes40[o].ro = build40(m+1, r)
	return o
}

func add40(old, l, r, i int32) int32 {
	pid40++
	o := pid40
	nodes40[o] = nodes40[old]
	if l == r {
		nodes40[o].sum++
		return o
	}
	m := (l + r) >> 1
	if i <= m {
		nodes40[o].lo = add40(nodes40[o].lo, l, m, i)
	} else {
		nodes40[o].ro = add40(nodes40[o].ro, m+1, r, i)
	}
	nodes40[o].sum = nodes40[nodes40[o].lo].sum + nodes40[nodes40[o].ro].sum
	return o
}

func kth40(o, old, l, r, k int32) int32 {
	if l == r {
		return l
	}
	cntL := nodes40[nodes40[o].lo].sum - nodes40[nodes40[old].lo].sum
	m := (l + r) >> 1
	if k <= cntL {
		return kth40(nodes40[o].lo, nodes40[old].lo, l, m, k)
	}
	return kth40(nodes40[o].ro, nodes40[old].ro, m+1, r, k-cntL)
}

func query40(o, old, l, r, i int32) int32 {
	if l == r {
		return nodes40[o].sum - nodes40[old].sum
	}
	m := (l + r) >> 1
	if i <= m {
		return query40(nodes40[o].lo, nodes40[old].lo, l, m, i)
	}
	return query40(nodes40[o].ro, nodes40[old].ro, m+1, r, i)
}

func cf840D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r, k int32
	Fscan(in, &n, &q)
	t := make([]int32, n+1)
	t[0] = build40(1, n)
	for i := int32(1); i <= n; i++ {
		Fscan(in, &k)
		t[i] = add40(t[i-1], 1, n, k)
	}
o:
	for ; q > 0; q-- {
		Fscan(in, &l, &r, &k)
		d := (r-l+1)/k + 1
		for rk := int32(1); rk <= r-l+1; rk += d {
			v := kth40(t[r], t[l-1], 1, n, rk)
			if query40(t[r], t[l-1], 1, n, v) >= d {
				Fprintln(out, v)
				continue o
			}
		}
		Fprintln(out, -1)
	}
}

//func main() { cf840D(os.Stdin, os.Stdout) }
