package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2108B(in io.Reader, out io.Writer) {
	var T, n, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		if x == 0 {
			if n == 1 {
				Fprintln(out, -1)
			} else {
				Fprintln(out, n+n%2*3)
			}
		} else if x == 1 {
			Fprintln(out, n+(1-n%2)*3)
		} else {
			ones := bits.OnesCount(uint(x))
			ex := max(n-ones, 0)
			Fprintln(out, x+ex+ex%2)
		}
	}
}

//func main() { cf2108B(bufio.NewReader(os.Stdin), os.Stdout) }
