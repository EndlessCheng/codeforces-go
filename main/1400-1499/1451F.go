package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1451F(in io.Reader, out io.Writer) {
	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n+m)
		for i := range n {
			for j := range m {
				Fscan(in, &v)
				a[i+j] ^= v
			}
		}
		if slices.Max(a) > 0 {
			Fprintln(out, "Ashish")
		} else {
			Fprintln(out, "Jeel")
		}
	}
}

//func main() { cf1451F(bufio.NewReader(os.Stdin), os.Stdout) }
