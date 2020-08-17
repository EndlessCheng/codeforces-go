package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1379B(in io.Reader, out io.Writer) {
	var t, l, r, m int64
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &l, &r, &m)
		ml, mr := m+l-r, m+r-l
		for a := l; a <= r; a++ {
			if na := mr - mr%a; na >= ml {
				if na < m {
					Fprintln(out, a, r, r+na-m)
				} else {
					Fprintln(out, a, l, l+na-m)
				}
				break
			}
		}
	}
}

//func main() { CF1379B(os.Stdin, os.Stdout) }
