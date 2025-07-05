package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1986G2(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	const mx = 500001
	ds := [mx][]uint32{}
	for i := uint32(1); i < mx; i++ {
		for j := i; j < mx; j += i {
			ds[j] = append(ds[j], i)
		}
	}

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 0
		p2i := make([][]uint32, n+1)
		i2p := make([][]uint32, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			g := gcd(v, i)
			p, i := v/g, i/g
			if i == 1 {
				ans--
			}
			p2i[p] = append(p2i[p], uint32(i))
			i2p[i] = append(i2p[i], uint32(p))
		}
		cnt := make([]uint32, n+1)
		for i := 1; i <= n; i++ {
			for pj := i; pj <= n; pj += i {
				for _, j := range p2i[pj] {
					cnt[j]++
				}
			}
			for _, pi := range i2p[i] {
				for _, j := range ds[pi] {
					ans += int(cnt[j])
				}
			}
			for pj := i; pj <= n; pj += i {
				for _, j := range p2i[pj] {
					cnt[j]--
				}
			}
		}
		Fprintln(out, ans/2)
	}
}

//func main() { cf1986G2(bufio.NewReader(os.Stdin), os.Stdout) }
