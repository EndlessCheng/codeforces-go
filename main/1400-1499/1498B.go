package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1498B(in io.Reader, out io.Writer) {
	var T, n, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &w)
		cnt := [20]int{}
		for i := 0; i < n; i++ {
			var v uint32
			Fscan(in, &v)
			cnt[bits.TrailingZeros32(v)]++
		}
		ans := 0
		for n > 0 {
			ans++
			left := w
			for i := 19; i >= 0; i-- {
				k := min(left>>i, cnt[i])
				cnt[i] -= k
				n -= k
				left -= k << i
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1498B(bufio.NewReader(os.Stdin), os.Stdout) }
