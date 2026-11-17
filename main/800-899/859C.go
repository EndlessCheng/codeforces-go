package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf859C(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var n, s int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
	}

	f := 0
	for i := n - 1; i >= 0; i-- {
		f = abs(f - a[i])
	}

	Fprint(out, (s-f)/2, (s+f)/2)
}

//func main() { cf859C(os.Stdin, os.Stdout) }
