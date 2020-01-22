package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1067B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, v, w, ct int
	Fscan(in, &n, &k)
	if n < 4 {
		Fprint(out, "No")
		return
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	vis := make([]bool, n)
	centers := map[int]bool{}
	for i, gi := range g {
		if len(gi) == 1 {
			vis[i] = true
			centers[gi[0]] = true
		}
	}
	for {
		k--
		newCenters := map[int]bool{}
		for v := range centers {
			cnt1, cnt2 := 0, 0
			for _, w := range g[v] {
				if vis[w] {
					cnt1++
				} else {
					if centers[w] {
						Fprint(out, "No")
						return
					}
					cnt2++
					ct = w
				}
			}
			if cnt1 < 3 || cnt2 > 1 {
				Fprint(out, "No")
				return
			}
			if cnt2 == 0 {
				if k == 0 {
					Fprint(out, "Yes")
				} else {
					Fprint(out, "No")
				}
				return
			}
			vis[v] = true
			newCenters[ct] = true
		}
		centers = newCenters
	}
}

//func main() {
//	CF1067B(os.Stdin, os.Stdout)
//}
