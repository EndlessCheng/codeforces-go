package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF837D(in io.Reader, out io.Writer) {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, k, ans int
	var v uint64
	Fscan(in, &n, &k)
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, k*25+1)
		for j := range f[i] {
			f[i][j] = -1
		}
	}
	f[0][0] = 0
	for i, s5 := 1, 0; i <= n; i++ {
		Fscan(in, &v)
		c2 := bits.TrailingZeros64(v)
		c5 := 0
		for ; v%5 == 0; v /= 5 {
			c5++
		}
		s5 += c5
		for j := min(i, k); j > 0; j-- {
			for p := min(s5, k*25); p >= c5; p-- {
				if f[j-1][p-c5] >= 0 {
					f[j][p] = max(f[j][p], f[j-1][p-c5]+c2)
				}
			}
		}
	}
	for i, c2 := range f[k] {
		ans = max(ans, min(i, c2))
	}
	Fprint(out, ans)
}

//func main() { CF837D(os.Stdin, os.Stdout) }
