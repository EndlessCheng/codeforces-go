package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	const mx = 40
	var n, k, v, ans, s int
	Fscan(in, &n, &k)
	cnt := [mx]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		for j := uint(v); j > 0; j &= j - 1 {
			cnt[bits.TrailingZeros(j)]++
		}
	}

	f := [mx + 1]int{}
	for i, c := range cnt {
		f[i+1] = f[i] + max(c, n-c)<<i
	}

	k++
	for i := mx - 1; i >= 0; i-- {
		if k>>i&1 > 0 {
			ans = max(ans, s+cnt[i]<<i+f[i])
			s += (n - cnt[i]) << i
		} else {
			s += cnt[i] << i
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
