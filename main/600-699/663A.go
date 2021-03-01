package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF663A(in io.Reader, out io.Writer) {
	var s string
	var x, y, n int
	a := []string{}
	for {
		Fscan(in, &s, &s)
		if s == "-" {
			y++
		} else {
			x++
		}
		if s == "=" {
			break
		}
		a = append(a, s)
	}
	Fscan(in, &n)
	if n+y > n*x || x > n+n*y {
		Fprint(out, "Impossible")
		return
	}
	Fprintln(out, "Possible")
	if n+y >= x {
		q := (n + y) / x
		c := (n + y) % x
		if c > 0 {
			Fprint(out, q+1)
			c--
		} else {
			Fprint(out, q)
		}
		for _, s := range a {
			if s == "+" {
				if c > 0 {
					Fprint(out, " + ", q+1)
					c--
				} else {
					Fprint(out, " + ", q)
				}
			} else {
				Fprint(out, " - 1")
			}
		}
	} else {
		q := (x - n) / y
		c := (x - n) % y
		Fprint(out, 1)
		for _, s := range a {
			if s == "-" {
				if c > 0 {
					Fprint(out, " - ", q+1)
					c--
				} else {
					Fprint(out, " - ", q)
				}
			} else {
				Fprint(out, " + 1")
			}
		}
	}
	Fprint(out, " = ", n)
}

//func main() { CF663A(os.Stdin, os.Stdout) }
