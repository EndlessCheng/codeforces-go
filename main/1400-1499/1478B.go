package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1478B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, q, d, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &q, &d)
		dp := make([]bool, d*10)
		dp[0] = true
		for i := range dp {
			for j := 0; j < d && 10*j+d <= i; j++ {
				dp[i] = dp[i] || dp[i-10*j-d]
			}
		}
		for ; q > 0; q-- {
			Fscan(in, &v)
			if v >= d*10 || dp[v] {
				Fprintln(out, "YES")
			} else {
				Fprintln(out, "NO")
			}
		}
	}
}

//func main() { CF1478B(os.Stdin, os.Stdout) }
