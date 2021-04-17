package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1509B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		m := strings.Count(s, "M")
		if m*3 != n {
			Fprintln(out, "NO")
			continue
		}
		c := 0
		for _, b := range s {
			if b == 'T' {
				c++
			} else {
				c--
				if c < 0 {
					Fprintln(out, "NO")
					continue o
				}
			}
		}
		c = 0
		for i := n - 1; i >= 0; i-- {
			if s[i] == 'T' {
				c++
			} else {
				c--
				if c < 0 {
					Fprintln(out, "NO")
					continue o
				}
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1509B(os.Stdin, os.Stdout) }
