package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf505E(in io.Reader, out io.Writer) {
	var n, m, k, dec int
	Fscan(in, &n, &m, &k, &dec)
	a := make([]struct{ h, grow int }, n)
	for i := range a {
		Fscan(in, &a[i].h, &a[i].grow)
	}
	ans := sort.Search((m+1)*1e9, func(mx int) bool {
		cnt := make([]int, m)
		left := m * k
		for _, p := range a {
			t := max((p.h+p.grow*m-mx+dec-1)/dec, 0)
			left -= t
			if left < 0 {
				return false
			}
			for i := range t {
				if d := (mx + i*dec) / p.grow; d < m {
					cnt[d]++
				}
			}
		}

		s := 0
		for i, c := range cnt {
			s += c
			if s > i*k {
				return false
			}
		}
		return true
	})
	Fprint(out, ans)
}

//func main() { cf505E(bufio.NewReader(os.Stdin), os.Stdout) }
