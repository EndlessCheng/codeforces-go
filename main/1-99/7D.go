package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF7D(in io.Reader, out io.Writer) {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var s []byte
	Fscan(bufio.NewReader(in), &s)
	n := len(s)
	t := append(make([]byte, 0, n*2+3), '^')
	for _, c := range s {
		t = append(t, '#', c)
	}
	t = append(t, '#', '$')
	halfLen := make([]int, n*2+1)
	halfLen[1] = 1
	for i, mid, r := 2, 1, 0; i < len(halfLen); i++ {
		hl := 1
		if i < r {
			hl = min(halfLen[mid*2-i], r-i)
		}
		for ; t[i-hl] == t[i+hl]; hl++ {
		}
		if i+hl > r {
			mid, r = i, i+hl
		}
		halfLen[i] = hl
	}

	ans := 0
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if halfLen[i+1] > i {
			dp[i] = dp[i/2] + 1
			ans += dp[i]
		}
	}
	Fprint(out, ans)
}

//func main() { CF7D(os.Stdin, os.Stdout) }
