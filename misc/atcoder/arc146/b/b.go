package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, k, m, ans int
	Fscan(in, &n, &k, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for bt := 30; bt >= 0; bt-- {
		t := ans | 1<<bt
		for i, x := range a {
			j := bits.Len(uint(t &^ x))
			msk := 1<<j - 1
			b[i] = t&msk - x&msk
		}

		slices.Sort(b)
		s := 0
		for _, x := range b[:m] {
			s += x
		}
		if s <= k {
			ans = t
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
