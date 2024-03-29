package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf1950G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([]int, n)
		a := make([]struct{ x, y int }, n)
		idx := map[string]int{}
		for i := range a {
			Fscan(in, &s, &t)
			if idx[s] == 0 {
				idx[s] = len(idx) + 1
			}
			a[i].x = idx[s] - 1
			if idx[t] == 0 {
				idx[t] = len(idx) + 1
			}
			a[i].y = idx[t] - 1
			for j, p := range a[:i] {
				if p.x == a[i].x || p.y == a[i].y {
					g[i] |= 1 << j
					g[j] |= 1 << i
				}
			}
		}

		f := make([][]int, 1<<n)
		for i := range f {
			f[i] = make([]int, n)
		}
		for j := range f[0] {
			f[1<<j][j] = 1
		}
		ans := 1
		for s, fs := range f {
			for _s := uint(s); _s > 0; _s &= _s - 1 {
				i := bits.TrailingZeros(_s)
				if fs[i] == 0 {
					continue
				}
				for cus, lb := g[i]&^s, 0; cus > 0; cus ^= lb {
					lb = cus & -cus
					ns := s | lb
					j := bits.TrailingZeros(uint(lb))
					f[ns][j] = max(f[ns][j], fs[i]+1)
					ans = max(ans, f[ns][j])
				}
			}
		}
		Fprintln(out, n-ans)
	}
}

//func main() { cf1950G(os.Stdin, os.Stdout) }
