package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1734D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n, k int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &s[i])
			s[i] += s[i-1]
		}
		l, r := k-1, k
		ls, rs := s[l], s[r]
		for l > 0 && r < n {
			if s[l-1] <= rs {
				ls = min(ls, s[l-1])
				l--
			} else if ls <= s[r+1] {
				rs = max(rs, s[r+1])
				r++
			} else {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1734D(os.Stdin, os.Stdout) }
