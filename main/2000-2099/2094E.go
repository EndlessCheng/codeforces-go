package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2094E(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		cnt := make([]int, 30)
		for i := range a {
			Fscan(in, &a[i])
			for m := uint(a[i]); m > 0; m &= m - 1 {
				cnt[bits.TrailingZeros(m)]++
			}
		}
		ans := 0
		for _, v := range a {
			s := 0
			for i, c := range cnt {
				if v>>i&1 == 0 {
					s += c << i
				} else {
					s += (n - c) << i
				}
			}
			ans = max(ans, s)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2094E(bufio.NewReader(os.Stdin), os.Stdout) }
