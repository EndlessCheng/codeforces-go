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
	var n, ans int
	Fscan(in, &n)
	cnt := make([]int, 30)
	a := make([]uint, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		for s := a[i]; s > 0; s &= s - 1 {
			cnt[bits.TrailingZeros(s)]++
		}
	}
	for _, v := range a {
		s := 0
		for i, c := range cnt {
			if v>>i&1 > 0 {
				c = n - c
			}
			s += c << i
		}
		ans = max(ans, s)
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
