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

	var n, k int
	var s string
	Fscan(in, &n, &k, &s)
	pow := func(x, n, mod int) (res int) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}

	dp := make([][]byte, k+1)
	for i := range dp {
		dp[i] = make([]byte, n)
	}
	var f func(int, int) byte
	f = func(k, st int) (res byte) {
		if k == 0 {
			return s[st]
		}
		dv := &dp[k][st]
		if *dv > 0 {
			return *dv
		}
		defer func() { *dv = res }()
		x, y := f(k-1, st), f(k-1, (st+pow(2, k-1, n))%n)
		if x == 'R' && y == 'S' || x == 'P' && y == 'R' || x == 'S' && y == 'P' {
			return x
		}
		return y
	}
	Fprint(out, string(f(k, 0)))
}

func main() { run(os.Stdin, os.Stdout) }
