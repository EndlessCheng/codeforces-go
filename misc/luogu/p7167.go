package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p7167(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, cur, left int
	Fscan(in, &n, &q)
	type ft struct{ d, c int }
	a := make([]ft, n)
	for i := range a {
		Fscan(in, &a[i].d, &a[i].c)
	}

	const mx = 17
	type pair struct{ to, s int }
	pa := make([][mx]pair, n+1)
	pa[n][0].to = n
	st := []int{n}
	for i := n - 1; i >= 0; i-- {
		for len(st) > 1 && a[i].d >= a[st[len(st)-1]].d {
			st = st[:len(st)-1]
		}
		pa[i][0] = pair{st[len(st)-1], a[i].c}
		st = append(st, i)
	}

	for i := 0; i < mx-1; i++ {
		for x := range pa {
			p := pa[x][i]
			q := pa[p.to][i]
			pa[x][i+1] = pair{q.to, p.s + q.s}
		}
	}

	for ; q > 0; q-- {
		Fscan(in, &cur, &left)
		cur--
		for k := mx - 1; k >= 0; k-- {
			p := pa[cur][k]
			if left > p.s {
				left -= p.s
				cur = p.to
			}
		}
		p := pa[cur][0]
		if left > p.s {
			cur = p.to
		}
		if cur == n {
			cur = -1
		}
		Fprintln(out, cur+1)
	}
}

//func main() { p7167(bufio.NewReader(os.Stdin), os.Stdout) }
