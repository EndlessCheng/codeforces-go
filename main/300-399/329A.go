package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF329A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, cr, cc int
	var s string
	Fscan(in, &n)
	r := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &s)
		for j, b := range s {
			if b == '.' {
				if r[j] == 0 {
					r[j] = i + 1
					cr++
				}
				if c[i] == 0 {
					c[i] = j + 1
					cc++
				}
			}
		}
	}
	if cr == n {
		for i, v := range r {
			Fprintln(out, v, i+1)
		}
	} else if cc == n {
		for i, v := range c {
			Fprintln(out, i+1, v)
		}
	} else {
		Fprint(out, -1)
	}
}

//func main() { CF329A(os.Stdin, os.Stdout) }
