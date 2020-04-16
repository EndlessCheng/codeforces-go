package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF833A(_r io.Reader, _w io.Writer) {
	r := bufio.NewReader(_r)
	w := bufio.NewWriter(_w)
	defer w.Flush()
	cbrt := func(a int64) int64 {
		r := int64(math.Round(math.Cbrt(float64(a))))
		if r*r*r == a {
			return r
		}
		return -1
	}
	var t, a, b int64
	for Fscan(r, &t); t > 0; t-- {
		Fscan(r, &a, &b)
		r := cbrt(a * b)
		if r == -1 || a%r != 0 || b%r != 0 {
			Fprintln(w, "No")
		} else {
			Fprintln(w, "Yes")
		}
	}
}

//func main() { CF833A(os.Stdin, os.Stdout) }
