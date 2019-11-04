package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol467C(reader io.Reader, writer io.Writer) {
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, sz, k int
	Fscan(in, &n, &sz, &k)
	p := make([]int64, n)
	for i := range p {
		Fscan(in, &p[i])
	}

	sum := make([]int64, n-sz+1)
	for _, v := range p[:sz] {
		sum[0] += v
	}
	for i, v := range p[sz:] {
		sum[i+1] = sum[i] + v - p[i]
	}

	dp := make([]int64, n+1)
	for i := 1; i <= k; i++ {
		for end := n; end >= sz; end-- {
			dp[end] = max(dp[end], dp[end-sz]+sum[end-sz])
		}
		maxVal := int64(0)
		for i, dpi := range dp {
			maxVal = max(maxVal, dpi)
			dp[i] = maxVal
		}
	}
	Fprint(out, dp[n])
}

//func main() {
//	Sol467C(os.Stdin, os.Stdout)
//}
