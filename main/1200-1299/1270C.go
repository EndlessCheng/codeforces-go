package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1270C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int64
	for Fscan(in, &T); T > 0; T-- {
		var s, xor int64
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			s += v
			xor ^= v
		}
		Fprintln(out, 2)
		Fprintln(out, xor, s+xor)
	}
}

//func main() { CF1270C(os.Stdin, os.Stdout) }
