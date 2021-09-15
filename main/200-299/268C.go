package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF268C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, m int
	Fscan(in, &n, &m)
	Fprintln(out, min(n, m)+1)
	for x, y := 0, m; x <= n && y >= 0; x, y = x+1, y-1 {
		Fprintln(out, x, y)
	}
}

//func main() { CF268C(os.Stdin, os.Stdout) }
