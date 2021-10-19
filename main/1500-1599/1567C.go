package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1567C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a, b := 1, 1
		for p10 := 1; n > 0; p10 *= 10 {
			a += n % 10 * p10
			n /= 10
			b += n % 10 * p10
			n /= 10
		}
		Fprintln(out, a*b-2)
	}
}

//func main() { CF1567C(os.Stdin, os.Stdout) }
