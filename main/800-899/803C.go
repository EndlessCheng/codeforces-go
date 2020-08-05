package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF803C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k, g int64
	Fscan(in, &n, &k)
	if k > 1e6 {
		Fprint(out, -1)
		return
	}
	s := k * (k + 1) / 2
	for d := int64(1); d*d <= n; d++ {
		if n%d == 0 {
			if d >= s {
				g = n / d
				break
			}
			if n/d >= s {
				g = d
			}
		}
	}
	if g == 0 {
		Fprint(out, -1)
		return
	}
	for i := int64(1); i < k; i++ {
		Fprint(out, i*g, " ")
		n -= i * g
	}
	Fprint(out, n)
}

//func main() { CF803C(os.Stdin, os.Stdout) }
