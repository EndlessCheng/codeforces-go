package main

import (
	. "fmt"
	"io"
	"math/bits"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	g := [62][]int{}
	var n, v int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		tz := bits.TrailingZeros(uint(v))
		g[tz] = append(g[tz], v)
	}
	for _, a := range g {
		sort.Ints(a)
		for i := len(a) - 1; i >= 0; i-- {
			Fprint(out, a[i], " ")
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
