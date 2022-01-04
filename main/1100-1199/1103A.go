package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1103A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(in, &s)
	c0, c1 := 0, 0
	for _, b := range s {
		if b == '0' {
			Fprintln(out, 1, c0%4+1)
			c0++
		} else {
			Fprintln(out, 3, c1%2*2+1)
			c1++
		}
	}
}

//func main() { CF1103A(os.Stdin, os.Stdout) }
