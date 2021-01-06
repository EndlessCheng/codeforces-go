package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1471A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		var mi, mx, x, v int64
		for Fscan(in, &n, &x); n > 0; n-- {
			Fscan(in, &v)
			mi += v
			mx += (v-1)/x + 1
		}
		Fprintln(out, (mi-1)/x+1, mx)
	}
}

//func main() { CF1471A(os.Stdin, os.Stdout) }
