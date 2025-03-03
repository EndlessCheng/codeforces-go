package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf82D(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	fr := make([][]int, n)
	for i := range fr {
		fr[i] = make([]int, n)
	}
	var f []int
	if n%2 > 0 {
		f = slices.Clone(a)
	} else {
		f = make([]int, n)
		for i, x := range a {
			f[i] = max(x, a[n-1])
		}
	}
	for i := n - 3 + n%2; i > 0; i -= 2 {
		b, c := a[i], a[i+1]
		for j, a := range a[:i] {
			r1 := f[j] + max(b, c)
			r2 := f[i] + max(a, c)
			r3 := f[i+1] + max(a, b)
			fr[i][j] = j
			if r2 < r1 {
				fr[i][j] = i
				r1 = r2
			}
			if r3 < r1 {
				fr[i][j] = i + 1
				r1 = r3
			}
			f[j] = r1
		}
	}
	Fprintln(out, f[0])

	j := 0
	for i := 1; i < n-2+n%2; i += 2 {
		j2 := fr[i][j]
		if j2 == j {
			Fprintln(out, i+1, i+2)
		} else if j2 == i {
			Fprintln(out, j+1, i+2)
		} else {
			Fprintln(out, j+1, i+1)
		}
		j = j2
	}
	if n%2 > 0 {
		Fprintln(out, j+1)
	} else {
		Fprintln(out, j+1, n)
	}
}

//func main() { cf82D(os.Stdin, os.Stdout) }
