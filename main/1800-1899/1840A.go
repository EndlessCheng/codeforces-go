package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1840A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		ans := []byte{}
		last := s[0]
		for i := 1; i < n; i++ {
			if s[i] == last {
				ans = append(ans, last)
				i++
				if i < n {
					last = s[i]
				}
			}
		}
		Fprintf(out, "%s\n", ans)
	}
}

//func main() { CF1840A(os.Stdin, os.Stdout) }
