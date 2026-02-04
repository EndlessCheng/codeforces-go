package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2180C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, k)
		for i := range a {
			a[i] = n
		}

		if k%2 == 0 {
			free := 0
			for i := bits.Len(uint(n)) - 1; i >= 0; i-- {
				if n>>i&1 > 0 {
					a[min(free, k-1)] ^= 1 << i
					free++
				} else {
					for j := range min(free&^1, k) {
						a[j] |= 1 << i
					}
				}
			}
		}

		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2180C(bufio.NewReader(os.Stdin), os.Stdout) }
