package main

import (
	. "fmt"
	"io"
	. "math"
)

// github.com/EndlessCheng/codeforces-go
func CF452B(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	mi, mx := float64(n), float64(m)
	if mi > mx {
		mi, mx = mx, mi
	}
	if mi == 0 {
		if m > n {
			Fprintf(out, "0 1\n0 %d\n0 0\n0 %d", m, m-1)
		} else {
			Fprintf(out, "1 0\n%d 0\n0 0\n%d 0", n, n-1)
		}
	} else if 2*Hypot(mi-1, mx) > mx+Hypot(mi, mx) {
		if m > n {
			Fprintf(out, "1 0\n%d %d\n0 0\n%d %d", n, m, n-1, m)
		} else {
			Fprintf(out, "0 1\n%d %d\n0 0\n%d %d", n, m, n, m-1)
		}
	} else {
		if m > n {
			Fprintf(out, "0 0\n%d %d\n%d 0\n0 %d", n, m, n, m)
		} else {
			Fprintf(out, "0 0\n%d %d\n0 %d\n%d 0", n, m, m, n)
		}
	}
}

//func main() { CF452B(os.Stdin, os.Stdout) }
