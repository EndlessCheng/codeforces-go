package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF1003C(_r io.Reader, out io.Writer) {
	const eps = 1e-8
	in := bufio.NewReader(_r)

	var n, d int
	Fscan(in, &n, &d)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	l, r := 1.0, 5e3
o:
	for t := int(math.Log2((r - l) / eps)); t > 0; t-- {
		x := (l + r) / 2
		s := make([]float64, n+1)
		for i, v := range a {
			s[i+1] = s[i] + float64(v) - x
		}
		minI := 0
		for i, v := range s[:n+1-d] {
			if v < s[minI] {
				minI = i
			}
			if s[i+d] >= s[minI] {
				l = x
				continue o
			}
		}
		r = x
	}
	Fprintf(out, "%.8f", (l+r)/2)
}

//func main() { CF1003C(os.Stdin, os.Stdout) }
