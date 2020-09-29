package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1426B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, a, b, c, d int
	for Fscan(in, &T); T > 0; T-- {
		ok := false
		for Fscan(in, &n, &m); n > 0; n-- {
			Fscan(in, &a, &b, &c, &d)
			ok = ok || b == c
		}
		if ok && m&1 == 0 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1426B(os.Stdin, os.Stdout) }
