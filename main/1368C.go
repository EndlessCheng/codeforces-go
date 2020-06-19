package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1368C(in io.Reader, _w io.Writer) {
	o := bufio.NewWriter(_w)
	defer o.Flush()
	p := Fprintln

	n := 0
	Fscan(in, &n)
	p(o, 3*n+4)
	p(o, 0, 0)
	p(o, 1, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			p(o, i+j, i+1)
		}
	}
	p(o, n, n+1)
	p(o, n+1, n+1)
}

//func main() { CF1368C(os.Stdin, os.Stdout) }
