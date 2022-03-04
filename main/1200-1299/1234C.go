package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1234C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		f := [2]bool{true}
		for i, v := range s {
			w := t[i]
			if v < '3' || w < '3' {
				f[0] = f[0] && v < '3'
				f[1] = f[1] && w < '3'
			} else {
				f[0], f[1] = f[1], f[0]
			}
		}
		if f[1] {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1234C(os.Stdin, os.Stdout) }
