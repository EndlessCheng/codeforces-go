package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1375C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, x, y int
	for Fscan(in, &T); T > 0; T-- {
		for Fscan(in, &n, &x); n > 1; n-- {
			Fscan(in, &y)
		}
		if x < y {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1375C(os.Stdin, os.Stdout) }
