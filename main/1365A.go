package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1365A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var t, n, m int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		r, c := 0, 0
		a := make([][]byte, n)
		for i := range a {
			a[i] = make([]byte, m)
			for j := range a[i] {
				Fscan(in, &a[i][j])
			}
			if !bytes.Contains(a[i], []byte{1}) {
				r++
			}
		}
	o:
		for j := 0; j < m; j++ {
			for _, r := range a {
				if r[j] > 0 {
					continue o
				}
			}
			c++
		}
		if c < r {
			r = c
		}
		if r&1 == 1 {
			Fprintln(out, "Ashish")
		} else {
			Fprintln(out, "Vivek")
		}
	}
}

//func main() { CF1365A(os.Stdin, os.Stdout) }
