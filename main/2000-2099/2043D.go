package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2043D(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var T, l, r, g int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r, &g)
		l = (l-1)/g + 1
		r /= g
		for d := range r - l + 1 {
			for i := range d + 1 {
				if gcd(l+i, r-(d-i)) == 1 {
					Fprintln(out, (l+i)*g, (r-(d-i))*g)
					continue o
				}
			}
		}
		Fprintln(out, -1, -1)
	}
}

//func main() { cf2043D(bufio.NewReader(os.Stdin), os.Stdout) }
