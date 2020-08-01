package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF940E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	var n, c int
	Fscan(in, &n, &c)
	s := int64(0)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s += int64(a[i])
	}

	dp := make([]int64, n+1)
	idQ := make([]int, n)
	l, r := 0, 0
	for i, v := range a {
		if l < r && i-idQ[l]+1 > c {
			l++
		}
		for ; l < r && a[idQ[r-1]] >= v; r-- {
		}
		idQ[r] = i
		r++
		if i+1 >= c {
			dp[i+1] = max(dp[i], dp[i+1-c]+int64(a[idQ[l]]))
		}
	}
	Fprint(out, s-dp[n])
}

func main() { CF940E(os.Stdin, os.Stdout) }
