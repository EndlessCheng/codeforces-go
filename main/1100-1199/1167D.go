package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1167D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var s string
	Fscan(in, &n, &s)
	for i, b := range s {
		if i&1 > 0 == (b == '(') {
			Fprint(out, 1)
		} else {
			Fprint(out, 0)
		}
	}
}

//func main() { CF1167D(os.Stdin, os.Stdout) }
