package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1965B(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		hb := 1 << (bits.Len(uint(k)) - 1)
		a := []any{k - hb, k + 1, hb<<1 | k}
		for i := 1; i <= n; i <<= 1 {
			if i != hb {
				a = append(a, i)
			}
		}
		Fprintln(out, len(a))
		Fprintln(out, a...)
	}
}

//func main() { cf1965B(bufio.NewReader(os.Stdin), os.Stdout) }
