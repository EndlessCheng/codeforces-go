package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf626E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)

	s := make([]int, n+1)
	for i, v := range a {
		s[i+1] = s[i] + v
	}

	f := func(i, k int) (int, int) {
		m := k*2 + 1
		return s[i+1] - s[i-k] + s[n] - s[n-k] - a[i]*m, m
	}

	mx, bestI, bestK := 0., 0, 0
	for i := range n {
		k := sort.Search(min(i, n-1-i), func(k int) bool {
			a, b := f(i, k)
			c, d := f(i, k+1)
			return a*d > b*c
		})
		a, b := f(i, k)
		skew := float64(a) / float64(b)
		if skew > mx {
			mx, bestI, bestK = skew, i, k
		}
	}

	Fprintln(out, bestK*2+1)
	for _, v := range a[bestI-bestK : bestI+1] {
		Fprint(out, v, " ")
	}
	for _, v := range a[n-bestK:] {
		Fprint(out, v, " ")
	}
}

//func main() { cf626E(bufio.NewReader(os.Stdin), os.Stdout) }
