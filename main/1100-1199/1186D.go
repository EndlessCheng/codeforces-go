package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF1186D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]float64, n)
	s := .0
	for i := range a {
		Fscan(in, &a[i])
		s += a[i] - math.Floor(a[i])
	}
	si := int(math.Round(s))
	for _, v := range a {
		f := math.Floor(v)
		if si > 0 && v != f {
			si--
			Fprintln(out, int(f)+1)
		} else {
			Fprintln(out, int(f))
		}
	}
}

//func main() { CF1186D(os.Stdin, os.Stdout) }
