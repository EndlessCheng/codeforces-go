package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF665D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mx int = 2e6
	np := [mx + 1]bool{}
	for i := 2; i <= mx; i++ {
		if !np[i] {
			for j := 2 * i; j <= mx; j += i {
				np[j] = true
			}
		}
	}

	var n, c1 int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] == 1 {
			c1++
		}
	}
	if c1 > 0 {
		for _, v := range a {
			if v > 1 && !np[v+1] {
				Fprintln(out, c1+1)
				Fprint(out, strings.Repeat("1 ", c1), v)
				return
			}
		}
		if c1 > 1 {
			Fprintln(out, c1)
			Fprint(out, strings.Repeat("1 ", c1))
			return
		}
	}
	for i, v := range a {
		for _, w := range a[i+1:] {
			if !np[v+w] {
				Fprint(out, "2\n", v, w)
				return
			}
		}
	}
	Fprint(out, "1\n", a[0])
}

//func main() { CF665D(os.Stdin, os.Stdout) }
