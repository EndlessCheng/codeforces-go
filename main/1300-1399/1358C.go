package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1358C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, x1, y1, x, y int64
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &x1, &y1, &x, &y)
		x -= x1
		y -= y1
		if x > y {
			x, y = y, x
		}
		Fprintln(out, x*y+1)
	}
}

//func main() { CF1358C(os.Stdin, os.Stdout) }
