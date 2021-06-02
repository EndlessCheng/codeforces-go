package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ fi, se int }

	var mx int
	Fscan(in, &mx)
	dp := make([]pair, 1<<mx)
	for i := range dp {
		Fscan(in, &dp[i].fi)
	}
	for i := 0; i < mx; i++ {
		for s := 0; s < 1<<mx; s++ {
			if s>>i&1 == 0 {
				s |= 1 << i
			}
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
	mx = 0
	for _, p := range dp[1:] {
		mx = max(mx, p.fi+p.se)
		Fprintln(out, mx)
	}
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
