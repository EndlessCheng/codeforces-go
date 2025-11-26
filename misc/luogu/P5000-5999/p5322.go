package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func p5322(in io.Reader, out io.Writer) {
	var s, n, mx, v int
	Fscan(in, &s, &n, &mx)
	groups := make([][]int, n)
	for range s {
		for i := range groups {
			Fscan(in, &v)
			v = v*2 + 1
			if v <= mx {
				groups[i] = append(groups[i], v)
			}
		}
	}

	f := make([]int, mx+1)
	for i, g := range groups {
		slices.Sort(g)
		type pair struct{ v, w int }
		ps := []pair{}
		for j, w := range g {
			if j == len(g)-1 || w != g[j+1] {
				ps = append(ps, pair{(j + 1) * (i + 1), w})
			}
		}
		for j := mx; j >= 0; j-- {
			for _, p := range ps {
				if p.w > j {
					break
				}
				f[j] = max(f[j], f[j-p.w]+p.v)
			}
		}
	}
	Fprint(out, f[mx])
}

//func main() { p5322(bufio.NewReader(os.Stdin), os.Stdout) }
