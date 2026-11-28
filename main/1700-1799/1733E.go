package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1733E(in io.Reader, out io.Writer) {
	var q, t, x, y int
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &t, &x, &y)
		if t < x+y {
			Fprintln(out, "NO")
			continue
		}

		a := make([]int, x+y+1)
		b := make([]int, x+y+1)
		a[0] = t - x - y
		b[0] = a[0] + 1
		for i := range x + y {
			for j := i; j >= 0; j-- {
				a[j+1] += a[j] / 2
				a[j] -= a[j] / 2
				b[j+1] += b[j] / 2
				b[j] -= b[j] / 2
			}
		}

		if a[x] != b[x] {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1733E(bufio.NewReader(os.Stdin), os.Stdout) }
