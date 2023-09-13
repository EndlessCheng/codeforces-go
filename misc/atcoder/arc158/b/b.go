package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	i := sort.SearchInts(a, 0)
	if i < n-6 {
		a = append(a[:i+3], a[n-3:]...)
	}
	if i > 6 {
		a = append(a[:3], a[i-3:]...)
	}
	mn, mx := 3., -3.
	for i, x := range a {
		for j := i + 1; j < len(a); j++ {
			y := a[j]
			for k := j + 1; k < len(a); k++ {
				z := a[k]
				res := float64(x+y+z) / float64(x*y*z)
				mn = math.Min(mn, res)
				mx = math.Max(mx, res)
			}
		}
	}
	Fprintf(out, "%.15f\n", mn)
	Fprintf(out, "%.15f\n", mx)
}

func main() { run(os.Stdin, os.Stdout) }
