package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1391D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, m int
	Fscanln(in, &n, &m)
	if n > 3 && m > 3 {
		Fprint(out, -1)
		return
	}
	if n == 1 || m == 1 {
		Fprint(out, 0)
		return
	}
	var a, s []byte
	if n > m {
		a = make([]byte, n)
		for i := range a {
			Fscanf(in, "%b\n", &a[i])
		}
	} else {
		n, m = m, n
		a = make([]byte, n)
		for j := 0; j < m; j++ {
			Fscan(in, &s)
			for i, b := range s {
				a[i] |= b & 1 << j
			}
		}
	}

	if m == 2 {
		from := [4][2]byte{{1, 2}, {0, 3}, {0, 3}, {1, 2}}
		dp := [4]int{}
		for i := byte(0); i < 4; i++ {
			dp[i] = bits.OnesCount8(a[0] ^ i)
		}
		for _, v := range a[1:] {
			d := dp
			for i := byte(0); i < 4; i++ {
				dp[i] = min(d[from[i][0]], d[from[i][1]]) + bits.OnesCount8(v^i)
			}
		}
		Fprint(out, min(min(dp[0], dp[1]), min(dp[2], dp[3])))
	} else {
		from := [8][2]byte{{2, 5}, {3, 4}, {0, 7}, {1, 6}, {1, 6}, {0, 7}, {3, 4}, {2, 5}}
		dp := [8]int{}
		for i := byte(0); i < 8; i++ {
			dp[i] = bits.OnesCount8(a[0] ^ i)
		}
		for _, v := range a[1:] {
			d := dp
			for i := byte(0); i < 8; i++ {
				dp[i] = min(d[from[i][0]], d[from[i][1]]) + bits.OnesCount8(v^i)
			}
		}
		ans := dp[0]
		for _, v := range dp[1:] {
			ans = min(ans, v)
		}
		Fprint(out, ans)
	}
}

//func main() { CF1391D(os.Stdin, os.Stdout) }
