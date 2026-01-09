package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2162C(in io.Reader, out io.Writer) {
	var T, a, b uint
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b)
		n, m := bits.Len(a), bits.Len(b)
		if m > n {
			Fprintln(out, -1)
		} else if m == n {
			Fprintln(out, 1, a^b)
		} else {
			hb := uint(1) << (n - 1)
			Fprintln(out, 2, a^b^hb, hb)
		}
	}
}

//func main() { cf2162C(bufio.NewReader(os.Stdin), os.Stdout) }
