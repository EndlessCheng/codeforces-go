package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1323A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n int
o:
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		for i, v := range a {
			if v%2 == 0 {
				Fprintln(out, 1)
				Fprintln(out, i+1)
				continue o
			}
		}
		if len(a) < 2 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, 2)
			Fprintln(out, 1, 2)
		}
	}
}

//func main() { CF1323A(os.Stdin, os.Stdout) }
