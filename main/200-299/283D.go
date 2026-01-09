package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf283D(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]uint, n+1)
	e2 := make([]int, n+1)
	for i := range n {
		Fscan(in, &a[i])
		e2[i] = bits.TrailingZeros(a[i])
		a[i] >>= e2[i]
	}
	a[n] = 1
	n++

	f := make([]int, n+1)
	for i, e := range e2 {
		f[i] = i
		for j := range i {
			if (e <= i-j-1 || e-e2[j] == i-j) && a[j]%a[i] == 0 {
				f[i] = min(f[i], f[j]+i-j-1)
			}
		}
	}
	Fprint(out, f[n-1])
}

//func main() { cf283D(bufio.NewReader(os.Stdin), os.Stdout) }
