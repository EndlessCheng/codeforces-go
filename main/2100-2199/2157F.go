package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2157F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const n = 250000
	const B = 60
	Fprintln(out, n-1)
	for i := 1; i < B; i++ {
		for j := n - 1; j > 0; j-- {
			if j%B == i {
				Fprintln(out, j, 1)
			}
		}
	}
	for i := 1; i < B; i++ {
		for j := n - 1; j > 0; j-- {
			if j%B == 0 && j/B%B == i {
				Fprintln(out, j, B)
			}
		}
	}
	for i := B * B; i < n; i += B * B {
		Fprintln(out, i, B*B)
	}
}

//func main() { cf2157F(nil, os.Stdout) }
