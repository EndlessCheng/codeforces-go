package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF797E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, p, k int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	blockSize := int(math.Sqrt(float64(n)))
	dp := make([][]int, blockSize)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int) int
	f = func(p int) (res int) {
		if p > n {
			return
		}
		dv := &dp[k][p]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		return 1 + f(p+a[p]+k)
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &p, &k)
		if k < blockSize {
			Fprintln(out, f(p))
		} else {
			c := 0
			for ; p <= n; p += a[p] + k {
				c++
			}
			Fprintln(out, c)
		}
	}
}

//func main() { CF797E(os.Stdin, os.Stdout) }
