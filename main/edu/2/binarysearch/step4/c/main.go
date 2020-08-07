package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	const eps = 1e-8
	in := bufio.NewReader(_r)
	type pair struct{ a, b int }

	var n, k int
	Fscan(in, &n, &k)
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].a, &a[i].b)
	}
	l, r := 0.0, 1e5
	for step := int(math.Log2((r - l) / eps)); step > 0; step-- {
		mid := (l + r) / 2
		b := make([]float64, n)
		for i, p := range a {
			b[i] = float64(p.a) - mid*float64(p.b)
		}
		sort.Float64s(b)
		s := 0.0
		for _, v := range b[n-k:] {
			s += v
		}
		if s < 0 {
			r = mid
		} else {
			l = mid
		}
	}
	Fprintf(out, "%.8f", (l+r)/2)
}

func main() { run(os.Stdin, os.Stdout) }
