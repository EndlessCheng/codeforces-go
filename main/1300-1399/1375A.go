package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1375A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		for Fscan(in, &n); n > 0; n-- {
			if Fscan(in, &v); v > 0 == (n&1 > 0) {
				v = -v
			}
			Fprint(out, -v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1375A(os.Stdin, os.Stdout) }
