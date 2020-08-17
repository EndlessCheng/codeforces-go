package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1359C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, a, b, c int64
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &c, &a, &b)
		d := 2*b - a - c
		if d <= 0 {
			Fprintln(out, 2)
			continue
		}
		x := (c - b) / d
		p, q := a*x+c*(x+1), 2*x+1
		v, w := a*(x+1)+c*(x+2), 2*x+3
		if p*w+q*v <= 2*b*q*w {
			Fprintln(out, q)
		} else {
			Fprintln(out, w)
		}
	}
}

//func main() { CF1359C(os.Stdin, os.Stdout) }
