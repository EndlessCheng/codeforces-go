package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf45I(in io.Reader, out io.Writer) {
	var n, neg, z int
	Fscan(in, &n)
	a := make([]int, n)
	mxI := -1
	for i := range a {
		Fscan(in, &a[i])
		if a[i] < 0 {
			neg++
			if mxI < 0 || a[i] > a[mxI] {
				mxI = i
			}
		} else if a[i] == 0 {
			z++
		}
	}
	if n == 1 {
		Fprint(out, a[0])
		return
	}
	if z == n || neg == 1 && z == n-1 {
		Fprint(out, 0)
		return
	}
	for i, v := range a {
		if v != 0 && (neg%2 == 0 || i != mxI) {
			Fprint(out, v, " ")
		}
	}
}

//func main() { cf45I(os.Stdin, os.Stdout) }
