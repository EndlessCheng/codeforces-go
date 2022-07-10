package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1697C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s, t string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		if strings.ReplaceAll(s, "b", "") != strings.ReplaceAll(t, "b", "") {
			Fprintln(out, "NO")
			continue
		}
		j := 0
		for i, c := range s {
			if c != 'b' {
				for t[j] == 'b' {
					j++
				}
				if i != j && c == 'c' == (i < j) {
					Fprintln(out, "NO")
					continue o
				}
				j++
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1697C(os.Stdin, os.Stdout) }
