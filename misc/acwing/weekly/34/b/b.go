package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, sum int
	Fscan(in, &n)
	for k := 2; k < n; k++ {
		for x := n; x > 0; x /= k {
			sum += x % k
		}
	}
	g := gcd(sum, n-2)
	Fprint(out, sum/g, "/", (n-2)/g)
}

func main() { run(os.Stdin, os.Stdout) }

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
