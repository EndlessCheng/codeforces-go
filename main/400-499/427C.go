package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF427C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1_000_000_007
	var n, m, v, w int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for Fscan(in, &m); m > 0; m-- {
		Fscan(in, &v, &w)
		g[v-1] = append(g[v-1], w-1)
	}

	sum, ways := int64(0), int64(1)
	dfn := make([]int, n)
	clock := 0
	stk := []int{}
	inStk := make([]bool, n)
	var f func(int) int
	f = func(v int) int {
		clock++
		dfn[v] = clock
		lowV := clock
		stk = append(stk, v)
		inStk[v] = true
		for _, w := range g[v] {
			if dfn[w] == 0 {
				lowW := f(w)
				lowV = min(lowV, lowW)
			} else if inStk[w] {
				lowV = min(lowV, dfn[w])
			}
		}
		if dfn[v] == lowV {
			mn, cnt := int(1e9), 0
			for {
				w := stk[len(stk)-1]
				stk = stk[:len(stk)-1]
				inStk[w] = false
				if a[w] < mn {
					mn, cnt = a[w], 1
				} else if a[w] == mn {
					cnt++
				}
				if w == v {
					break
				}
			}
			sum += int64(mn)
			ways = ways * int64(cnt) % mod
		}
		return lowV
	}
	for i, ts := range dfn {
		if ts == 0 {
			f(i)
		}
	}
	Fprint(out, sum, ways)
}

//func main() { CF427C(os.Stdin, os.Stdout) }
