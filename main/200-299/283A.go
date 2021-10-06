package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF283A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var q, op, x, y, tot int64
	a, s := []int64{0}, []int64{0}
	for Fscan(in, &q); q > 0; q-- {
		if Fscan(in, &op); op == 1 {
			Fscan(in, &x, &y)
			tot += x * y
			s[x-1] += y
		} else if op == 2 {
			Fscan(in, &x)
			tot += x
			s = append(s, 0)
			a = append(a, x)
		} else {
			tot -= s[len(s)-1] + a[len(a)-1]
			s[len(s)-2] += s[len(s)-1]
			s = s[:len(s)-1]
			a = a[:len(a)-1]
		}
		Fprintf(out, "%.8f\n", float64(tot)/float64(len(a)))
	}
}

//func main() { CF283A(os.Stdin, os.Stdout) }
