package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF441C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k int
	Fscan(in, &n, &m, &k)
	avg := n * m / k
	r := avg + n*m%k
	Fprint(out, r, " ")
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if r == 0 {
				r = avg
				Fprintln(out)
				Fprint(out, r, " ")
			}
			if i&1 > 0 {
				Fprint(out, i, j, " ")
			} else {
				Fprint(out, i, m+1-j, " ")
			}
			r--
		}
	}
}

//func main() { CF441C(os.Stdin, os.Stdout) }
