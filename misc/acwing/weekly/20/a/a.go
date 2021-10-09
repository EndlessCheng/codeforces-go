package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var T, a, b, c, d, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c, &d, &k)
		x, y := (a+c-1)/c, (b+d-1)/d
		if x+y > k {
			Fprintln(out, -1)
		} else {
			Fprintln(out, x, y)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
