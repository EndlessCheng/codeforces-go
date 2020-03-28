package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1183B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, k, v int
	for Fscan(in, &t); t > 0; t-- {
		mi, mx := int(1e9), 0
		for Fscan(in, &n, &k); n > 0; n-- {
			Fscan(in, &v)
			if v < mi {
				mi = v
			}
			if v > mx {
				mx = v
			}
		}
		if mi+2*k >= mx {
			Fprintln(out, mi+k)
		} else {
			Fprintln(out, -1)
		}
	}
}

//func main() { CF1183B(os.Stdin, os.Stdout) }
