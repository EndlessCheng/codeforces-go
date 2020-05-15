package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1353A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, m int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		if n == 1 {
			Fprintln(out, 0)
		} else if n == 2 {
			Fprintln(out, m)
		} else {
			Fprintln(out, 2*m)
		}
	}
}

//func main() { CF1353A(os.Stdin, os.Stdout) }
