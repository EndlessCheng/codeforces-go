package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1496A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	var s []byte
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		if 2*k == n {
			Fprintln(out, "NO")
			continue
		}
		for i, b := range s[:k] {
			if b != s[n-1-i] {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1496A(os.Stdin, os.Stdout) }
