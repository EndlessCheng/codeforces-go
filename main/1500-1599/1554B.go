package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1554B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ fi, se int }

	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		mx := bits.Len(uint(n))
		dp := make([]pair, 1<<mx)
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			dp[v] = pair{i, dp[v].fi}
		}
		for i := 0; i < mx; i++ {
			for s := 0; s < 1<<mx; s++ {
				s |= 1 << i
				p, q := dp[s], dp[s^1<<i]
				if q.se > p.fi {
					dp[s] = q
				} else if q.fi > p.fi {
					dp[s] = pair{q.fi, p.fi}
				} else if q.fi > p.se {
					dp[s].se = q.fi
				}
			}
		}
		ans := int64(-1e18)
		for s, p := range dp {
			if p.se > 0 {
				if v := int64(p.fi)*int64(p.se) - int64(k)*int64(s); v > ans {
					ans = v
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1554B(os.Stdin, os.Stdout) }
