package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF235B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var p, l, dp float64
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &p)
		dp += (l*2 + 1) * p
		l = (l + 1) * p
	}
	Fprintf(out, "%.15f", dp)
}

//func main() { CF235B(os.Stdin, os.Stdout) }
