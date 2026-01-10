package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1988C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n&(n-1) == 0 {
			Fprintln(out, 1)
			Fprintln(out, n)
			continue
		}
		Fprintln(out, bits.OnesCount(uint(n))+1)
		for i := bits.Len(uint(n)) - 1; i >= 0; i-- {
			if n>>i&1 > 0 {
				Fprint(out, n^1<<i, " ")
			}
		}
		Fprintln(out, n)
	}
}

//func main() { cf1988C(os.Stdin, os.Stdout) }
