package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1381B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, 2*n)
		pos := []int{0}
		for i := range a {
			Fscan(in, &a[i])
			if a[i] > a[pos[len(pos)-1]] {
				pos = append(pos, i)
			}
		}
		pos = append(pos, 2*n)
		a = nil
		for i := 1; i < len(pos); i++ {
			a = append(a, pos[i]-pos[i-1])
		}
		dp := make([]bool, n+1)
		dp[0] = true
		for _, v := range a {
			for s := n; s >= v; s-- {
				dp[s] = dp[s] || dp[s-v]
			}
		}
		if dp[n] {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1381B(os.Stdin, os.Stdout) }
