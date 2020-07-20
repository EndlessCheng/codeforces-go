package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1379A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const a = "abacaba"
	var T, n int
	var s string
O:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
	o:
		for r := 7; r <= n; r++ {
			for l := r - 7; l < r; l++ {
				if s[l] != '?' && s[l] != a[l-r+7] {
					continue o
				}
			}
			t, f := s[:r-7]+a+s[r:], false
			for i := 7; i <= n; i++ {
				if t[i-7:i] == a {
					if f {
						continue o
					}
					f = true
				}
			}
			if f {
				Fprintln(out, "Yes")
				Fprintln(out, strings.ReplaceAll(t, "?", "d"))
				continue O
			}
		}
		Fprintln(out, "No")
	}
}

//func main() { CF1379A(os.Stdin, os.Stdout) }
