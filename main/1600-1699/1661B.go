package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1661B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, v int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		if v == 0 {
			Fprint(out, "0 ")
			continue
		}
		ans := 15
		for i := 0; i < 15; i++ {
			ans = min(ans, i+15-bits.TrailingZeros(uint(v+i)))
		}
		Fprint(out, ans, " ")
	}
}

func CF1661B_BFS(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx = 32768
	g := [mx][]int{}
	for w := range g {
		v := (w + 1) % mx
		g[v] = append(g[v], w)
		v = w * 2 % mx
		g[v] = append(g[v], w)
	}
	dis := [mx]int{-1}
	q := []int{0}
	for k := 1; len(q) > 0; k++ {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, w := range g[v] {
				if dis[w] == 0 {
					dis[w] = k
					q = append(q, w)
				}
			}
		}
	}
	dis[0] = 0

	var n, v int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		Fprint(out, dis[v], " ")
	}
}

//func main() { CF1661B(os.Stdin, os.Stdout) }
