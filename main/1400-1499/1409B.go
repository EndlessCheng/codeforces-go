package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1409B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	f := func(a, b, x, y, n int64) int64 {
		d := min(a-x, n)
		a -= d
		n -= d
		d = min(b-y, n)
		b -= d
		n -= d
		return a * b
	}

	var T, a, b, x, y, n int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &x, &y, &n)
		Fprintln(out, min(f(a, b, x, y, n), f(b, a, y, x, n)))
	}
}

//func main() { CF1409B(os.Stdin, os.Stdout) }
