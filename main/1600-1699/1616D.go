package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1616D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n int
	var v, x int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s := make([]int64, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			s[i] = s[i-1] + v
		}
		Fscan(in, &x)
		ans, mx, pre := n, int64(0), 0
		for i := range s {
			s[i] -= int64(i) * x
			if i >= pre+2 {
				if s[i-2] > mx {
					mx = s[i-2]
				}
				if mx > s[i] {
					mx, pre = -1e18, i
					ans--
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1616D(os.Stdin, os.Stdout) }
