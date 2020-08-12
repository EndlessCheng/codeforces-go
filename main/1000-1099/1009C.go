package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1009C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, s, x, d int64
	Fscan(in, &n, &m)
	k := n / 2
	for ; m > 0; m-- {
		Fscan(in, &x, &d)
		s += x * n
		if d > 0 {
			s += d * n * (n - 1) / 2
		} else {
			s += d * k * (k + n&1)
		}
	}
	Fprintf(out, "%.15f", float64(s)/float64(n))
}

//func main() { CF1009C(os.Stdin, os.Stdout) }
