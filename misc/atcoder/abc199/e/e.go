package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n, &m)
	a := make([]struct{ x, y, z int }, m)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y, &a[i].z)
	}

	N := uint(1 << n)
	dp := make([]int, N)
	dp[0] = 1
	for s := uint(0); s < N-1; s++ {
		o := bits.OnesCount(s)
	o:
		for t, lb := s^(N-1), uint(0); t > 0; t ^= lb {
			lb = t & -t
			for _, p := range a {
				if o < p.x && bits.OnesCount((s|lb)&(1<<p.y-1)) > p.z {
					continue o
				}
			}
			dp[s|lb] += dp[s]
		}
	}
	Fprint(out, dp[N-1])
}

func main() { run(os.Stdin, os.Stdout) }
