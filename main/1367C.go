package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1367C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, k int
	var s []byte
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k, &s)
		c, p := 0, -1
		for i, b := range s {
			if b == '0' && (p == -1 || i-p > k) {
				c++
				p = i
			} else if b == '1' {
				if p >= 0 && s[p] == '0' && i-p <= k {
					c--
				}
				p = i
			}
		}
		Fprintln(out, c)
	}
}

//func main() { CF1367C(os.Stdin, os.Stdout) }
