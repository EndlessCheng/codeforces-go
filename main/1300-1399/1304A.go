package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1304A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func() (ans int) {
		var x, y, a, b int
		Fscan(in, &x, &y, &a, &b)
		y -= x
		a += b
		if y%a != 0 {
			return -1
		}
		return y / a
	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		Fprintln(out, solve())
	}
}

//func main() { CF1304A(os.Stdin, os.Stdout) }
