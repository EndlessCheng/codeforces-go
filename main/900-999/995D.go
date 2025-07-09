package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf995D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, r, z, g, s int
	Fscan(in, &n, &r)
	m := 1 << n
	a := make([]int, m)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
	}
	Fprintf(out, "%.6f\n", float64(s)/float64(m))
	for range r {
		Fscan(in, &z, &g)
		s += g - a[z]
		a[z] = g
		Fprintf(out, "%.6f\n", float64(s)/float64(m))
	}
}

//func main() { cf995D(bufio.NewReader(os.Stdin), os.Stdout) }
