package main

import (
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, k, v, ans int
	Fscan(in, &n, &k)

	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, k*25+1) // 25=log(1e18)/log(5)
		for j := range f[i] {
			f[i][j] = -1e9
		}
	}
	f[0][0] = 0
	for ; n > 0; n-- {
		Fscan(in, &v)
		c2 := bits.TrailingZeros(uint(v))
		c5 := 0
		for ; v%5 == 0; v /= 5 {
			c5++
		}
		for i := k; i > 0; i-- {
			for j := i * 25; j >= c5; j-- {
				f[i][j] = max(f[i][j], f[i-1][j-c5]+c2)
			}
		}
	}

	for c5, c2 := range f[k] {
		ans = max(ans, min(c5, c2))
	}
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
