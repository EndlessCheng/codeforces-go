package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1632C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b)
		c1, c2 := 0, 0
		y := b
		for a|y != y {
			y++
			c1++
		}
		if a != y {
			c1++
		}
		for a|b != b {
			a++
			c2++
		}
		if a != b {
			c2++
		}
		if c2 < c1 {
			c1 = c2
		}
		Fprintln(out, c1)
	}
}

//func main() { CF1632C(os.Stdin, os.Stdout) }
