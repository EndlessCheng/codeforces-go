package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1657C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		c := 0
		for len(s) > 1 {
			if s[0] == '(' {
				s = s[2:]
			} else {
				i := strings.IndexByte(s[1:], ')')
				if i < 0 {
					break
				}
				s = s[i+2:]
			}
			c++
		}
		Fprintln(out, c, len(s))
	}
}

//func main() { CF1657C(os.Stdin, os.Stdout) }
