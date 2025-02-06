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
	var n, k, ans, ban int
	Fscan(in, &n, &k)
	a := make([]struct{ v, w int }, n)
	for i := range a {
		Fscan(in, &a[i].v, &a[i].w)
	}
	k++
	for i := bits.Len(uint(k)) - 1; i >= 0; i-- {
		if k>>i&1 > 0 {
			ban ^= 1 << i
			s := 0
			for _, p := range a {
				if p.v&ban == 0 {
					s += p.w
				}
			}
			ans = max(ans, s)
		}
		ban ^= 1 << i
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
