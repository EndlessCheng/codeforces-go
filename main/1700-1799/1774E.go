package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1774E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, d, v, w, m, ans int
	Fscan(in, &n, &d)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	tp := make([]int, n)
	for msk := 1; msk < 3; msk <<= 1 {
		for Fscan(in, &m); m > 0; m-- {
			Fscan(in, &v)
			tp[v-1] |= msk
		}
	}

	st := []int{-1}
	var f func(int)
	f = func(v int) {
		st = append(st, v)
		for _, w := range g[v] {
			if w != st[len(st)-2] {
				f(w)
				tp[v] |= tp[w]
			}
		}
		if len(st) > d+1 && tp[v] > 0 {
			tp[st[len(st)-d-1]] = 3
		}
		st = st[:len(st)-1]
	}
	f(0)

	for _, v := range tp[1:] {
		ans += v>>1 + v&1
	}
	Fprint(out, ans*2)
}

//func main() { CF1774E(os.Stdin, os.Stdout) }
