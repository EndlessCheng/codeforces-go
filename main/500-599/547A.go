package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF547A(in io.Reader, out io.Writer) {
	var m, h1, h2, a1, a2, x1, x2, y1, y2, p, q, d1, d2 int64
	Fscan(in, &m, &h1, &a1, &x1, &y1, &h2, &a2, &x2, &y2)
	for i := int64(1); i <= 2*m; i++ {
		if h1 = (h1*x1 + y1) % m; h1 == a1 {
			if p == 0 {
				p = i
			} else if d1 == 0 {
				d1 = i - p
			}
		}
		if h2 = (h2*x2 + y2) % m; h2 == a2 {
			if q == 0 {
				q = i
			} else if d2 == 0 {
				d2 = i - q
			}
		}
	}

	if p == 0 || q == 0 {
		Fprint(out, -1)
		return
	}

	for m *= 2; +m > 0; m-- {
		if p == q {
			Fprint(out, p)
			return
		}
		if p < q {
			p += d1
		} else {
			q += d2
		}
	}
	Fprint(out, -1)
}

//func main() { CF547A(os.Stdin, os.Stdout) }
