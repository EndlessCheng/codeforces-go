package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1870B(in io.Reader, out io.Writer) {
	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		var or, x1, x2 int
		for range m {
			Fscan(in, &v)
			or |= v
		}
		for _, v := range a {
			x1 ^= v
			x2 ^= v | or
		}
		Fprintln(out, min(x1, x2), max(x1, x2))
	}
}

//func main() { cf1870B(bufio.NewReader(os.Stdin), os.Stdout) }
