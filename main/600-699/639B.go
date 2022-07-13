package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF639B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, d, h int
	Fscan(in, &n, &d, &h)
	if d > h*2 || d == 1 && n > 2 {
		Fprint(out, -1)
		return
	}
	if n == 2 {
		Fprint(out, 1, 2)
		return
	}
	if h == 1 {
		for i := 2; i <= n; i++ {
			Fprintln(out, 1, i)
		}
		return
	}
	for i := 1; i <= h; i++ {
		Fprintln(out, i, i+1)
	}
	if h < d {
		Fprintln(out, 1, h+2)
		for i := h + 2; i <= d; i++ {
			Fprintln(out, i, i+1)
		}
	}
	for i := d + 2; i <= n; i++ {
		Fprintln(out, 2, i)
	}
}

//func main() { CF639B(os.Stdin, os.Stdout) }
