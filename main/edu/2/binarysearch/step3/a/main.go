package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	const eps = 1e-8
	in := bufio.NewReader(_r)
	type pair struct{ x, v float64 }
	var n int
	Fscan(in, &n)
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].v)
	}
	l, r := 0.0, 1e9
	for step := int(math.Log2((r - l) / eps)); step > 0; step-- {
		t := (l + r) / 2
		minR, maxL := 1e9, -1e9
		for _, p := range a {
			minR = math.Min(minR, p.x+p.v*t)
			maxL = math.Max(maxL, p.x-p.v*t)
		}
		if maxL <= minR {
			r = t
		} else {
			l = t
		}
	}
	Fprintf(out, "%.8f", (l+r)/2)
}

func main() { run(os.Stdin, os.Stdout) }
