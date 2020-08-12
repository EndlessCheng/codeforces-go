package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1067B(_r io.Reader, _w io.Writer) (ans bool) {
	in := bufio.NewReader(_r)
	defer func() {
		if ans {
			Fprint(_w, "Yes")
		} else {
			Fprint(_w, "No")
		}
	}()

	var n, k, v, w int
	Fscan(in, &n, &k)
	if n < 4 {
		return false
	}
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	vis := make([]bool, n+1)
	centers := map[int]int{}
	for i, gi := range g {
		if len(gi) == 1 {
			vis[i] = true
			centers[gi[0]]++
		}
	}
	for {
		k--
		newCenters := map[int]int{}
		for v, cnt1 := range centers {
			if cnt1 < 3 {
				return false
			}
			ct := 0
			for _, w := range g[v] {
				if !vis[w] {
					if ct != 0 || centers[w] > 0 {
						return false
					}
					ct = w
				}
			}
			if ct == 0 {
				return k == 0
			}
			vis[v] = true
			newCenters[ct]++
		}
		centers = newCenters
	}
}

//func main() {
//	CF1067B(os.Stdin, os.Stdout)
//}
