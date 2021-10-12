package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1504B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s, t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		for ; n > 0 && s[n-1] == t[n-1]; n-- {
		}
		c := [2]int{}
		for i, b := range s[:n] {
			if c[0] != c[1] && b == t[i] != (s[i-1] == t[i-1]) {
				break
			}
			c[b&1]++
		}
		if c[0] != c[1] {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

//func main() { CF1504B(os.Stdin, os.Stdout) }
