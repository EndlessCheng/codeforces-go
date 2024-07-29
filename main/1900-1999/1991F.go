package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

func cf1991F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, l int
	Fscan(in, &n, &q)
	a := make([]int32, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ok := func(l, r int) bool {
		if r-l < 6 {
			return false
		}
		b := slices.Clone(a[l:r])
		slices.Sort(b)
		for i := 2; i < len(b); i++ {
			for j := 0; j < i-1; j++ {
				if b[j]+b[i-1] <= b[i] {
					continue
				}
				c := append(slices.Clone(b[j+1:i-1]), b[i+1:]...)
				for k := 2; k < len(c); k++ {
					if c[k-2]+c[k-1] > c[k] {
						return true
					}
				}
			}
		}
		return false
	}
	minR := make([]int, n)
	r := 6
	for i := range minR {
		for r <= n && !ok(i, r) {
			r++
		}
		minR[i] = r
	}

	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		if r < minR[l-1] {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

//func main() { cf1991F(bufio.NewReader(os.Stdin), os.Stdout) }
