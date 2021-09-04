package main

import (
	. "fmt"
	"io"
	"math"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		for i := 0; i < k; i++ {
			if n&1 == 0 {
				n /= 2
			}
			if n%5 == 0 {
				n /= 5
			}
		}
		Fprintln(out, n*int(math.Pow10(k)))
	}
}

func main() { run(os.Stdin, os.Stdout) }
