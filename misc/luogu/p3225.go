package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p3225(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	G := [1000][]int{}
	var m int
	for T := 1; ; T++ {
		Fscan(in, &m)
		if m == 0 {
			break
		}
		n := 0
		for ; m > 0; m-- {
			var v, w int
			Fscan(in, &v, &w)
			n = max(n, v, w)
			v--
			w--
			G[v] = append(G[v], w)
			G[w] = append(G[w], v)
		}
		g := G[:n]

		isCut := make([]bool, n)
		comps := [][]int{}
		st := []int{}
		dfn := make([]int, n)
		clock := 0
		var tarjan func(int, int) int
		tarjan = func(v, fa int) int {
			clock++
			dfn[v] = clock
			lowV := clock
			childCnt := 0
			st = append(st, v)
			for _, w := range g[v] {
				if dfn[w] == 0 {
					childCnt++
					lowW := tarjan(w, v)
					lowV = min(lowV, lowW)
					if lowW >= dfn[v] {
						isCut[v] = true
						cc := []int{v}
						for {
							u := st[len(st)-1]
							st = st[:len(st)-1]
							cc = append(cc, u)
							if u == w {
								break
							}
						}
						comps = append(comps, cc)
					}
				} else {
					lowV = min(lowV, dfn[w])
				}
			}
			if fa < 0 && childCnt == 1 {
				isCut[v] = false
			}
			return lowV
		}
		tarjan(0, -1)

		ans, ways := 0, 1
		for _, cc := range comps {
			cutCnt := 0
			for _, v := range cc {
				if isCut[v] {
					cutCnt++
				}
			}
			sz := len(cc)
			if cutCnt == 0 {
				ans += 2
				ways *= sz * (sz - 1) / 2
			} else if cutCnt == 1 {
				ans++
				ways *= sz - 1
			}
		}
		Fprintf(out, "Case %d: %d %d\n", T, ans, ways)
		clear(g)
	}
}

//func main() { p3225(bufio.NewReader(os.Stdin), os.Stdout) }
