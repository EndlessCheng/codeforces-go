package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf262B(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var n, k, s int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	slices.Sort(a)
	mn := int(1e9)
	for i, v := range a {
		mn = min(mn, abs(v))
		if k > 0 && v < 0 {
			a[i] = -v
			k--
		}
		s += a[i]
	}

	if k%2 > 0 {
		s -= mn * 2
	}
	Fprint(out, s)
}

//func main() { cf262B(bufio.NewReader(os.Stdin), os.Stdout) }
