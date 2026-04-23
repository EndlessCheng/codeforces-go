package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2119C(in io.Reader, out io.Writer) {
	var T, n, l, r, k uint
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &l, &r, &k)
		if n%2 > 0 {
			Fprintln(out, l)
		} else if n == 2 || 1<<bits.Len(l) > r {
			Fprintln(out, -1)
		} else if k < n-1 {
			Fprintln(out, l)
		} else {
			Fprintln(out, 1<<bits.Len(l))
		}
	}
}

//func main() { cf2119C(bufio.NewReader(os.Stdin), os.Stdout) }
