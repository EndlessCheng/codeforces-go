package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf946G(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i] -= i
	}

	var f0, f1 []int
	for i := 1; i < n; i++ {
		v := a[i] + 1
		j := sort.SearchInts(f1, v+1)
		if j < len(f1) {
			f1[j] = v
		} else {
			f1 = append(f1, v)
		}

		v = a[i-1]
		j = sort.SearchInts(f0, v+1)
		if j < len(f0) {
			f0[j] = v
		} else {
			f0 = append(f0, v)
		}

		if j < len(f1) {
			f1[j] = min(f1[j], f0[j])
		}
	}

	Fprint(out, n-1-max(len(f0), len(f1)))
}

//func main() { cf946G(bufio.NewReader(os.Stdin), os.Stdout) }
