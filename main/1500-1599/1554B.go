package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func cf1554B(in io.Reader, out io.Writer) {
	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		mx := bits.Len(uint(n))
		type pair struct{ mx, mx2 int }
		f := make([]pair, 1<<mx)
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			f[v] = pair{i, f[v].mx}
		}
		for i := range mx {
			for s := 0; s < 1<<mx; s++ {
				s |= 1 << i
				p, q := f[s], f[s^1<<i]
				if q.mx > p.mx {
					p.mx2 = max(p.mx, q.mx2)
					p.mx = q.mx
				} else if q.mx > p.mx2 {
					p.mx2 = q.mx
				}
				f[s] = p
			}
		}
		ans := int(-1e18)
		for s, p := range f {
			if p.mx2 > 0 {
				ans = max(ans, p.mx*p.mx2-k*s)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1554B(bufio.NewReader(os.Stdin), os.Stdout) }
