package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF264B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mx = 1e5 + 1
	ds := [mx][]int{1: {1}}
	for i := 2; i < mx; i++ {
		for j := i; j < mx; j += i {
			ds[j] = append(ds[j], i)
		}
	}

	var n, v, ans int
	dp := [mx]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		max := 0
		for _, d := range ds[v] {
			if dp[d] > max {
				max = dp[d]
			}
		}
		max++
		for _, d := range ds[v] {
			dp[d] = max
		}
		if max > ans {
			ans = max
		}
	}
	Fprint(out, ans)
}

//func main() { CF264B(os.Stdin, os.Stdout) }
