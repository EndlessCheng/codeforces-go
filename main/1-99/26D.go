package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf26D(in io.Reader, out io.Writer) {
	var n, m, k int
	Fscan(in, &n, &m, &k)
	if n+k < m {
		Fprint(out, 0)
		return
	}
	if k >= m {
		Fprint(out, 1)
		return
	}
	ans := 1.
	for i := range k + 1 {
		ans *= float64(m-i) / float64(n+1+i)
	}
	Fprintf(out, "%.6f", 1-ans)
}

//func main() { cf26D(os.Stdin, os.Stdout) }
