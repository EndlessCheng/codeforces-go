package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf93E(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, k)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	const mx int = 2e5
	dp := [mx][100]int{}
	var f func(int, int) int
	f = func(x, i int) int {
		if x == 0 || i < 0 {
			return 0
		}
		if x < mx && dp[x][i] > 0 {
			return dp[x][i] - 1
		}
		res := x/a[i] + f(x, i-1) - f(x/a[i], i-1)
		if x < mx {
			dp[x][i] = res + 1
		}
		return res
	}
	Fprint(out, n-f(n, k-1))
}

//func main() { cf93E(os.Stdin, os.Stdout) }
