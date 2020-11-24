package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	for {
		var n, m, ans int
		Fscan(in, &n, &m)
		if n == 0 {
			break
		}
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		c := make([]int, n)
		for i := range c {
			Fscan(in, &c[i])
		}
		dp := make([]bool, m+1)
		dp[0] = true
		for i, v := range a {
			used := make([]int, m+1)
			for j := v; j <= m; j++ {
				if !dp[j] && dp[j-v] && used[j-v] < c[i] {
					dp[j] = true
					used[j] = used[j-v] + 1
				}
			}
		}
		for _, b := range dp[1:] {
			if b {
				ans++
			}
		}
		Fprintln(out, ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }
