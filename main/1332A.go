package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1332A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, a, b, c, d, x, y, x1, y1, x2, y2 int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &a, &b, &c, &d, &x, &y, &x1, &y1, &x2, &y2)
		d1, d2 := b-a, d-c
		x += d1
		y += d2
		if (d1 == 0 && (a == 0 || x1 < x || x < x2) || d1 != 0 && x1 <= x && x <= x2) &&
			(d2 == 0 && (c == 0 || y1 < y || y < y2) || d2 != 0 && y1 <= y && y <= y2) {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

//func main() { CF1332A(os.Stdin, os.Stdout) }
