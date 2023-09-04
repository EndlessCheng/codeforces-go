package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const eps = 1e-8

	var v0, x, y float64
	Fscan(in, &v0, &x, &y)
	f := func(t float64) float64 {
		return t + y/(v0+t*x)
	}
	l, r := 0., float64(y)+1
	for step := int(math.Log((r-l)/eps) / math.Log(1.5)); step > 0; step-- {
		m1 := l + (r-l)/3
		m2 := r - (r-l)/3
		v1, v2 := f(m1), f(m2)
		if v1 < v2 {
			r = m2
		} else {
			l = m1
		}
	}
	Fprintf(out, "%.15f\n", f((l+r)/2))
}

func main() { run(os.Stdin, os.Stdout) }
