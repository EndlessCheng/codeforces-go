package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1593D1(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var T, n, v0, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &v0)
		g := 0
		for ; n > 1; n-- {
			Fscan(in, &v)
			g = gcd(g, abs(v-v0))
		}
		if g == 0 {
			g = -1
		}
		Fprintln(out, g)
	}
}

//func main() { CF1593D1(os.Stdin, os.Stdout) }
